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

func NewValidator() *Validator {
	return newValidator(nil)
}

func GetResult(v *Validator) (code string, message string, messages []string) {
	if !v.hasError {
		code = SUCCESS
	} else if v.code == "" {
		code = FAIL
	} else {
		code = v.code
	}
	messages = v.messages
	if len(messages) > 0 {
		message = messages[0]
	}
	return
}

func SetHasError(v *Validator) {
	v.hasError = true
}

func SetAll(v *Validator) {
	v.allRule = true
}

func SetFieldName(v *Validator, fieldName string) {
	v.fieldInfo.fieldName = fieldName
}

func SetOmitNil(v *Validator) {
	OmitNil()(v.fieldInfo)
}

func CheckRule(v *Validator, isNil bool, opts []RuleOption, formatArgs func() []string, checkNil func() bool, checkNonNil func() bool) {
	v.beforeCheckRule(opts)
	v.checkRule(isNil, messageDefault, formatArgs, checkNil, checkNonNil)
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

type DemoBase struct {
	Id string
}

func (d DemoBase) Checklist(v *Validator) {
	v.String(&d.Id, "id").Required().NotBlank()
}

type DemoUser struct {
	DemoBase
	Name   string
	Attach *DemoAttach
}

func (d DemoUser) Checklist(v *Validator) {
	v.Embed(&d.DemoBase)
	v.String(&d.Name, "name").NotBlank()
	v.Struct(d.Attach, "attach").Dive()
}

type DemoAttach struct {
	DemoBase
	Images   []DemoImage
	Score    map[string]int
	Profiles map[string]DemoProfile
}

func (d DemoAttach) Checklist(v *Validator) {
	v.Embed(&d.DemoBase)
	v.Slice(d.Images, "images").Dive(func(i DemoImage) {
		v.Struct(&i, "not work").Dive()
	})
	v.Map(d.Score, "score").Dive(nil, func(value int) {
		v.Int(&value, "not work").Max(100)
	})
	v.Map(d.Profiles, "profiles").Dive(func(k string) {
		v.String(&k, "not work").Max(4)
	}, func(value DemoProfile) {
		v.Struct(&value, "not work").Dive()
	})
}

type DemoImage struct {
	Url string
	*DemoBase
}

func (d DemoImage) Checklist(v *Validator) {
	v.String(&d.Url, "url").NotBlank()
	v.Embed(d.DemoBase)
}

type DemoProfile struct {
	Name string
}

func (s DemoProfile) Checklist(v *Validator) {
	v.String(&s.Name, "name").NotBlank()
}

type DemoNode struct {
	Name string
	Next *DemoNode
}

func (d *DemoNode) Checklist(v *Validator) {
	v.String(&d.Name, "name").NotBlank()
	v.Struct(d.Next, "next").Dive()
}

type DemoNode2 struct {
	Name string
	*DemoNode2
}

func (d *DemoNode2) Checklist(v *Validator) {
	v.String(&d.Name, "name").NotBlank()
	v.Embed(d.DemoNode2)
}
