<p align="center">
  <img src="media/logo.jpg" width="400">
  <p align="center">Golang utility library, With JavaScript / Python additional functions!</p>
  <p align="center">
    <img src="https://img.shields.io/github/go-mod/go-version/xjh22222228/gosh" />
    <img src="https://img.shields.io/github/v/release/xjh22222228/gosh" />
    <img src="https://img.shields.io/github/license/xjh22222228/gosh" />
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
- [strings](#strings)
  - [Reverse](#Reverse)
  - [StartsWith](#StartsWith)
  - [EndsWith](#EndsWith)
  - [Concat](#Concat)
  - [Slice](#Slice)
  - [Search](#Search)
  - [Includes](#Includes)
  - [Trim](#Trim)


# strings

## Reverse
Syntax: `Reverse(str string) string`

```golang
package main

import (
    "fmt"
    "github.com/xjh22222228/gosh/gstrings"
)

func main()  {
    v := gstrings.Reverse("Hello World")
    fmt.Println(v) // => dlroW olleH
}
```


## StartsWith
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/startsWith](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/startsWith)

Syntax: `StartsWith(str, searchString string[, position int]) bool`

- str - string
- searchString - The characters to be searched for at the start of this string.
- position - [Optional], The position in this string at which to begin searching for searchString. Defaults to 0.


```golang
package main

import (
    "fmt"
    "github.com/xjh22222228/gosh/gstrings"
)

func main()  {
    str := "Saturday night plans"
    fmt.Println(gstrings.StartsWith(str, "Sat")) // => true
    fmt.Println(gstrings.StartsWith(str, "Sat", 3)) // => false
}
```


## EndsWith
https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/endsWith

Syntax: `EndsWith(str, searchString string[, length int]) bool`

- str - string
- searchString - The characters to be searched for at the end of str.
- length - [Optional], If provided, it is used as the length of str. Defaults to str.length.

```golang
package main

import (
    "fmt"
    "github.com/xjh22222228/gosh/gstrings"
)

func main()  {
    str := "Cats are the best!"
    fmt.Println(gstrings.EndsWith(str, "best", 17)) // => true

    str2 := "Is this a question"
    fmt.Println(gstrings.EndsWith(str2, "?")) // => false
}
```


## Concat
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/concat](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/concat)

Syntax: `Concat(str2, [, ...strN]) string`

```golang
package main

import (
    "fmt"
    "github.com/xjh22222228/gosh/gstrings"
)

func main()  {
    fmt.Println(gstrings.Concat("a", "b", "c")) // => abc
}
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
package main

import (
    "fmt"
    "github.com/xjh22222228/gosh/gstrings"
    "math"
)

func main()  {
    str := "The quick brown fox jumps over the lazy dog."
    fmt.Println(gstrings.Slice(str, 31))      // => the lazy dog.
    fmt.Println(gstrings.Slice(str, 4, 19))   // => quick brown fox
    fmt.Println(gstrings.Slice(str, -4))      // => dog.
    fmt.Println(gstrings.Slice(str, -9, -5))  // => lazy
}
```


## Search
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/search](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/search)


Syntax: `Search(str, regexp string) (int, error)`

- str - string
- regexp - A regular expression

```golang
package main

import (
    "fmt"
    "github.com/xjh22222228/gosh/gstrings"
)

func main()  {
    str := "hey JudE"
    fmt.Println(gstrings.Search(str, "[A-Z]")) // => 4 <nil>
    fmt.Println(gstrings.Search(str, "[.]"))   // => -1 not found
}
```


## Includes
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/includes](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/includes)

Syntax: `Includes(str, searchString string[, position int]) bool`

- str - string
- searchString - A string to be searched for within str.
- position - [Optional], The position within the string at which to begin searching for searchString. (Defaults to 0.)

```golang
package main

import (
    "fmt"
    "github.com/xjh22222228/gosh/gstrings"
)

func main()  {
    fmt.Println(gstrings.Includes("Blue Whale", "blue")) // => false
    fmt.Println(gstrings.Includes("Blue Whale", "Blue")) // => true
    fmt.Println(gstrings.Includes("To be", "To be", 1))  // => false
}
```


## Trim
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/Trim](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/Trim)

The **Trim()** method removes whitespace from both ends of a string. Whitespace in this context is all the whitespace characters (space, tab, no-break space, etc.) and all the line terminator characters (LF, CR, etc.).

Syntax: `Trim(str string) string`

```golang
package main

import (
    "fmt"
    "github.com/xjh22222228/gosh/gstrings"
)

func main()  {
    fmt.Println(gstrings.Trim("   Hello world!   "))  // => Hello world!

    str := `


abc
`
	fmt.Println(gstrings.Trim(str))  // => abc
}
```




## LICENSE
[MIT](./LICENSE)

