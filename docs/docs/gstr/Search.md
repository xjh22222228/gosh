---
id: Search
title: Search
sidebar_label: Search
---

## Search
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/search](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/search)

The **Search()** method executes a search for a match between a regular expression and this String object.

Syntax: `Search(str, regexp string) int`

- str - string
- regexp - A regular expression

```js
str := "hey JudE"
fmt.Println(gstr.Search(str, "[A-Z]")) // => 4
fmt.Println(gstr.Search(str, "[.]"))   // => -1
```