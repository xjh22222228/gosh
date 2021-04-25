---
id: Clamp
title: Clamp
---


## Clamp
`func Clamp(number, lower, upper int) int`

Clamps `number` within the inclusive `lower` and `upper` bounds.:

```js
gmath.Clamp(-10, -5, 5) // => -5
gmath.Clamp(10, -5, 5)  // => 5
```


## ClampFloat

```js
gmath.ClampFloat(-10, -5, 5) // => -5.0
gmath.ClampFloat(10, -5, 5)  // => 5.0
```
