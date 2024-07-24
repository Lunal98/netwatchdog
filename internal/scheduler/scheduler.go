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
package scheduler

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/go-co-op/gocron/v2"
	"github.com/rs/zerolog/log"
)

type Scheduler struct {
	scheduler gocron.Scheduler
	once      sync.Once
	started   bool
}

func (s *Scheduler) init() {
	s.once.Do(func() {
		var err error
		s.scheduler, err = gocron.NewScheduler()
		if err != nil {
			log.Fatal().Err(err).Msg("Error starting internal scheduler, exiting")
		}
		s.started = false

	})
}
func (s *Scheduler) Addjob(function any, duration time.Duration) (uuid.UUID, error) {
	s.init()
	job, err := s.scheduler.NewJob(
		gocron.DurationJob(duration),
		gocron.NewTask(function),
	)
	return job.ID(), err
}
func (s *Scheduler) SetRemediator(eventListenerFunc func(jobID uuid.UUID, jobName string, err error)) {
	//gocron.EventListener
	s.init()
	gocron.AfterJobRunsWithError(eventListenerFunc)
}
func (s *Scheduler) Start(ctx context.Context) {
	if s.scheduler == nil {
		log.Warn().Msg("Scheduler.Start() Was called without it being initialized, See if any checks have been added to it.")
		return
	}
	s.scheduler.Start()
	s.started = true

	<-ctx.Done()
	s.scheduler.StopJobs()
}
func (s *Scheduler) Stop() {

	if s.started {
		s.scheduler.StopJobs()
	}
	s.started = false
}