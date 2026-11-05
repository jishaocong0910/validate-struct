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

// 校验结构体
type checkStruct[T any, V validatable[T]] struct {
	v *Validator
	s V
}

// Required 限制不能为nil
func (c *checkStruct[T, V]) Required(opts ...RuleOption) *checkStruct[T, V] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageRequired, nil, func() bool {
		return false
	}, nil)
	return c
}

// Custom 自定义校验
func (c *checkStruct[T, V]) Custom(successIfNil bool, custom func(s T) bool, opts ...RuleOption) *checkStruct[T, V] {
	c.v.beforeCheckRule(opts)
	c.v.checkRule(c.v.fieldInfo.virtualDive || c.s == nil, messageDefault, nil, func() bool {
		return successIfNil
	}, func() bool {
		return custom(*c.s)
	})
	return c
}

// Dive 下沉校验结构体的字段
func (c *checkStruct[T, V]) Dive() {
	if c.v.shouldStop() {
		return
	}
	allowed, after := c.v.beforeDiveStruct(c.s)
	if allowed {
		c.s.Checklist(c.v)
		if after != nil {
			after()
		}
	}
}
