package main

import (
	"time"

	"github.com/Lunal98/netwatchdog/internal/scheduler"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type CheckRoutine struct {
	checker  Checker
	period   time.Duration
	name     string
	priority int
	UUID     uuid.UUID
}

type NwdCore struct {
	checks    []CheckRoutine
	scheduler scheduler.Scheduler
}

func (nwd *NwdCore) AddCheck(checker Checker, period time.Duration, name string, priority int) error {
	uuid, err := nwd.scheduler.Addjob(checker.Check, 30*time.Second)
	if err != nil {
		log.Error().Err(err).Str("check name", name).Msg("Error Adding Check")
	}
	nwd.checks = append(
		nwd.checks,
		CheckRoutine{
			checker:  checker,
			period:   period,
			name:     name,
			priority: priority,
			UUID:     uuid,
		},
	)
	return nil

}
func (nwd *NwdCore) Start() {
	nwd.scheduler.Start()

}
