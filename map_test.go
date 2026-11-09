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

func TestCheckMap_Required(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Map((map[string]int)(nil), "param").Required()
	}, "param is required")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int{}, "param").Required()
	})
}

func TestCheckMap_NotEmpty(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Map(map[string]int{}, "param").NotEmpty()
	}, "param must not be empty")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1}, "param").NotEmpty()
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map((map[string]int)(nil), "param").NotEmpty()
	})
}

func TestCheckMap_Length(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1}, "param").Length(2)
	}, "param must have a length of 2")
	checklistFail(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1, "b": 2, "c": 3}, "param").Length(2)
	}, "param must have a length of 2")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1, "b": 2}, "param").Length(2)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int(nil), "param").Length(2)
	})
}

func TestCheckMap_Min(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1}, "param").Min(2)
	}, "param must have a length of at least 2")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1, "b": 2}, "param").Min(2)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int(nil), "param").Min(1)
	})
}

func TestCheckMap_Max(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1, "b": 2, "c": 3}, "param").Max(2)
	}, "param must have a length of at most 2")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1, "b": 2}, "param").Max(2)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int(nil), "param").Max(-1)
	})
}

func TestCheckMap_Range(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1}, "param").Range(2, 3)
	}, "param must have a length between 2 and 3, inclusive")
	checklistFail(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, "param").Range(2, 3)
	}, "param must have a length between 2 and 3, inclusive")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1, "b": 2, "c": 3}, "param").Range(2, 3)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int(nil), "param").Range(1, 3)
	})
}

func TestCheckMap_Gt(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1, "b": 2}, "param").Gt(2)
	}, "param must have a length greater than 2")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1, "b": 2, "c": 3}, "param").Gt(2)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int(nil), "param").Gt(0)
	})
}

func TestCheckMap_Lt(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1, "b": 2, "c": 3}, "param").Lt(2)
	}, "param must have a length less than 2")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1}, "param").Lt(2)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int(nil), "param").Lt(0)
	})
}

func TestCheckMap_RangeExclusive(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1}, "param").RangeExclusive(2, 5)
	}, "param must have a length between 2 and 5, exclusive")
	checklistFail(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6}, "param").RangeExclusive(2, 5)
	}, "param must have a length between 2 and 5, exclusive")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int{"a": 1, "b": 2, "c": 3}, "param").RangeExclusive(2, 5)
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int(nil), "param").RangeExclusive(0, 5)
	})
}

func TestCheckMap_Custom(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Map(map[string]int{}, "param").Custom(func(t map[string]int) bool {
			return false
		})
	}, "param is invalid")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int{}, "param").Custom(func(t map[string]int) bool {
			return true
		})
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Map(map[string]int(nil), "param").Custom(func(t map[string]int) bool {
			return true
		})
	})
}

func TestCheckMap_Stop(t *testing.T) {
	r := require.New(t)
	{
		v := validate.NewValidator()
		validate.SetAll(v)
		v.Map(map[string]int{"aaa": 5, "bbb": 5, "a": 11, "b": 11}, "param").Dive(func(k string) {
			v.String(&k, "").Lt(3)
		}, func(value int) {
			v.Int(&value, "").Lt(10)
		})
		_, _, messages := validate.GetResult(v)
		r.Len(messages, 4)
	}
	{
		v := validate.NewValidator()
		v.Map(map[string]int{"aaa": 5, "bbb": 5}, "param").Dive(func(k string) {
			v.String(&k, "").Lt(3)
		}, func(value int) {
			v.Int(&value, "").Lt(10)
		})
		_, _, messages := validate.GetResult(v)
		r.Len(messages, 1)
	}
	{
		v := validate.NewValidator()
		v.Map(map[string]int{"a": 11, "b": 11}, "param").Dive(func(k string) {
			v.String(&k, "").Lt(3)
		}, func(value int) {
			v.Int(&value, "").Lt(10)
		})
		_, _, messages := validate.GetResult(v)
		r.Len(messages, 1)
	}
	{
		v := validate.NewValidator()
		validate.SetHasError(v)
		v.Map(map[string]int{"a": 11, "b": 11}, "param").Dive(func(k string) {
			v.String(&k, "").Lt(3)
		}, func(value int) {
			v.Int(&value, "").Lt(10)
		})
		_, _, messages := validate.GetResult(v)
		r.Nil(messages)
	}
}
