---
id: Shuffle
title: Shuffle
sidebar_label: Shuffle
---


## Shuffle
Syntax: `Shuffle(str string) string`

Shuffle string order：
```js
fmt.Println(gstrings.Shuffle("世界你好 hello world"))  // => 世 rhldwe l好你oo界l
fmt.Println(gstrings.Shuffle("123456789"))            // => 648317529
fmt.Println(gstrings.Shuffle("abcdefghijk"))          // => hgebfdkcjai
```
