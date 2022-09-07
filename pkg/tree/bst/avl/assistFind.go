package avl

func (node *avlNode) get(k int) *avlNode {
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

func (node *avlNode) max() *avlNode {
	if node == nil {
		return nil
	}
	if node.right == nil {
		return node
	}
	return node.right.max()
}

func (node *avlNode) min() *avlNode {
	if node == nil {
		return nil
	}
	if node.left == nil {
		return node
	}
	return node.left.min()
}
