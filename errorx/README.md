# errorx

`errorx` 提供统一的错误码注册、错误创建、错误包装与错误码提取能力。

## 特性

- 统一错误码常量：
  - `CodeOK = 0`
  - `CodeUnknown = 1000`
- 通过注册中心维护错误码与默认文案
- 支持创建业务错误并通过 `errors.As` 链路提取错误码
- 支持包装底层错误，保留原始错误链（`Unwrap`）

## 快速开始

```go
package main

import (
	"errors"
	"fmt"

	"github.com/cn-hideyoshi/HUtil/errorx"
)

const (
	CodeUserNotFound = 2001
)

func main() {
	// 注册错误码（建议在 init 或应用启动阶段执行一次）
	errorx.MustRegister(CodeUserNotFound, "user.not_found", "user not found")

	// 基于注册信息创建错误
	err1 := errorx.New(CodeUserNotFound)
	fmt.Println(err1.Error())      // user not found
	fmt.Println(errorx.Code(err1)) // 2001

	// 包装底层错误
	base := errors.New("sql: no rows")
	err2 := errorx.Wrap(base, CodeUserNotFound, "query user failed")
	fmt.Println(err2.Error())      // query user failed: sql: no rows
	fmt.Println(errorx.Is(err2, CodeUserNotFound)) // true

	// nil -> CodeOK
	fmt.Println(errorx.Code(nil)) // 0
}
```

## API

- `Register(code int, key, defaultMsg string) error`
  - 注册错误码，重复注册会返回错误。
  - `defaultMsg` 会作为该错误码的默认消息。
  - 当前实现中 `key` 参数会被接收但不参与存储，可用于兼容上层调用约定。
- `MustRegister(code int, key, defaultMsg string)`
  - 与 `Register` 相同，但失败时 panic。
- `GetMeta(code int) (*CodeMeta, bool)`
  - 按错误码查询元信息。
- `ListCodes() []*CodeMeta`
  - 返回已注册错误码列表，按 `Code` 升序排序。
- `New(code int) *Error`
  - 根据已注册错误码创建错误；若未注册会 panic。
- `NewWithMessage(code int, msg string) *Error`
  - 直接用自定义消息创建错误（无需预注册）。
- `Wrap(err error, code int, msg string) *Error`
  - 包装已有错误；当 `err == nil` 时返回 `nil`。
- `Code(err error) int`
  - 提取错误码：
    - `err == nil` 返回 `CodeOK`
    - 非 `errorx.Error` 返回 `CodeUnknown`
- `Is(err error, code int) bool`
  - 判断错误码是否匹配。

## 使用建议

- 在应用启动阶段统一注册业务错误码，避免运行时遗漏导致 `New` panic。
- 推荐在边界层使用 `Wrap` 保留底层错误细节，在上层通过 `Code`/`Is` 做流程分支。
