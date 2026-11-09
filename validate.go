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
	"strings"
)

const (
	SUCCESS = "SUCCESS"
	FAIL    = "FAIL"
)

type Validatable interface {
	Checklist(v *Validator)
}

type validatable[T any] interface {
	*T
	Validatable
}

func Check[T any, V validatable[T]](ctx context.Context, target V, opts ...ValidatorOption) (code string, message string, messages []string) {
	validator := newValidator(ctx)
	for _, o := range opts {
		o(validator)
	}
	validator.checkRoot(target)

	if !validator.hasError {
		code = SUCCESS
	} else {
		code = validator.code
		if code == "" {
			code = FAIL
		}
	}

	messages = validator.messages
	if len(messages) > 0 {
		message = messages[0]
	}
	return
}

type Validator struct {
	ctx context.Context
	// 错误码，一般情况只有“SUCCESS”和"FALSE"，使用自定义错误码选项 [Code] 时会有其他值。
	code string
	// 是否校验所有规则，使用选项 [AllRules] 时为true
	allRules bool
	// 是否已有校验错误，在不校验全部规则时，通过此此值判断是否停止校验
	hasError bool
	// 错误消息，一般情况下只可能有一条，使用选项 [AllRules] 时会有多条
	messages []string
	// 当前校验字段的信息
	fieldInfo *FieldInfo
	// 用于递归保护
	activeDiveKeys map[any]struct{}
}

func (v *Validator) Ctx() context.Context {
	return v.ctx
}

func (v *Validator) Bool(b *bool, fieldName string) *checkBool {
	v.beforeCheckField(fieldName)
	return &checkBool{v: v, b: b}
}

func (v *Validator) Byte(b *byte, fieldName string) *checkNumber[byte] {
	v.beforeCheckField(fieldName)
	return &checkNumber[byte]{v: v, n: b, formatter: formatUint8}
}

func (v *Validator) Int(i *int, fieldName string) *checkNumber[int] {
	v.beforeCheckField(fieldName)
	return &checkNumber[int]{v: v, n: i, formatter: formatInt}
}

func (v *Validator) Int8(i *int8, fieldName string) *checkNumber[int8] {
	v.beforeCheckField(fieldName)
	return &checkNumber[int8]{v: v, n: i, formatter: formatInt8}
}

func (v *Validator) Int16(i *int16, fieldName string) *checkNumber[int16] {
	v.beforeCheckField(fieldName)
	return &checkNumber[int16]{v: v, n: i, formatter: formatInt16}
}

func (v *Validator) Int32(i *int32, fieldName string) *checkNumber[int32] {
	v.beforeCheckField(fieldName)
	return &checkNumber[int32]{v: v, n: i, formatter: formatInt32}
}

func (v *Validator) Int64(i *int64, fieldName string) *checkNumber[int64] {
	v.beforeCheckField(fieldName)
	return &checkNumber[int64]{v: v, n: i, formatter: formatInt64}
}

func (v *Validator) Uint(u *uint, fieldName string) *checkNumber[uint] {
	v.beforeCheckField(fieldName)
	return &checkNumber[uint]{v: v, n: u, formatter: formatUint}
}

func (v *Validator) Uint8(u *uint8, fieldName string) *checkNumber[uint8] {
	v.beforeCheckField(fieldName)
	return &checkNumber[uint8]{v: v, n: u, formatter: formatUint8}
}

func (v *Validator) Uint16(u *uint16, fieldName string) *checkNumber[uint16] {
	v.beforeCheckField(fieldName)
	return &checkNumber[uint16]{v: v, n: u, formatter: formatUint16}
}

func (v *Validator) Uint32(u *uint32, fieldName string) *checkNumber[uint32] {
	v.beforeCheckField(fieldName)
	return &checkNumber[uint32]{v: v, n: u, formatter: formatUint32}
}

func (v *Validator) Uint64(u *uint64, fieldName string) *checkNumber[uint64] {
	v.beforeCheckField(fieldName)
	return &checkNumber[uint64]{v: v, n: u, formatter: formatUint64}
}

func (v *Validator) Float32(f *float32, fieldName string) *checkNumber[float32] {
	v.beforeCheckField(fieldName)
	return &checkNumber[float32]{v: v, n: f, formatter: formatFloat32}
}

func (v *Validator) Float64(f *float64, fieldName string) *checkNumber[float64] {
	v.beforeCheckField(fieldName)
	return &checkNumber[float64]{v: v, n: f, formatter: formatFloat64}
}

func (v *Validator) String(s *string, fieldName string) *checkString {
	v.beforeCheckField(fieldName)
	return &checkString{v: v, s: s}
}

func (v *Validator) Struct[T any, V validatable[T]](s V, fieldName string) *checkStruct[T, V] {
	v.beforeCheckField(fieldName)
	return &checkStruct[T, V]{v: v, s: s}
}

func (v *Validator) Slice[T any](s []T, fieldName string) *checkSlice[T] {
	v.beforeCheckField(fieldName)
	return &checkSlice[T]{v: v, s: s}
}

func (v *Validator) Map[K comparable, V any](m map[K]V, fieldName string) *checkMap[K, V] {
	v.beforeCheckField(fieldName)
	return &checkMap[K, V]{v: v, m: m}
}

func (v *Validator) Any[T any](a *T, fieldName string) *checkAny[T] {
	v.beforeCheckField(fieldName)
	return &checkAny[T]{v: v, a: a}
}

func (v *Validator) Embed[T any, V validatable[T]](e V) {
	c := &checkEmbed[T, V]{v: v, e: e}
	c.dive()
}

func (v *Validator) checkRoot[T any, V validatable[T]](r V) {
	var key any
	if r == nil {
		key = (*T)(nil)
		r = V(new(T))
		v.fieldInfo.virtualDive = true
	} else {
		key = r
	}
	v.activeDiveKeys[key] = struct{}{}
	r.Checklist(v)
}

func (v *Validator) shouldStop() bool {
	return v.hasError && !v.allRules
}

func (v *Validator) beforeCheckField(fieldName string) {
	if v.fieldInfo.isElem {
		v.fieldInfo = &FieldInfo{
			fieldName:   v.fieldInfo.fieldName,
			isElem:      true,
			virtualDive: v.fieldInfo.virtualDive,
		}
	} else {
		v.fieldInfo = &FieldInfo{
			fieldNamePrefix: v.fieldInfo.fieldNamePrefix,
			fieldName:       v.fieldInfo.fieldNamePrefix + fieldName,
			virtualDive:     v.fieldInfo.virtualDive,
		}
	}
}

func (v *Validator) beforeDiveStruct[T any, V validatable[T]](value V) (allowed bool, after func()) {
	if value == nil {
		return false, nil
	}

	if _, ok := v.activeDiveKeys[value]; ok {
		return false, nil
	}

	prev := v.fieldInfo
	v.fieldInfo = &FieldInfo{
		fieldNamePrefix: v.fieldInfo.fieldName + ".",
		virtualDive:     v.fieldInfo.virtualDive,
	}
	v.activeDiveKeys[value] = struct{}{}
	return true, func() {
		v.fieldInfo = prev
		delete(v.activeDiveKeys, value)
	}
}

func (v *Validator) beforeDiveEmbed[T any, V validatable[T]](value V) (target V, after func()) {
	var key any
	if value == nil {
		key = (*T)(nil)
	} else {
		key = value
	}

	if _, ok := v.activeDiveKeys[key]; ok {
		return nil, nil
	}

	prev := v.fieldInfo
	v.fieldInfo = &FieldInfo{
		fieldNamePrefix: v.fieldInfo.fieldNamePrefix,
		virtualDive:     v.fieldInfo.virtualDive,
	}
	if value == nil {
		value = V(new(T))
		v.fieldInfo.virtualDive = true
	}
	v.activeDiveKeys[key] = struct{}{}
	return value, func() {
		v.fieldInfo = prev
		delete(v.activeDiveKeys, key)
	}
}

func (v *Validator) beforeDiveSliceMap(elemName string) (after func()) {
	prev := v.fieldInfo
	v.fieldInfo = &FieldInfo{
		fieldName: v.fieldInfo.fieldName + elemName,
		isElem:    true,
	}
	return func() {
		v.fieldInfo = prev
	}
}

func (v *Validator) checkRule(opts []RuleOption, messageBuilder messageBuilder, formatArgs func() []string, rule func() bool) {
	if v.shouldStop() {
		return
	}
	v.fieldInfo.code = ""
	v.fieldInfo.messageBuilder = nil
	for _, opt := range opts {
		opt(v.fieldInfo)
	}
	if !rule() {
		var args []string
		if formatArgs != nil {
			args = formatArgs()
		}
		v.fail(messageBuilder, args)
	}
}

func (v *Validator) fail(messageBuilder messageBuilder, args []string) {
	v.hasError = true
	if !v.allRules {
		v.code = v.fieldInfo.code
	}
	v.fieldInfo.args = args
	var message string
	if v.fieldInfo.messageBuilder == nil {
		if messageBuilder != nil {
			message = messageBuilder(v.fieldInfo)
		}
	} else {
		message = v.fieldInfo.messageBuilder(v.fieldInfo)
	}
	v.messages = append(v.messages, message)
}

func newValidator(ctx context.Context) *Validator {
	if ctx == nil {
		ctx = context.Background()
	}
	return &Validator{ctx: ctx, fieldInfo: &FieldInfo{}, activeDiveKeys: make(map[any]struct{})}
}

// FieldInfo 字段信息
type FieldInfo struct {
	// 字段名前缀
	fieldNamePrefix string
	// 是否为切片元素、map的key或value
	isElem bool
	// 虚拟下沉（结构体字段、内嵌结构体为临时创建的零值）
	virtualDive bool
	// 字段名
	fieldName string
	// 限制值的字符串形式，校验每个规则方法时会修改此字段
	args []string
	// 错误消息函数，校验每个规则方法时会修改此字段
	messageBuilder messageBuilder
	// 校验时的错误码，校验每个规则方法时会修改此字段
	code string
}

func (f *FieldInfo) FieldName() string {
	return f.fieldName
}

func (f *FieldInfo) Arg(i int) string {
	arg := ""
	if i < len(f.args) {
		arg = f.args[i]
	}
	return arg
}

func (f *FieldInfo) Args() string {
	var builder strings.Builder
	if len(f.args) > 0 {
		builder.WriteString(f.args[0])
		for i := 1; i < len(f.args)-1; i++ {
			builder.WriteString(", ")
			builder.WriteString(f.args[i])
		}
		if len(f.args) > 1 {
			if len(f.args) > 2 {
				builder.WriteString(",")
			}
			builder.WriteString(" or ")
			builder.WriteString(f.args[len(f.args)-1])
		}
	}
	return builder.String()
}
