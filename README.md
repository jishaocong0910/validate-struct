# Gverify

一款用于Golang的结构体校验工具。它不通过标签来指定校验规则，而是由代码指定。使用标签的校验工具需要修改结构体，在一些代码生成器生成的结构体上使用时，可能导致代码覆盖或冲突问题。Gverify可实现**零代码入侵**，支持自定义错误码和错误消息，不使用反射，处理速度理论上比使用标签的校验工具快。

[![Go Reference](https://pkg.go.dev/badge/github.com/jishaocong0910/gverify.svg)](https://pkg.go.dev/github.com/jishaocong0910/gverify)
[![Go Report Card](https://goreportcard.com/badge/github.com/jishaocong0910/gverify)](https://goreportcard.com/report/github.com/jishaocong0910/gverify)
![coverage](https://raw.githubusercontent.com/jishaocong0910/gverify/badges/.badges/main/coverage.svg)

# 安装

```shell
go get github.com/jishaocong0910/gverify
```

# 用法与例子

```go
package main

import (
    "context"
    "fmt"
    "regexp"

    validate "github.com/jishaocong0910/validate-struct"
)

type Book struct {
    Title       string
    Isbn        string
    Description *string
    Stock       int
    Price       float64
    Language    *string
    Author      *Author
    Categories  []*Category
}

func (b *Book) Checklist(vc *validate.VContext) {
    v.String( &b.Title, "title").NotBlank().Max(10)
    v.String( &b.Isbn, "isbn").NotBlank().Regex(regexp.MustCompile(`^[0-9]{13}$`))
    v.String( b.Description, "description").Required()
    v.Int( &b.Stock, "stock").Range(0, 100)
    v.Float64( &b.Price, "price").Gt(0)
    v.String( b.Language, "language").Enum([]string{"zh-cn", "en-US", "ja-JP"})
    v.Struct( b.Author, "author").Required().Dive()
    v.Slice( b.Categories, "categories").NotEmpty().Dive(func(e *Category) {
        v.Struct( e, "").Dive()
    })
}

type Author struct {
    Name    string
    Resumes *string
}

func (a Author) Checklist(vc *validate.VContext) {
    v.String( &a.Name, "name").NotBlank()
    v.String( &a.Name, "resumes").NotBlank()
}

type Category struct {
    Id   string
    Sort int
}

func (c Category) Checklist(vc *validate.VContext) {
    v.String( &c.Id, "id").NotBlank()
    v.Int( &c.Sort, "sort").Max(999)
}

func main() {
    b := Book{Author: &Author{}, Categories: []*Category{{Sort: 127}, {Id: "c1", Sort: 1000}}}
    code, _, msgs := validate.Check(context.Background(), &b, validate.All())
    if code != validate.SUCCESS {
        for i, msg := range msgs {
            fmt.Println(i, msg)
        }
    }
    // Output:
    // 0 title must not be blank
    // 1 isbn must not be blank
    // 2 isbn's format is illegal
    // 3 description is required
    // 4 price must be greater than 0
    // 5 language must be zh-cn, en-US or ja-JP
    // 6 author.name must not be blank
    // 7 author.resumes must not be blank
    // 8 categories[0].id must not be blank
    // 9 categories[1].sort must not be greater than 999
}
```

# 校验入口

***校验函数***：`validate.Check`，<u>只能校验非nil结构体</u>。

<table>
    <th>入参</th>
    <th>描述</th>
    <tr>
        <td width=170px>ctx context.Context</td>
        <td>Context</td>
    </tr>
    <tr>
        <td>s validate.Verifiable</td>
        <td>待校验的结构体</td>
    </tr>
    <tr>
        <td>opts ...validate.Option</td>
        <td>选项</td>
    </tr>
    
</table>

<table>
    <th>返回</th>
    <th>描述</th>
    <tr>
        <td width=160px>code string</td>
        <td>错误码，"SUCCESS"为校验成功，"FAIL"为校验失败，可使用常量<i>validate.SUCCESS</i>和<i>validate.FAIL</i>进行比较。</td>
    </tr>
    <tr>
        <td>msg string</td>
        <td>首个错误消息</td>
    </tr>
    <tr>
        <td>msgs []string</td>
        <td>所有错误消息</td>
    </tr>
</table>

| 选项      | 描述                              |
|---------|---------------------------------|
| validate.All | 校验所有校验规则，不使用此选项时，首个校验错误出现既停止校验。 |

# Checklist方法

结构体通过实现*validate.Verifiable*接口的*Checklist*方法，在其中编写字段的校验规则。校验字段时，根据字段类型，调用对应的***字段类型函数***，传入<i>Checklist</i>方法的<i>vc</i>参数、字段值的指针和字段名称，再链式调用***规则方法***，可指定一些选项。

*e.g.*

```go
func (b *Book) Checklist(vc *validate.VContext) {
    v.String( &b.Title, "title", validate.Omitempty()).NotBlank().Max(10, validate.Code("TITLE_TOO_LONG"))
}
```

| 字段类型函数      | 对应类型           |
|-------------|----------------|
| validate.Bool    | bool           |
| validate.Byte    | byte           |
| validate.Int     | int            |
| validate.Int8    | int8           |
| validate.Int16   | int16          |
| validate.Int32   | int32          |
| validate.Int64   | int64          |
| validate.Uint    | uint           |
| validate.Uint8   | uint8          |
| validate.Uint16  | uint16         |
| validate.Uint32  | uint32         |
| validate.Uint64  | uint64         |
| validate.Float32 | float32        |
| validate.Float64 | float64        |
| validate.String  | string         |
| validate.Slice   | slice          |
| validate.Map     | map            |
| validate.Struct  | struct         |
| validate.Embed   | 内嵌struct，且非nil |
| validate.Any     | any            |

<table>
    <tr>
        <th>规则方法</th>
        <th>描述</th>
    </tr>
    <tr>
        <td>Required</td>
        <td>必填，即不能为nil。</td>
    </tr>
    <tr>
        <td>NotBlank</td>
        <td>string不能为空白。nil值处理：视为空字符串。</td>
    </tr>
    <tr>
        <td>Regex</td>
        <td>string必须匹配正则表达式。nil值处理：视为空字符串。</td>
    </tr>
    <tr>
        <td>NotEmpty</td>
        <td>slice和map长度必须大于0。nil值处理：视为0长度。</td>
    </tr>
    <tr>
        <td>Length</td>
        <td>string字符长度、slice和map长度的固定值。nil值处理：视为0长度。</td>
    </tr>
    <tr>
        <td>Min</td>
        <td>数值、string字符长度、slice和map长度的最小值。nil值处理：数值视为无穷小，长度视为0。</td>
    </tr>
    <tr>
        <td>Max</td>
        <td>数值、string字符长度、slice和map长度的最大值。nil值处理：数值视为无穷小，长度视为0。</td>
    </tr>
    <tr>
        <td>Range</td>
        <td>数值、string字符长度、slice和map长度的范围值，包含边界。nil值处理：数值视为无穷小，长度视为0。</td>
    </tr>
    <tr>
        <td>Gt</td>
        <td>数值、string字符长度、slice和map长度，必须大于指定值。nil值处理：数值视为无穷小，长度视为0。</td>
    </tr>
    <tr>
        <td>Lt</td>
        <td>数值、string字符长度、slice和map长度，必须小于指定值。nil值处理：数值视为无穷小，长度视为0。</td>
    </tr>
    <tr>
        <td>Within</td>
        <td>数值、string字符长度、slice和map长度的范围值，不包含边界。nil值处理：数值视为无穷小，长度视为0。</td>
    </tr>
    <tr>
        <td>Enum</td>
        <td>枚举值</td>
    </tr>
    <tr>
        <td>Custom</td>
        <td>自定义校验</td>
    </tr>
    <tr>
        <td>Dive</td>
        <td>下沉校验struct字段、slice元素、map的key和value。</td>
    </tr>
</table>

| 字段类型函数选项      | 描述       |
|---------------|----------|
| validate.Omitempty | 值为nil时忽略 |
| validate.Amend     | 修正字段值    |

| 规则方法选项   | 描述      |
|----------|---------|
| validate.Code | 自定义错误码  |
| validate.Msg  | 自定义错误消息 |

# 自定义错误码&错误消息

<i>validate.Code</i>只在不使用<i>validate.All()</i>选项的情况下有效。

<i>validate.Msg</i>选项须传入一个函数，通过函数参数<i> f </i>的<i>Msg</i>方法设置错误消息，同时参数<i> f </i>还提供了一些动态参数用于构建错误消息。

*e.g.*

```go
func (b *Book) Checklist(vc *validate.VContext) {
    v.String( &b.Title, "title").Max(10, validate.Msg(func(f *validate.FieldInfo) {
        f.Msg("%s exceeds max length %s", f.FieldName(), f.Confine(0))
    }))
}
```

<table>
    <th width="130px">动态参数方法</th>
    <th>描述</th>
    <tr>
        <td>FieldName</td>
        <td>返回具有层级关系的字段名称。例如：<i>title</i>、<i>author.name</i>、<i>categoryId[2].sort</i>。</td>
    </tr>
    <tr>
        <td>Confine</td>
        <td>返回<i>规则方法</i>的指定索引的限制值的字符串形式。例如：对于<i>Max(10)</i>，<i>vc.Confine(0)</i>返回<i>"10"</i>；对于<i>Range(5, 15)</i>，<i>vc.Confine(0)</i>返回<i>"5"</i>，<i>vc.Confine(1)</i>返回<i>"15"</i>。</td>
    </tr>
    <tr>
        <td>Confines</td>
        <td>返回<i>规则方法</i>的所有限制值的字符串形式，用","拼接，若数量超过两个，则最后一个用"or"拼接。例如：对于<i>Enum("zh-cn", "en-US")</i>，返回"zh-cn, en-US"；对于<i>Enum("zh-cn", "en-US", "ja-JP")</i>，返回"zh-cn, en-US or ja-JP"。</td>
    </tr>
</table>

# 元素的字段名称

在slice和map的<i>Dive</i>方法内校验元素值时，会预设字段名称并且无法更改，slice元素为<i><字段名称>[<索引>]</i>，对于map，key为<i><字段名称>$key</i>，value为<i><字段名称>$value</i>。

*e.g.*

```go
package main

import (
    "context"
    "fmt"

    validate "github.com/jishaocong0910/validate-struct"
)

type Demo struct {
    MySlice []int
    MyMap   map[string]int
}

func (d Demo) Checklist(vc *validate.VContext) {
    v.Slice( d.MySlice, "mySlice").Dive(func(t int) {
        // fieldName参数无效
        v.Int( &t, "elem").Gt(5)
    })
    v.Map( d.MyMap, "myMap").Dive(func(k string) {
        // fieldName参数无效
        v.String( &k, "mapKey").NotBlank()
    }, func(v int) {
        // fieldName参数无效
        v.Int( &v, "mapValue").Gt(5)
    })
}

func main() {
    d := &Demo{
        MySlice: []int{4},
        MyMap:   map[string]int{"": 4},
    }
    code, _, msgs := validate.Check(context.Background(), d, validate.All())
    if code != validate.SUCCESS {
        for i, msg := range msgs {
            fmt.Println(i, msg)
        }
    }
    // Output:
    // 0 mySlice[0] must be greater than 5
    // 1 myMap$key must not be blank
    // 2 myMap$value must be greater than 5
}
```

# 零代码入侵

零代码入侵是Gverify的特色，可在不修改结构体文件的情况下实现校验，原理是<i>Checklist</i>方法可以与任意文件下实现，有以下三种场景，一些场景可实现零代码入侵。

* 在结构体所在文件

对于非代码生成器生成的结构体，可直接在结构体所在文件实现<i>validate.Verifiable</i>接口，上文中的代码示例都是这种风格。

* 在同目录其他文件

为了避免代码生成器覆盖文件，可以在同包下另一个文件实现<i>validate.Verifiable</i>接口。

*目录结构*

```
.
├─ gen_struct.go
└─ gen_struct_checklist.go
```

*gen_struct.go*

```go
package demo

type GenStruct struct {
    Id   string
    Name string
}
```

*gen_struct_checklist.go*

```go
package demo

import validate "github.com/jishaocong0910/validate-struct"

func (g GenStruct) Checklist(vc *validate.VContext) {
    v.String( &g.Id, "id").NotBlank()
    v.String( &g.Name, "name").NotBlank()
}
```

* 在其他目录的文件

也可以在其他包实现`validate.Verifiable`接口。由于Golang语法禁止类型的方法放在非定义它的包下，因此需增加一个内嵌原结构体的新的结构体，通过新的结构体来校验。

*目录结构*

```
.
├─ model
│  └─ gen_struct.go
└─ check
   └─ gen_struct_checklist.go
```

*gen_struct.go*

```go
package model

type GenStruct struct {
    Id   string
    Name string
}
```

*gen_struct_checklist.go*

```go
package check

import (
    "demo/model"

    validate "github.com/jishaocong0910/validate-struct"
)

type GenStruct struct {
    model.GenStruct
}

func (g GenStruct) Checklist(vc *validate.VContext) {
    v.String( &g.Id, "id").NotBlank()
    v.String( &g.Name, "name").NotBlank()
}
```

*校验示例*

```go
package main

import (
    "context"
    "fmt"

    "demo/check"
    "demo/model"

    validate "github.com/jishaocong0910/validate-struct"
)

func main() {
    g := model.GenStruct{}
    code, _, msgs := validate.Check_(context.Background(), &check.GenStruct{g}, validate.All())
    if code != validate.SUCCESS {
        for i, msg := range msgs {
            fmt.Println(i, msg)
        }
    }
    // Output:
    // 0 id must not be blank
    // 1 name must not be blank
}
```