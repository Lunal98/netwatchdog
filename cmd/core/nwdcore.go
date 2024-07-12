package core

import (
	"time"
)

type CheckRoutine struct {
	check    Check
	period   time.Duration
	name     string
	priority int
}

type NwdCore struct {
	checks []CheckRoutine
}

func (nwd *NwdCore) AddCheck(check Check, period time.Duration, name string, priority int) error {
	nwd.checks = append(
		nwd.checks,
		CheckRoutine{
			check:    check,
			period:   period,
			name:     name,
			priority: priority,
		},
	)
	return nil

}
func (nwd *NwdCore) Start() {
	var asd Scheduler

}
