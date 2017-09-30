# go-natural

go-natural is a small library for performing "natural" string comparison in Go and on the beach.

# Install

```
go get github.com/halleknast/go-natural
```
or (preferrably)
```
glide get github.com/halleknast/go-natural
```

# Function

The library contains a single function `Natural` with signature:

```
Natural(left, right string) int
```

Like any comparator, `Natural` evaluates to
* zero if `left = right`,
* a negative value if `left < right`, and
* a positive value if `left > right`.

Natural comparison is similar to "ordinary" (i.e. character by character) string comparison.
The only difference is that numbers are compared in their entirety even if they span multiple charcters.
For example, `"2" < "10"` even though `"1" < "2"`.

See [the tests](https://github.com/halleknast/go-natural/blob/master/natural_test.go)
for a complete set of examples.

Natural comparison is useful whenever strings are sorted for readability.
It also allows correct comparison of e.g. semantic version strings.
For instance, `"v1.2.3" < "v1.10.1"` as one would expect.

## Formal definition

The precise definition of natural order is the following:

Strings are equal if and only if they consist of the same bytes in the same order.
If they are different, the order is determined using the following rule:

Write the compared strings `left` and `right` as their longest common prefix
concatenated by different (non-empty) suffixes:
```
left  = <prefix><suffix1>
right = <prefix><suffix2>
```

Extract any prefix number from the suffixes into a separate group,
such that none of the suffixes start with a number character:
```
left  = <prefix><number1><suffix1>
right = <prefix><number2><suffix2>
```

If both number groups are non-empty,
the string with the greater number is "greater".
If the numbers are equal but have a different number of leading zeros,
the number with the most leading zeros is "greater".

If either of the number groups are empty,
the order is simply defined according to the usual definition (i.e. byte comparison).

## Sorting

For sorting a string slice in natural order,
`Naturally` is provided as a convenience implementation of `sort.Interface`.

For example, the program
```
package main

import (
	"sort"

	"github.com/halleknast/go-natural"
)

func main() {
	ss := []string{
		"1.2.3",
		"0.2.3",
		"1.10.3",
		"1.2.10",
	}
	sort.Sort(strcmp.Naturally(ss))

	for _, s := range ss {
		println(s)
	}
}
```
prints
```
0.2.3
1.2.3
1.2.10
1.10.3
```

## Limitations

Strings are currently indexed as single bytes.
This means that multibyte characters will not compare correctly.
This can be (and might get) fixed by using `utf8.DecodeRuneInString`.

Dash is never interpreted as "minus",
nor is such a feature planned.
