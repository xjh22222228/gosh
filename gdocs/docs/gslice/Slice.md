---
id: Slice
title: Slice
sidebar_label: Slice
---


## Slice
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/slice](https://developer.mozilla.
org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/slice)

Syntax: `func Slice(s []string, beginIndex int[, endIndex int]) []string`

The **Slice()** method returns a shallow copy of a portion of an array into a new array object selected from start to end (end not included) where start and end represent the index of items in that array. The original array will not be modified.


```js
animals := []string{"ant", "bison", "camel", "duck", "elephant"}
fmt.Println(gslice.Slice(animals, 2))     // => [camel duck elephant]
fmt.Println(gslice.Slice(animals, 2, 4))  // => [camel duck]
fmt.Println(gslice.Slice(animals, 1, 5))  // => [bison camel duck elephant]
fmt.Println(gslice.Slice(animals, -1))    // => [elephant]
```