# seq2linq
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/seq2linq.svg)](https://pkg.go.dev/github.com/solsw/seq2linq)

[<img src="https://api.gitsponsors.com/api/badge/img?id=427105928" height="20">](https://api.gitsponsors.com/api/badge/link?p=JuJstBNp7ndvJE51saddkORQC9tJKbhThJOER++0kJb1kqonUPOnKXTv2w4yRhJ9ukTgSIu3Uvj+vYYAKMdEQECKTFSCouvgBUkFNTNJ8aOJKxIwMtLdUqa8v2k+kPZy)

**seq2linq** is Go implementation of .NET's 
[LINQ to Objects](https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/linq-to-objects).
(See also: [Language Integrated Query](https://en.wikipedia.org/wiki/Language_Integrated_Query),
[LINQ](https://learn.microsoft.com/en-us/dotnet/csharp/programming-guide/concepts/linq/),
[Enumerable Class](https://learn.microsoft.com/dotnet/api/system.linq.enumerable).)

---

## Installation

```
go get github.com/solsw/seq2linq
```

## Examples

Examples of **seq2linq** usage are in the `Example...` functions in test files
(see [Examples](https://pkg.go.dev/github.com/solsw/seq2linq#pkg-examples)).

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
