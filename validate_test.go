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

	"github.com/jishaocong0910/validate-struct"
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
		validate.CheckRule(v, false, []validate.RuleOption{validate.Code("MY_CODE"), validate.Message(func(b *validate.FieldInfo) string {
			return fmt.Sprintf("%s %s", b.FieldName(), b.Args())
		})}, func() []string {
			return []string{"a", "b", "c"}
		}, func() bool {
			return true
		}, func() bool {
			return false
		})
		code, message, _ := validate.GetResult(v)
		r.Equal("MY_CODE", code)
		r.Equal("param a, b, or c", message)
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
		code, message, _ := validate.GetResult(v)
		r.Equal(validate.FAIL, code)
		r.Equal("param is invalid", message)
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
		u := validate.DemoProfile{
			Name: "name",
		}
		code, message, messages := validate.Check(nil, &u)
		r.Equal(validate.SUCCESS, code)
		r.Equal("", message)
		r.Nil(messages)
	}
	{
		u := validate.DemoUser{
			Attach: &validate.DemoAttach{
				Images:   []validate.DemoImage{{}, {}},
				Score:    map[string]int{"math": 101},
				Profiles: map[string]validate.DemoProfile{"123456": {}},
			},
		}
		code, _, messages := validate.Check(nil, &u, validate.All())
		r.Equal(validate.FAIL, code)
		r.Len(messages, 12)
		r.Equal("id must not be blank", messages[0])
		r.Equal("name must not be blank", messages[1])
		r.Equal("attach.id must not be blank", messages[2])
		r.Equal("attach.images[0].url must not be blank", messages[3])
		r.Equal("attach.images[0].id is required", messages[4])
		r.Equal("attach.images[0].id must not be blank", messages[5])
		r.Equal("attach.images[1].url must not be blank", messages[6])
		r.Equal("attach.images[1].id is required", messages[7])
		r.Equal("attach.images[1].id must not be blank", messages[8])
		r.Equal("attach.score$value must be at most 100", messages[9])
		r.Equal("attach.profiles$key must have a length of at most 4", messages[10])
		r.Equal("attach.profiles$value.name must not be blank", messages[11])
	}
}
