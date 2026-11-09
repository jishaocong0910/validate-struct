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

type checkString struct {
	v *Validator
	s *string
}

func (c *checkString) isNil() bool {
	return c.s == nil || c.v.fieldInfo.virtualDive
}

func (c *checkString) Required(opts ...RuleOption) *checkString {
	c.v.checkRule(opts, messageRequired, nil, func() bool {
		return !c.isNil()
	})
	return c
}

func (c *checkString) NotBlank(opts ...RuleOption) *checkString {
	c.v.checkRule(opts, messageNotBlank, nil, func() bool {
		if c.isNil() {
			return true
		}
		for _, r := range *c.s {
			if !unicode.IsSpace(r) {
				return true
			}
		}
		return false
	})
	return c
}

func (c *checkString) Regex(r *regexp.Regexp, opts ...RuleOption) *checkString {
	c.v.checkRule(opts, messageRegex, nil, func() bool {
		return c.isNil() || r.MatchString(*c.s)
	})
	return c
}

func (c *checkString) Length(l int, opts ...RuleOption) *checkString {
	c.v.checkRule(opts, messageLength, func() []string {
		return formatInt(l)
	}, func() bool {
		return c.isNil() || utf8.RuneCountInString(*c.s) == l
	})
	return c
}

func (c *checkString) Min(min int, opts ...RuleOption) *checkString {
	c.v.checkRule(opts, messageLengthMin, func() []string {
		return formatInt(min)
	}, func() bool {
		return c.isNil() || utf8.RuneCountInString(*c.s) >= min
	})
	return c
}

func (c *checkString) Max(max int, opts ...RuleOption) *checkString {
	c.v.checkRule(opts, messageLengthMax, func() []string {
		return formatInt(max)
	}, func() bool {
		return c.isNil() || utf8.RuneCountInString(*c.s) <= max
	})
	return c
}

func (c *checkString) Range(min, max int, opts ...RuleOption) *checkString {
	c.v.checkRule(opts, messageLengthRange, func() []string {
		return formatInt(min, max)
	}, func() bool {
		if c.isNil() {
			return true
		}
		l := utf8.RuneCountInString(*c.s)
		return l >= min && l <= max
	})
	return c
}

func (c *checkString) Gt(min int, opts ...RuleOption) *checkString {
	c.v.checkRule(opts, messageLengthGt, func() []string {
		return formatInt(min)
	}, func() bool {
		return c.isNil() || utf8.RuneCountInString(*c.s) > min
	})
	return c
}

func (c *checkString) Lt(max int, opts ...RuleOption) *checkString {
	c.v.checkRule(opts, messageLengthLt, func() []string {
		return formatInt(max)
	}, func() bool {
		return c.isNil() || utf8.RuneCountInString(*c.s) < max
	})
	return c
}

func (c *checkString) RangeExclusive(min, max int, opts ...RuleOption) *checkString {
	c.v.checkRule(opts, messageLengthRangeExclusive, func() []string {
		return formatInt(min, max)
	}, func() bool {
		if c.isNil() {
			return true
		}
		l := utf8.RuneCountInString(*c.s)
		return l > min && l < max
	})
	return c
}

func (c *checkString) Enum(enums []string, opts ...RuleOption) *checkString {
	c.v.checkRule(opts, messageEnum, func() []string {
		var args []string
		for _, e := range enums {
			args = append(args, "\""+e+"\"")
		}
		return args
	}, func() bool {
		return c.isNil() || slices.Contains(enums, *c.s)
	})
	return c
}

func (c *checkString) Custom(custom func(s string) bool, opts ...RuleOption) *checkString {
	c.v.checkRule(opts, messageDefault, nil, func() bool {
		return c.isNil() || custom == nil || custom(*c.s)
	})
	return c
}
