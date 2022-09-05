package bst

type Action func(int, string)

type BST interface {
	IsEmpty() bool
	Size() int
	Add(int, string)
	Get(int) (string, error)
	Delete(int)

	// IsBST 是否是二分搜索树
	IsBST() bool

	Dfs(Action)

	Bfs(Action)

	PreOrder(Action)

	InOrder(Action)

	PostOrder(Action)
}
