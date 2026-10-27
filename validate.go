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

import (
	"context"
	"reflect"
	"strconv"
)

const (
	SUCCESS = "SUCCESS"
	FAIL    = "FAIL"
)

type Validatable interface {
	Checklist(v *Validator)
}

func Check[V Validatable](ctx context.Context, v V, opts ...StructOption) (code string, msg string, msgs []string) {
	if !reflect.Indirect(reflect.ValueOf(v)).IsValid() {
		return
	}
	if ctx == nil {
		ctx = context.Background()
	}

	validator := &Validator{ctx: ctx, fieldInfo: &FieldInfo{}}
	for _, o := range opts {
		o(validator)
	}
	v.Checklist(validator)

	if !validator.hasWrong {
		code = SUCCESS
	} else {
		code = validator.code
		if code == "" {
			code = FAIL
		}
	}
	msgs = validator.msgs
	if len(msgs) > 0 {
		msg = msgs[0]
	}
	return
}

type Validator struct {
	ctx context.Context
	// 错误码，一般情况只有“SUCCESS”和"FALSE"，使用自定义错误码选项 [Code] 时会有其他值。
	code string
	// 是否校验所有规则，使用选项 [All] 时为true
	all bool
	// 是否已有校验错误，在不校验全部规则时，通过此此值判断是否停止校验
	hasWrong bool
	// 错误消息，一般情况下只可能有一条，使用选项 [All] 时会有多条
	msgs []string
	// 当前校验字段的信息
	fieldInfo *FieldInfo
}

func (v *Validator) Bool(b *bool, fieldName string, opts ...FieldOption) *checkBool {
	v.beforeCheckField(fieldName, opts)
	return &checkBool{v: v.beforeCheckField(fieldName, opts), b: b}
}

func (v *Validator) Byte(b *byte, fieldName string, opts ...FieldOption) *checkNumber[byte] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[byte]{v: v.beforeCheckField(fieldName, opts), n: b, nts: uint8ToStr}
}

func (v *Validator) Int(i *int, fieldName string, opts ...FieldOption) *checkNumber[int] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[int]{v: v.beforeCheckField(fieldName, opts), n: i, nts: intToStr}
}

func (v *Validator) Int8(i *int8, fieldName string, opts ...FieldOption) *checkNumber[int8] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[int8]{v: v.beforeCheckField(fieldName, opts), n: i, nts: int8ToStr}
}

func (v *Validator) Int16(i *int16, fieldName string, opts ...FieldOption) *checkNumber[int16] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[int16]{v: v.beforeCheckField(fieldName, opts), n: i, nts: int16ToStr}
}

func (v *Validator) Int32(i *int32, fieldName string, opts ...FieldOption) *checkNumber[int32] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[int32]{v: v.beforeCheckField(fieldName, opts), n: i, nts: int32ToStr}
}

func (v *Validator) Int64(i *int64, fieldName string, opts ...FieldOption) *checkNumber[int64] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[int64]{v: v.beforeCheckField(fieldName, opts), n: i, nts: int64ToStr}
}

func (v *Validator) Uint(u *uint, fieldName string, opts ...FieldOption) *checkNumber[uint] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[uint]{v: v.beforeCheckField(fieldName, opts), n: u, nts: uintToStr}
}

func (v *Validator) Uint8(u *uint8, fieldName string, opts ...FieldOption) *checkNumber[uint8] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[uint8]{v: v.beforeCheckField(fieldName, opts), n: u, nts: uint8ToStr}
}

func (v *Validator) Uint16(u *uint16, fieldName string, opts ...FieldOption) *checkNumber[uint16] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[uint16]{v: v.beforeCheckField(fieldName, opts), n: u, nts: uint16ToStr}
}

func (v *Validator) Uint32(u *uint32, fieldName string, opts ...FieldOption) *checkNumber[uint32] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[uint32]{v: v.beforeCheckField(fieldName, opts), n: u, nts: uint32ToStr}
}

func (v *Validator) Uint64(u *uint64, fieldName string, opts ...FieldOption) *checkNumber[uint64] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[uint64]{v: v.beforeCheckField(fieldName, opts), n: u, nts: uint64ToStr}
}

func (v *Validator) Float32(f *float32, fieldName string, opts ...FieldOption) *checkNumber[float32] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[float32]{v: v.beforeCheckField(fieldName, opts), n: f, nts: float32ToStr}
}

func (v *Validator) Float64(f *float64, fieldName string, opts ...FieldOption) *checkNumber[float64] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[float64]{v: v.beforeCheckField(fieldName, opts), n: f, nts: float64ToStr}
}

func (v *Validator) String(s *string, fieldName string, opts ...FieldOption) *checkString {
	v.beforeCheckField(fieldName, opts)
	return &checkString{v: v.beforeCheckField(fieldName, opts), s: s}
}

func (v *Validator) Struct[V Validatable](s *V, fieldName string, opts ...FieldOption) *checkStruct[V] {
	v.beforeCheckField(fieldName, opts)
	return &checkStruct[V]{v: v.beforeCheckField(fieldName, opts), s: s}
}

func (v *Validator) Slice[T any](s []T, fieldName string, opts ...FieldOption) *checkSlice[T] {
	v.beforeCheckField(fieldName, opts)
	return &checkSlice[T]{v: v.beforeCheckField(fieldName, opts), s: s}
}

func (v *Validator) Map[K comparable, V any](m map[K]V, fieldName string, opts ...FieldOption) *checkMap[K, V] {
	v.beforeCheckField(fieldName, opts)
	return &checkMap[K, V]{v: v.beforeCheckField(fieldName, opts), m: m}
}

func (v *Validator) Any[T any](a *T, fieldName string, opts ...FieldOption) *checkAny[T] {
	v.beforeCheckField(fieldName, opts)
	return &checkAny[T]{v: v.beforeCheckField(fieldName, opts), a: a}
}

func (v *Validator) Embed[V Validatable](s *V) {
	c := &checkEmbed[V]{v: v.beforeCheckEmbed(), s: s}
	c.dive()
}

func (v *Validator) interrupt() bool {
	return v.hasWrong && !v.all
}

func (v *Validator) beforeCheckField(fieldName string, opts []FieldOption) {
	if v.fieldInfo.hasElemName {
		v.fieldInfo = &FieldInfo{
			fieldName:   v.fieldInfo.fieldName,
			hasElemName: true,
		}
	} else {
		v.fieldInfo = &FieldInfo{
			fieldNamePrefix: v.fieldInfo.fieldNamePrefix,
			fieldName:       v.fieldInfo.fieldNamePrefix + fieldName,
		}
	}
	for _, o := range opts {
		o(v.fieldInfo)
	}
}

func (v *Validator) beforeDiveStruct() {
	v.fieldInfo = &FieldInfo{
		fieldNamePrefix: v.fieldInfo.fieldName + ".",
	}
}

func (v *Validator) beforeDiveSliceMap(elemName string) {
	v.fieldInfo = &FieldInfo{
		fieldName:   v.fieldInfo.fieldName + elemName,
		hasElemName: true,
	}
}

func (v *Validator) beforeCheckEmbed() *Validator {
	v.fieldInfo = &FieldInfo{
		fieldNamePrefix: v.fieldInfo.fieldNamePrefix,
	}
	return v
}

func (v *Validator) predicate[C comparable, A any, T *A | []A | map[C]A](t T, opts []RuleOption, mbf msgBuildFunc, confineFunc func() []string, predicateNil func() bool, predicateNoNil func() bool) {
	if v.interrupt() {
		return
	}
	for _, opt := range opts {
		opt(v.fieldInfo)
	}
	var fail bool
	if t == nil {
		if !v.fieldInfo.omitempty {
			fail = !predicateNil()
		}
	} else if predicateNoNil != nil {
		fail = !predicateNoNil()
	}
	if fail {
		var confines []string
		if confineFunc != nil {
			confines = confineFunc()
		}
		v.fail(mbf, confines)
	}
}

func (v *Validator) fail(mbf msgBuildFunc, confines []string) {
	v.hasWrong = true
	if !v.all {
		v.code = v.fieldInfo.code
	}
	v.fieldInfo.confines = confines
	var msg string
	if v.fieldInfo.mbf == nil {
		if mbf != nil {
			mbf(v.fieldInfo)
		}
	} else {
		v.fieldInfo.mbf(v.fieldInfo)
	}
	msg = v.fieldInfo.msg
	v.msgs = append(v.msgs, msg)
}

func intToStr(ii ...int) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(int64(i), 10))
	}
	return result
}

func int8ToStr(ii ...int8) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(int64(i), 10))
	}
	return result
}

func int16ToStr(ii ...int16) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(int64(i), 10))
	}
	return result
}

func int32ToStr(ii ...int32) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(int64(i), 10))
	}
	return result
}

func int64ToStr(ii ...int64) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(i, 10))
	}
	return result
}

func uintToStr(uu ...uint) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(uint64(u), 10))
	}
	return result
}

func uint8ToStr(uu ...uint8) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(uint64(u), 10))
	}
	return result
}

func uint16ToStr(uu ...uint16) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(uint64(u), 10))
	}
	return result
}

func uint32ToStr(uu ...uint32) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(uint64(u), 10))
	}
	return result
}

func uint64ToStr(uu ...uint64) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(u, 10))
	}
	return result
}

func float32ToStr(ff ...float32) []string {
	result := make([]string, 0, len(ff))
	for _, f := range ff {
		result = append(result, strconv.FormatFloat(float64(f), 'f', -1, 32))
	}
	return result
}

func float64ToStr(ff ...float64) []string {
	result := make([]string, 0, len(ff))
	for _, f := range ff {
		result = append(result, strconv.FormatFloat(f, 'f', -1, 32))
	}
	return result
}
