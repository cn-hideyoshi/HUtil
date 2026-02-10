# HUtil
a golang util library
## 简介
HUtil 是一个轻量的 Golang 工具库集合，提供常用的通用函数与工具包，目标是简单、直观、易复用。

## 目录
- `slice`：切片相关工具

## 快速开始
安装：
```bash
go get github.com/cn-hideyoshi/HUtil
```

使用示例：
```go
package main

import (
	"fmt"

	"github.com/cn-hideyoshi/HUtil/slice"
)

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(slice.InSlice(nums, 3)) // true
	fmt.Println(slice.InSlice(nums, 5)) // false
}
```

## 维护与贡献
欢迎提交 PR 或 Issue 来完善工具集。
