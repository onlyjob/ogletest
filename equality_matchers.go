// Copyright 2011 Aaron Jacobs. All Rights Reserved.
// Author: aaronjjacobs@gmail.com (Aaron Jacobs)
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

package ogletest

import (
	"reflect"
)

// Equals returns a matcher that matches any value v such that v == x, with the
// exception that if x is a numeric type, Equals(x) will match equivalent
// numeric values of any type.
func Equals(x interface{}) Matcher {
	return &equalsMatcher{x}
}

type equalsMatcher struct {
	expected interface{}
}

func checkAgainstInt(e int64, v reflect.Value) (res MatchResult, err string) {
	res = MATCH_FALSE

	switch v.Kind() {
	default:
		res = MATCH_UNDEFINED
		err = "which is not numeric"

	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		if (e == v.Int()) {
			res = MATCH_TRUE
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if (uint64(e) == v.Uint()) {
			res = MATCH_TRUE
		}
	}

	return
}

func (m *equalsMatcher) Matches(candidate interface{}) (MatchResult, string) {
	expectedValue := reflect.ValueOf(m.expected)
	candidateValue := reflect.ValueOf(candidate)

	switch expectedValue.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return checkAgainstInt(expectedValue.Int(), candidateValue)
	}

	return MATCH_UNDEFINED, "TODO"
}

func (m *equalsMatcher) Description() string {
	return "TODO"
}
