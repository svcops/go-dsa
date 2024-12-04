package skiplist

type SkipList interface {
	length() int
	max() int
	min() int
	add(int)
	delete(int)
}
