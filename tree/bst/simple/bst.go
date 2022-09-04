package simple

type SimpleBST struct {
	size int
	root *node
}

// 是否为空
func (bst SimpleBST) IsEmpty() bool {
	return bst.size == 0
}

func (bst SimpleBST) Size() int {
	return bst.size
}

func (bst SimpleBST) Add(k int, v string) {

}

func (bst SimpleBST) Get(k int) (int, error) {
	return 0, nil
}

func (bst SimpleBST) Delete(k int) {

}

func (bst SimpleBST) IsBalanced() bool {
	return true
}
