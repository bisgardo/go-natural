package strcmp

import (
	"bytes"
	"strings"
	"testing"
)

// blackHole is a package-scope variable that benchmarks can
// assign a value to in order to prevent that value from being
// optimized away by the compiler.
var blackHole interface{}

func Benchmark__default_string_comparison_equal(b *testing.B) {
	var v int

	a54_1 := a(53) + "a"
	a54_2 := a(53) + "a"
	for n := 0; n < b.N; n++ {
		v = strings.Compare(a54_1, a54_2)
	}

	blackHole = v
}

func Benchmark__default_string_comparison_nonequal_first(b *testing.B) {
	var v int

	aa53 := "a" + a(53)
	ab53 := "b" + a(53)
	for n := 0; n < b.N; n++ {
		v = strings.Compare(aa53, ab53)
	}

	blackHole = v
}

func Benchmark__default_string_comparison_nonequal_last(b *testing.B) {
	var v int

	a53a := a(53) + "a"
	a53b := a(53) + "b"
	for n := 0; n < b.N; n++ {
		v = strings.Compare(a53a, a53b)
	}

	blackHole = v
}

func Benchmark__natural_letter_only_string_comparison_equal(b *testing.B) {
	var v int

	a54 := a(54)
	for n := 0; n < b.N; n++ {
		v = Natural(a54, a54)
	}

	blackHole = v
}

func Benchmark__natural_letter_only_string_comparison_nonequal_first(b *testing.B) {
	var v int

	aa53 := "a" + a(53)
	ba53 := "b" + a(53)
	for n := 0; n < b.N; n++ {
		v = Natural(aa53, ba53)
	}

	blackHole = v
}

func Benchmark__natural_letter_only_string_comparison_nonequal_last(b *testing.B) {
	var v int

	a53a := a(53) + "a"
	a53b := a(53) + "b"
	for n := 0; n < b.N; n++ {
		v = Natural(a53a, a53b)
	}

	blackHole = v
}

func Benchmark__natural_single_digit_number(b *testing.B) {
	var v int

	l := a(10) + "1" + a(10)
	r := a(10) + "2" + a(10)
	for n := 0; n < b.N; n++ {
		v = Natural(l, r)
	}

	blackHole = v
}

func Benchmark__natural_single_digit_number_and_larger_10_digit_number(b *testing.B) {
	var v int

	l := "1"
	r := "0000000002"
	for n := 0; n < b.N; n++ {
		v = Natural(l, r)
	}

	blackHole = v
}

func Benchmark__natural_single_digit_number_and_smalller_10_digit_number(b *testing.B) {
	var v int

	l := "0000000001"
	r := "2"
	for n := 0; n < b.N; n++ {
		v = Natural(l, r)
	}

	blackHole = v
}

func Benchmark__natural_small_single_digit_number_and_big_10_digit_number(b *testing.B) {
	var v int

	l := "2000000001"
	r := "3"
	for n := 0; n < b.N; n++ {
		v = Natural(l, r)
	}

	blackHole = v
}

func Benchmark__natural_small_single_digit_number_and_big_10_digit_number_shortcut(b *testing.B) {
	var v int

	l := "2000000001"
	r := "2"
	for n := 0; n < b.N; n++ {
		v = Natural(l, r)
	}

	blackHole = v
}

func a(count int) string {
	b := bytes.Buffer{}
	for i := 0; i < count; i++ {
		b.WriteByte('a')
	}
	return b.String()
}
