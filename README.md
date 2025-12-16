# seq2linq
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/seq2linq.svg)](https://pkg.go.dev/github.com/solsw/seq2linq)

**seq2linq** is Go implementation of .NET's 
[LINQ to Objects](https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/linq-to-objects)
over [sequences of pairs of values](https://pkg.go.dev/iter#Seq2).
(See also: [Language Integrated Query](https://en.wikipedia.org/wiki/Language_Integrated_Query),
[LINQ](https://learn.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/),
[Enumerable Class](https://learn.microsoft.com/dotnet/api/system.linq.enumerable).)

For Go implementation of
[LINQ to Objects](https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/linq-to-objects)
over [sequences of individual values](https://pkg.go.dev/iter#Seq)
see [**go2linq**](https://github.com/solsw/go2linq) package.

---

## Installation

```
go get github.com/solsw/seq2linq
```

### Quick and easy example:

```go
package main

import (
	"fmt"

	"github.com/solsw/iterhelper"
	"github.com/solsw/seq2linq"
)

func main() {
	seq2, _ := iterhelper.Var2[int, string](1, "one", 2, "two", 3, "three",
		4, "four", 5, "five", 6, "six", 7, "seven", 8, "eight", 9, "nine", 10, "ten")
	filter, _ := seq2linq.Where(
		seq2,
		func(i int, s string) bool { return i%3 == 2 || s[0] == 't' },
	)
	projection, _ := seq2linq.Select(
		filter,
		func(i int, s string) (string, int) { return string(s[0]) + string(s[len(s)-1]), i * i },
	)
	for s, i := range projection {
		fmt.Println(s, i)
	}
}
```

The previous code outputs the following:
```
to 4
te 9
fe 25
et 64
tn 100
```
