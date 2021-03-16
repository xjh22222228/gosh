---
id: Min
title: Min
---


## Min
`func Min(n ...float64) float64`

The `Min()` function returns the smallest of the zero or more numbers given as input parameters:

```js
nums := []float64{-10, 1, 2, 3, 4, 5}
gmath.Min(nums...) // => -10

gmath.Min()        // => 0
gmath.Min(5)       // => 5
gmath.Min(-10, -5) // => -10
```
