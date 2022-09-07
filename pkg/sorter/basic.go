package sorter

import "golang.org/x/exp/slices"

// type Iterm[T int | float32 | float64 | string] []T

// Iterm Reference: https://go.dev/doc/tutorial/generics
type Iterm interface {
	int | float32 | float64 | string
}

func Len[T Iterm](arr []T) int {
	return len(arr)
}

func Swap[T Iterm](arr []T, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func IsSorted[T Iterm](arr []T) bool {
	if arr == nil || len(arr) <= 1 {
		return true
	}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func Equals[T Iterm](before, sorted []T) bool {
	if before == nil || sorted == nil {
		return false
	}
	if len(before) != len(sorted) {
		return false
	}

	slices.Sort(before)

	for i := 0; i < len(before); i++ {
		if before[i] != sorted[i] {
			return false
		}
	}
	return true
}
