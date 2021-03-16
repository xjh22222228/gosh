---
id: Splice
title: Splice
sidebar_label: Splice
---


## Splice
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/splice](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/splice)

Syntax: `func Splice(elems *[]string, start, deleteCount int[, item ...string]) []string`

The **Splice()** method changes the contents of an slice by removing or replacing existing elements and/or adding new elements in place.

Example-1
```js
months := []string{"Jan", "March", "April", "June"}
deleteItem := gslice.Splice(&months, 1, 0, "Feb")
fmt.Println(months, deleteItem) // => [Jan Feb March April June] []
```

Example-2
```js
months := []string{"Jan", "March", "April", "June"}
deleteItem := gslice.Splice(&months, 4, 1, "May")
fmt.Println(months, deleteItem)  // => [Jan March April June May] []
```

Example-3
```js
myFish := []string{"angel", "clown", "drum", "mandarin", "sturgeon"}
deleteItem := gslice.Splice(&myFish, 3, 1)
fmt.Println(myFish, deleteItem)   // => [angel clown drum sturgeon] [mandarin]
```
