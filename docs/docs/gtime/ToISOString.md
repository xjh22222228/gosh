---
id: ToISOString
title: ToISOString
---


## ToISOString
`func ToISOString(t time.Time) string`

The `ToISOString()` method returns a string in simplified extended ISO format (ISO 8601), which is always 24 or 27 characters long (`YYYY-MM-DDTHH:mm:ss.sssZ` or `±YYYYYY-MM-DDTHH:mm:ss.sssZ`, respectively). The timezone is always zero UTC offset, as denoted by the suffix `"Z"`.

```js
t := time.Now()
ToISOString(t)
// => 2021-03-22T21:05:05.077Z
```

