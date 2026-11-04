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

func SetHasError(v *Validator) {
	v.hasError = true
}

func SetAll(v *Validator) {
	v.all = true
}

func SetFieldName(v *Validator, fieldName string) {
	v.fieldInfo.fieldName = fieldName
}

func SetOmitNil(v *Validator) {
	OmitNil()(v.fieldInfo)
}

func CheckRule[T any](v *Validator, t *T, opts []RuleOption, formatArgs func() []string, checkNil func() bool, checkNonNil func() bool) {
	v.beforeCheckRule(opts)
	v.checkRule[any, T](t, messageDefault, formatArgs, checkNil, checkNonNil)
}

type Number = number

func FormatInt(args ...int) []string {
	return formatInt(args...)
}

func FormatInt8(args ...int8) []string {
	return formatInt8(args...)
}

func FormatInt16(args ...int16) []string {
	return formatInt16(args...)
}

func FormatInt32(args ...int32) []string {
	return formatInt32(args...)
}

func FormatInt64(args ...int64) []string {
	return formatInt64(args...)
}

func FormatUint(args ...uint) []string {
	return formatUint(args...)
}

func FormatUint8(args ...uint8) []string {
	return formatUint8(args...)
}

func FormatUint16(args ...uint16) []string {
	return formatUint16(args...)
}

func FormatUint32(args ...uint32) []string {
	return formatUint32(args...)
}

func FormatUint64(args ...uint64) []string {
	return formatUint64(args...)
}

func FormatFloat32(args ...float32) []string {
	return formatFloat32(args...)
}

func FormatFloat64(args ...float64) []string {
	return formatFloat64(args...)
}
