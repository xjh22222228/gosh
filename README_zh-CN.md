<p align="center">
  <a href="https://xjh22222228.github.io/gosh">
    <img src="media/logo.svg" width="350">
  </a>
  <p align="center">Golang实用程序库，带有其他功能，例如JavaScript / Python！</p>
  <p align="center">
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


## 文档
[文档地址](https://xiejiahe.gitee.io/gosh)


## 例子
```golang
package main

import (
    "fmt"
    "github.com/xjh22222228/gosh/gstrings"
)

func main() {
    s := gstrings.Reverse("Hello World")
    fmt.Println(s) // => dlroW olleH
}
```

