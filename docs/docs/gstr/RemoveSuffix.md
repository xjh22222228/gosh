---
id: RemoveSuffix
title: RemoveSuffix
sidebar_label: RemoveSuffix
---


## RemoveSuffix
[https://docs.python.org/3/library/stdtypes.html#str.removesuffix](https://docs.python.org/3/library/stdtypes.html#str.removesuffix)

If the string ends with the suffix string and that suffix is not empty, return string[:-len(suffix)]. Otherwise, return a copy of the original string:

```js
RemoveSuffix("MiscTests", "Tests")      // => "Misc"
RemoveSuffix("TmpDirMixin", "Tests")    // => "TmpDirMixin"
```