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

type messageBuilder func(f *FieldInfo) string

func messageDefault(f *FieldInfo) string {
	return fmt.Sprintf("%s is invalid", f.fieldName)
}

func messageRequired(f *FieldInfo) string {
	return fmt.Sprintf("%s is required", f.fieldName)
}

func messageMin(f *FieldInfo) string {
	return fmt.Sprintf("%s must be at least %s", f.fieldName, f.Arg(0))
}

func messageMax(f *FieldInfo) string {
	return fmt.Sprintf("%s must be at most %s", f.fieldName, f.Arg(0))
}

func messageBetweenInclusive(f *FieldInfo) string {
	return fmt.Sprintf("%s must be between %s and %s, inclusive", f.fieldName, f.Arg(0), f.Arg(1))
}

func messageGt(f *FieldInfo) string {
	return fmt.Sprintf("%s must be greater than %s", f.fieldName, f.Arg(0))
}

func messageLt(f *FieldInfo) string {
	return fmt.Sprintf("%s must be less than %s", f.fieldName, f.Arg(0))
}

func messageBetweenExclusive(f *FieldInfo) string {
	return fmt.Sprintf("%s must be between %s and %s, exclusive", f.fieldName, f.Arg(0), f.Arg(1))
}

func messageEnum(f *FieldInfo) string {
	return fmt.Sprintf("%s must be %s", f.fieldName, f.Args())
}

func messageNotBlank(f *FieldInfo) string {
	return fmt.Sprintf("%s must not be blank", f.fieldName)
}

func messageRegex(f *FieldInfo) string {
	return fmt.Sprintf("%s has an invalid format", f.fieldName)
}

func messageNotEmpty(f *FieldInfo) string {
	return fmt.Sprintf("%s must not be empty", f.fieldName)
}

func messageLength(f *FieldInfo) string {
	return fmt.Sprintf("%s must have a length of %s", f.fieldName, f.Arg(0))
}

func messageLengthMin(f *FieldInfo) string {
	return fmt.Sprintf("%s must have a length of at least %s", f.fieldName, f.Arg(0))
}

func messageLengthMax(f *FieldInfo) string {
	return fmt.Sprintf("%s must have a length of at most %s", f.fieldName, f.Arg(0))
}

func messageLengthBetweenInclusive(f *FieldInfo) string {
	return fmt.Sprintf("%s must have a length between %s and %s, inclusive", f.fieldName, f.Arg(0), f.Arg(1))
}

func messageLengthGt(f *FieldInfo) string {
	return fmt.Sprintf("%s must have a length greater than %s", f.fieldName, f.Arg(0))
}

func messageLengthLt(f *FieldInfo) string {
	return fmt.Sprintf("%s must have a length less than %s", f.fieldName, f.Arg(0))
}

func messageLengthBetweenExclusive(f *FieldInfo) string {
	return fmt.Sprintf("%s must have a length between %s and %s, exclusive", f.fieldName, f.Arg(0), f.Arg(1))
}
