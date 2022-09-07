package sort

import (
	"math/rand"
	"testing"
)

func generateArr(size int, simple bool, t *testing.T) []int {
	if simple {
		simpleArr := []int{5, 4, 3, 2, 1}
		t.Logf("generate slice(len=%d) success...", len(simpleArr))
		return simpleArr
	}

	data := make([]int, size)
	randMax := size * 1000
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(randMax)
	}
	t.Logf("generate slice(len=%d) success...", size)
	return data
}
