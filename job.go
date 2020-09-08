package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Job struct {
	ID     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}

type IndexData struct {
	Jobs []*Job
}

// static list of existing jobs with respective statuses
// it will be normally on redis
var listJobs = []*Job{}

func createJobList() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 15; i++ {
		listJobs = append(listJobs, &Job{ID: uuid.New(), Status: randomStatus()})
	}
}

func randomStatus() string {
	statuses := []string{"waiting", "running", "failed", "succeeded"}
	return statuses[rand.Intn(len(statuses))]
}
