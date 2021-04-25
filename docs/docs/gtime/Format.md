---
id: Format
title: Format
---


## Constants
```js
const (
    FormatDate = "YYYY-MM-DD"
    FormatTime = "HH:mm:ss"
    FormatYear = "YYYY"
    FormatDateTime = "YYYY-MM-DD HH:mm:ss"
)
```


## Format
`func Format(t time.Time, formatStr string) string`

Format the date in a commonly used format：
```js
gtime.Format(time.Now(), "YYYY-MM-DD HH:mm:ss")
// => "2021-03-12 21:25:37"
```


## List of all available formats

| Format| Output           | Description              |
| ------------------------ |------------------ |
| YY   | 18                | Two-digit year     |
| YYYY | 2018              | Four-digit year|
| M    | 1-12              | The month, beginning at 1 |
| MM   | 01-12             | The month, 2-digits |
| MMM  | Jan-Dec           | The abbreviated month name |
| MMMM | January-December  | The full month name |
| D    | 1-31              | The day of the month |
| DD   | 01-31             | The day of the month, 2-digits |
| d    | 0-6               | The day of the week, with Sunday as 0 |
| dd   | Su-Sa             | The min name of the day of the week |
| ddd  | Sun-Sat           | The short name of the day of the week |
| dddd | Sunday-Saturday   | The name of the day of the week |
| H    | 0-23              | The hour |
| HH   | 00-23             | The hour, 2-digits |
| h    | 1-12              | The hour, 12-hour clock |
| hh   | 01-12             | The hour, 12-hour clock, 2-digits |
| m    | 0-59              | The minute |
| mm   | 00-59             | The minute, 2-digits |
| s    | 0-59              | The second |
| ss   | 00-59             | The second, 2-digits |
| sss  | 000-999           | The millisecond, 3-digits |
| A    | AM PM             |  |
| a    | am pm             |  |



## Supported languages

| Language               | Name          |
| ---------------------- |-------------- |
| English                | glocale.ENUS   |
| Chinese (Simplified)   | glocale.ZHCN   |
| 繁体中文（中国香港）      | glocale.ZHHK   |
| 繁体中文（中国香港）      | glocale.ZHTW   |


```js
t := gtime.SetLocale(glocale.ZHCN)
t.Format(time.now(), "YYYY年MM月DD日 HH时mm分ss秒")
```

