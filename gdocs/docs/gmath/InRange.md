---
id: InRange
title: InRange
---


## InRange
`func InRange(number[, start float64 = 0], end float64) bool`

Checks if n is between start and up to, but not including, end. If end is not specified, it's set to start with start then set to 0. If start is greater than end the params are swapped to support negative ranges:

```js
gmath.InRange(3, 2, 4)    // => true
gmath.InRange(-3, -2, -6) // => true
gmath.InRange(4, 8)       // => true
gmath.InRange(1.2, 2)     // => true
gmath.InRange(4, 2)       // => false
gmath.InRange(2, 2)       // => false
gmath.InRange(5.2, 4)     // => false
```
