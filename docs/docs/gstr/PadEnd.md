---
id: PadEnd
title: PadEnd
sidebar_label: PadEnd
---


## PadEnd
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/padEnd](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/padEnd)

The **PadEnd()** method pads the current string with a given string (repeated, if needed) so that the resulting string reaches a given length. The padding is applied from the end of the current string.


Syntax: `PadStart(str string, targetLength int[, padString string]) string`

- str - string
- targetLength - int
  - The length of the resulting string once the current str has been padded. If the value is lower than **len(str)**, the current string will be returned as-is.
- padString - string, [Optional]
  - The string to pad the current str with. If padString is too long to stay within targetLength, it will be truncated: for left-to-right languages the left-most part and for right-to-left languages the right-most will be applied. The default value for this parameter is " ".


```js
str := "abc"
fmt.Println(gstr.PadEnd(str, 10))          // => "abc"
fmt.Println(gstr.PadEnd(str, 10, "foo"))   // => "abcfoofoof"
fmt.Println(gstr.PadEnd(str, 6, "12345"))  // => "abc123"
fmt.Println(gstr.PadEnd(str, 1))           // => "abc"
```