---
id: IndexOf
title: IndexOf
sidebar_label: IndexOf
---


## IndexOf
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/indexOf](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/indexOf)

The **IndexOf()** method returns the index within the calling String object of the first occurrence of the specified value, starting the search at fromIndex. Returns -1 if the value is not found.

Syntax: `IndexOf(str, searchStr string, [, fromIndex int]) int`

```js
str := "hello world"
fmt.Println(gstr.IndexOf(str, ""))      // => 0
fmt.Println(gstr.IndexOf(str, "", 0))   // => 0
fmt.Println(gstr.IndexOf(str, "", 3))   // => 3
fmt.Println(gstr.IndexOf(str, "", 8))   // => 8
fmt.Println(gstr.IndexOf(str, "", 11))  // => 11
fmt.Println(gstr.IndexOf(str, "", 13))  // => 11
fmt.Println(gstr.IndexOf(str, "", 22))  // => 11

str = "Blue Whale"
fmt.Println(gstr.IndexOf(str, "Blue"))      // => 0
fmt.Println(gstr.IndexOf(str, "Blute"))     // => -1
fmt.Println(gstr.IndexOf(str, "Whale", 0))  // => 5
fmt.Println(gstr.IndexOf(str, "Whale", 5))  // => 5
fmt.Println(gstr.IndexOf(str, "", -1))      // => 0
fmt.Println(gstr.IndexOf(str, "", 9))       // => 9
fmt.Println(gstr.IndexOf(str, "", 10))      // => 10
fmt.Println(gstr.IndexOf(str, "", 11))      // => 10
fmt.Println(gstr.IndexOf(str, "blue"))      // => -1
```