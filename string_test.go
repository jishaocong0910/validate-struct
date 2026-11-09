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
	"regexp"
	"testing"

	validate "github.com/jishaocong0910/validate-struct"
	"github.com/stretchr/testify/require"
)

func TestCheckString_Required(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Required()
	}, "param is required")
	checklistSuccess(r, func(v *validate.Validator) {
		v.String(new("a"), "param").Required()
	})
}

func TestCheckString_NotBlank(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.String(new(""), "param").NotBlank()
	}, "param must not be blank")
	checklistSuccess(r, func(v *validate.Validator) {
		v.String(new("a"), "param").NotBlank()
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").NotBlank()
	})
}

func TestCheckString_Regex(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.String(new("aaa"), "param").Regex(regexp.MustCompile(`\d+`))
	}, "param has an invalid format")
	checklistSuccess(r, func(v *validate.Validator) {
		v.String(new("123"), "param").Regex(regexp.MustCompile(`\d+`))
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Regex(regexp.MustCompile(`\d+`))
	})
}

func TestCheckString_Custom(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.String(new("a"), "param").Custom(func(t string) bool {
			return false
		})
	}, "param is invalid")
	checklistSuccess(r, func(v *validate.Validator) {
		v.String(new("a"), "param").Custom(func(t string) bool {
			return true
		})
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Custom(func(t string) bool {
			return false
		})
	})
}

func TestCheckString_Length(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.String(new("中"), "param").Length(2)
	}, "param must have a length of 2")
	checklistFail(r, func(v *validate.Validator) {
		v.String(new("中文中"), "param").Length(2)
	}, "param must have a length of 2")
	checklistSuccess(r, func(v *validate.Validator) {
		v.String(new("中文"), "param").Length(2)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Length(1)
	})
}

func TestCheckString_Min(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.String(new("中"), "param").Min(2)
	}, "param must have a length of at least 2")
	checklistSuccess(r, func(v *validate.Validator) {
		v.String(new("中文"), "param").Min(2)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Min(1)
	})
}

func TestCheckString_Max(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.String(new("中文中"), "param").Max(2)
	}, "param must have a length of at most 2")
	checklistSuccess(r, func(v *validate.Validator) {
		v.String(new("中文"), "param").Max(2)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Max(-1)
	})
}

func TestCheckString_Range(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.String(new("中"), "param").Range(2, 3)
	}, "param must have a length between 2 and 3, inclusive")
	checklistFail(r, func(v *validate.Validator) {
		v.String(new("中文中文"), "param").Range(2, 3)
	}, "param must have a length between 2 and 3, inclusive")
	checklistSuccess(r, func(v *validate.Validator) {
		v.String(new("中文中"), "param").Range(2, 3)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Range(1, 3)
	})
}

func TestCheckString_Gt(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.String(new("中文"), "param").Gt(2)
	}, "param must have a length greater than 2")
	checklistSuccess(r, func(v *validate.Validator) {
		v.String(new("中文中"), "param").Gt(2)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Gt(0)
	})
}

func TestCheckString_Lt(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.String(new("中文"), "param").Lt(2)
	}, "param must have a length less than 2")
	checklistSuccess(r, func(v *validate.Validator) {
		v.String(new("中"), "param").Lt(2)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Lt(0)
	})
}

func TestCheckString_RangeExclusive(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.String(new("中"), "param").RangeExclusive(2, 5)
	}, "param must have a length between 2 and 5, exclusive")
	checklistFail(r, func(v *validate.Validator) {
		v.String(new("中文中文中文"), "param").RangeExclusive(2, 5)
	}, "param must have a length between 2 and 5, exclusive")
	checklistSuccess(r, func(v *validate.Validator) {
		v.String(new("中文中"), "param").RangeExclusive(2, 5)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").RangeExclusive(0, 5)
	})
}

func TestCheckString_Enum(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.String(new("d"), "param").Enum([]string{"a", "b", "c"})
	}, `param must be "a", "b", or "c"`)
	checklistSuccess(r, func(v *validate.Validator) {
		v.String(new("c"), "param").Enum([]string{"a", "b", "c"})
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Enum([]string{"a", "b", "c"})
	})
}
