---
id: Currency
title: Currency
---


## Currency
`func Currency(num float64[, symbol string = "$"]) string`

Convert standard currency numbers:
```js
Currency(123)            // => "$123.00"
Currency(1234.12)        // => "$1,234.12"
Currency(12345.56)       // => "$12,345.56"
Currency(123456.567)     // => "$123,456.57"
Currency(1234567, "￥")  // => "￥1,234,567.00"
Currency(12345678)       // => "$12,345,678.00"
Currency(123456789)      // => "$123,456,789.00"
```
