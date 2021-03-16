---
id: Trim
title: Trim
sidebar_label: Trim
---

## Trim
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/Trim](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/Trim)

The **Trim()** method removes whitespace from both ends of a string. Whitespace in this context is all the whitespace characters (space, tab, no-break space, etc.) and all the line terminator characters (LF, CR, etc.).

Syntax: `Trim(str string) string`

```js
fmt.Println(gstrings.Trim("   Hello world!   "))  // => Hello world!

str := `


abc
`
fmt.Println(gstrings.Trim(str))  // => abc
```