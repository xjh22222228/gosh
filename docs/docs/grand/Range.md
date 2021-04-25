---
id: Range
title: Range
---


## Range
`func Range(start, stop int[, step int = 1]) []int`

The range() function returns a sequence of numbers, increments by 1 (by default), and stops before a specified number:

```js
grand.Range(0, 10)     // => [0 1 2 3 4 5 6 7 8 9]
grand.Range(0, 10, 2)  // => [0 2 4 6 8]
grand.Range(0, 10, 3)  // => [0 3 6 9]
grand.Range(0, 10, 0)  // => []
```
