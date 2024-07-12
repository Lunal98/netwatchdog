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
package core

import (
	"context"
	"fmt"
	"time"

	"github.com/Lunal98/netwatchdog/cmd/check"

	"github.com/go-co-op/gocron/v2"
	"github.com/rs/zerolog/log"
)

func Init() {
	cf := NewCheckfactory()
	cf.AddCheck(&check.InterfaceCheck{CheckName: "eth0"})
	asdasd := check.InterfaceCheck{CheckName: "asd123"}
	cks := cf.GetAll()
	for _, check := range cks {
		fmt.Println(check.GetCheckName())
	}
	ctx := context.TODO()
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal().Err(err).Msg("Error starting internal scheduler, exiting")
	}
	var jobs []gocron.Job
	for _, check := range cf.GetAll() {
		job, err := scheduler.NewJob(
			gocron.DurationJob(30*time.Second),
			gocron.NewTask(check.Check, ctx),
		)
		if err != nil {

			log.Error().Err(err).Msg(fmt.Sprintf("Error starting job: %s", check.GetCheckName()))
		}
		jobs = append(jobs, job)
		scheduler.Start()
		time.Sleep(10 * time.Second)
		job2, err := scheduler.NewJob(
			gocron.DurationJob(30*time.Second),
			gocron.NewTask(asdasd.Check, ctx),
		)
		if err != nil {

			log.Error().Err(err).Msg(fmt.Sprintf("Error starting job: %s", check.GetCheckName()))
		}
		select {
		case <-time.After(2 * time.Minute):
		}

		err = scheduler.Shutdown()

	}

}
