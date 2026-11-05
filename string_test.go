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
	testFail(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Required()
	}, "param is required")
	testSuccess(r, func(v *validate.Validator) {
		v.String(new("a"), "param").Required()
	})
}

func TestCheckString_NotBlank(t *testing.T) {
	r := require.New(t)
	testFail(r, func(v *validate.Validator) {
		v.String(new(""), "param").NotBlank()
	}, "param must not be blank")
	testSuccess(r, func(v *validate.Validator) {
		v.String(new("a"), "param").NotBlank()
	})
	testFail(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").NotBlank()
	}, "param must not be blank")
}

func TestCheckString_Regex(t *testing.T) {
	r := require.New(t)
	testFail(r, func(v *validate.Validator) {
		v.String(new("aaa"), "param").Regex(regexp.MustCompile(`\d+`))
	}, "param has an invalid format")
	testSuccess(r, func(v *validate.Validator) {
		v.String(new("123"), "param").Regex(regexp.MustCompile(`\d+`))
	})
	testFail(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Regex(regexp.MustCompile(`\d+`))
	}, "param has an invalid format")
	testSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Regex(regexp.MustCompile(`\d*`))
	})
}

func TestCheckString_Custom(t *testing.T) {
	r := require.New(t)
	testFail(r, func(v *validate.Validator) {
		v.String(new("a"), "param").Custom(true, func(t string) bool {
			return false
		})
	}, "param is invalid")
	testSuccess(r, func(v *validate.Validator) {
		v.String(new("a"), "param").Custom(false, func(t string) bool {
			return true
		})
	})
	testFail(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Custom(false, func(t string) bool {
			return true
		})
	}, "param is invalid")
	testSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Custom(true, func(t string) bool {
			return false
		})
	})
}

func TestCheckString_Length(t *testing.T) {
	r := require.New(t)
	testFail(r, func(v *validate.Validator) {
		v.String(new("中"), "param").Length(2)
	}, "param must have a length of 2")
	testFail(r, func(v *validate.Validator) {
		v.String(new("中文中"), "param").Length(2)
	}, "param must have a length of 2")
	testSuccess(r, func(v *validate.Validator) {
		v.String(new("中文"), "param").Length(2)
	})
	testFail(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Length(1)
	}, "param must have a length of 1")
	testSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Length(0)
	})
}

func TestCheckString_Min(t *testing.T) {
	r := require.New(t)
	testFail(r, func(v *validate.Validator) {
		v.String(new("中"), "param").Min(2)
	}, "param must have a length of at least 2")
	testSuccess(r, func(v *validate.Validator) {
		v.String(new("中文"), "param").Min(2)
	})
	testFail(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Min(1)
	}, "param must have a length of at least 1")
	testSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Min(0)
	})
}

func TestCheckString_Max(t *testing.T) {
	r := require.New(t)
	testFail(r, func(v *validate.Validator) {
		v.String(new("中文中"), "param").Max(2)
	}, "param must have a length of at most 2")
	testSuccess(r, func(v *validate.Validator) {
		v.String(new("中文"), "param").Max(2)
	})
	testFail(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Max(-1)
	}, "param must have a length of at most -1")
	testSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Max(0)
	})
}

func TestCheckString_BetweenInclusive(t *testing.T) {
	r := require.New(t)
	testFail(r, func(v *validate.Validator) {
		v.String(new("中"), "param").BetweenInclusive(2, 3)
	}, "param must have a length between 2 and 3, inclusive")
	testFail(r, func(v *validate.Validator) {
		v.String(new("中文中文"), "param").BetweenInclusive(2, 3)
	}, "param must have a length between 2 and 3, inclusive")
	testSuccess(r, func(v *validate.Validator) {
		v.String(new("中文中"), "param").BetweenInclusive(2, 3)
	})
	testFail(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").BetweenInclusive(1, 3)
	}, "param must have a length between 1 and 3, inclusive")
	testFail(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").BetweenInclusive(-3, -1)
	}, "param must have a length between -3 and -1, inclusive")
	testSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").BetweenInclusive(0, 3)
	})
	testSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").BetweenInclusive(-3, 0)
	})
}

func TestCheckString_Gt(t *testing.T) {
	r := require.New(t)
	testFail(r, func(v *validate.Validator) {
		v.String(new("中文"), "param").Gt(2)
	}, "param must have a length greater than 2")
	testSuccess(r, func(v *validate.Validator) {
		v.String(new("中文中"), "param").Gt(2)
	})
	testFail(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Gt(0)
	}, "param must have a length greater than 0")
	testSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Gt(-1)
	})
}

func TestCheckString_Lt(t *testing.T) {
	r := require.New(t)
	testFail(r, func(v *validate.Validator) {
		v.String(new("中文"), "param").Lt(2)
	}, "param must have a length less than 2")
	testSuccess(r, func(v *validate.Validator) {
		v.String(new("中"), "param").Lt(2)
	})
	testFail(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Lt(0)
	}, "param must have a length less than 0")
	testSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Lt(1)
	})
}

func TestCheckString_BetweenExclusive(t *testing.T) {
	r := require.New(t)
	testFail(r, func(v *validate.Validator) {
		v.String(new("中"), "param").BetweenExclusive(2, 5)
	}, "param must have a length between 2 and 5, exclusive")
	testFail(r, func(v *validate.Validator) {
		v.String(new("中文中文中文"), "param").BetweenExclusive(2, 5)
	}, "param must have a length between 2 and 5, exclusive")
	testSuccess(r, func(v *validate.Validator) {
		v.String(new("中文中"), "param").BetweenExclusive(2, 5)
	})
	testFail(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").BetweenExclusive(0, 5)
	}, "param must have a length between 0 and 5, exclusive")
	testFail(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").BetweenExclusive(-5, 0)
	}, "param must have a length between -5 and 0, exclusive")
	testSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").BetweenExclusive(-1, 5)
	})
	testSuccess(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").BetweenExclusive(-5, 1)
	})
}

func TestCheckString_Enum(t *testing.T) {
	r := require.New(t)
	testFail(r, func(v *validate.Validator) {
		v.String(new("d"), "param").Enum([]string{"a", "b", "c"})
	}, `param must be "a", "b", or "c"`)
	testSuccess(r, func(v *validate.Validator) {
		v.String(new("c"), "param").Enum([]string{"a", "b", "c"})
	})
	testFail(r, func(v *validate.Validator) {
		v.String((*string)(nil), "param").Enum([]string{"a", "b", "c"})
	}, `param must be "a", "b", or "c"`)
}
