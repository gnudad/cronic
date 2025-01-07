package main

import (
	"github.com/go-co-op/gocron/v2"
)

func NewScheduler(cronic *Cronic) gocron.Scheduler {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}
	for key, job := range cronic.Jobs {
		job.key = key
		j, err := s.NewJob(
			job.JobDefinition(),
			job.Task(),
		)
		if err != nil {
			panic(err)
		}
		job.jobID = j.ID()
	}
	return s
}
