# validate-struct

[English](#english) | [中文](#中文)

<a id="english"></a>

validate-struct is a Go library for defining struct validation rules in code. It provides a small set of general-purpose rules, with custom functions for application-specific checks. Validation logic can live separately from model definitions, including structs generated from protobuf.

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jishaocong0910/validate-struct)
[![Go Reference](https://pkg.go.dev/badge/github.com/jishaocong0910/validate-struct.svg)](https://pkg.go.dev/github.com/jishaocong0910/validate-struct)
![coverage](https://raw.githubusercontent.com/jishaocong0910/validate-struct/badges/.badges/main/coverage.svg)

## Installation

```shell
go get github.com/jishaocong0910/validate-struct
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/jishaocong0910/validate-struct"
)

type Book struct {
	Title       string
	Isbn        string
	Description *string
	Stock       int
	Price       float64
	Language    string
	Author      *Author
	Categories  []*Category
}

func (b *Book) Checklist(v *validate.Validator) {
	v.String(&b.Title, "title").NotBlank().Max(10)
	v.String(&b.Isbn, "isbn").NotBlank().Regex(regexp.MustCompile(`^[0-9]{13}$`))
	v.String(b.Description, "description").Required()
	v.Int(&b.Stock, "stock").Range(0, 100)
	v.Float64(&b.Price, "price").Gt(0)
	v.String(&b.Language, "language").Enum([]string{"zh-cn", "en-US", "ja-JP"})
	v.Struct(b.Author, "author").Required().Dive()
	v.Slice(b.Categories, "categories").NotEmpty().Dive(func(e *Category) {
		v.Struct(e, "").Dive()
	})
}

type Author struct {
	Name    string
	Resumes *string
}

func (a Author) Checklist(v *validate.Validator) {
	v.String(&a.Name, "name").NotBlank()
	v.String(a.Resumes, "resumes").Required().NotBlank()
}

type Category struct {
	Id   string
	Sort int
}

func (c Category) Checklist(v *validate.Validator) {
	v.String(&c.Id, "id").NotBlank()
	v.Int(&c.Sort, "sort").Max(999)
}

func main() {
	b := Book{Author: &Author{}, Categories: []*Category{{Sort: 127}, {Id: "c1", Sort: 1000}}}
	code, _, messages := validate.Check(context.Background(), &b, validate.AllRules())
	if code != validate.SUCCESS {
		for i, msg := range messages {
			fmt.Println(i, msg)
		}
	}
	// Output:
	// 0 title must not be blank
	// 1 isbn must not be blank
	// 2 isbn has an invalid format
	// 3 description is required
	// 4 price must be greater than 0
	// 5 language must be "zh-cn", "en-US", or "ja-JP"
	// 6 author.name must not be blank
	// 7 author.resumes is required
	// 8 categories[0].id must not be blank
	// 9 categories[1].sort must be at most 999
}
```

## Validation Function

Call `validate.Check` to validate a struct.

Parameters:

| Parameter | Type | Description |
| --- | --- | --- |
| `ctx` | `context.Context` | The context available to validation rules. A `nil` context defaults to `context.Background()`. |
| `target` | `V` | A pointer to the struct to validate. It must implement `validate.Validatable`. |
| `opts` | `...ValidatorOption` | Options that configure validation behavior. |

Return values:

| Return value | Type | Description |
| --- | --- | --- |
| `code` | `string` | `validate.SUCCESS` if validation succeeds; otherwise, `validate.FAIL` unless a custom error code is set. |
| `message` | `string` | The first error message. |
| `messages` | `[]string` | The list of error messages. |

By default, `Check` stops validation at the first failed rule and returns at most one error message. Pass `validate.AllRules()` to check every rule and collect all error messages.

## Implementing Validatable

To make a struct validatable, implement `validate.Validatable` and define its rules in `Checklist`.

For generated types, you can add the method in a separate file in the same package. Alternatively, embed the generated type in a wrapper struct in your own package and implement `Checklist` on the wrapper. Both approaches keep validation logic outside generated files.

```go
type Validatable interface {
    Checklist(v *Validator)
}
```

## Validator

Use `validate.Validator` to select fields and apply rules for their types.

| Method | Description |
| --- | --- |
| `Ctx` | Returns the context passed to `Check`. |
| `Bool`, `Byte`, `Int`, `Int8`, `Int16`, `Int32`, `Int64`, `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`, `Float32`, `Float64`, `String`, `Struct`, `Slice`, `Map`, `Any` | Select a field and return a checker with rules for that type. Pass a pointer for all types except slices and maps. |
| `Embed` | Validates an embedded struct that implements `validate.Validatable`. |

## Field Names

Each field takes a name to use in error messages. Within a slice or map `Dive` callback, the validator supplies the name automatically and ignores the name passed to the field method:

- Slice elements use `fieldName[index]`.
- Map keys use `fieldName$key`.
- Map values use `fieldName$value`.

```go
package main

import (
	"context"
	"fmt"

	"github.com/jishaocong0910/validate-struct"
)

type Product struct {
	Name       string
	Tags       []int
	Attributes map[string]int
}

func (p Product) Checklist(v *validate.Validator) {
	v.String(&p.Name, "name").NotBlank()
	v.Slice(p.Tags, "tags").Dive(func(t int) {
		v.Int(&t, "useless").Gt(5)
	})
	v.Map(p.Attributes, "attributes").Dive(func(key string) {
		v.String(&key, "useless").NotBlank()
	}, func(value int) {
		v.Int(&value, "useless").Gt(5)
	})
}

func main() {
	d := &Product{
		Tags:       []int{4},
		Attributes: map[string]int{"": 4},
	}
	code, _, messages := validate.Check(context.Background(), d, validate.AllRules())
	if code != validate.SUCCESS {
		for i, msg := range messages {
			fmt.Println(i, msg)
		}
	}
	// Output:
	// 0 name must not be blank
	// 1 tags[0] must be greater than 5
	// 2 attributes$key must not be blank
	// 3 attributes$value must be greater than 5
}
```

## Validation Rules

Chain rule methods to apply multiple checks to a field. All rules except `Required` skip `nil` values. `Dive` returns no value and must be the last call in a chain.

| Method | Description |
| --- | --- |
| `Required` | Requires a non-nil value. |
| `NotBlank` | Rejects empty strings and strings containing only whitespace. |
| `Regex` | Checks that a string matches a regular expression. |
| `NotEmpty` | Requires a length greater than zero. |
| `Length` | Requires an exact length. |
| `Min` | Sets an inclusive lower bound on a value or length. |
| `Max` | Sets an inclusive upper bound on a value or length. |
| `Range` | Checks that a value or length falls within a range, including both bounds. |
| `Gt` | Requires a value or length greater than the given bound. |
| `Lt` | Requires a value or length less than the given bound. |
| `RangeExclusive` | Checks that a value or length falls within a range, excluding both bounds. |
| `Enum` | Checks that a value is one of the allowed values. |
| `Custom` | Applies a custom validation function. |
| `Dive` | Validates nested struct fields, slice elements, or map keys and values. |

## Error Codes

Pass `validate.Code(...)` to a rule to specify the code returned when that rule fails. Custom error codes apply only in the default mode; they are ignored when `validate.AllRules()` is enabled.

## Error Messages

Pass `validate.Message(...)` to a rule to override its default error message. The callback, `func(f *validate.FieldInfo) string`, builds the message from the field name and rule arguments available through `FieldInfo`.

| Method | Description |
| --- | --- |
| `FieldName` | Returns the full field path. Examples: `title`, `author.name`, and `categoryId[2].sort`. |
| `Arg` | Returns the rule argument at the given index as a string. For `Max(10)`, `Arg(0)` returns `10`. For `Range(5, 15)`, `Arg(0)` returns `5` and `Arg(1)` returns `15`. |
| `Args` | Formats the rule arguments as a list, using ` or ` between two values and `, or ` before the last of three or more values. String enum values are enclosed in double quotes. For `Enum([]string{"zh-cn", "en-US"})`, the result is `"zh-cn" or "en-US"`. For `Enum([]string{"zh-cn", "en-US", "ja-JP"})`, it is `"zh-cn", "en-US", or "ja-JP"`. |

For example:

```go
v.String(&b.Title, "title").Max(10, validate.Message(func(f *validate.FieldInfo) string {
	return fmt.Sprintf("%s exceeds max length %s", f.FieldName(), f.Arg(0))
}))
```

---

## 中文

[English](#english) | [中文](#中文)

用于Go的编程式结构体校验工具，提供基础、通用的校验规则，并通过自定义函数扩展业务校验逻辑。校验规则使用Go代码编写，可与模型定义分开维护，适合protobuf等工具生成的结构体。

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jishaocong0910/validate-struct)
[![Go Reference](https://pkg.go.dev/badge/github.com/jishaocong0910/validate-struct.svg)](https://pkg.go.dev/github.com/jishaocong0910/validate-struct)
![coverage](https://raw.githubusercontent.com/jishaocong0910/validate-struct/badges/.badges/main/coverage.svg)

## 安装

```shell
go get github.com/jishaocong0910/validate-struct
```

## 快速开始

```go
package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/jishaocong0910/validate-struct"
)

type Book struct {
	Title       string
	Isbn        string
	Description *string
	Stock       int
	Price       float64
	Language    string
	Author      *Author
	Categories  []*Category
}

func (b *Book) Checklist(v *validate.Validator) {
	v.String(&b.Title, "title").NotBlank().Max(10)
	v.String(&b.Isbn, "isbn").NotBlank().Regex(regexp.MustCompile(`^[0-9]{13}$`))
	v.String(b.Description, "description").Required()
	v.Int(&b.Stock, "stock").Range(0, 100)
	v.Float64(&b.Price, "price").Gt(0)
	v.String(&b.Language, "language").Enum([]string{"zh-cn", "en-US", "ja-JP"})
	v.Struct(b.Author, "author").Required().Dive()
	v.Slice(b.Categories, "categories").NotEmpty().Dive(func(e *Category) {
		v.Struct(e, "").Dive()
	})
}

type Author struct {
	Name    string
	Resumes *string
}

func (a Author) Checklist(v *validate.Validator) {
	v.String(&a.Name, "name").NotBlank()
	v.String(a.Resumes, "resumes").Required().NotBlank()
}

type Category struct {
	Id   string
	Sort int
}

func (c Category) Checklist(v *validate.Validator) {
	v.String(&c.Id, "id").NotBlank()
	v.Int(&c.Sort, "sort").Max(999)
}

func main() {
	b := Book{Author: &Author{}, Categories: []*Category{{Sort: 127}, {Id: "c1", Sort: 1000}}}
	code, _, messages := validate.Check(context.Background(), &b, validate.AllRules())
	if code != validate.SUCCESS {
		for i, msg := range messages {
			fmt.Println(i, msg)
		}
	}
	// Output:
	// 0 title must not be blank
	// 1 isbn must not be blank
	// 2 isbn has an invalid format
	// 3 description is required
	// 4 price must be greater than 0
	// 5 language must be "zh-cn", "en-US", or "ja-JP"
	// 6 author.name must not be blank
	// 7 author.resumes is required
	// 8 categories[0].id must not be blank
	// 9 categories[1].sort must be at most 999
}
```

## 校验函数

调用`validate.Check`校验结构体。

参数：

| 参数     | 类型                 | 说明                                                           |
|----------|----------------------|----------------------------------------------------------------|
| `ctx`    | `context.Context`    | 传递给校验器的上下文；为 `nil` 时使用 `context.Background()`。 |
| `target` | `V`                  | 待校验结构体的指针，结构体需实现 `validate.Validatable` 接口。 |
| `opts`   | `...ValidatorOption` | 校验器选项。                                                   |

返回值：

| 返回值     | 类型       | 说明                                                        |
|------------|------------|-------------------------------------------------------------|
| `code`     | `string`   | 成功时为 `validate.SUCCESS`，失败时默认为 `validate.FAIL`。 |
| `message`  | `string`   | 首个错误消息。                                              |
| `messages` | `[]string` | 错误消息列表。                                              |

默认情况下，校验到首个错误时会停止校验，`messages`最多只有1个元素，使用`validate.AllRules`选项可校验所有规则，`messages`将返回所有错误消息。

## 实现Validatable

待校验结构体需实现`validate.Validatable`接口，在`Checklist`方法中编写字段的校验规则。对于会被代码生成器覆盖的结构体，可以在同包的独立文件中实现`Checklist`，也可以在包外定义内嵌原模型的结构体并实现该方法，使校验规则独立于模型维护。

```go
type Validatable interface {
    Checklist(v *Validator)
}
```

## 校验器

`validate.Validator`提供了各种类型的校验入口。

| 方法                                                                                                                                                                     | 描述                                                                           |
|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------|
| `Ctx`                                                                                                                                                                    | 获取入口函数透传的上下文。                                                     |
| `Bool`、`Byte`、`Int`、`Int8`、`Int16`、`Int32`、`Int64`、`Uint`、`Uint8`、`Uint16`、`Uint32`、`Uint64`、`Float32`、`Float64`、`String`、`Struct`、`Slice`、`Map`、`Any` | 用于绑定校验字段，并获取对应类型的规则方法。除了`Slice`和`Map`都需要传入指针。 |
| `Embed`                                                                                                                                                                  | 用于校验实现了`validate.Validatable`接口的内嵌结构体。                         |

## 字段名称

绑定校验字段须指定字段名称，用于在错误消息中显示。而slice和map的`Dive`方法内校验元素值时，会预设字段名称并且无法更改，slice元素为`fieldName[index]`，对于map，key为`fieldName$key`，value为`fieldName$value`。

```go
package main

import (
	"context"
	"fmt"

	"github.com/jishaocong0910/validate-struct"
)

type Product struct {
	Name       string
	Tags       []int
	Attributes map[string]int
}

func (p Product) Checklist(v *validate.Validator) {
	v.String(&p.Name, "name").NotBlank()
	v.Slice(p.Tags, "tags").Dive(func(t int) {
		v.Int(&t, "useless").Gt(5)
	})
	v.Map(p.Attributes, "attributes").Dive(func(key string) {
		v.String(&key, "useless").NotBlank()
	}, func(value int) {
		v.Int(&value, "useless").Gt(5)
	})
}

func main() {
	d := &Product{
		Tags:       []int{4},
		Attributes: map[string]int{"": 4},
	}
	code, _, messages := validate.Check(context.Background(), d, validate.AllRules())
	if code != validate.SUCCESS {
		for i, msg := range messages {
			fmt.Println(i, msg)
		}
	}
	// Output:
	// 0 name must not be blank
	// 1 tags[0] must be greater than 5
	// 2 attributes$key must not be blank
	// 3 attributes$value must be greater than 5
}
```

## 规则方法

提供了以下规则方法，其中除了`Required`的规则都会忽略`nil`值，可链式调用指定多个规则。`Dive`没有返回值，调用后不能继续追加规则。

| 方法               | 描述                                                     |
|--------------------|----------------------------------------------------------|
| `Required`         | 必填，即不能为nil。                                      |
| `NotBlank`         | 字符串不能为空白。                                       |
| `Regex`            | 字符串必须匹配正则表达式。                               |
| `NotEmpty`         | 长度必须大于0。                                          |
| `Length`           | 长度必须等于指定值。                                     |
| `Min`              | 数值或长度的最小值。                                     |
| `Max`              | 数值或长度的最大值。                                     |
| `Range` | 数值或长度必须在指定范围内，包含边界。                   |
| `Gt`               | 数值或长度必须大于指定值。                               |
| `Lt`               | 数值或长度必须小于指定值。                               |
| `RangeExclusive` | 数值或长度必须在指定范围内，不包含边界。                 |
| `Enum`             | 必须是枚举选项中的一个。                                 |
| `Custom`           | 自定义校验。                                             |
| `Dive`             | 递归校验struct字段、slice字段元素、map字段的key和value。 |

## 错误码

规则方法使用`validate.Code`选项可自定义错误码。仅在不使用`validate.AllRules`选项的情况下有效。

## 错误消息

规则方法使用`validate.Message`选项可自定义错误消息。通过传入`func(f *validate.FieldInfo) string`函数构建错误消息，`validate.FieldInfo`提供了一些方法返回动态参数。

| 方法        | 描述                                                                                                                                                                                                                                                                                                         |
|-------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `FieldName` | 返回具有层级关系的字段名称。例如：title、author.name、categoryId[2].sort。                                                                                                                                                                                                                                   |
| `Arg`       | 返回规则方法的指定索引的限制值的字符串形式。例如：对于`Max(10)`，`Arg(0)`返回`10`；对于`Range(5, 15)`，`Arg(0)`返回`5`，`Arg(1)`返回`15`。                                                                                                                                                        |
| `Args`      | 返回规则方法的所有限制值的字符串形式。两个值用` or `连接；超过两个值时，前面的值用`, `连接，最后一个值前用`, or `连接。字符串枚举值包含双引号。例如：对于`Enum([]string{"zh-cn", "en-US"})`，返回`"zh-cn" or "en-US"`；对于`Enum([]string{"zh-cn", "en-US", "ja-JP"})`，返回`"zh-cn", "en-US", or "ja-JP"`。 |

*e.g.*

```go
v.String(&b.Title, "title").Max(10, validate.Message(func(f *validate.FieldInfo) string {
	return fmt.Sprintf("%s exceeds max length %s", f.FieldName(), f.Arg(0))
}))
```