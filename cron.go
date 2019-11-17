package main

import (
	"github.com/jasonlvhit/gocron"
)

type Cron struct {
	Second bool
	Minute bool
	Hour bool
	Day bool
	Week bool
	Every uint64
	At string
	Actions map[string]func()
}

func NewCron(c Cron) *Cron {
	return &Cron{
		Second:     c.Second,
		Minute:     c.Minute,
		Hour:       c.Hour,
		Day:        c.Day,
		Week:       c.Week,
		Every:      c.Every,
		At: c.At,
		Actions: 	c.Actions,
	}
}

func StartCron(c Cron) {
	sch := gocron.NewScheduler()
	evr := sch.Every(c.Every)

	if c.Second {
		evr = evr.Second()
	} else if c.Minute {
		evr = evr.Minute()
	} else if c.Hour {
		evr = evr.Hour()
	} else if c.Day {
		evr = evr.Day()
	}

	if c.At == "" {
		evr.DoSafely(c.Actions)
	} else {
		evr.At(c.At).DoSafely(c.Actions)
	}

	<- sch.Start()
}

func removeTask(task func()){
	gocron.Remove(task)
}

func clearAllJobs() {
	gocron.Clear()
}


