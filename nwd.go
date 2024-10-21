/*
Copyright © 2024 Alex Bedo <alex98hun@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"context"
	"sync"
	"time"

	"github.com/Lunal98/netwatchdog/internal/checker"
	"github.com/Lunal98/netwatchdog/internal/remediationhelper"
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
	for {
		checkid := nwd.scheduler.Start(nwd.ctx)
		if checkid == uuid.Nil {
			log.Info().Msg("context is done, exiting normally")
			return
		}
		nwd.remediatefault()

	}

}
func (nwd *NwdCore) remediatefault() {
	checks := nwd.getCheckers()
	for {
		primaryfault := findTopPriority(nwd)
		if primaryfault == uuid.Nil {
			return
		}
		checks[primaryfault].Remediate(remediationhelper.GetDebugHelper())

	}
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
	lastresult := runAllChecks(nwd)
	if len(lastresult) == 0 {
		return uuid.Nil
	}
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
func runAllChecks(nwd *NwdCore) map[uuid.UUID]error {

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
