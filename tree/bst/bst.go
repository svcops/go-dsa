package bst

type BST interface {
	IsEmpty() bool
	Size() int
	Add(int, string)
	Get(int) (int, error)
	Delete(int)
	IsBalanced() bool
}
