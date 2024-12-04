package unbalanced

import "go-dsa/pkg/tree/bst"

// Dfs 深度遍历
func (tree *unbalancedBST) Dfs(ac bst.Action) {
	// 默认中序遍历
	tree.root.inOrder(ac)
}

// Bfs 广度遍历
func (tree *unbalancedBST) Bfs(ac bst.Action) {
	if tree.IsEmpty() {
		return
	}
	queue := newQueue()
	queue.add(tree.root)
	for !queue.isEmpty() {
		node := queue.poll()
		ac(node.k, node.v)
		if node.left != nil {
			queue.add(node.left)
		}
		if node.right != nil {
			queue.add(node.right)
		}
	}
}

// PreOrder 先序遍历
func (tree *unbalancedBST) PreOrder(ac bst.Action) {
	tree.root.preOrder(ac)
}

// InOrder 中序遍历
func (tree *unbalancedBST) InOrder(ac bst.Action) {
	tree.root.inOrder(ac)
}

// PostOrder 后序遍历
func (tree *unbalancedBST) PostOrder(ac bst.Action) {
	tree.root.postOrder(ac)
}
