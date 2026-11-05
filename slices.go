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

import "strconv"

// 校验切片
type checkSlice[T any] struct {
	v *Validator
	s []T
}

// Required 限制不能为nil
func (c *checkSlice[T]) Required(opts ...RuleOption) *checkSlice[T] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageRequired, nil, func() bool {
		return false
	}, nil)
	return c
}

// NotEmpty 限制不能为空
func (c *checkSlice[T]) NotEmpty(opts ...RuleOption) *checkSlice[T] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageNotEmpty, nil, func() bool {
		return false
	}, func() bool {
		return len(c.s) != 0
	})
	return c
}

// Length 限制长度范围
func (c *checkSlice[T]) Length(l int, opts ...RuleOption) *checkSlice[T] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageLength, func() []string {
		return formatInt(l)
	}, func() bool {
		return l == 0
	}, func() bool {
		return len(c.s) == l
	})
	return c
}

// Min 限制长度最小值
func (c *checkSlice[T]) Min(min int, opts ...RuleOption) *checkSlice[T] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageLengthMin, func() []string {
		return formatInt(min)
	}, func() bool {
		return min <= 0
	}, func() bool {
		return len(c.s) >= min
	})
	return c
}

// Max 限制长度最大值
func (c *checkSlice[T]) Max(max int, opts ...RuleOption) *checkSlice[T] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageLengthMax, func() []string {
		return formatInt(max)
	}, func() bool {
		return max >= 0
	}, func() bool {
		return len(c.s) <= max
	})
	return c
}

// BetweenInclusive 限制长度范围（包含边界）
func (c *checkSlice[T]) BetweenInclusive(min, max int, opts ...RuleOption) *checkSlice[T] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageLengthBetweenInclusive, func() []string {
		return formatInt(min, max)
	}, func() bool {
		return min <= 0 && max >= 0
	}, func() bool {
		l := len(c.s)
		return l >= min && l <= max
	})
	return c
}

// Gt 限制长度必须大于指定值
func (c *checkSlice[T]) Gt(min int, opts ...RuleOption) *checkSlice[T] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageLengthGt, func() []string {
		return formatInt(min)
	}, func() bool {
		return min < 0
	}, func() bool {
		return len(c.s) > min
	})
	return c
}

// Lt 限制长度必须小于指定值
func (c *checkSlice[T]) Lt(max int, opts ...RuleOption) *checkSlice[T] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageLengthLt, func() []string {
		return formatInt(max)
	}, func() bool {
		return max > 0
	}, func() bool {
		return len(c.s) < max
	})
	return c
}

// BetweenExclusive 限制长度范围（不包含边界）
func (c *checkSlice[T]) BetweenExclusive(min, max int, opts ...RuleOption) *checkSlice[T] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageLengthBetweenExclusive, func() []string {
		return formatInt(min, max)
	}, func() bool {
		return min < 0 && max > 0
	}, func() bool {
		l := len(c.s)
		return l > min && l < max
	})
	return c
}

// Custom 自定义校验
func (c *checkSlice[T]) Custom(successIfNil bool, custom func(s []T) bool, opts ...RuleOption) *checkSlice[T] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageDefault, nil, func() bool {
		return successIfNil
	}, func() bool {
		return custom(c.s)
	})
	return c
}

// Dive 下沉校验元素
func (c *checkSlice[T]) Dive(checkElem func(t T)) {
	if c.v.shouldStop() {
		return
	}
	if c.s != nil && checkElem != nil {
		for i, t := range c.s {
			if c.v.shouldStop() {
				break
			}
			after := c.v.beforeDiveSliceMap("[" + strconv.Itoa(i) + "]")
			checkElem(t)
			after()
		}
	}
}
