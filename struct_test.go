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

func TestCheckStruct_Required(t *testing.T) {
	r := require.New(t)
	testFail(r, func(v *validate.Validator) {
		v.Struct((*validate.DemoUser)(nil), "param").Required()
	}, "param is required")
	testSuccess(r, func(v *validate.Validator) {
		v.Struct(&validate.DemoUser{}, "param").Required()
	})
}

func TestCheckStruct_Custom(t *testing.T) {
	r := require.New(t)
	testFail(r, func(v *validate.Validator) {
		v.Struct(&validate.DemoUser{}, "param").Custom(true, func(t validate.DemoUser) bool {
			return false
		})
	}, "param is invalid")
	testSuccess(r, func(v *validate.Validator) {
		v.Struct(&validate.DemoUser{}, "param").Custom(false, func(t validate.DemoUser) bool {
			return true
		})
	})
	testFail(r, func(v *validate.Validator) {
		v.Struct((*validate.DemoUser)(nil), "param").Custom(false, func(t validate.DemoUser) bool {
			return true
		})
	}, "param is invalid")
	testSuccess(r, func(v *validate.Validator) {
		v.Struct((*validate.DemoUser)(nil), "param").Custom(true, func(t validate.DemoUser) bool {
			return false
		})
	})
}

func TestCheckStruct_Dive(t *testing.T) {
	r := require.New(t)
	testSuccess(r, func(v *validate.Validator) {
		v.Struct(&validate.DemoBase{Id: "id"}, "param").Dive()
	})
	testFail(r, func(v *validate.Validator) {
		v.Struct(&validate.DemoBase{}, "param").Dive()
	}, "param.id must not be blank")
	testFail(r, func(v *validate.Validator) {
		v.Struct((*validate.DemoBase)(nil), "param").Dive()
	}, "param.id is required")
	testFail(r, func(v *validate.Validator) {
		v.Struct((*validate.DemoCommunication)(nil), "param").Dive()
	}, "param.address.city is required")
}

func TestCheckStruct_ShouldStop(t *testing.T) {
	r := require.New(t)
	testFail(r, func(v *validate.Validator) {
		v.Struct(&validate.DemoUser{Id: "id"}, "param").Dive()
	}, "param.name must not be blank")
}
