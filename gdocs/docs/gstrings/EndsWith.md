---
id: EndsWith
title: EndsWith
sidebar_label: EndsWith
---

## EndsWith
https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/endsWith

Syntax: `EndsWith(str, searchString string[, length int]) bool`

- str - string
- searchString - The characters to be searched for at the end of str.
- length - [Optional], If provided, it is used as the length of str. Defaults to str.length.

```js
str := "Cats are the best!"
fmt.Println(gstrings.EndsWith(str, "best", 17)) // => true

str2 := "Is this a question"
fmt.Println(gstrings.EndsWith(str2, "?")) // => false
```
