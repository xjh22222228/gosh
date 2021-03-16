---
id: Center
title: Center
sidebar_label: Center
---


## Center
[https://www.w3schools.com/python/ref_string_center.asp](https://www.w3schools.com/python/ref_string_center.asp)

The **center()** method will center align the string, using a specified character (space is default) as the fill character.

Syntax: `Center(str string, width int[, fillChar rune]) string`
- width - Required. The length of the returned string
- fillChar - [Optional], The character to fill the missing space on each side. Default is " " (space)


```js
fmt.Println(gstrings.Center("[GOSH]", 6)) // => [GOSH]

fmt.Println(gstrings.Center("[GOSH]", 7)) // => " [GOSH]"

fmt.Println(gstrings.Center("[GOSH]", 7, '*')) // => *[GOSH]

fmt.Println(gstrings.Center("[GOSH]", 8, '*')) // => *[GOSH]*
```