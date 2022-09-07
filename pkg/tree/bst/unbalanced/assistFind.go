package unbalanced

func (node *bstNode) get(k int) *bstNode {
	if node == nil {
		return nil
	}
	if k == node.k {
		return node
	} else if k < node.k {
		return node.left.get(k)
	} else {
		return node.right.get(k)
	}
}

func (node *bstNode) max() *bstNode {
	if node == nil {
		return nil
	}
	if node.right == nil {
		return node
	}
	return node.right.max()
}

func (node *bstNode) min() *bstNode {
	if node == nil {
		return nil
	}
	if node.left == nil {
		return node
	}
	return node.left.min()
}
