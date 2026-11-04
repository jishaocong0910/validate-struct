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
	c.v.checkRule[any, T](c.s, messageRequired, nil, func() bool {
		return false
	}, nil)
	return c
}

// NotEmpty 限制不能为空
func (c *checkSlice[T]) NotEmpty(opts ...RuleOption) *checkSlice[T] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule[any, T](c.s, messageNotEmpty, nil, func() bool {
		return false
	}, func() bool {
		return len(c.s) != 0
	})
	return c
}

// Length 限制长度范围
func (c *checkSlice[T]) Length(l int, opts ...RuleOption) *checkSlice[T] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule[any, T](c.s, messageLength, func() []string {
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
	c.v.checkRule[any, T](c.s, messageLengthMin, func() []string {
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
	c.v.checkRule[any, T](c.s, messageLengthMax, func() []string {
		return formatInt(max)
	}, func() bool {
		return max >= 0
	}, func() bool {
		return len(c.s) <= max
	})
	return c
}

// Range 限制长度范围（包含边界）
func (c *checkSlice[T]) Range(min, max int, opts ...RuleOption) *checkSlice[T] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule[any, T](c.s, messageLengthRange, func() []string {
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
	c.v.checkRule[any, T](c.s, messageLengthGt, func() []string {
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
	c.v.checkRule[any, T](c.s, messageLengthLt, func() []string {
		return formatInt(max)
	}, func() bool {
		return max > 0
	}, func() bool {
		return len(c.s) < max
	})
	return c
}

// Within 限制长度范围（不包含边界）
func (c *checkSlice[T]) Within(min, max int, opts ...RuleOption) *checkSlice[T] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule[any, T](c.s, messageLengthWithin, func() []string {
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
	c.v.checkRule[any, T](c.s, messageDefault, nil, func() bool {
		return successIfNil
	}, func() bool {
		return custom(c.s)
	})
	return c
}

// Dive 下沉校验元素
func (c *checkSlice[T]) Dive(checkElem func(t T)) {
	if c.v.stop() {
		return
	}
	if c.s != nil && checkElem != nil {
		for i, t := range c.s {
			if c.v.stop() {
				break
			}
			fi := c.v.fieldInfo
			c.v.beforeDiveSliceMap("[" + strconv.Itoa(i) + "]")
			checkElem(t)
			c.v.fieldInfo = fi
		}
	}
}
