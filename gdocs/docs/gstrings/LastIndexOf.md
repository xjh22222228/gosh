---
id: LastIndexOf
title: LastIndexOf
sidebar_label: LastIndexOf
---


## LastIndexOf
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/lastIndexOf](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/lastIndexOf)

The **LastIndexOf()** method returns the index within the calling String object of the last occurrence of the specified value, searching backwards from fromIndex. Returns -1 if the value is not found.

Syntax: `LastIndexOf(str, searchStr string, [, fromIndex int]) int`

```js
str := "canal"
fmt.Println(gstrings.LastIndexOf(str, "a"))      // => 3
fmt.Println(gstrings.LastIndexOf(str, "a", 2))   // => 1
fmt.Println(gstrings.LastIndexOf(str, "a", 0))   // => -1
fmt.Println(gstrings.LastIndexOf(str, "x"))      // => -1
fmt.Println(gstrings.LastIndexOf(str, "c", -5))  // => 0
fmt.Println(gstrings.LastIndexOf(str, "c", 0))   // => 0
fmt.Println(gstrings.LastIndexOf(str, ""))       // => 5
fmt.Println(gstrings.LastIndexOf(str, "", 2))    // => 2

fmt.Println(gstrings.LastIndexOf("Brave new world", "w"))      // => 10
fmt.Println(gstrings.LastIndexOf("Brave new world", "new"))    // => 6
```