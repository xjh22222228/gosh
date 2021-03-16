---
id: RemovePrefix
title: RemovePrefix
sidebar_label: RemovePrefix
---


## RemovePrefix
[https://docs.python.org/3/library/stdtypes.html#str.removeprefix](https://docs.python.org/3/library/stdtypes.html#str.removeprefix)

If the string starts with the prefix string, return string[len(prefix):]. Otherwise, return a copy of the original string:

```js
RemovePrefix("TestHook", "Test")      // => "Hook"
RemovePrefix("BaseTestCase", "Test")  // => "BaseTestCase"
```