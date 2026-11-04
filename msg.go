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

import "fmt"

type msgBuilder func(f *FieldInfo) string

func msgBuilderDefault(f *FieldInfo) string {
	return fmt.Sprintf("%s is illegal", f.fieldName)
}

func msgBuilderRequired(f *FieldInfo) string {
	return fmt.Sprintf("%s is required", f.fieldName)
}

func msgBuilderMin(f *FieldInfo) string {
	return fmt.Sprintf("%s must not be less than %s", f.fieldName, f.Arg(0))
}

func msgBuilderMax(f *FieldInfo) string {
	return fmt.Sprintf("%s must not be greater than %s", f.fieldName, f.Arg(0))
}

func msgBuilderRange(f *FieldInfo) string {
	return fmt.Sprintf("%s must be %s to %s", f.fieldName, f.Arg(0), f.Arg(1))
}

func msgBuilderGt(f *FieldInfo) string {
	return fmt.Sprintf("%s must be greater than %s", f.fieldName, f.Arg(0))
}

func msgBuilderLt(f *FieldInfo) string {
	return fmt.Sprintf("%s must be less than %s", f.fieldName, f.Arg(0))
}

func msgBuilderWithin(f *FieldInfo) string {
	return fmt.Sprintf("%s must be greater than %s and less than %s", f.fieldName, f.Arg(0), f.Arg(1))
}

func msgBuilderEnum(f *FieldInfo) string {
	return fmt.Sprintf("%s must be %s", f.fieldName, f.Args())
}

func msgBuilderNotBlank(f *FieldInfo) string {
	return fmt.Sprintf("%s must not be blank", f.fieldName)
}

func msgBuilderRegex(f *FieldInfo) string {
	return fmt.Sprintf("%s's format is illegal", f.fieldName)
}

func msgBuilderNotEmpty(f *FieldInfo) string {
	return fmt.Sprintf("%s must not be empty", f.fieldName)
}

func msgBuilderLength(f *FieldInfo) string {
	return fmt.Sprintf("%s's length must be %s", f.fieldName, f.Arg(0))
}

func msgBuilderLengthMin(f *FieldInfo) string {
	return fmt.Sprintf("%s's length must not be less than %s", f.fieldName, f.Arg(0))
}

func msgBuilderLengthMax(f *FieldInfo) string {
	return fmt.Sprintf("%s's length must not be greater than %s", f.fieldName, f.Arg(0))
}

func msgBuilderLengthRange(f *FieldInfo) string {
	return fmt.Sprintf("%s's length must be %s to %s", f.fieldName, f.Arg(0), f.Arg(1))
}

func msgBuilderLengthGt(f *FieldInfo) string {
	return fmt.Sprintf("%s's length must be greater than %s", f.fieldName, f.Arg(0))
}

func msgBuilderLengthLt(f *FieldInfo) string {
	return fmt.Sprintf("%s's length must be less than %s", f.fieldName, f.Arg(0))
}

func msgBuilderLengthWithin(f *FieldInfo) string {
	return fmt.Sprintf("%s's length must be greater than %s and less than %s", f.fieldName, f.Arg(0), f.Arg(1))
}
