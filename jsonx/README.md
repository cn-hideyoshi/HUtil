# jsonx

`jsonx` 是对标准库 `encoding/json` 的轻量封装，提供更便捷的 JSON 编解码、严格解析、格式化与常用工具方法。

## 功能
- 基础编解码：`Marshal`、`Unmarshal`
- Must 风格：`MustMarshal`、`MustUnmarshal`、`MustMarshalToString`
- 字符串编码：`MarshalToString`
- 对象转换：`ToMap`
- 字段读取：`Get`
- Map 合并：`Merge`
- JSON 美化与压缩：`Format`、`Compress`
- 严格解析（禁止未知字段）：`UnmarshalStrict`
- 深拷贝：`DeepCopy`

## 安装
```bash
go get github.com/cn-hideyoshi/HUtil
```

## 导入
```go
import "github.com/cn-hideyoshi/HUtil/jsonx"
```

## 示例
```go
package main

import (
	"fmt"

	"github.com/cn-hideyoshi/HUtil/jsonx"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := User{Name: "Tom", Age: 18}

	// 编码
	s, _ := jsonx.MarshalToString(u)
	fmt.Println(s) // {"name":"Tom","age":18}

	// 解码
	var u2 User
	_ = jsonx.Unmarshal([]byte(s), &u2)

	// 转 map
	m, _ := jsonx.ToMap(u2)
	fmt.Println(m["name"]) // Tom

	// 读取字段
	v, _ := jsonx.Get([]byte(s), "age")
	fmt.Println(v) // 18

	// 美化与压缩
	pretty, _ := jsonx.Format(s)
	minified, _ := jsonx.Compress(pretty)
	fmt.Println(len(pretty) > len(minified)) // true
}
```

## API 说明
- `Marshal(v any) ([]byte, error)`：序列化为 JSON 字节。
- `Unmarshal(data []byte, v any) error`：反序列化 JSON 到目标对象。
- `MustMarshal(v any) []byte`：`Marshal` 失败时 `panic`。
- `MustUnmarshal(data []byte, v any)`：`Unmarshal` 失败时 `panic`。
- `MustMarshalToString(v any) string`：序列化为字符串，失败时 `panic`。
- `MarshalToString(v any) (string, error)`：序列化为字符串。
- `ToMap(v any) (map[string]any, error)`：将对象转为 `map[string]any`。
- `Get(data []byte, key string) (any, error)`：读取顶层 key 对应的值。
- `Merge(dst, src map[string]any) map[string]any`：将 `src` 字段覆盖合并到 `dst`。
- `Format(input string) (string, error)`：JSON 字符串格式化（2 空格缩进）。
- `Compress(input string) (string, error)`：JSON 字符串压缩为单行。
- `UnmarshalStrict(data []byte, v any) error`：严格模式解析，禁止未知字段。
- `DeepCopy[T any](src T) (T, error)`：基于 JSON 的泛型深拷贝。
