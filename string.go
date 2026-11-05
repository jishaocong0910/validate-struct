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
	"regexp"
	"slices"
	"unicode"
	"unicode/utf8"
)

// 校验字符串类型
type checkString struct {
	v *Validator
	s *string
}

// Required 限制不能为nil
func (c *checkString) Required(opts ...RuleOption) *checkString {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageRequired, nil, func() bool {
		return false
	}, nil)
	return c
}

// NotBlank 限制不能为空白
func (c *checkString) NotBlank(opts ...RuleOption) *checkString {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageNotBlank, nil, func() bool {
		return false
	}, func() bool {
		for _, r := range *c.s {
			if !unicode.IsSpace(r) {
				return true
			}
		}
		return false
	})
	return c
}

// Regex 限制必须匹配正则
func (c *checkString) Regex(r *regexp.Regexp, opts ...RuleOption) *checkString {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageRegex, nil, func() bool {
		return r.MatchString("")
	}, func() bool {
		return r.MatchString(*c.s)
	})
	return c
}

// Length 限制字符长度
func (c *checkString) Length(l int, opts ...RuleOption) *checkString {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageLength, func() []string {
		return formatInt(l)
	}, func() bool {
		return l == 0
	}, func() bool {
		return utf8.RuneCountInString(*c.s) == l
	})
	return c
}

// Min 限制字符长度最小值
func (c *checkString) Min(min int, opts ...RuleOption) *checkString {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageLengthMin, func() []string {
		return formatInt(min)
	}, func() bool {
		return min <= 0
	}, func() bool {
		return utf8.RuneCountInString(*c.s) >= min
	})
	return c
}

// Max 限制字符长度最大值
func (c *checkString) Max(max int, opts ...RuleOption) *checkString {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageLengthMax, func() []string {
		return formatInt(max)
	}, func() bool {
		return max >= 0
	}, func() bool {
		return utf8.RuneCountInString(*c.s) <= max
	})
	return c
}

// BetweenInclusive 限制字符长度范围（包含边界）
func (c *checkString) BetweenInclusive(min, max int, opts ...RuleOption) *checkString {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageLengthBetweenInclusive, func() []string {
		return formatInt(min, max)
	}, func() bool {
		return min <= 0 && max >= 0
	}, func() bool {
		l := utf8.RuneCountInString(*c.s)
		return l >= min && l <= max
	})
	return c
}

// Gt 限制字符长度必须大于指定值
func (c *checkString) Gt(min int, opts ...RuleOption) *checkString {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageLengthGt, func() []string {
		return formatInt(min)
	}, func() bool {
		return min < 0
	}, func() bool {
		return utf8.RuneCountInString(*c.s) > min
	})
	return c
}

// Lt 限制字符长度必须小于指定值
func (c *checkString) Lt(max int, opts ...RuleOption) *checkString {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageLengthLt, func() []string {
		return formatInt(max)
	}, func() bool {
		return max > 0
	}, func() bool {
		return utf8.RuneCountInString(*c.s) < max
	})
	return c
}

// BetweenExclusive 限制字符串长度范围（不包含边界）
func (c *checkString) BetweenExclusive(min, max int, opts ...RuleOption) *checkString {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageLengthBetweenExclusive, func() []string {
		return formatInt(min, max)
	}, func() bool {
		return min < 0 && max > 0
	}, func() bool {
		l := utf8.RuneCountInString(*c.s)
		return l > min && l < max
	})
	return c
}

// Enum 限制值必须在指定枚举值中
func (c *checkString) Enum(enums []string, opts ...RuleOption) *checkString {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageEnum, func() []string {
		var args []string
		for _, e := range enums {
			args = append(args, "\""+e+"\"")
		}
		return args
	}, func() bool {
		return false
	}, func() bool {
		return slices.Contains(enums, *c.s)
	})
	return c
}

// Custom 自定义校验
func (c *checkString) Custom(successIfNil bool, custom func(s string) bool, opts ...RuleOption) *checkString {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageDefault, nil, func() bool {
		return successIfNil
	}, func() bool {
		return custom(*c.s)
	})
	return c
}
