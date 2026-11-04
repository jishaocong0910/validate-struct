// Copyright 2024-present jishaocong0910
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

import "context"

func NewValidator() *Validator {
	return &Validator{ctx: context.Background(), fieldInfo: &FieldInfo{}}
}

func GetResult(v *Validator) (code string, msg string, msgs []string) {
	if !v.hasError {
		code = SUCCESS
	} else if v.code == "" {
		code = FAIL
	} else {
		code = v.code
	}
	msgs = v.msgs
	if len(msgs) > 0 {
		msg = msgs[0]
	}
	return
}

func SetHasWrong(v *Validator) {
	v.hasError = true
}

func SetAll(v *Validator) {
	v.all = true
}

func SetFieldName(v *Validator, fieldName string) {
	v.fieldInfo.fieldName = fieldName
}

func SetOmittable(v *Validator) {
	Omitempty()(v.fieldInfo)
}

func CheckRule[T any](v *Validator, t *T, opts []RuleOption, confineFunc func() []string, checkNil func() bool, checkNonNil func() bool) {
	v.checkRule[int, T](t, opts, msgBuilderDefault, confineFunc, checkNil, checkNonNil)
}

type Number = number

type CheckOption = RuleOption

func IntToConfine(confines ...int) []string {
	return intToStr(confines...)
}

func Int8ToConfine(confines ...int8) []string {
	return int8ToStr(confines...)
}

func Int16ToConfine(confines ...int16) []string {
	return int16ToStr(confines...)
}

func Int32ToConfine(confines ...int32) []string {
	return int32ToStr(confines...)
}

func Int64ToConfine(confines ...int64) []string {
	return int64ToStr(confines...)
}

func UintToConfine(confines ...uint) []string {
	return uintToStr(confines...)
}

func Uint8ToConfine(confines ...uint8) []string {
	return uint8ToStr(confines...)
}

func Uint16ToConfine(confines ...uint16) []string {
	return uint16ToStr(confines...)
}

func Uint32ToConfine(confines ...uint32) []string {
	return uint32ToStr(confines...)
}

func Uint64ToConfine(confines ...uint64) []string {
	return uint64ToStr(confines...)
}

func Float32ToConfine(confines ...float32) []string {
	return float32ToStr(confines...)
}

func Float64ToConfine(confines ...float64) []string {
	return float64ToStr(confines...)
}
