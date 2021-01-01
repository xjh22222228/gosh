<p align="center">
  <img src="media/logo.svg" width="350">
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
  - [TrimStart](#TrimStart)
  - [TrimEnd](#TrimEnd)
  - [Center](#Center)
  - [PadStart](#PadStart)
  - [PadEnd](#PadEnd)
- [random](#random)
  - [Random](#Random)


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




## TrimStart
[https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/TrimLeft](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/TrimLeft)

The **TrimStart()** method removes whitespace from the beginning of a string.

Syntax: `TrimStart(str string) string`


```golang
package main

import (
    "fmt"
    "github.com/xjh22222228/gosh/gstrings"
)

func main() {
    fmt.Println(gstrings.TrimStart(`

GOSH`))  // => "GOSH"


    fmt.Println(gstrings.TrimStart("   Hello world!   ")) // => "Hello world!   "
}
```




## TrimEnd
[https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/trimEnd](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/trimEnd)

The **TrimEnd()** method removes whitespace from the end of a string. 

Syntax: `TrimEnd(str string) string`

```golang
package main

import (
    "fmt"
    "github.com/xjh22222228/gosh/gstrings"
)

func main() {
    fmt.Println(gstrings.TrimEnd("  GOSH  "))  // => "  GOSH"

    fmt.Println(gstrings.TrimEnd(`GOSH

    `))  // => "GOSH"
}
```








## Center
[https://www.w3schools.com/python/ref_string_center.asp](https://www.w3schools.com/python/ref_string_center.asp)

The **center()** method will center align the string, using a specified character (space is default) as the fill character.

Syntax: `Center(str string, width int[, fillChar rune]) string`
- width - Required. The length of the returned string
- fillChar - [Optional], The character to fill the missing space on each side. Default is " " (space)


```golang
package main

import (
    "fmt"
    "github.com/xjh22222228/gosh/gstrings"
)

func main() {
    fmt.Println(gstrings.Center("[GOSH]", 6)) // => [GOSH]

    fmt.Println(gstrings.Center("[GOSH]", 7)) // => " [GOSH]"

    fmt.Println(gstrings.Center("[GOSH]", 7, '*')) // => *[GOSH]

    fmt.Println(gstrings.Center("[GOSH]", 8, '*')) // => *[GOSH]*
}
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
package main

import (
    "fmt"
    "github.com/xjh22222228/gosh/gstrings"
)

func main() {
    str := "abc"
    fmt.Println(gstrings.PadStart(str, 10))           // => "       abc"
    fmt.Println(gstrings.PadStart(str, 10, "foo"))    // => "foofoofabc"
    fmt.Println(gstrings.PadStart(str, 6), "12345")   // => "   abc 12345"
    fmt.Println(gstrings.PadStart(str, 8, "0"))       // => "00000abc"
    fmt.Println(gstrings.PadStart(str, 1))            // => "abc"
}
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
package main

import (
    "fmt"
    "github.com/xjh22222228/gosh/gstrings"
)

func main() {
    str := "abc"
    fmt.Println(gstrings.PadEnd(str, 10))          // => "abc"
    fmt.Println(gstrings.PadEnd(str, 10, "foo"))   // => "abcfoofoof"
    fmt.Println(gstrings.PadEnd(str, 6, "12345"))  // => "abc123"
    fmt.Println(gstrings.PadEnd(str, 1))           // => "abc"
}
```








---
# random
## Random
The **Random()** function returns a floating-point, pseudo-random number in the range 0 to less than 1 (inclusive of 0, but not 1).

```golang
package main

import (
    "fmt"
    "github.com/xjh22222228/gosh/grandom"
)

func main() {
    fmt.Println(grandom.Random()) // => 0.9023581008768102
    fmt.Println(grandom.Random()) // => 0.0903486953039843
}
```






## LICENSE
[MIT](./LICENSE)

