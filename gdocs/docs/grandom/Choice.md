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
fmt.Println(grandom.Choice(str)) // => G
fmt.Println(grandom.Choice(str)) // => S
fmt.Println(grandom.Choice(str)) // => S
fmt.Println(grandom.Choice(str)) // => O
fmt.Println(grandom.Choice(str)) // => 好
```

## ChoiceInt
`func ChoiceInt(nums []int) int`

```js
nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
fmt.Println(grandom.ChoiceInt(nums)) // => 4
fmt.Println(grandom.ChoiceInt(nums)) // => 3
fmt.Println(grandom.ChoiceInt(nums)) // => 2
fmt.Println(grandom.ChoiceInt(nums)) // => 3
fmt.Println(grandom.ChoiceInt(nums)) // => 3
```

## ChoiceStr
`func ChoiceStr(s []string) string`

```js
s := []string{"G", "O", "S", "H", "你", "好", "，", "世", "界"}
fmt.Println(grandom.ChoiceStr(s)) // => "G"
fmt.Println(grandom.ChoiceStr(s)) // => "世"
fmt.Println(grandom.ChoiceStr(s)) // => "你"
fmt.Println(grandom.ChoiceStr(s)) // => "O"
fmt.Println(grandom.ChoiceStr(s)) // => "好"
```
