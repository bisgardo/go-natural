# go-natural

go-natural a small library for performing "natural" string comparison in Go (and on the beach).
It consists of the single function `Natural` with signature:

```
Natural(left, right string) int
```

Like any other comparator, `Natural` returns
* zero if `left = right`,
* a negative value if `left < right`, and
* a positive value if `left > right`.

Equality means byte-for-byte equality as usual.

Inequality is also similar to the usual definition,
but with the exception that numbers are compared in their entirety - not by charcter/byte.

See [the tests](https://github.com/halleknast/go-natural/natural_test.go) for examples.

## Formal definition

The precise definition of the function is the following:

Write `left` and `right` as the longest common prefix
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
nor is it intended to be in the future.
