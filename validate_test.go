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

package validate_test

import (
	"fmt"
	"testing"

	validate "github.com/jishaocong0910/validate-struct"
	"github.com/stretchr/testify/require"
)

func TestCheckRule(t *testing.T) {
	r := require.New(t)
	{
		// 成功
		v := validate.NewValidator()
		validate.CheckRule(v, false, nil, nil, func() bool {
			return false
		}, func() bool {
			return true
		})
		code, _, _ := validate.GetResult(v)
		r.Equal(validate.SUCCESS, code)
	}
	{
		// 指定code和msg
		v := validate.NewValidator()
		validate.SetFieldName(v, "param")
		validate.CheckRule(v, false, []validate.RuleOption{validate.Code("MY_CODE"), validate.Msg(func(b *validate.FieldInfo) string {
			return fmt.Sprintf("%s %s", b.FieldName(), b.Args())
		})}, func() []string {
			return []string{"a", "b", "c"}
		}, func() bool {
			return true
		}, func() bool {
			return false
		})
		code, msg, _ := validate.GetResult(v)
		r.Equal("MY_CODE", code)
		r.Equal("param a, b, or c", msg)
	}
	{
		// 默认code和msg
		v := validate.NewValidator()
		validate.SetFieldName(v, "param")
		validate.CheckRule(v, false, nil, nil, func() bool {
			return true
		}, func() bool {
			return false
		})
		code, msg, _ := validate.GetResult(v)
		r.Equal(validate.FAIL, code)
		r.Equal("param is illegal", msg)
	}
	{
		// nil导致错误
		v := validate.NewValidator()
		validate.SetFieldName(v, "param")
		validate.CheckRule(v, true, nil, nil, func() bool {
			return false
		}, func() bool {
			return true
		})
		code, _, _ := validate.GetResult(v)
		r.Equal(validate.FAIL, code)
	}
	{
		// 忽略nil
		v := validate.NewValidator()
		validate.SetOmitNil(v)
		validate.CheckRule(v, true, nil, nil, func() bool {
			return false
		}, func() bool {
			return false
		})
		code, _, _ := validate.GetResult(v)
		r.Equal(validate.SUCCESS, code)
	}
	{
		// 中断
		v := validate.NewValidator()
		validate.SetHasError(v)
		validate.CheckRule(v, true, []validate.RuleOption{validate.Code("MY_CODE")}, nil, func() bool {
			return false
		}, func() bool {
			return false
		})
		code, _, _ := validate.GetResult(v)
		r.Equal(validate.FAIL, code)
	}
}

func TestFieldTypeFunc(t *testing.T) {
	r := require.New(t)
	v := validate.NewValidator()
	r.NotNil(v.Bool(new(true), "param"))
	r.NotNil(v.Byte(new(byte(1)), "param"))
	r.NotNil(v.Int(new(1), "param"))
	r.NotNil(v.Int8(new(int8(1)), "param"))
	r.NotNil(v.Int16(new(int16(1)), "param"))
	r.NotNil(v.Int32(new(int32(1)), "param"))
	r.NotNil(v.Int64(new(int64(1)), "param"))
	r.NotNil(v.Uint(new(uint(1)), "param"))
	r.NotNil(v.Uint8(new(uint8(1)), "param"))
	r.NotNil(v.Uint16(new(uint16(1)), "param"))
	r.NotNil(v.Uint32(new(uint32(1)), "param"))
	r.NotNil(v.Uint64(new(uint64(1)), "param"))
	r.NotNil(v.Float32(new(float32(0.1)), "param"))
	r.NotNil(v.Float64(new(0.1), "param"))
}

func TestCheck(t *testing.T) {
	r := require.New(t)
	{
		u := demoUser{
			id:   "1147",
			name: "name",
		}
		code, msg, msgs := validate.Check(nil, &u)
		r.Equal(validate.SUCCESS, code)
		r.Equal("", msg)
		r.Nil(msgs)
	}
	{
		var u *demoUser
		code, msg, msgs := validate.Check(nil, u)
		r.Equal("", code)
		r.Equal("", msg)
		r.Nil(msgs)
	}
}

func TestDive(t *testing.T) {
	r := require.New(t)
	u := demoUser{
		attach: &demoAttach{
			images:   []demoImage{{}, {}},
			score:    map[string]int{"math": 101},
			profiles: map[string]demoProfile{"123456": {}},
		},
	}

	code, _, msgs := validate.Check(nil, &u, validate.All())
	r.Equal(validate.FAIL, code)
	r.Len(msgs, 10)
	r.Equal("id must not be blank", msgs[0])
	r.Equal("name must not be blank", msgs[1])
	r.Equal("attach.id must not be blank", msgs[2])
	r.Equal("attach.images[0].url must not be blank", msgs[3])
	r.Equal("attach.images[0].id must not be blank", msgs[4])
	r.Equal("attach.images[1].url must not be blank", msgs[5])
	r.Equal("attach.images[1].id must not be blank", msgs[6])
	r.Equal("attach.score$value must not be greater than 100", msgs[7])
	r.Equal("attach.profiles$key's length must not be greater than 4", msgs[8])
	r.Equal("attach.profiles$value.name must not be blank", msgs[9])
}

type demoBase struct {
	id string
}

func (d demoBase) Checklist(v *validate.Validator) {
	v.String(&d.id, "id").NotBlank()
}

type demoUser struct {
	demoBase
	name   string
	attach *demoAttach
}

func (d demoUser) Checklist(v *validate.Validator) {
	v.Embed(&d.demoBase)
	v.String(&d.name, "name").NotBlank()
	v.Struct(d.attach, "attach").Dive()
}

type demoAttach struct {
	demoBase
	images   []demoImage
	score    map[string]int
	profiles map[string]demoProfile
}

func (d demoAttach) Checklist(v *validate.Validator) {
	v.Embed(&d.demoBase)
	v.Slice(d.images, "images").Dive(func(i demoImage) {
		v.Struct(&i, "not work").Dive()
	})
	v.Map(d.score, "score").Dive(nil, func(value int) {
		v.Int(&value, "not work").Max(100)
	})
	v.Map(d.profiles, "profiles").Dive(func(k string) {
		v.String(&k, "not work").Max(4)
	}, func(value demoProfile) {
		v.Struct(&value, "not work").Dive()
	})
}

type demoImage struct {
	url string
	demoBase
}

func (d demoImage) Checklist(v *validate.Validator) {
	v.String(&d.url, "url").NotBlank()
	v.Embed(&d.demoBase)
}

type demoProfile struct {
	Name string
}

func (s demoProfile) Checklist(v *validate.Validator) {
	v.String(&s.Name, "name").NotBlank()
}
