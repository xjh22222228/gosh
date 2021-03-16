<p align="center">
  <a href="https://gosh.vercel.app/">
    <img src="media/logo.svg" width="350">
  </a>
  <p align="center">Golang实用程序库，带有其他功能，例如JavaScript / Python！</p>
  <p align="center">
    <a href="https://gosh.vercel.app/">
      <img src="https://img.shields.io/badge/在线-文档-red.svg?longCache=true&style=flat-square">
    </a>
    <a href="README.md">
      <img src="https://img.shields.io/badge/lang-English-red.svg?longCache=true&style=flat-square">
    </a>
    <img src="https://img.shields.io/github/go-mod/go-version/xjh22222228/gosh" />
    <img src="https://img.shields.io/github/v/release/xjh22222228/gosh" />
    <img src="https://img.shields.io/github/license/xjh22222228/gosh" />
    <img src="https://img.shields.io/badge/Coverage-100%25-brightgreen.svg" />
  </p>
<p>


## 安装
```bash
go get -d github.com/xjh22222228/gosh
```


## 入门
```golang
package main

import (
    "fmt"
    "time"
    "github.com/xjh22222228/gosh/gstr"
    "github.com/xjh22222228/gosh/gslice"
    "github.com/xjh22222228/gosh/gtime"
)

func main() {
    s := gstr.Reverse("Hello World")
    fmt.Println(s) // => dlroW olleH
    

    months := []string{"Jan", "March", "April", "June"}
    deleteItem := gslice.Splice(&months, 4, 1, "May")
    fmt.Println(months, deleteItem)  // => [Jan March April June May] []
    
    // 格式化日期
    fmt.Println(gtime.Format(time.Now(), "YYYY-MM-DD HH:mm:ss"))
    // => 2021-03-12 21:25:37
}
```

可以在运行时[文档中](https://gosh.vercel.app/)找到更多详细信息。


## 贡献
感谢您的帮助！
