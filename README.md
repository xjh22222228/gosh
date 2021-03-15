<p align="center">
  <a href="https://gosh.vercel.app/">
    <img src="media/logo.svg" width="350">
  </a>
  <p align="center">Golang utility library, With additional functions such as JavaScript/Python!</p>
  <p align="center">
    <a href="https://gosh.vercel.app/">
      <img src="https://img.shields.io/badge/online-docs-red.svg?longCache=true&style=flat-square">
    </a>
    <a href="README_zh-CN.md">
      <img src="https://img.shields.io/badge/lang-%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87-red.svg?longCache=true&style=flat-square">
    </a>
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



## Getting Started
```golang
package main

import (
    "fmt"
    "time"
    "github.com/xjh22222228/gosh/gstrings"
    "github.com/xjh22222228/gosh/gslice"
    "github.com/xjh22222228/gosh/gtime"
)

func main() {
    s := gstrings.Reverse("Hello World")
    fmt.Println(s) // => dlroW olleH


    months := []string{"Jan", "March", "April", "June"}
    deleteItem := gslice.Splice(&months, 4, 1, "May")
    fmt.Println(months, deleteItem)  // => [Jan March April June May] []
    
    // Format date
    fmt.Println(gtime.Format(time.Now(), "YYYY-MM-DD HH:mm:ss"))
    // => 2021-03-12 21:25:37
}
```

More in-depth info can be found in the runtime [documentation](https://gosh.vercel.app/).


## Contributing
We appreciate your help!
