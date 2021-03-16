---
id: Fill
title: Fill
sidebar_label: Fill
---


## Fill
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/fill](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/fill)


Syntax: `func Fill(elems []string, value string, start int, end ...int) []string`


The **Fill()** method changes all elements in an array to a static value, from a start index (default 0) to an end index (default array.length). It returns the modified array.

```js
elems := []string{"1", "2", "3", "4"}
fmt.Println(gslice.Fill(elems, "0", 2, 4))  // => [1 2 0 0]
fmt.Println(gslice.Fill(elems, "6", 0))     // => [6 6 6 6]
```

## FillInt

```js
elems := []int{1, 2, 3, 4}
fmt.Println(gslice.FillInt(elems, 0, 2, 4))  // => [1 2 0 0]
fmt.Println(gslice.FillInt(elems, 6, 0))     // => [6 6 6 6]
```

## FillAny

```js
elems := []interface{}{1, "2", 3, 4}
fmt.Println(gslice.FillAny(elems, 0, 2, 4))  // => [1 2 0 0]
fmt.Println(gslice.FillAny(elems, 6, 0))     // => [6 6 6 6]
```
