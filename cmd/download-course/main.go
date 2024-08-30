package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kodeyeen/coursedl/internal/api"
	"golang.org/x/sync/errgroup"
)

const (
	maxWorkerCnt = 3
)

var (
	outputDir = filepath.Join("resources")
)

// Grokking Modern Behavioral Interview
var courseFlag = flag.String("course", "", "course name")

func main() {
	flag.Parse()

	ctx := context.Background()
	client := api.NewClient(http.DefaultClient)

	courses, err := client.GetCourses(ctx)
	if err != nil {
		log.Fatalf("Failed to get courses: %v", err)
	}

	course, exists := courses.CourseByTitle(*courseFlag)
	if !exists {
		log.Fatal("No such course")
	}

	lessons, err := client.GetLessons(ctx, course.Slug)
	if err != nil {
		log.Fatalf("Failed to fetch course (%s) lessons: %v\n", course.Slug, err)
	}

	log.Printf("Lesson count: %d\n", len(lessons.Data))

	courseID := course.Courses[0]
	var lessonDocs []*api.LessonDocument

	for _, lesson := range lessons.Data {
		lessonDocs = append(lessonDocs, lesson.Documents...)
	}

	log.Printf("Document count: %d\n", len(lessonDocs))

	jobs := make(chan *Job)
	results := make(chan *Result, len(lessonDocs))

	g, ctx := errgroup.WithContext(ctx)

	for range maxWorkerCnt {
		g.Go(func() error {
			return worker(ctx, client, jobs, results)
		})
	}

loop:
	for _, lessonDoc := range lessonDocs {
		select {
		case <-ctx.Done():
			break loop
		default:
		}

		jobs <- &Job{
			DocumentID: lessonDoc.DocumentID,
			CourseID:   courseID,
		}
	}

	close(jobs)

	err = g.Wait()
	close(results)
	if err != nil {
		log.Fatalf("Failed to get all documents: %v\n", err)
	}

	// SAVE

	// save courses

	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	err = saveAsJSON(courses, filepath.Join(outputDir, "courses.json"))
	if err != nil {
		log.Fatalf("Failed to save courses: %v", err)
	}

	// course and documents dir

	courseDir := filepath.Join(outputDir, course.Slug)

	err = remkdir(courseDir)
	if err != nil {
		log.Fatalf("Failed to create course (%s) directory: %v\n", course.Slug, err)
	}

	docsDir := filepath.Join(courseDir, "documents")

	err = os.Mkdir(docsDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create course (%s) documents directory: %v\n", course.Slug, err)
	}

	// save lessons

	err = saveAsJSON(lessons, filepath.Join(courseDir, "lessons.json"))
	if err != nil {
		log.Fatalf("Failed to save course (%s) lessons: %v\n", course.Slug, err)
	}

	// save documents

	for result := range results {
		err = saveAsJSON(result.Document, filepath.Join(docsDir, result.Document.ID()+".json"))
		if err != nil {
			log.Printf("Failed to save document %v: %v\n", result.Document, err)
			continue
		}
	}

	log.Println("Success")
}

func worker(ctx context.Context, client *api.Client, jobs <-chan *Job, results chan<- *Result) error {
	for job := range jobs {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		log.Println("WORKING")

		doc, err := client.GetDocument(ctx, job.DocumentID, job.CourseID)
		if err != nil {
			return fmt.Errorf("failed to get document %s: %w", job.DocumentID, err)
		}

		results <- &Result{
			Document: doc,
		}
	}

	return nil
}

func remkdir(name string) error {
	err := os.RemoveAll(name)
	if err != nil {
		return err
	}

	return os.Mkdir(name, os.ModePerm)
}

func saveAsJSON(v any, name string) error {
	ouf, err := os.Create(name)
	if err != nil {
		return err
	}
	defer ouf.Close()

	enc := json.NewEncoder(ouf)
	enc.SetIndent("", "    ")

	return enc.Encode(v)
}
