package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kodeyeen/coursedl/internal/api"
	"golang.org/x/sync/errgroup"
)

type Job struct {
	DocumentID string
	CourseID   string
}

type Result struct {
	Document *api.Document
}

type WorkerPool struct {
	maxWorkerCnt int
	jobs         chan *Job
	results      chan *Result
	client       *api.Client
	g            *errgroup.Group
	ctx          context.Context
}

func NewWorkerPool(maxWorkerCnt, jobCnt int) *WorkerPool {
	g, ctx := errgroup.WithContext(context.Background())

	return &WorkerPool{
		maxWorkerCnt: maxWorkerCnt,
		jobs:         make(chan *Job),
		results:      make(chan *Result, jobCnt),
		g:            g,
		ctx:          ctx,
	}
}

func (wp *WorkerPool) Start(ctx context.Context) {
	for range wp.maxWorkerCnt {
		wp.g.Go(func() error {
			return wp.Worker(ctx)
		})
	}
}

func (wp *WorkerPool) AddJob(job *Job) error {
	select {
	case <-wp.ctx.Done():
		return wp.ctx.Err()
	default:
	}

	wp.jobs <- job

	return nil
}

func (wp *WorkerPool) Worker(ctx context.Context) error {
	for job := range wp.jobs {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		log.Println("WORKING")

		doc, err := wp.client.GetDocument(ctx, job.DocumentID, job.CourseID)
		if err != nil {
			return fmt.Errorf("failed to get document %s: %w", job.DocumentID, err)
		}

		wp.results <- &Result{
			Document: doc,
		}
	}

	return nil
}
