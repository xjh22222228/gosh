---
id: Choice
title: Choice
---


## Choice
[https://docs.python.org/zh-cn/3/library/random.html#random.choice](https://docs.python.org/zh-cn/3/library/random.html#random.choice)

`func Choice(str string) string`

Return a random element from a string:

```js
str := "GOSH你好世界"
fmt.Println(grand.Choice(str)) // => G
fmt.Println(grand.Choice(str)) // => S
fmt.Println(grand.Choice(str)) // => S
fmt.Println(grand.Choice(str)) // => O
fmt.Println(grand.Choice(str)) // => 好
```

## ChoiceInt
`func ChoiceInt(nums []int) int`

```js
nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
fmt.Println(grand.ChoiceInt(nums)) // => 4
fmt.Println(grand.ChoiceInt(nums)) // => 3
fmt.Println(grand.ChoiceInt(nums)) // => 2
fmt.Println(grand.ChoiceInt(nums)) // => 3
fmt.Println(grand.ChoiceInt(nums)) // => 3
```

## ChoiceStr
`func ChoiceStr(s []string) string`

```js
s := []string{"G", "O", "S", "H", "你", "好", "，", "世", "界"}
fmt.Println(grand.ChoiceStr(s)) // => "G"
fmt.Println(grand.ChoiceStr(s)) // => "世"
fmt.Println(grand.ChoiceStr(s)) // => "你"
fmt.Println(grand.ChoiceStr(s)) // => "O"
fmt.Println(grand.ChoiceStr(s)) // => "好"
```
