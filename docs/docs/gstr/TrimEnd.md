---
id: TrimEnd
title: TrimEnd
sidebar_label: TrimEnd
---


## TrimEnd
[https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/trimEnd](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/trimEnd)

The **TrimEnd()** method removes whitespace from the end of a string. 

Syntax: `TrimEnd(str string) string`

```js
fmt.Println(gstr.TrimEnd("  GOSH  "))
// => "  GOSH"

fmt.Println(gstr.TrimEnd(`GOSH

`))
// => "GOSH"
```