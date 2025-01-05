package main

import (
	"fmt"
	"os/exec"

	"github.com/go-co-op/gocron/v2"
)

func LoadScheduler(config Config) gocron.Scheduler {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	for _, job := range config.Jobs {
		j, err := s.NewJob(
			gocron.CronJob(job.Cron, true),
			gocron.NewTask(
				func() {
					out, err := exec.Command("sh", "-c", job.Cmd).Output()
					if err != nil {
						panic(err)
					}
					fmt.Print(string(out))
				},
			),
		)
		if err != nil {
			panic(err)
		}
		job.JobID = j.ID()
	}
	return s
}
