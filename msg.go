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

type msgBuilder func(f *FieldInfo)

func msgBuilderDefault(f *FieldInfo) {
	f.Msg("%s is illegal", f.fieldName)
}

func msgBuilderRequired(f *FieldInfo) {
	f.Msg("%s is required", f.fieldName)
}

func msgBuilderMin(f *FieldInfo) {
	f.Msg("%s must not be less than %s", f.fieldName, f.Confine(0))
}

func msgBuilderMax(f *FieldInfo) {
	f.Msg("%s must not be greater than %s", f.fieldName, f.Confine(0))
}

func msgBuilderRange(f *FieldInfo) {
	f.Msg("%s must be %s to %s", f.fieldName, f.Confine(0), f.Confine(1))
}

func msgBuilderGt(f *FieldInfo) {
	f.Msg("%s must be greater than %s", f.fieldName, f.Confine(0))
}

func msgBuilderLt(f *FieldInfo) {
	f.Msg("%s must be less than %s", f.fieldName, f.Confine(0))
}

func msgBuilderWithin(f *FieldInfo) {
	f.Msg("%s must be greater than %s and less than %s", f.fieldName, f.Confine(0), f.Confine(1))
}

func msgBuilderEnum(f *FieldInfo) {
	f.Msg("%s must be %s", f.fieldName, f.Confines())
}

func msgBuilderNotBlank(f *FieldInfo) {
	f.Msg("%s must not be blank", f.fieldName)
}

func msgBuilderRegex(f *FieldInfo) {
	f.Msg("%s's format is illegal", f.fieldName)
}

func msgBuilderNotEmpty(f *FieldInfo) {
	f.Msg("%s must not be empty", f.fieldName)
}

func msgBuilderLength(f *FieldInfo) {
	f.Msg("%s's length must be %s", f.fieldName, f.Confine(0))
}

func msgBuilderLengthMin(f *FieldInfo) {
	f.Msg("%s's length must not be less than %s", f.fieldName, f.Confine(0))
}

func msgBuilderLengthMax(f *FieldInfo) {
	f.Msg("%s's length must not be greater than %s", f.fieldName, f.Confine(0))
}

func msgBuilderLengthRange(f *FieldInfo) {
	f.Msg("%s's length must be %s to %s", f.fieldName, f.Confine(0), f.Confine(1))
}

func msgBuilderLengthGt(f *FieldInfo) {
	f.Msg("%s's length must be greater than %s", f.fieldName, f.Confine(0))
}

func msgBuilderLengthLt(f *FieldInfo) {
	f.Msg("%s's length must be less than %s", f.fieldName, f.Confine(0))
}

func msgBuilderLengthWithin(f *FieldInfo) {
	f.Msg("%s's length must be greater than %s and less than %s", f.fieldName, f.Confine(0), f.Confine(1))
}
