---
id: Max
title: Max
---


## Max
`func Max(n ...float64) float64`

The `Max()` function returns the largest of the zero or more numbers given as input parameters:

```js
nums := []float64{-10, 1, 2, 3, 4, 5}
gmath.Max(nums...) // => 5

gmath.Max()        // => 0
gmath.Max(5)       // => 5
gmath.Max(-10, -5) // => -5
```
