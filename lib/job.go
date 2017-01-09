package lib

import (
	"fmt"
	"time"
)

type Job struct {
	Id uint64
	Payload string
}

func NewJob(Id uint64, Payload string) *Job {
	return &Job{Id, Payload}
}

func (j *Job) Process() error {
	// stub implementation
	fmt.Println(fmt.Sprintf("Processing job with ID #%d", j.Id))
	time.Sleep(time.Second)

	return nil
}
