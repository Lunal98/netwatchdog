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
package remediationhelper

import (
	"fmt"
	"time"
)

func GetDebugHelper() Helper {
	var helper debughelper
	return helper
}

type debughelper struct {
}

// ResetInterface implements Helper.
func (d debughelper) ResetInterface() {
	fmt.Println("restarting interface")
	for range 4 {
		time.Sleep(2 * time.Second)
		fmt.Print(".")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("restarted interface")
}

// Restart implements Helper.
func (d debughelper) Restart() {
	fmt.Println("restarting computer")
	for range 6 {
		time.Sleep(2 * time.Second)
		fmt.Print(".")
		time.Sleep(2 * time.Second)
	}
	fmt.Println("restarted computer")
}
