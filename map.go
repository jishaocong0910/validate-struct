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

// 校验map类型
type checkMap[K comparable, V any] struct {
	v *Validator
	m map[K]V
}

// Required 限制不能为nil
func (c *checkMap[K, V]) Required(opts ...RuleOption) *checkMap[K, V] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.parentIsNil || c.m == nil, messageRequired, nil, func() bool {
		return false
	}, nil)
	return c
}

// NotEmpty 限制不能为空
func (c *checkMap[K, V]) NotEmpty(opts ...RuleOption) *checkMap[K, V] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.parentIsNil || c.m == nil, messageNotEmpty, nil, func() bool {
		return false
	}, func() bool {
		return len(c.m) > 0
	})
	return c
}

// Length 限制长度必须为指定值
func (c *checkMap[K, V]) Length(l int, opts ...RuleOption) *checkMap[K, V] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.parentIsNil || c.m == nil, messageLength, func() []string {
		return formatInt(l)
	}, func() bool {
		return l == 0
	}, func() bool {
		return len(c.m) == l
	})
	return c
}

// Min 限制最小长度
func (c *checkMap[K, V]) Min(min int, opts ...RuleOption) *checkMap[K, V] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.parentIsNil || c.m == nil, messageLengthMin, func() []string {
		return formatInt(min)
	}, func() bool {
		return min <= 0
	}, func() bool {
		return len(c.m) >= min
	})
	return c
}

// Max 限制最大长度
func (c *checkMap[K, V]) Max(max int, opts ...RuleOption) *checkMap[K, V] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.parentIsNil || c.m == nil, messageLengthMax, func() []string {
		return formatInt(max)
	}, func() bool {
		return max >= 0
	}, func() bool {
		return len(c.m) <= max
	})
	return c
}

// BetweenInclusive 限制长度范围（包含边界）
func (c *checkMap[K, V]) BetweenInclusive(min, max int, opts ...RuleOption) *checkMap[K, V] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.parentIsNil || c.m == nil, messageLengthBetweenInclusive, func() []string {
		return formatInt(min, max)
	}, func() bool {
		return min <= 0 && max >= 0
	}, func() bool {
		l := len(c.m)
		return l >= min && l <= max
	})
	return c
}

// Gt 限制长度必须大于指定值
func (c *checkMap[K, V]) Gt(min int, opts ...RuleOption) *checkMap[K, V] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.parentIsNil || c.m == nil, messageLengthGt, func() []string {
		return formatInt(min)
	}, func() bool {
		return min < 0
	}, func() bool {
		return len(c.m) > min
	})
	return c
}

// Lt 限制长度必须小于指定值
func (c *checkMap[K, V]) Lt(max int, opts ...RuleOption) *checkMap[K, V] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.parentIsNil || c.m == nil, messageLengthLt, func() []string {
		return formatInt(max)
	}, func() bool {
		return max > 0
	}, func() bool {
		return len(c.m) < max
	})
	return c
}

// BetweenExclusive 限制长度范围（不包含边界）
func (c *checkMap[K, V]) BetweenExclusive(min, max int, opts ...RuleOption) *checkMap[K, V] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.parentIsNil || c.m == nil, messageLengthBetweenExclusive, func() []string {
		return formatInt(min, max)
	}, func() bool {
		return min < 0 && max > 0
	}, func() bool {
		l := len(c.m)
		return l > min && l < max
	})
	return c
}

// Custom 自定义校验
func (c *checkMap[K, V]) Custom(successIfNil bool, custom func(m map[K]V) bool, opts ...RuleOption) *checkMap[K, V] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.parentIsNil || c.m == nil, messageDefault, nil, func() bool {
		return successIfNil
	}, func() bool {
		return custom(c.m)
	})
	return c
}

// Dive 下沉校验key和value
func (c *checkMap[K, V]) Dive(checkKey func(key K), checkValue func(value V)) {
	if c.v.shouldStop() {
		return
	}
	if c.m != nil {
		for k, v := range c.m {
			if checkKey != nil {
				if c.v.shouldStop() {
					break
				}
				fi := c.v.fieldInfo
				c.v.beforeDiveSliceMap("$key")
				checkKey(k)
				c.v.fieldInfo = fi
			}
			if checkValue != nil {
				if c.v.shouldStop() {
					break
				}
				fi := c.v.fieldInfo
				c.v.beforeDiveSliceMap("$value")
				checkValue(v)
				c.v.fieldInfo = fi
			}
		}
	}
}
