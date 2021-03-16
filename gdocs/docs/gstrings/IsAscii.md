---
id: IsAscii
title: IsAscii
sidebar_label: IsAscii
---


## IsAscii
[https://docs.python.org/3/library/stdtypes.html#str.isascii](https://docs.python.org/3/library/stdtypes.html#str.isascii)

Return `true` if the string is empty or all characters in the string are ASCII, `false` otherwise. ASCII characters have code points in the range U+0000-U+007F.

```js
str := "\u0001\u0002\u0003\u0004\u0005\u0006\a\b\t\n\v\f\n\u000E\u000F\u0010\u0011\u0012\u0013\u0014\u0015\u0016\u0017\u0018\u0019\u001A\u001B\u001C\u001D\u001E\u001F !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	
fmt.Println(IsAscii(str))      // => true
fmt.Println(IsAscii(""))      // => false
fmt.Println(IsAscii("\u0080")) // => false
fmt.Println(IsAscii("好"))     // => false
fmt.Println(IsAscii("a好"))    // => false
```