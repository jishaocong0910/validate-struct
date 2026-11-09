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

type checkSlice[T any] struct {
	v *Validator
	s []T
}

func (c *checkSlice[T]) isNil() bool {
	return c.s == nil || c.v.fieldInfo.virtualDive
}

func (c *checkSlice[T]) Required(opts ...RuleOption) *checkSlice[T] {
	c.v.checkRule(opts, messageRequired, nil, func() bool {
		return !c.isNil()
	})
	return c
}

func (c *checkSlice[T]) NotEmpty(opts ...RuleOption) *checkSlice[T] {
	c.v.checkRule(opts, messageNotEmpty, nil, func() bool {
		return c.isNil() || len(c.s) != 0
	})
	return c
}

func (c *checkSlice[T]) Length(l int, opts ...RuleOption) *checkSlice[T] {
	c.v.checkRule(opts, messageLength, func() []string {
		return formatInt(l)
	}, func() bool {
		return c.isNil() || len(c.s) == l
	})
	return c
}

func (c *checkSlice[T]) Min(min int, opts ...RuleOption) *checkSlice[T] {
	c.v.checkRule(opts, messageLengthMin, func() []string {
		return formatInt(min)
	}, func() bool {
		return c.isNil() || len(c.s) >= min
	})
	return c
}

func (c *checkSlice[T]) Max(max int, opts ...RuleOption) *checkSlice[T] {
	c.v.checkRule(opts, messageLengthMax, func() []string {
		return formatInt(max)
	}, func() bool {
		return c.isNil() || len(c.s) <= max
	})
	return c
}

func (c *checkSlice[T]) Range(min, max int, opts ...RuleOption) *checkSlice[T] {
	c.v.checkRule(opts, messageLengthRange, func() []string {
		return formatInt(min, max)
	}, func() bool {
		return c.isNil() || len(c.s) >= min && len(c.s) <= max
	})
	return c
}

func (c *checkSlice[T]) Gt(min int, opts ...RuleOption) *checkSlice[T] {
	c.v.checkRule(opts, messageLengthGt, func() []string {
		return formatInt(min)
	}, func() bool {
		return c.isNil() || len(c.s) > min
	})
	return c
}

func (c *checkSlice[T]) Lt(max int, opts ...RuleOption) *checkSlice[T] {
	c.v.checkRule(opts, messageLengthLt, func() []string {
		return formatInt(max)
	}, func() bool {
		return c.isNil() || len(c.s) < max
	})
	return c
}

func (c *checkSlice[T]) RangeExclusive(min, max int, opts ...RuleOption) *checkSlice[T] {
	c.v.checkRule(opts, messageLengthRangeExclusive, func() []string {
		return formatInt(min, max)
	}, func() bool {
		return c.isNil() || len(c.s) > min && len(c.s) < max
	})
	return c
}

func (c *checkSlice[T]) Custom(custom func(s []T) bool, opts ...RuleOption) *checkSlice[T] {
	c.v.checkRule(opts, messageDefault, nil, func() bool {
		return c.isNil() || custom == nil || custom(c.s)
	})
	return c
}

func (c *checkSlice[T]) Dive(checkElem func(t T)) {
	if c.v.shouldStop() {
		return
	}
	if c.s != nil && checkElem != nil && !c.v.fieldInfo.virtualDive {
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
