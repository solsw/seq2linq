# seq2linq — operator reference
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/seq2linq.svg)](https://pkg.go.dev/github.com/solsw/seq2linq)
[![GitHub](https://img.shields.io/badge/github--green?logo=github)](https://github.com/solsw/seq2linq)

**seq2linq** is a Go implementation of .NET's
[LINQ to Objects](https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/linq-to-objects)
over [sequences of pairs of values](https://pkg.go.dev/iter#Seq2) ([`iter.Seq2`](https://pkg.go.dev/iter#Seq2)).

This page is a categorized index of every operator in the package.
Full signatures, parameter descriptions and examples are on
[pkg.go.dev](https://pkg.go.dev/github.com/solsw/seq2linq).
For the getting-started guide and a runnable example, see the
[project README](https://github.com/solsw/seq2linq#readme).

For the same operators over [sequences of individual values](https://pkg.go.dev/iter#Seq)
([`iter.Seq`](https://pkg.go.dev/iter#Seq)) see the
[**go2linq**](https://github.com/solsw/go2linq) package.

## Conventions

Every operator is a generic function whose first parameter is (usually) the
input sequence and whose last result is an `error`:

```go
func Where[K, V any](seq2 iter.Seq2[K, V], pred func(K, V) bool) (iter.Seq2[K, V], error)
```

- **Deferred operators** return a new [`iter.Seq2`](https://pkg.go.dev/iter#Seq2);
  the input sequence is only iterated when the result is ranged over.
- **Immediate operators** consume the input sequence and return a value
  (a single pair, a `bool`, …).
- The trailing `error` reports invalid arguments — a `nil` input sequence,
  a `nil` predicate/selector, a negative count, an empty sequence where a value
  is required, and so on. The sentinel errors are exported (`ErrNilInput`,
  `ErrEmptyInput`, `ErrIndexOutOfRange`, `ErrNegativeCount`, `ErrNoMatch`,
  `ErrMultiElements`, …) so callers can match them with
  [`errors.Is`](https://pkg.go.dev/errors#Is).
- Operators named with an `Idx` suffix (`WhereIdx`, `SelectIdx`, …) pass the
  element's zero-based index to the callback.
- Operators named with an `Eq` suffix (`DistinctEq`, `ContainsEq`, …) take a
  custom equality function; the plain form relies on Go comparability.

## Filtering

| Operator | .NET | Description |
| --- | --- | --- |
| [`Where`](https://pkg.go.dev/github.com/solsw/seq2linq#Where) | `Where` | Filters pairs by a predicate. |
| [`WhereIdx`](https://pkg.go.dev/github.com/solsw/seq2linq#WhereIdx) | `Where` (indexed) | Filters pairs by a predicate that also receives the index. |
| [`OfType`](https://pkg.go.dev/github.com/solsw/seq2linq#OfType) | `OfType` | Keeps only the pairs assignable to the requested key/value types. |

## Projection

| Operator | .NET | Description |
| --- | --- | --- |
| [`Select`](https://pkg.go.dev/github.com/solsw/seq2linq#Select) | `Select` | Projects each pair into a new pair. |
| [`SelectIdx`](https://pkg.go.dev/github.com/solsw/seq2linq#SelectIdx) | `Select` (indexed) | Projects each pair, using its index. |
| [`SelectMany`](https://pkg.go.dev/github.com/solsw/seq2linq#SelectMany) | `SelectMany` | Projects each pair into a sequence and flattens the results. |
| [`SelectManyIdx`](https://pkg.go.dev/github.com/solsw/seq2linq#SelectManyIdx) | `SelectMany` (indexed) | Flattening projection that also receives the index. |
| [`Zip`](https://pkg.go.dev/github.com/solsw/seq2linq#Zip) | `Zip` | Merges two sequences pairwise using a result selector. |

## Partitioning

| Operator | .NET | Description |
| --- | --- | --- |
| [`Skip`](https://pkg.go.dev/github.com/solsw/seq2linq#Skip) | `Skip` | Skips the first *count* pairs. |
| [`SkipLast`](https://pkg.go.dev/github.com/solsw/seq2linq#SkipLast) | `SkipLast` | Skips the last *count* pairs. |
| [`SkipWhile`](https://pkg.go.dev/github.com/solsw/seq2linq#SkipWhile) | `SkipWhile` | Skips pairs while a predicate holds, then yields the rest. |
| [`SkipWhileIdx`](https://pkg.go.dev/github.com/solsw/seq2linq#SkipWhileIdx) | `SkipWhile` (indexed) | `SkipWhile` whose predicate receives the index. |
| [`Take`](https://pkg.go.dev/github.com/solsw/seq2linq#Take) | `Take` | Yields the first *count* pairs. |
| [`TakeLast`](https://pkg.go.dev/github.com/solsw/seq2linq#TakeLast) | `TakeLast` | Yields the last *count* pairs. |
| [`TakeWhile`](https://pkg.go.dev/github.com/solsw/seq2linq#TakeWhile) | `TakeWhile` | Yields pairs while a predicate holds. |
| [`TakeWhileIdx`](https://pkg.go.dev/github.com/solsw/seq2linq#TakeWhileIdx) | `TakeWhile` (indexed) | `TakeWhile` whose predicate receives the index. |

## Concatenation

| Operator | .NET | Description |
| --- | --- | --- |
| [`Concat`](https://pkg.go.dev/github.com/solsw/seq2linq#Concat) | `Concat` | Concatenates any number of sequences. |
| [`Append`](https://pkg.go.dev/github.com/solsw/seq2linq#Append) | `Append` | Appends a single pair to the end of a sequence. |
| [`Prepend`](https://pkg.go.dev/github.com/solsw/seq2linq#Prepend) | `Prepend` | Prepends a single pair to the start of a sequence. |

## Set operations

| Operator | .NET | Description |
| --- | --- | --- |
| [`Distinct`](https://pkg.go.dev/github.com/solsw/seq2linq#Distinct) | `Distinct` | Removes duplicate pairs (using Go comparability). |
| [`DistinctEq`](https://pkg.go.dev/github.com/solsw/seq2linq#DistinctEq) | `Distinct` | Removes duplicate pairs using a custom equality function. |
| [`DistinctBy`](https://pkg.go.dev/github.com/solsw/seq2linq#DistinctBy) | `DistinctBy` | Removes pairs with duplicate keys produced by a key selector. |
| [`DistinctByEq`](https://pkg.go.dev/github.com/solsw/seq2linq#DistinctByEq) | `DistinctBy` | `DistinctBy` with a custom key equality function. |

## Sorting

| Operator | .NET | Description |
| --- | --- | --- |
| [`OrderByLs`](https://pkg.go.dev/github.com/solsw/seq2linq#OrderByLs) | `OrderBy` | Stably sorts the pairs using a supplied *less* function. |
| [`Reverse`](https://pkg.go.dev/github.com/solsw/seq2linq#Reverse) | `Reverse` | Reverses the order of the pairs. |

## Quantifiers

| Operator | .NET | Description |
| --- | --- | --- |
| [`All`](https://pkg.go.dev/github.com/solsw/seq2linq#All) | `All` | Reports whether every pair satisfies a predicate. |
| [`Any`](https://pkg.go.dev/github.com/solsw/seq2linq#Any) | `Any` | Reports whether the sequence contains any pair. |
| [`AnyPred`](https://pkg.go.dev/github.com/solsw/seq2linq#AnyPred) | `Any` (predicate) | Reports whether any pair satisfies a predicate. |
| [`Contains`](https://pkg.go.dev/github.com/solsw/seq2linq#Contains) | `Contains` | Reports whether the sequence contains a given pair. |
| [`ContainsEq`](https://pkg.go.dev/github.com/solsw/seq2linq#ContainsEq) | `Contains` | `Contains` using a custom equality function. |

## Element operations

| Operator | .NET | Description |
| --- | --- | --- |
| [`ElementAt`](https://pkg.go.dev/github.com/solsw/seq2linq#ElementAt) | `ElementAt` | Returns the pair at a given index; errors if out of range. |
| [`ElementAtOrDefault`](https://pkg.go.dev/github.com/solsw/seq2linq#ElementAtOrDefault) | `ElementAtOrDefault` | Returns the pair at a given index or the zero pair. |
| [`First`](https://pkg.go.dev/github.com/solsw/seq2linq#First) | `First` | Returns the first pair; errors if empty. |
| [`FirstPred`](https://pkg.go.dev/github.com/solsw/seq2linq#FirstPred) | `First` (predicate) | Returns the first pair matching a predicate. |
| [`FirstOrDefault`](https://pkg.go.dev/github.com/solsw/seq2linq#FirstOrDefault) | `FirstOrDefault` | Returns the first pair or the zero pair. |
| [`FirstOrDefaultPred`](https://pkg.go.dev/github.com/solsw/seq2linq#FirstOrDefaultPred) | `FirstOrDefault` (predicate) | Returns the first matching pair or the zero pair. |
| [`Last`](https://pkg.go.dev/github.com/solsw/seq2linq#Last) | `Last` | Returns the last pair; errors if empty. |
| [`LastPred`](https://pkg.go.dev/github.com/solsw/seq2linq#LastPred) | `Last` (predicate) | Returns the last pair matching a predicate. |
| [`LastOrDefault`](https://pkg.go.dev/github.com/solsw/seq2linq#LastOrDefault) | `LastOrDefault` | Returns the last pair or the zero pair. |
| [`LastOrDefaultPred`](https://pkg.go.dev/github.com/solsw/seq2linq#LastOrDefaultPred) | `LastOrDefault` (predicate) | Returns the last matching pair or the zero pair. |
| [`Single`](https://pkg.go.dev/github.com/solsw/seq2linq#Single) | `Single` | Returns the only pair; errors if empty or if more than one. |
| [`SinglePred`](https://pkg.go.dev/github.com/solsw/seq2linq#SinglePred) | `Single` (predicate) | Returns the only pair matching a predicate. |
| [`SingleOrDefault`](https://pkg.go.dev/github.com/solsw/seq2linq#SingleOrDefault) | `SingleOrDefault` | Returns the only pair or a supplied default. |
| [`SingleOrDefaultPred`](https://pkg.go.dev/github.com/solsw/seq2linq#SingleOrDefaultPred) | `SingleOrDefault` (predicate) | Returns the only matching pair or a supplied default. |
| [`SingleOrZero`](https://pkg.go.dev/github.com/solsw/seq2linq#SingleOrZero) | `SingleOrDefault` | Returns the only pair or the zero pair. |

## Generation

| Operator | .NET | Description |
| --- | --- | --- |
| [`Repeat`](https://pkg.go.dev/github.com/solsw/seq2linq#Repeat) | `Repeat` | Produces a sequence that repeats one pair *count* times. |
| [`DefaultIfEmpty`](https://pkg.go.dev/github.com/solsw/seq2linq#DefaultIfEmpty) | `DefaultIfEmpty` | Yields the source, or a single zero pair if it is empty. |
| [`DefaultIfEmptyDef`](https://pkg.go.dev/github.com/solsw/seq2linq#DefaultIfEmptyDef) | `DefaultIfEmpty` (default) | Yields the source, or a single supplied pair if it is empty. |
| [`InfiniteSequence`](https://pkg.go.dev/github.com/solsw/seq2linq#InfiniteSequence) | — | Produces an unbounded sequence from a start pair and a *next* function. |

## Type conversion

| Operator | .NET | Description |
| --- | --- | --- |
| [`Cast`](https://pkg.go.dev/github.com/solsw/seq2linq#Cast) | `Cast` | Casts each pair to the requested key/value types. |
| [`OfType`](https://pkg.go.dev/github.com/solsw/seq2linq#OfType) | `OfType` | Filters to the pairs assignable to the requested key/value types. |

---

See also:
[Language Integrated Query](https://en.wikipedia.org/wiki/Language_Integrated_Query),
[LINQ](https://learn.microsoft.com/dotnet/csharp/programming-guide/concepts/linq/),
[Enumerable Class](https://learn.microsoft.com/dotnet/api/system.linq.enumerable).
