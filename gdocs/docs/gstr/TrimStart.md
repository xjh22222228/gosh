---
id: TrimStart
title: TrimStart
sidebar_label: TrimStart
---

## TrimStart
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/TrimLeft](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/TrimLeft)

The **TrimStart()** method removes whitespace from the beginning of a string.

Syntax: `TrimStart(str string) string`


```js
fmt.Println(gstr.TrimStart(`

GOSH`))
// => "GOSH"


fmt.Println(gstr.TrimStart("   Hello world!   "))
// => "Hello world!   "
```