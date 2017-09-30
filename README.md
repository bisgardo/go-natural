# go-natural

go-natural a small library for performing "natural" string comparison in Go (and on the beach).

# Install

```
go get github.com/halleknast/go-natural
```
or (preferrably)
```
glide get github.com/halleknast/go-natural
```

# Function

It consists of the single function `Natural` with signature:

```
Natural(left, right string) int
```

Like any other comparator, `Natural` returns
* zero if `left = right`,
* a negative value if `left < right`, and
* a positive value if `left > right`.

Natural comparison is similar to "ordinary" (i.e. byte to byte) comparison.
The only difference is that numbers are compared in their entirety even if they span multiple charcters.
For example, `"2" < "10"` even though `"1" < "2"`.


See [the tests](https://github.com/halleknast/go-natural/blob/master/natural_test.go)
for a complete set of examples.

## Formal definition

The precise definition of natural comparison is the following:

Strings are equal if and only if they consist of the same bytes in the same order.
If they are different, the order is determined as follows:

Write the compared strings `left` and `right` as the longest common prefix
followed by (non-empty) suffixes:
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
the string with the greater number is "larger".
If the numbers are equal but have a different number of leading zeros,
the number with the most leading zeros is "larger".

If either of the number groups are empty,
the order is simply defined according to the usual definition (i.e. byte comparison).

## Limitation

Strings are currently compared byte per byte.
This means that multibyte characters will not compare correctly.
This can (and might) be fixed by using `utf8.DecodeRuneInString`.

Dash is never interpreted as "minus",
nor is such a feature planned.
