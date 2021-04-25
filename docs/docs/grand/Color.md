---
id: Color
title: Color
---


## Constants
```js
const (
	ColorRGB colorType = iota + 1
	ColorRGBA
	ColorHEX
	ColorNAME
	ColorHSL
	ColorHSLA
)
```


## Color
`func Color(t colorType) string`

Randomly generate colors, `RGB`, `RGBA`, `COLOR`, `HEX`:

```js
grand.Color(grand.ColorRGB)    // => rgb(247,39,65)
grand.Color(grand.ColorRGBA)   // => rgba(124,236,25,0.55)
grand.Color(grand.ColorHEX)    // => #c5e814
grand.Color(grand.ColorNAME)   // => Red
grand.Color(grand.ColorHSL)    // => hsl(16,46%,52%)
grand.Color(grand.ColorHSLA)   // => hsla(266,42%,73%,0.26)
```
