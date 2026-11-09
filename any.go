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

type checkAny[T any] struct {
	v *Validator
	a *T
}

func (c *checkAny[T]) isNil() bool {
	return c.a == nil || c.v.fieldInfo.virtualDive
}

func (c *checkAny[T]) Required(opts ...RuleOption) *checkAny[T] {
	c.v.checkRule(opts, messageRequired, nil, func() bool {
		return !c.isNil()
	})
	return c
}

func (c *checkAny[T]) Custom(custom func(a T) bool, opts ...RuleOption) *checkAny[T] {
	c.v.checkRule(opts, messageDefault, nil, func() bool {
		return c.isNil() || custom == nil || custom(*c.a)
	})
	return c
}
