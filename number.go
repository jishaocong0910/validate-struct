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

import "slices"

import "strconv"

type number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type numberFormatter[T number] func(t ...T) []string

type checkNumber[T number] struct {
	v         *Validator
	n         *T
	formatter numberFormatter[T]
}

func (c *checkNumber[T]) isNil() bool {
	return c.n == nil || c.v.fieldInfo.virtualDive
}

func (c *checkNumber[T]) Required(opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule(opts, messageRequired, nil, func() bool {
		return !c.isNil()
	})
	return c
}

func (c *checkNumber[T]) Min(min T, opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule(opts, messageMin, func() []string {
		return c.formatter(min)
	}, func() bool {
		return c.isNil() || *c.n >= min
	})
	return c
}

func (c *checkNumber[T]) Max(max T, opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule(opts, messageMax, func() []string {
		return c.formatter(max)
	}, func() bool {
		return c.isNil() || *c.n <= max
	})
	return c
}

func (c *checkNumber[T]) BetweenInclusive(min, max T, opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule(opts, messageBetweenInclusive, func() []string {
		return c.formatter(min, max)
	}, func() bool {
		return c.isNil() || *c.n >= min && *c.n <= max
	})
	return c
}

func (c *checkNumber[T]) Gt(min T, opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule(opts, messageGt, func() []string {
		return c.formatter(min)
	}, func() bool {
		return c.isNil() || *c.n > min
	})
	return c
}

func (c *checkNumber[T]) Lt(max T, opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule(opts, messageLt, func() []string {
		return c.formatter(max)
	}, func() bool {
		return c.isNil() || *c.n < max
	})
	return c
}

func (c *checkNumber[T]) BetweenExclusive(min, max T, opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule(opts, messageBetweenExclusive, func() []string {
		return c.formatter(min, max)
	}, func() bool {
		return c.isNil() || *c.n > min && *c.n < max
	})
	return c
}

func (c *checkNumber[T]) Enum(enums []T, opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule(opts, messageEnum, func() []string {
		return c.formatter(enums...)
	}, func() bool {
		return c.isNil() || slices.Contains(enums, *c.n)
	})
	return c
}

func (c *checkNumber[T]) Custom(custom func(n T) bool, opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule(opts, messageDefault, nil, func() bool {
		return c.isNil() || custom == nil || custom(*c.n)
	})
	return c
}

func formatInt(ii ...int) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(int64(i), 10))
	}
	return result
}

func formatInt8(ii ...int8) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(int64(i), 10))
	}
	return result
}

func formatInt16(ii ...int16) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(int64(i), 10))
	}
	return result
}

func formatInt32(ii ...int32) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(int64(i), 10))
	}
	return result
}

func formatInt64(ii ...int64) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(i, 10))
	}
	return result
}

func formatUint(uu ...uint) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(uint64(u), 10))
	}
	return result
}

func formatUint8(uu ...uint8) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(uint64(u), 10))
	}
	return result
}

func formatUint16(uu ...uint16) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(uint64(u), 10))
	}
	return result
}

func formatUint32(uu ...uint32) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(uint64(u), 10))
	}
	return result
}

func formatUint64(uu ...uint64) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(u, 10))
	}
	return result
}

func formatFloat32(ff ...float32) []string {
	result := make([]string, 0, len(ff))
	for _, f := range ff {
		result = append(result, strconv.FormatFloat(float64(f), 'f', -1, 32))
	}
	return result
}

func formatFloat64(ff ...float64) []string {
	result := make([]string, 0, len(ff))
	for _, f := range ff {
		result = append(result, strconv.FormatFloat(f, 'f', -1, 64))
	}
	return result
}
