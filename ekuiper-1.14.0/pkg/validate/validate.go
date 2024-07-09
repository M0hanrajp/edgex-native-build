// Copyright 2024 EMQ Technologies Co., Ltd.
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

package validate

import (
	"fmt"
	"strings"
)

var invalidRuleChars = []string{
	"/", "#", "%",
}

func ValidateID(id string) error {
	if id != strings.TrimSpace(id) {
		return fmt.Errorf("ruleID: %v should be trimed", id)
	}
	for _, invalidChar := range invalidRuleChars {
		if strings.Contains(id, invalidChar) {
			return fmt.Errorf("ruleID:%s contains invalidChar:%v", id, invalidChar)
		}
	}
	return nil
}
