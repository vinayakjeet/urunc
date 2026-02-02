// Copyright (c) 2023-2026, Nubificus LTD
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package urunce2etesting

import "fmt"

func nerdctlTestCases(kvmGroup ...int64) []containerTestArgs {
	tests, err := loadTestCases("nerdctl")
	if err != nil {
		fmt.Printf("Failed to load nerdctl test cases: %v\n", err)
		return nil
	}

	var group int64
	if len(kvmGroup) > 0 {
		group = kvmGroup[0]
	}

	// Update the test case that needs the group ID
	for i := range tests {
		if tests[i].Name == "Hvt-rumprun-UserGroup" {
			tests[i].Groups = []int64{group}
			break
		}
	}

	return tests
}

func ctrTestCases() []containerTestArgs {
	tests, err := loadTestCases("ctr")
	if err != nil {
		fmt.Printf("Failed to load ctr test cases: %v\n", err)
		return nil
	}
	return tests
}

func crictlTestCases(kvmGroup ...int64) []containerTestArgs {
	tests, err := loadTestCases("crictl")
	if err != nil {
		fmt.Printf("Failed to load crictl test cases: %v\n", err)
		return nil
	}

	var group int64
	if len(kvmGroup) > 0 {
		group = kvmGroup[0]
	}

	for i := range tests {
		if tests[i].Name == "Qemu-unikraft-UserGroup" {
			tests[i].Groups = []int64{group}
			break
		}
	}

	return tests
}

func dockerTestCases(kvmGroup ...int64) []containerTestArgs {
	tests, err := loadTestCases("docker")
	if err != nil {
		fmt.Printf("Failed to load docker test cases: %v\n", err)
		return nil
	}

	var group int64
	if len(kvmGroup) > 0 {
		group = kvmGroup[0]
	}

	for i := range tests {
		if tests[i].Name == "Spt-mirage-UserGroup" {
			tests[i].Groups = []int64{group}
			break
		}
	}

	return tests
}
