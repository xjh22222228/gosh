---
id: StartsWith
title: StartsWith
sidebar_label: StartsWith
---

## StartsWith
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/startsWith](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/startsWith)

Syntax: `StartsWith(str, searchString string[, position int]) bool`

- str - string
- searchString - The characters to be searched for at the start of this string.
- position - [Optional], The position in this string at which to begin searching for searchString. Defaults to 0.


```js
str := "Saturday night plans"
fmt.Println(gstr.StartsWith(str, "Sat"))    // => true
fmt.Println(gstr.StartsWith(str, "Sat", 3)) // => false
```
