// Package generic
// @Description: 泛型测试
// @Reference: https://github.com/mattn/go-generics-example

package generic

import (
	"fmt"
	"golang.org/x/exp/slices"
	"testing"
)

func Test_Slice(t *testing.T) {
	a := []string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
		"END",
	}
	a = a[:12]
	fmt.Println(cap(a)) // 13
	a = slices.Clip(a)
	fmt.Println(cap(a)) // 12

	a = []string{
		"foo",
		"foo",
		"bar",
		"baz",
		"foo",
		"baz",
	}
	a = slices.Compact(a)
	fmt.Println(a) // [foo bar baz foo baz]

	fmt.Println(slices.Contains(a, "zoo")) // false

	a = slices.Delete(a, 1, 3)
	fmt.Println(a) // [foo foo baz]

	fmt.Println(slices.Index(a, "baz")) // 2

	fmt.Println(cap(a)) // 6
	a = slices.Grow(a, 5)
	fmt.Println(cap(a)) // 12 WTF?

	a = slices.Insert(a, 0, "noo")
	fmt.Println(a) // [noo foo foo baz]
}

/* *****  *****  *****  *****  *****  *****  *****  *****  *****  ***** */

type addable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | complex64 | complex128 | string
}

func add[T addable](a, b T) T {
	return a + b
}

func Test_ADD(t *testing.T) {
	t.Log(add(1, 2))
	t.Log(add("foo", "bar"))
}

func makeComparableFunc[T comparable]() func(lhs, rhs T) bool {
	return func(lhs, rhs T) bool {
		return lhs == rhs
	}
}
func Test_Anonymous(t *testing.T) {
	equal := makeComparableFunc[int]()
	fmt.Println(equal(1, 2))
	fmt.Println(equal(2, 2))
}

/* *****  *****  *****  *****  *****  *****  *****  *****  *****  ***** */
