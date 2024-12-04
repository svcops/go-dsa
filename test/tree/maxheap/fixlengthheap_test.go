package maxheap

import (
	fixed "go-dsa/pkg/tree/maxheap/fixedlength"
	"math/rand"
	"testing"
)

const capacity = 10

func Test_Add(t *testing.T) {
	heap := fixed.CreateFixedLengthHeap(capacity)
	values := make([]float64, capacity+1)
	for i := 0; i < capacity+1; i++ {
		add, _ := heap.Add(values[i])
		t.Log("Add success ", add)
		if i == capacity {
			_, e2 := heap.Add(values[i])
			if e2 == nil {
				t.Errorf("method Add implementation error")
			} else {
				t.Log("Can't add |", e2.Error())
			}
		}
	}
	t.Log(heap)
}

func Test_Max(t *testing.T) {
	heap := fixed.CreateFixedLengthHeap(capacity + 1)

	for i := 0; i < capacity; i++ {
		// values[i] = float64(rand.Intn(capacity))
		v := float64(rand.Intn(capacity))
		ar, _ := heap.Add(v)
		t.Log("Add value", v, "; success = ", ar)
	}

	endV := float64(capacity + 1)
	ar, _ := heap.Add(endV)
	t.Log("Add end value", endV, "; success = ", ar)

	max, _ := heap.Max()
	t.Log("Max value is ", max)
}

func Test_ExtractMax(t *testing.T) {
	heap := fixed.CreateFixedLengthHeap(capacity)

	for i := 0; i < capacity; i++ {
		ar, _ := heap.Add(float64(i + 1))
		t.Log("Add value", i+1, "; success = ", ar)
	}

	for !heap.IsEmpty() {
		max, _ := heap.ExtractMax()
		t.Log("Extract Max; Value is :", max)
	}

	t.Log("After Extract , Size is", heap.Size())
}
