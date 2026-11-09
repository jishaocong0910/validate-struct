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

package validate_test

import (
	"testing"

	validate "github.com/jishaocong0910/validate-struct"
	"github.com/stretchr/testify/require"
)

func TestCheckNumber_Required(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Int((*int)(nil), "param").Required()
	}, "param is required")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Int(new(5), "param").Required()
	})
}

func TestCheckNumber_Min(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Int(new(5), "param").Min(6)
	}, "param must be at least 6")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Int(new(5), "param").Min(5)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Int((*int)(nil), "param").Min(0)
	})
}

func TestCheckNumber_Max(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Int(new(5), "param").Max(4)
	}, "param must be at most 4")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Int(new(5), "param").Max(5)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Int((*int)(nil), "param").Max(0)
	})
}

func TestCheckNumber_BetweenInclusive(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Int(new(5), "param").BetweenInclusive(1, 4)
	}, "param must be between 1 and 4, inclusive")
	checklistFail(r, func(v *validate.Validator) {
		v.Int(new(5), "param").BetweenInclusive(6, 9)
	}, "param must be between 6 and 9, inclusive")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Int(new(5), "param").BetweenInclusive(1, 9)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Int((*int)(nil), "param").BetweenInclusive(0, 9)
	})
}

func TestCheckNumber_Gt(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Int(new(5), "param").Gt(5)
	}, "param must be greater than 5")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Int(new(5), "param").Gt(4)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Int((*int)(nil), "param").Gt(0)
	})
}

func TestCheckNumber_Lt(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Int(new(5), "param").Lt(5)
	}, "param must be less than 5")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Int(new(5), "param").Lt(6)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Int((*int)(nil), "param").Lt(0)
	})
}

func TestCheckNumber_BetweenExclusive(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Int(new(5), "param").BetweenExclusive(1, 5)
	}, "param must be between 1 and 5, exclusive")
	checklistFail(r, func(v *validate.Validator) {
		v.Int(new(5), "param").BetweenExclusive(5, 9)
	}, "param must be between 5 and 9, exclusive")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Int(new(5), "param").BetweenExclusive(1, 9)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Int((*int)(nil), "param").BetweenExclusive(-5, 5)
	})
}

func TestCheckNumber_Enum(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Int(new(5), "param").Enum([]int{1})
	}, "param must be 1")
	checklistFail(r, func(v *validate.Validator) {
		v.Int(new(5), "param").Enum([]int{1, 2})
	}, "param must be 1 or 2")
	checklistFail(r, func(v *validate.Validator) {
		v.Int(new(5), "param").Enum([]int{1, 2, 3})
	}, "param must be 1, 2, or 3")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Int(new(5), "param").Enum([]int{3, 4, 5})
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Int((*int)(nil), "param").Enum([]int{1, 2, 3})
	})
}

func TestCheckNumber_Custom(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Int(new(5), "param").Custom(func(t int) bool {
			return false
		})
	}, "param is invalid")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Int(new(5), "param").Custom(func(t int) bool {
			return true
		})
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Int((*int)(nil), "param").Custom(func(t int) bool {
			return false
		})
	})
}

func TestFormatNumber(t *testing.T) {
	r := require.New(t)
	testFormatNumber(r, validate.FormatInt)
	testFormatNumber(r, validate.FormatInt8)
	testFormatNumber(r, validate.FormatInt16)
	testFormatNumber(r, validate.FormatInt32)
	testFormatNumber(r, validate.FormatInt64)
	testFormatNumber(r, validate.FormatUint)
	testFormatNumber(r, validate.FormatUint8)
	testFormatNumber(r, validate.FormatUint16)
	testFormatNumber(r, validate.FormatUint32)
	testFormatNumber(r, validate.FormatUint64)
	testFormatNumber(r, validate.FormatFloat32)
	testFormatNumber(r, validate.FormatFloat64)
}

func testFormatNumber[T validate.Number](r *require.Assertions, formatNumber func(t ...T) []string) {
	args := formatNumber(1, 20, 100)
	r.Len(args, 3)
	r.Contains(args, "1")
	r.Contains(args, "20")
	r.Contains(args, "100")
}
