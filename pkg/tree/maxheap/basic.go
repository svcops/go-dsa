package maxheap

type Heap interface {
	Size() int
	IsEmpty() bool
	Add(float64) (bool, error)
	Max() (float64, error)
	ExtractMax() (float64, error)
}

type HeapError string

func (h HeapError) Error() string {
	return string(h)
}
