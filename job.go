package main

import (
	"fmt"
	"os/exec"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
)

type Job struct {
	key   string
	Name  string
	Cron  string
	Cmd   string
	jobID uuid.UUID
}

func (job Job) JobDefinition() gocron.JobDefinition {
	return gocron.CronJob(job.Cron, true)

}

func (job Job) Task() gocron.Task {
	return gocron.NewTask(
		func() {
			out, err := exec.Command("sh", "-c", job.Cmd).Output()
			if err != nil {
				panic(err)
			}
			fmt.Print(string(out))
		},
	)
}
