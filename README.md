<p align="center">
  <img src="media/logo.svg" width="350">
  <p align="center">Golang utility library, With additional functions such as JavaScript/Python!</p>
  <p align="center">
    <img src="https://img.shields.io/github/go-mod/go-version/xjh22222228/gosh" />
    <img src="https://img.shields.io/github/v/release/xjh22222228/gosh" />
    <img src="https://img.shields.io/github/license/xjh22222228/gosh" />
    <img src="https://img.shields.io/badge/Coverage-100%25-brightgreen.svg" />
  </p>
<p>


## Installation
```bash
go get -d github.com/xjh22222228/gosh
```


## Demo
```golang
package main

import (
    "fmt"
    "github.com/xjh22222228/gosh/gstrings"
)

func main() {
    s := gstrings.Reverse("Hello World")
    fmt.Println(s) // => dlroW olleH
}
```



## API
- [gstrings](#gstrings)
  - [Reverse](#Reverse)
  - [StartsWith](#StartsWith)
  - [EndsWith](#EndsWith)
  - [Concat](#Concat)
  - [Slice](#Slice)
  - [Search](#Search)
  - [Includes](#Includes)
  - [Trim](#Trim)
  - [TrimStart](#TrimStart)
  - [TrimEnd](#TrimEnd)
  - [Center](#Center)
  - [PadStart](#PadStart)
  - [PadEnd](#PadEnd)
  - [IndexOf](#IndexOf)
  - [LastIndexOf](#LastIndexOf)
  - [Shuffle](#Shuffle)
- [grandom](#grandom)
  - [Random](#Random)
  - [RandBool](#RandBool)
  - [RandInt](#RandInt)
  - [Shuffle](#Shuffle)
  - [ShuffleInt](#ShuffleInt)
  - [ShuffleAny](#ShuffleAny)
  - [Choice](#Choice)


# gstrings

## Reverse
Syntax: `Reverse(str string) string`

The `reverse()` method reverses an string in place. 
```golang
gstrings.Reverse("Hello World")
// => dlroW olleH
```


## StartsWith
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/startsWith](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/startsWith)

Syntax: `StartsWith(str, searchString string[, position int]) bool`

- str - string
- searchString - The characters to be searched for at the start of this string.
- position - [Optional], The position in this string at which to begin searching for searchString. Defaults to 0.


```golang
str := "Saturday night plans"
fmt.Println(gstrings.StartsWith(str, "Sat"))    // => true
fmt.Println(gstrings.StartsWith(str, "Sat", 3)) // => false
```


## EndsWith
https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/endsWith

Syntax: `EndsWith(str, searchString string[, length int]) bool`

- str - string
- searchString - The characters to be searched for at the end of str.
- length - [Optional], If provided, it is used as the length of str. Defaults to str.length.

```golang
str := "Cats are the best!"
fmt.Println(gstrings.EndsWith(str, "best", 17)) // => true

str2 := "Is this a question"
fmt.Println(gstrings.EndsWith(str2, "?")) // => false
```


## Concat
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/concat](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/concat)

Syntax: `Concat(str2, [, ...strN]) string`

```golang
fmt.Println(gstrings.Concat("a", "b", "c")) // => abc
```



## Slice
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/slice](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/slice)

Syntax: `Slice(beginIndex int[, endIndex int]) string`

- beginIndex
  - The zero-based index at which to begin extraction. If negative, it is treated as str.length + beginIndex. (For example, if beginIndex is -3, it is treated as str.length - 3.) If beginIndex is not a number after Number(beginIndex), it is treated as 0.
  - If beginIndex is greater than or equal to str.length, an empty string is returned.
- endIndex - [Optional]
  - The zero-based index before which to end extraction. The character at this index will not be included.
  - If endIndex is omitted or undefined, or greater than str.length, slice() extracts to the end of the string. If negative, it is treated as str.length + endIndex. (For example, if endIndex is -3, it is treated as str.length - 3.) If it is not undefined and not a number after Number(endIndex), an empty string is returned.
  - If endIndex is specified and startIndex is negative, endIndex should be negative, otherwise an empty string is returned. (For example, slice(-3, 0) returns "".)
  - If endIndex is specified, and startIndex and endIndex are both positive or negative, endIndex should be greater than startIndex, otherwise an empty string is returned. (For example, slice(-1, -3) or slice(3, 1) returns "".)

```golang
str := "The quick brown fox jumps over the lazy dog."
fmt.Println(gstrings.Slice(str, 31))      // => the lazy dog.
fmt.Println(gstrings.Slice(str, 4, 19))   // => quick brown fox
fmt.Println(gstrings.Slice(str, -4))      // => dog.
fmt.Println(gstrings.Slice(str, -9, -5))  // => lazy
```


## Search
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/search](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/search)

The **Search()** method executes a search for a match between a regular expression and this String object.

Syntax: `Search(str, regexp string) int`

- str - string
- regexp - A regular expression

```golang
str := "hey JudE"
fmt.Println(gstrings.Search(str, "[A-Z]")) // => 4
fmt.Println(gstrings.Search(str, "[.]"))   // => -1
```


## Includes
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/includes](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/includes)

Syntax: `Includes(str, searchString string[, position int]) bool`

- str - string
- searchString - A string to be searched for within str.
- position - [Optional], The position within the string at which to begin searching for searchString. (Defaults to 0.)

```golang
fmt.Println(gstrings.Includes("Blue Whale", "blue")) // => false
fmt.Println(gstrings.Includes("Blue Whale", "Blue")) // => true
fmt.Println(gstrings.Includes("To be", "To be", 1))  // => false
```


## Trim
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/Trim](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/Trim)

The **Trim()** method removes whitespace from both ends of a string. Whitespace in this context is all the whitespace characters (space, tab, no-break space, etc.) and all the line terminator characters (LF, CR, etc.).

Syntax: `Trim(str string) string`

```golang
fmt.Println(gstrings.Trim("   Hello world!   "))  // => Hello world!

str := `


abc
`
fmt.Println(gstrings.Trim(str))  // => abc
```




## TrimStart
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/TrimLeft](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/TrimLeft)

The **TrimStart()** method removes whitespace from the beginning of a string.

Syntax: `TrimStart(str string) string`


```golang
fmt.Println(gstrings.TrimStart(`

GOSH`))
// => "GOSH"


fmt.Println(gstrings.TrimStart("   Hello world!   "))
// => "Hello world!   "
```




## TrimEnd
[https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/trimEnd](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/trimEnd)

The **TrimEnd()** method removes whitespace from the end of a string. 

Syntax: `TrimEnd(str string) string`

```golang
fmt.Println(gstrings.TrimEnd("  GOSH  "))
// => "  GOSH"

fmt.Println(gstrings.TrimEnd(`GOSH

`))
// => "GOSH"
```








## Center
[https://www.w3schools.com/python/ref_string_center.asp](https://www.w3schools.com/python/ref_string_center.asp)

The **center()** method will center align the string, using a specified character (space is default) as the fill character.

Syntax: `Center(str string, width int[, fillChar rune]) string`
- width - Required. The length of the returned string
- fillChar - [Optional], The character to fill the missing space on each side. Default is " " (space)


```golang
fmt.Println(gstrings.Center("[GOSH]", 6)) // => [GOSH]

fmt.Println(gstrings.Center("[GOSH]", 7)) // => " [GOSH]"

fmt.Println(gstrings.Center("[GOSH]", 7, '*')) // => *[GOSH]

fmt.Println(gstrings.Center("[GOSH]", 8, '*')) // => *[GOSH]*
```



## PadStart
[https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/padStart](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/padStart)

The **PadStart()** method pads the current string with another string (multiple times, if needed) until the resulting string reaches the given length. The padding is applied from the start of the current string.

Syntax: `PadStart(str string, targetLength int[, padString string]) string`

- str - string
- targetLength - int
  - The length of the resulting string once the current str has been padded. If the value is less than **len(str)**, then str is returned as-is.
- padString - string, [Optional]
  - The string to pad the current str with. If **padString** is too long to stay within the **targetLength**, it will be truncated from the end. The default value is " ".


```golang
str := "abc"
fmt.Println(gstrings.PadStart(str, 10))           // => "       abc"
fmt.Println(gstrings.PadStart(str, 10, "foo"))    // => "foofoofabc"
fmt.Println(gstrings.PadStart(str, 6), "12345")   // => "   abc 12345"
fmt.Println(gstrings.PadStart(str, 8, "0"))       // => "00000abc"
fmt.Println(gstrings.PadStart(str, 1))            // => "abc"
```



## PadEnd
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/padEnd](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/padEnd)

The **PadEnd()** method pads the current string with a given string (repeated, if needed) so that the resulting string reaches a given length. The padding is applied from the end of the current string.


Syntax: `PadStart(str string, targetLength int[, padString string]) string`

- str - string
- targetLength - int
  - The length of the resulting string once the current str has been padded. If the value is lower than **len(str)**, the current string will be returned as-is.
- padString - string, [Optional]
  - The string to pad the current str with. If padString is too long to stay within targetLength, it will be truncated: for left-to-right languages the left-most part and for right-to-left languages the right-most will be applied. The default value for this parameter is " ".


```golang
str := "abc"
fmt.Println(gstrings.PadEnd(str, 10))          // => "abc"
fmt.Println(gstrings.PadEnd(str, 10, "foo"))   // => "abcfoofoof"
fmt.Println(gstrings.PadEnd(str, 6, "12345"))  // => "abc123"
fmt.Println(gstrings.PadEnd(str, 1))           // => "abc"
```





## IndexOf
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/indexOf](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/indexOf)

The **IndexOf()** method returns the index within the calling String object of the first occurrence of the specified value, starting the search at fromIndex. Returns -1 if the value is not found.

Syntax: `IndexOf(str, searchStr string, [, fromIndex int]) int`

```golang
str := "hello world"
fmt.Println(gstrings.IndexOf(str, ""))      // => 0
fmt.Println(gstrings.IndexOf(str, "", 0))   // => 0
fmt.Println(gstrings.IndexOf(str, "", 3))   // => 3
fmt.Println(gstrings.IndexOf(str, "", 8))   // => 8
fmt.Println(gstrings.IndexOf(str, "", 11))  // => 11
fmt.Println(gstrings.IndexOf(str, "", 13))  // => 11
fmt.Println(gstrings.IndexOf(str, "", 22))  // => 11

str = "Blue Whale"
fmt.Println(gstrings.IndexOf(str, "Blue"))      // => 0
fmt.Println(gstrings.IndexOf(str, "Blute"))     // => -1
fmt.Println(gstrings.IndexOf(str, "Whale", 0))  // => 5
fmt.Println(gstrings.IndexOf(str, "Whale", 5))  // => 5
fmt.Println(gstrings.IndexOf(str, "", -1))      // => 0
fmt.Println(gstrings.IndexOf(str, "", 9))       // => 9
fmt.Println(gstrings.IndexOf(str, "", 10))      // => 10
fmt.Println(gstrings.IndexOf(str, "", 11))      // => 10
fmt.Println(gstrings.IndexOf(str, "blue"))      // => -1
```




## LastIndexOf
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/lastIndexOf](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/lastIndexOf)

The **LastIndexOf()** method returns the index within the calling String object of the last occurrence of the specified value, searching backwards from fromIndex. Returns -1 if the value is not found.

Syntax: `LastIndexOf(str, searchStr string, [, fromIndex int]) int`

```golang
str := "canal"
fmt.Println(gstrings.LastIndexOf(str, "a"))      // => 3
fmt.Println(gstrings.LastIndexOf(str, "a", 2))   // => 1
fmt.Println(gstrings.LastIndexOf(str, "a", 0))   // => -1
fmt.Println(gstrings.LastIndexOf(str, "x"))      // => -1
fmt.Println(gstrings.LastIndexOf(str, "c", -5))  // => 0
fmt.Println(gstrings.LastIndexOf(str, "c", 0))   // => 0
fmt.Println(gstrings.LastIndexOf(str, ""))       // => 5
fmt.Println(gstrings.LastIndexOf(str, "", 2))    // => 2

fmt.Println(gstrings.LastIndexOf("Brave new world", "w"))      // => 10
fmt.Println(gstrings.LastIndexOf("Brave new world", "new"))    // => 6
```



## Shuffle
Syntax: `Shuffle(str string) string`

Shuffle string order：
```golang
fmt.Println(gstrings.Shuffle("世界你好 hello world"))  // => 世 rhldwe l好你oo界l
fmt.Println(gstrings.Shuffle("123456789"))            // => 648317529
fmt.Println(gstrings.Shuffle("abcdefghijk"))          // => hgebfdkcjai
```









---
# grandom
## Random
Syntax: `Random() float64`

The **Random()** function returns a floating-point, pseudo-random number in the range 0 to less than 1 (inclusive of 0, but not 1):

```golang
fmt.Println(grandom.Random()) // => 0.9023581008768102
fmt.Println(grandom.Random()) // => 0.0903486953039843
```



## RandBool
Syntax: `RandBool() bool`.

Randomly select one from `true` or `false`:
```golang
fmt.Println(grandom.RandBool()) // => false
fmt.Println(grandom.RandBool()) // => true
```



## RandInt
[https://www.w3schools.com/python/ref_random_randint.asp](https://www.w3schools.com/python/ref_random_randint.asp)

Syntax: `RandInt(min, max int) int`.

Return a number between `min` and `max` (both included):
```golang
fmt.Println(grandom.RandInt(-1000, 1000))  // => -352
fmt.Println(grandom.RandInt(0, 1000))      // => 191
fmt.Println(grandom.RandInt(-10000, -10))  // => -1356
```







## Shuffle
Syntax: `Shuffle([]string)`

Shuffle the list (rearrange the order of the list items):
```golang
l := []string{"a", "b", "c", "d", "e"}
grandom.Shuffle(l)
fmt.Println(l) // => [b e a c d]
```


## ShuffleInt
See [Shuffle](#Shuffle).

```golang
l := []int{1, 2, 3, 4, 5, 6}
grandom.ShuffleInt(l)
fmt.Println(l) // => [6 4 2 5 1 3]
```


## ShuffleAny
See [Shuffle](#Shuffle).

```golang
l := []interface{}{"a", false, true, 1, "e", 2, nil}
grandom.ShuffleAny(l)
fmt.Println(l) // => [<nil> 1 e a false true 2]
```



## Choice
[https://docs.python.org/zh-cn/3/library/random.html#random.choice](https://docs.python.org/zh-cn/3/library/random.html#random.choice)

`func Choice(str string) string`

Return a random element from a string:

```golang
str := "GOSH你好世界"
fmt.Println(grandom.Choice(str)) // => G
fmt.Println(grandom.Choice(str)) // => S
fmt.Println(grandom.Choice(str)) // => S
fmt.Println(grandom.Choice(str)) // => O
fmt.Println(grandom.Choice(str)) // => 好
```









