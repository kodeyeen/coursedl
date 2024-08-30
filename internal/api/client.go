package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	baseURL   = "https://www.designgurus.io"
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:129.0) Gecko/20100101 Firefox/129.0"
)

type Client struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) GetCourses(ctx context.Context) (*Courses, error) {
	coursesURL := fmt.Sprintf("%s/_next/data/N96LXRk0ealEXbchXFTgO/en/courses.json", baseURL)

	req, err := http.NewRequest(http.MethodGet, coursesURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get courses (status code %d)", resp.StatusCode)
	}

	var courses Courses

	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&courses)
	if err != nil {
		return nil, err
	}

	return &courses, nil
}

func (c *Client) GetLessons(ctx context.Context, courseSlug string) (*Lessons, error) {
	lessonsURL := fmt.Sprintf("%s/api/course/getLessonsList/%s", baseURL, courseSlug)

	req, err := http.NewRequest(http.MethodGet, lessonsURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get lessons (status code %d)", resp.StatusCode)
	}

	var lessons Lessons

	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&lessons)
	if err != nil {
		return nil, err
	}

	return &lessons, nil
}

func (c *Client) GetDocument(ctx context.Context, documentID, courseID string) (*Document, error) {
	documentURL := fmt.Sprintf("%s/api/document/%s?courseId=%s", baseURL, documentID, courseID)

	req, err := http.NewRequest(http.MethodGet, documentURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get document (status code %d)", resp.StatusCode)
	}

	var document Document

	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&document)
	if err != nil {
		return nil, err
	}

	return &document, nil
}
