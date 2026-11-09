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

func (c *checkMap[K, V]) isNil() bool {
	return c.m == nil || c.v.fieldInfo.virtualDive
}

// Required 限制不能为nil
func (c *checkMap[K, V]) Required(opts ...RuleOption) *checkMap[K, V] {
	c.v.checkRule(opts, messageRequired, nil, func() bool {
		return !c.isNil()
	})
	return c
}

// NotEmpty 限制不能为空
func (c *checkMap[K, V]) NotEmpty(opts ...RuleOption) *checkMap[K, V] {
	c.v.checkRule(opts, messageNotEmpty, nil, func() bool {
		return c.isNil() || len(c.m) > 0
	})
	return c
}

// Length 限制长度必须为指定值
func (c *checkMap[K, V]) Length(l int, opts ...RuleOption) *checkMap[K, V] {
	c.v.checkRule(opts, messageLength, func() []string {
		return formatInt(l)
	}, func() bool {
		return c.isNil() || len(c.m) == l
	})
	return c
}

// Min 限制最小长度
func (c *checkMap[K, V]) Min(min int, opts ...RuleOption) *checkMap[K, V] {
	c.v.checkRule(opts, messageLengthMin, func() []string {
		return formatInt(min)
	}, func() bool {
		return c.isNil() || len(c.m) >= min
	})
	return c
}

// Max 限制最大长度
func (c *checkMap[K, V]) Max(max int, opts ...RuleOption) *checkMap[K, V] {
	c.v.checkRule(opts, messageLengthMax, func() []string {
		return formatInt(max)
	}, func() bool {
		return c.isNil() || len(c.m) <= max
	})
	return c
}

// BetweenInclusive 限制长度范围（包含边界）
func (c *checkMap[K, V]) BetweenInclusive(min, max int, opts ...RuleOption) *checkMap[K, V] {
	c.v.checkRule(opts, messageLengthBetweenInclusive, func() []string {
		return formatInt(min, max)
	}, func() bool {
		return c.isNil() || len(c.m) >= min && len(c.m) <= max
	})
	return c
}

// Gt 限制长度必须大于指定值
func (c *checkMap[K, V]) Gt(min int, opts ...RuleOption) *checkMap[K, V] {
	c.v.checkRule(opts, messageLengthGt, func() []string {
		return formatInt(min)
	}, func() bool {
		return c.isNil() || len(c.m) > min
	})
	return c
}

// Lt 限制长度必须小于指定值
func (c *checkMap[K, V]) Lt(max int, opts ...RuleOption) *checkMap[K, V] {
	c.v.checkRule(opts, messageLengthLt, func() []string {
		return formatInt(max)
	}, func() bool {
		return c.isNil() || len(c.m) < max
	})
	return c
}

// BetweenExclusive 限制长度范围（不包含边界）
func (c *checkMap[K, V]) BetweenExclusive(min, max int, opts ...RuleOption) *checkMap[K, V] {
	c.v.checkRule(opts, messageLengthBetweenExclusive, func() []string {
		return formatInt(min, max)
	}, func() bool {
		return c.isNil() || len(c.m) > min && len(c.m) < max
	})
	return c
}

// Custom 自定义校验
func (c *checkMap[K, V]) Custom(custom func(m map[K]V) bool, opts ...RuleOption) *checkMap[K, V] {
	c.v.checkRule(opts, messageDefault, nil, func() bool {
		return c.isNil() || custom == nil || custom(c.m)
	})
	return c
}

// Dive 下沉校验key和value
func (c *checkMap[K, V]) Dive(checkKey func(key K), checkValue func(value V)) {
	if c.v.shouldStop() {
		return
	}
	if c.m != nil && !c.v.fieldInfo.virtualDive {
		for k, v := range c.m {
			if checkKey != nil {
				if c.v.shouldStop() {
					break
				}
				after := c.v.beforeDiveSliceMap("$key")
				checkKey(k)
				after()
			}
			if checkValue != nil {
				if c.v.shouldStop() {
					break
				}
				after := c.v.beforeDiveSliceMap("$value")
				checkValue(v)
				after()
			}
		}
	}
}
