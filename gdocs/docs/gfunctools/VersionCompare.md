---
id: VersionCompare
title: VersionCompare
---


## VersionCompare
`func VersionCompare(v1, v2 string) int`

Compare two version number strings, Prefix `v` optional:

The `v1` is greater than the `v2` returns `1`.

The `v2` is greater than the `v1` return `-1`.

Return 0 if both sides are equal.

```js
gfunctools.VersionCompare("v0.0.1", "v0.0.2")  // -1
gfunctools.VersionCompare("0.0.1", "0.0.2")    // -1
gfunctools.VersionCompare("V0.0.1", "V0.0.2")  // -1
gfunctools.VersionCompare("V0.0.2", "0.0.2")   // 0
gfunctools.VersionCompare("1.0.2", "0.0.2")    // 1
gfunctools.VersionCompare("V1", "0.0.2")       // 1
```
