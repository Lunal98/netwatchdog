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

func init() {
	time.Sleep(time.Second)
	var check check.InterfaceCheck
	ctx := context.TODO()
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal().Err(err).Msg("Error starting internal scheduler, exiting")
	}
	job, err := scheduler.NewJob(
		gocron.DurationJob(30*time.Second),
		gocron.NewTask(check.Check(ctx)),
	)
	if err != nil {

		log.Error().Err(err).Msg(fmt.Sprintf("Error starting job: %d", check.CheckName))
	}

}
