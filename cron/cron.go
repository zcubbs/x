package cron

import (
	"context"
	"github.com/robfig/cron/v3"
	"github.com/zcubbs/log"
	"github.com/zcubbs/log/structuredlogger"
)

type Job struct {
	CronPattern string
	Name        string
	Task        func(ctx context.Context)
	log         log.Logger
}

type JobOption func(job *Job)

func WithLogger(logger log.Logger) JobOption {
	return func(job *Job) {
		job.log = logger
	}
}

func NewJob(name, cronPattern string, task func(ctx context.Context), options ...JobOption) *Job {
	job := &Job{
		CronPattern: cronPattern,
		Task:        task,
		Name:        name,
	}

	for _, option := range options {
		option(job)
	}

	if job.log == nil {
		job.log = log.NewLogger(
			structuredlogger.StdLoggerType,
			"cron",
			structuredlogger.JSONFormat,
		)
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
