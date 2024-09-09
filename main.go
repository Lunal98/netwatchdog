package main

import (
	"context"
	"sync"
	"time"

	"github.com/Lunal98/netwatchdog/internal/checker"
	"github.com/Lunal98/netwatchdog/internal/remediationhandler"
	"github.com/Lunal98/netwatchdog/internal/scheduler"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type CheckRoutine struct {
	checker  checker.Checker
	period   time.Duration
	name     string
	priority int
	UUID     uuid.UUID
}

type NwdCore struct {
	checks    map[uuid.UUID]CheckRoutine
	scheduler scheduler.Scheduler
	ctx       context.Context
}

func (nwd *NwdCore) AddCheck(checker checker.Checker, period time.Duration, name string, priority int) error {
	uuid, err := nwd.scheduler.Addjob(checker.Check, 30*time.Second)
	if err != nil {
		log.Error().Err(err).Str("check name", name).Msg("Error Adding Check")
	}
	nwd.checks[uuid] = CheckRoutine{
		checker:  checker,
		period:   period,
		name:     name,
		priority: priority,
		UUID:     uuid,
	}
	return nil

}
func (nwd *NwdCore) Start() {
	if nwd.ctx == nil {
		nwd.ctx = context.Background()
	}
	nwd.scheduler.SetRemediator(nwd.handle)
	nwd.scheduler.Start(nwd.ctx)

}
func (nwd *NwdCore) handle(jobID uuid.UUID, jobName string, err error) {
	remediationhandler.Handle(jobID, jobName, err, nwd.getCheckers())
}
func (nwd *NwdCore) getCheckers() map[uuid.UUID]checker.Checker {
	checkers := make(map[uuid.UUID]checker.Checker)
	for uuid, check := range nwd.checks {
		checkers[uuid] = check.checker
	}
	return checkers
}
func findTopPriority(nwd *NwdCore) uuid.UUID {
	maxprio := -1
	var maxpriouuid uuid.UUID
	lastresult := RunAllChecks(nwd)
	for uuid, checkR := range nwd.checks {
		if lastresult[uuid] != nil {

		}
		if checkR.priority > maxprio {
			maxpriouuid = uuid
			maxprio = checkR.priority
		}

	}
	return maxpriouuid
}
func RunAllChecks(nwd *NwdCore) map[uuid.UUID]error {

	lastresult := make(map[uuid.UUID]error)
	var wg sync.WaitGroup
	for uuid, checkR := range nwd.checks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lastresult[uuid] = checkR.checker.Check(nwd.ctx)
		}()
	}
	wg.Wait()
	return lastresult
}
