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
	return formatInt(confines...)
}

func Int8ToConfine(confines ...int8) []string {
	return formatInt8(confines...)
}

func Int16ToConfine(confines ...int16) []string {
	return formatInt16(confines...)
}

func Int32ToConfine(confines ...int32) []string {
	return formatInt32(confines...)
}

func Int64ToConfine(confines ...int64) []string {
	return formatInt64(confines...)
}

func UintToConfine(confines ...uint) []string {
	return formatUint(confines...)
}

func Uint8ToConfine(confines ...uint8) []string {
	return formatUint8(confines...)
}

func Uint16ToConfine(confines ...uint16) []string {
	return formatUint16(confines...)
}

func Uint32ToConfine(confines ...uint32) []string {
	return formatUint32(confines...)
}

func Uint64ToConfine(confines ...uint64) []string {
	return formatUint64(confines...)
}

func Float32ToConfine(confines ...float32) []string {
	return formatFloat32(confines...)
}

func Float64ToConfine(confines ...float64) []string {
	return formatFloat64(confines...)
}
