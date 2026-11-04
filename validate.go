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
	"fmt"
	"reflect"
	"strings"
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

	if !validator.hasError {
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
	hasError bool
	// 错误消息，一般情况下只可能有一条，使用选项 [All] 时会有多条
	msgs []string
	// 当前校验字段的信息
	fieldInfo *FieldInfo
}

func (v *Validator) Bool(b *bool, fieldName string, opts ...FieldOption) *checkBool {
	v.beforeCheckField(fieldName, opts)
	return &checkBool{v: v, b: b}
}

func (v *Validator) Byte(b *byte, fieldName string, opts ...FieldOption) *checkNumber[byte] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[byte]{v: v, n: b, numberFormatter: uint8ToStr}
}

func (v *Validator) Int(i *int, fieldName string, opts ...FieldOption) *checkNumber[int] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[int]{v: v, n: i, numberFormatter: intToStr}
}

func (v *Validator) Int8(i *int8, fieldName string, opts ...FieldOption) *checkNumber[int8] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[int8]{v: v, n: i, numberFormatter: int8ToStr}
}

func (v *Validator) Int16(i *int16, fieldName string, opts ...FieldOption) *checkNumber[int16] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[int16]{v: v, n: i, numberFormatter: int16ToStr}
}

func (v *Validator) Int32(i *int32, fieldName string, opts ...FieldOption) *checkNumber[int32] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[int32]{v: v, n: i, numberFormatter: int32ToStr}
}

func (v *Validator) Int64(i *int64, fieldName string, opts ...FieldOption) *checkNumber[int64] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[int64]{v: v, n: i, numberFormatter: int64ToStr}
}

func (v *Validator) Uint(u *uint, fieldName string, opts ...FieldOption) *checkNumber[uint] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[uint]{v: v, n: u, numberFormatter: uintToStr}
}

func (v *Validator) Uint8(u *uint8, fieldName string, opts ...FieldOption) *checkNumber[uint8] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[uint8]{v: v, n: u, numberFormatter: uint8ToStr}
}

func (v *Validator) Uint16(u *uint16, fieldName string, opts ...FieldOption) *checkNumber[uint16] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[uint16]{v: v, n: u, numberFormatter: uint16ToStr}
}

func (v *Validator) Uint32(u *uint32, fieldName string, opts ...FieldOption) *checkNumber[uint32] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[uint32]{v: v, n: u, numberFormatter: uint32ToStr}
}

func (v *Validator) Uint64(u *uint64, fieldName string, opts ...FieldOption) *checkNumber[uint64] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[uint64]{v: v, n: u, numberFormatter: uint64ToStr}
}

func (v *Validator) Float32(f *float32, fieldName string, opts ...FieldOption) *checkNumber[float32] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[float32]{v: v, n: f, numberFormatter: float32ToStr}
}

func (v *Validator) Float64(f *float64, fieldName string, opts ...FieldOption) *checkNumber[float64] {
	v.beforeCheckField(fieldName, opts)
	return &checkNumber[float64]{v: v, n: f, numberFormatter: float64ToStr}
}

func (v *Validator) String(s *string, fieldName string, opts ...FieldOption) *checkString {
	v.beforeCheckField(fieldName, opts)
	return &checkString{v: v, s: s}
}

func (v *Validator) Struct[V Validatable](s *V, fieldName string, opts ...FieldOption) *checkStruct[V] {
	v.beforeCheckField(fieldName, opts)
	return &checkStruct[V]{v: v, s: s}
}

func (v *Validator) Slice[T any](s []T, fieldName string, opts ...FieldOption) *checkSlice[T] {
	v.beforeCheckField(fieldName, opts)
	return &checkSlice[T]{v: v, s: s}
}

func (v *Validator) Map[K comparable, V any](m map[K]V, fieldName string, opts ...FieldOption) *checkMap[K, V] {
	v.beforeCheckField(fieldName, opts)
	return &checkMap[K, V]{v: v, m: m}
}

func (v *Validator) Any[T any](a *T, fieldName string, opts ...FieldOption) *checkAny[T] {
	v.beforeCheckField(fieldName, opts)
	return &checkAny[T]{v: v, a: a}
}

func (v *Validator) Embed[V Validatable](s *V) {
	v.beforeCheckEmbed()
	c := &checkEmbed[V]{v: v, s: s}
	c.dive()
}

func (v *Validator) shouldStop() bool {
	return v.hasError && !v.all
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

func (v *Validator) beforeCheckEmbed() {
	v.fieldInfo = &FieldInfo{
		fieldNamePrefix: v.fieldInfo.fieldNamePrefix,
	}
}

func (v *Validator) checkRule[C comparable, A any, T *A | []A | map[C]A](t T, opts []RuleOption, mbf msgBuilder, confineFunc func() []string, checkNil func() bool, checkNonNil func() bool) {
	if v.shouldStop() {
		return
	}
	for _, opt := range opts {
		opt(v.fieldInfo)
	}
	var fail bool
	if t == nil {
		if !v.fieldInfo.omitempty {
			fail = !checkNil()
		}
	} else if checkNonNil != nil {
		fail = !checkNonNil()
	}
	if fail {
		var confines []string
		if confineFunc != nil {
			confines = confineFunc()
		}
		v.fail(mbf, confines)
	}
}

func (v *Validator) fail(mbf msgBuilder, confines []string) {
	v.hasError = true
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

// FieldInfo 字段信息
type FieldInfo struct {
	// 字段名前缀
	fieldNamePrefix string
	// 是否为切片元素、map的key或value
	hasElemName bool
	// nil时忽略校验
	omitempty bool
	// 字段名
	fieldName string
	// 限制值的字符串形式，校验每个规则方法时会修改此字段
	confines []string
	// 错误消息函数，校验每个规则方法时会修改此字段
	mbf msgBuilder
	// 校验时的错误码，校验每个规则方法时会修改此字段
	code string
	// 校验失败的错误消息，校验每个规则方法时会修改此字段
	msg string
}

// FieldName 返回具有层级关系的字段名称。例如：title、author.name、categoryId[2].sort。
func (f *FieldInfo) FieldName() string {
	return f.fieldName
}

// Confine 返回规则方法的指定索引的限制值的字符串形式。例如：对于Max(10)，vc.Confine(0)返回10；对于Range(5, 15)，vc.Confine(0)返回5，vc.Confine(1)返回15。
func (f *FieldInfo) Confine(i int) string {
	confine := ""
	if i < len(f.confines) {
		confine = f.confines[i]
	}
	return confine
}

// Confines 返回规则方法的所有限制值的字符串形式，用","拼接，若数量超过两个，则最后一个用"or"拼接。例如：对于Enum("zh-cn", "en-US")，返回"zh-cn, en-US"；对于Enum("zh-cn", "en-US", "ja-JP")，返回"zh-cn, en-US or ja-JP"。
func (f *FieldInfo) Confines() string {
	var builder strings.Builder
	l := len(f.confines)
	if l > 0 {
		builder.WriteString(f.confines[0])
	}
	for i := 1; i < l-1; i++ {
		builder.WriteString(", ")
		builder.WriteString(f.confines[i])
	}
	if l > 1 {
		builder.WriteString(" or ")
		builder.WriteString(f.confines[l-1])
	}
	return builder.String()
}

// Msg 自定义错误消息
func (f *FieldInfo) Msg(msg string, args ...any) {
	f.msg = fmt.Sprintf(msg, args...)
}
