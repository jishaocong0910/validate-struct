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
	v               *Validator
	n               *T
	numberFormatter numberFormatter[T]
}

func (c *checkNumber[T]) Required(opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule[int, T](c.n, opts, msgBuilderRequired, nil, func() bool {
		return false
	}, nil)
	return c
}

func (c *checkNumber[T]) Min(min T, opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule[int, T](c.n, opts, msgBuilderMin, func() []string {
		return c.numberFormatter(min)
	}, func() bool {
		return false
	}, func() bool {
		return *c.n >= min
	})
	return c
}

func (c *checkNumber[T]) Max(max T, opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule[int, T](c.n, opts, msgBuilderMax, func() []string {
		return c.numberFormatter(max)
	}, func() bool {
		return true
	}, func() bool {
		return *c.n <= max
	})
	return c
}

func (c *checkNumber[T]) Range(min, max T, opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule[int, T](c.n, opts, msgBuilderRange, func() []string {
		return c.numberFormatter(min, max)
	}, func() bool {
		return false
	}, func() bool {
		return *c.n >= min && *c.n <= max
	})
	return c
}

func (c *checkNumber[T]) Gt(min T, opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule[int, T](c.n, opts, msgBuilderGt, func() []string {
		return c.numberFormatter(min)
	}, func() bool {
		return false
	}, func() bool {
		return *c.n > min
	})
	return c
}

func (c *checkNumber[T]) Lt(max T, opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule[int, T](c.n, opts, msgBuilderLt, func() []string {
		return c.numberFormatter(max)
	}, func() bool {
		return true
	}, func() bool {
		return *c.n < max
	})
	return c
}

func (c *checkNumber[T]) Within(min, max T, opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule[int, T](c.n, opts, msgBuilderWithin, func() []string {
		return c.numberFormatter(min, max)
	}, func() bool {
		return false
	}, func() bool {
		return *c.n > min && *c.n < max
	})
	return c
}

func (c *checkNumber[T]) Enum(enums []T, opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule[int, T](c.n, opts, msgBuilderEnum, func() []string {
		return c.numberFormatter(enums...)
	}, func() bool {
		return false
	}, func() bool {
		return slices.Contains(enums, *c.n)
	})
	return c
}

func (c *checkNumber[T]) Custom(successIfNil bool, custom func(n T) bool, opts ...RuleOption) *checkNumber[T] {
	c.v.checkRule[int, T](c.n, opts, msgBuilderDefault, nil,
		func() bool {
			return successIfNil
		}, func() bool {
			return custom(*c.n)
		})
	return c
}

func intToStr(ii ...int) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(int64(i), 10))
	}
	return result
}

func int8ToStr(ii ...int8) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(int64(i), 10))
	}
	return result
}

func int16ToStr(ii ...int16) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(int64(i), 10))
	}
	return result
}

func int32ToStr(ii ...int32) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(int64(i), 10))
	}
	return result
}

func int64ToStr(ii ...int64) []string {
	result := make([]string, 0, len(ii))
	for _, i := range ii {
		result = append(result, strconv.FormatInt(i, 10))
	}
	return result
}

func uintToStr(uu ...uint) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(uint64(u), 10))
	}
	return result
}

func uint8ToStr(uu ...uint8) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(uint64(u), 10))
	}
	return result
}

func uint16ToStr(uu ...uint16) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(uint64(u), 10))
	}
	return result
}

func uint32ToStr(uu ...uint32) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(uint64(u), 10))
	}
	return result
}

func uint64ToStr(uu ...uint64) []string {
	result := make([]string, 0, len(uu))
	for _, u := range uu {
		result = append(result, strconv.FormatUint(u, 10))
	}
	return result
}

func float32ToStr(ff ...float32) []string {
	result := make([]string, 0, len(ff))
	for _, f := range ff {
		result = append(result, strconv.FormatFloat(float64(f), 'f', -1, 32))
	}
	return result
}

func float64ToStr(ff ...float64) []string {
	result := make([]string, 0, len(ff))
	for _, f := range ff {
		result = append(result, strconv.FormatFloat(f, 'f', -1, 32))
	}
	return result
}
