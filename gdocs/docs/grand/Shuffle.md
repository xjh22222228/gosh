---
id: Shuffle
title: Shuffle
---

## Shuffle
Syntax: `Shuffle([]string)`

Shuffle the list (rearrange the order of the list items):
```js
l := []string{"a", "b", "c", "d", "e"}
grand.Shuffle(l)
fmt.Println(l) // => [b e a c d]
```


## ShuffleInt

```js
l := []int{1, 2, 3, 4, 5, 6}
grand.ShuffleInt(l)
fmt.Println(l) // => [6 4 2 5 1 3]
```


## ShuffleAny

```js
l := []interface{}{"a", false, true, 1, "e", 2, nil}
grand.ShuffleAny(l)
fmt.Println(l) // => [<nil> 1 e a false true 2]
```