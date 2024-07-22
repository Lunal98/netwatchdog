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
package check

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Lunal98/netwatchdog/cmd/remediationhelper"
)

type InterfaceCheck struct {
	CheckName string
	counter   int
	once      sync.Once
}

func (I *InterfaceCheck) Check(ctx context.Context) error {
	I.once.Do(func() {
		I.counter = 5
	})
	if I.counter == 0 {
		return fmt.Errorf("run out of test checks, please purchase a new license to continue")
	} else {
		I.counter--
	}
	fmt.Printf("(%s):Test check has been run\n", I.GetCheckName())
	return nil
}
func (I *InterfaceCheck) GetCheckName() string {
	return I.CheckName
}

func (I *InterfaceCheck) Remediate(remhelp remediationhelper.Helper) {
	remhelp.ResetInterface()
	time.Sleep(30 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := I.Check(ctx)
	if err == nil {
		return
	}
	remhelp.Restart()
}
