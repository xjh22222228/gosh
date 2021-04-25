---
id: Includes
title: Includes
sidebar_label: Includes
---

## Includes
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/includes](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/includes)

Syntax: `Includes(str, searchString string[, position int]) bool`

- str - string
- searchString - A string to be searched for within str.
- position - [Optional], The position within the string at which to begin searching for searchString. (Defaults to 0.)

```js
fmt.Println(gstr.Includes("Blue Whale", "blue")) // => false
fmt.Println(gstr.Includes("Blue Whale", "Blue")) // => true
fmt.Println(gstr.Includes("To be", "To be", 1))  // => false
```