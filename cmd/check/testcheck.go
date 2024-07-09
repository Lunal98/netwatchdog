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
)

type InterfaceCheck struct {
	CheckName string
}

func (I *InterfaceCheck) Check(ctx context.Context) error {
	fmt.Println("Test check has been run")
	return nil
}
func (I *InterfaceCheck) GetCheckName() string {
	return I.CheckName
}
