package unbalanced

import (
	"go-dsa/pkg/tree/bst"
	"go-dsa/pkg/tree/bst/unbalanced"
	template "go-dsa/test/tree/bst"
	"testing"
)

func creator() bst.BST {
	return unbalanced.NewUnbalancedBST()
}

func Test_Add(t *testing.T) {
	template.TestAdd(t, creator)
}

func Test_Get(t *testing.T) {
	template.TestGet(t, creator)
}

func Test_IsBST(t *testing.T) {
	template.TestIsBST(t, creator)
}

func Test_Bfs(t *testing.T) {
	template.TestBfs(t, creator)
}

func Test_Dfs(t *testing.T) {
	template.TestDfs(t, creator)
}

func Test_PreOrder(t *testing.T) {
	template.TestPreOrder(t, creator)
}

func Test_InOrder(t *testing.T) {
	template.TestInOrder(t, creator)
}

func Test_PostOrder(t *testing.T) {
	template.TestPostOrder(t, creator)
}

func Test_Delete(t *testing.T) {
	template.TestDelete(t, creator)
}

func Test_RandomDelete(t *testing.T) {
	size := 10000 * 10
	max := size * 100
	template.TestRandomDelete(t, creator, size, max)
}
