package cron

import (
	"context"
	"github.com/robfig/cron/v3"
	"github.com/zcubbs/x/logwrapper"
)

type Job struct {
	CronPattern string
	Name        string
	Task        func(ctx context.Context)
	log         logwrapper.Logger
}

type JobOption func(job *Job)

func WithLogger(logger logwrapper.Logger) JobOption {
	return func(job *Job) {
		job.log = logger
	}
}

func NewJob(cronPattern string, task func(ctx context.Context), options ...JobOption) *Job {
	job := &Job{
		CronPattern: cronPattern,
		Task:        task,
	}

	for _, option := range options {
		option(job)
	}

	if job.log == nil {
		job.log = logwrapper.NewLogrusLogger("cron")
	}

	return job
}

func (job *Job) Start() {
	ctx := context.Background()

	if job.CronPattern == "" {
		job.log.Info("no cron pattern provided, not starting cron job")
		return
	}

	if job.CronPattern == "-" {
		job.log.Info("running cron job once")
		job.Task(ctx)
		job.log.Info("cron job finished")
		return
	}

	c := cron.New(cron.WithSeconds()) // cron with second-level precision
	_, err := c.AddFunc(job.CronPattern, func() {
		job.Task(ctx)
	})
	if err != nil {
		job.log.Error("cannot create cron job: %v", err)
	}

	job.log.Info("starting cron job: %s", job.Name)
	c.Start()
}
