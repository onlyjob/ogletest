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
	"github.com/jacobsa/oglematchers"
	"path"
	"runtime"
)

func getCallerForAlias() (fileName string, lineNumber int) {
	_, fileName, lineNumber, _ = runtime.Caller(2)
	fileName = path.Base(fileName)
	return
}

// ExpectEq(e, a) is equivalent to ExpectThat(a, oglematchers.Equals(e)).
func ExpectEq(expected, actual interface{}) ExpectationResult {
	return ExpectThat(actual, oglematchers.Equals(expected)).SetCaller(getCallerForAlias())
}

// ExpectNe(e, a) is equivalent to ExpectThat(a, oglematchers.Not(oglematchers.Equals(e))).
func ExpectNe(expected, actual interface{}) ExpectationResult {
	return ExpectThat(actual, oglematchers.Not(oglematchers.Equals(expected))).SetCaller(getCallerForAlias())
}

// ExpectLt(x, y) is equivalent to ExpectThat(x, oglematchers.LessThan(y)).
func ExpectLt(x, y interface{}) ExpectationResult {
	return ExpectThat(x, oglematchers.LessThan(y)).SetCaller(getCallerForAlias())
}

// ExpectLe(x, y) is equivalent to ExpectThat(x, oglematchers.LessOrEqual(y)).
func ExpectLe(x, y interface{}) ExpectationResult {
	return ExpectThat(x, oglematchers.LessOrEqual(y)).SetCaller(getCallerForAlias())
}

// ExpectGt(x, y) is equivalent to ExpectThat(x, oglematchers.GreaterThan(y)).
func ExpectGt(x, y interface{}) ExpectationResult {
	return ExpectThat(x, oglematchers.GreaterThan(y)).SetCaller(getCallerForAlias())
}

// ExpectGe(x, y) is equivalent to ExpectThat(x, oglematchers.GreaterOrEqual(y)).
func ExpectGe(x, y interface{}) ExpectationResult {
	return ExpectThat(x, oglematchers.GreaterOrEqual(y)).SetCaller(getCallerForAlias())
}

// ExpectTrue(b) is equivalent to ExpectThat(b, oglematchers.Equals(true)).
func ExpectTrue(b interface{}) ExpectationResult {
	return ExpectThat(b, oglematchers.Equals(true)).SetCaller(getCallerForAlias())
}

// ExpectFalse(b) is equivalent to ExpectThat(b, oglematchers.Equals(false)).
func ExpectFalse(b interface{}) ExpectationResult {
	return ExpectThat(b, oglematchers.Equals(false)).SetCaller(getCallerForAlias())
}
