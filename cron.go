package main

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

type Cron struct {
	Second  bool
	Minute  bool
	Hour    bool
	Day     bool
	Week    bool
	Every   uint64
	At      string
	Actions map[string]func()
}

func NewCron(c Cron) *Cron {
	return &Cron{
		Second:  c.Second,
		Minute:  c.Minute,
		Hour:    c.Hour,
		Day:     c.Day,
		Week:    c.Week,
		Every:   c.Every,
		At:      c.At,
		Actions: c.Actions,
	}
}

func (c *Cron) StartCron() {
	sch := gocron.NewScheduler()
	job := sch.Every(c.Every)

	if c.Second {
		job = job.Second()
	} else if c.Minute {
		job = job.Minute()
	} else if c.Hour {
		job = job.Hour()
	} else if c.Day {
		job = job.Day()
	}

	if c.At == "" {
		job.DoSafely(func(a string) { fmt.Println(a) }, "horse")
	} else {
		job.At(c.At).DoSafely(c.Actions)
	}

	<-sch.Start()
}

func removeTask(task func()) {
	gocron.Remove(task)
}

func clearAllJobs() {
	gocron.Clear()
}
