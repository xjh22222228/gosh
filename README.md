<p align="center">
  <a href="https://xjh22222228.github.io/gosh">
    <img src="media/logo.svg" width="350">
  </a>
  <p align="center">Golang utility library, With additional functions such as JavaScript/Python!</p>
  <p align="center">
    <img src="https://img.shields.io/github/go-mod/go-version/xjh22222228/gosh" />
    <img src="https://img.shields.io/github/v/release/xjh22222228/gosh" />
    <img src="https://img.shields.io/github/license/xjh22222228/gosh" />
    <img src="https://img.shields.io/badge/Coverage-100%25-brightgreen.svg" />
  </p>
<p>


## Installation
```bash
go get -d github.com/xjh22222228/gosh
```


## Documentation
[Documentation](https://xjh22222228.github.io/gosh)


## Demo
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

