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

func TestCheckBool_Required(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Bool((*bool)(nil), "param").Required()
	}, "param is required")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Bool(new(false), "param").Required()
	})
}

func TestCheckBool_Custom(t *testing.T) {
	r := require.New(t)
	checklistFail(r, func(v *validate.Validator) {
		v.Bool(new(true), "param").Custom(func(t bool) bool {
			return false
		})
	}, "param is invalid")
	checklistSuccess(r, func(v *validate.Validator) {
		v.Bool(new(true), "param").Custom(func(t bool) bool {
			return true
		})
	})
	checklistSuccess(r, func(v *validate.Validator) {
		v.Bool(nil, "param").Custom(func(t bool) bool {
			return false
		})
	})
}
