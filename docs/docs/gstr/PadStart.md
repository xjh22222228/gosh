---
id: PadStart
title: PadStart
sidebar_label: PadStart
---


## PadStart
[https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/padStart](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/padStart)

The **PadStart()** method pads the current string with another string (multiple times, if needed) until the resulting string reaches the given length. The padding is applied from the start of the current string.

Syntax: `PadStart(str string, targetLength int[, padString string]) string`

- str - string
- targetLength - int
  - The length of the resulting string once the current str has been padded. If the value is less than **len(str)**, then str is returned as-is.
- padString - string, [Optional]
  - The string to pad the current str with. If **padString** is too long to stay within the **targetLength**, it will be truncated from the end. The default value is " ".


```js
str := "abc"
fmt.Println(gstr.PadStart(str, 10))           // => "       abc"
fmt.Println(gstr.PadStart(str, 10, "foo"))    // => "foofoofabc"
fmt.Println(gstr.PadStart(str, 6), "12345")   // => "   abc 12345"
fmt.Println(gstr.PadStart(str, 8, "0"))       // => "00000abc"
fmt.Println(gstr.PadStart(str, 1))            // => "abc"
```