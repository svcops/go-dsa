// Package avl
// @Description: 队列辅助广度遍历

package avl

type queue []*avlNode

func newQueue() *queue {
	q := make(queue, 0)
	return &q
}

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *queue) add(node *avlNode) {
	*q = append(*q, node)
}

func (q *queue) poll() *avlNode {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}
