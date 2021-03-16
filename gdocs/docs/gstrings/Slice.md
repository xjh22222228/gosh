---
id: Slice
title: Slice
sidebar_label: Slice
---

## Slice
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/slice](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/slice)

Syntax: `Slice(beginIndex int[, endIndex int]) string`

- beginIndex
  - The zero-based index at which to begin extraction. If negative, it is treated as str.length + beginIndex. (For example, if beginIndex is -3, it is treated as str.length - 3.) If beginIndex is not a number after Number(beginIndex), it is treated as 0.
  - If beginIndex is greater than or equal to str.length, an empty string is returned.
- endIndex - [Optional]
  - The zero-based index before which to end extraction. The character at this index will not be included.
  - If endIndex is omitted or undefined, or greater than str.length, slice() extracts to the end of the string. If negative, it is treated as str.length + endIndex. (For example, if endIndex is -3, it is treated as str.length - 3.) If it is not undefined and not a number after Number(endIndex), an empty string is returned.
  - If endIndex is specified and startIndex is negative, endIndex should be negative, otherwise an empty string is returned. (For example, slice(-3, 0) returns "".)
  - If endIndex is specified, and startIndex and endIndex are both positive or negative, endIndex should be greater than startIndex, otherwise an empty string is returned. (For example, slice(-1, -3) or slice(3, 1) returns "".)

```js
str := "The quick brown fox jumps over the lazy dog."
fmt.Println(gstrings.Slice(str, 31))      // => the lazy dog.
fmt.Println(gstrings.Slice(str, 4, 19))   // => quick brown fox
fmt.Println(gstrings.Slice(str, -4))      // => dog.
fmt.Println(gstrings.Slice(str, -9, -5))  // => lazy
```