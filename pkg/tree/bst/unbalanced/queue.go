// Package unbalanced
// @Description: 队列辅助广度遍历

package unbalanced

type queue []*bstNode

func newQueue() *queue {
	q := make(queue, 0)
	return &q
}

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *queue) add(node *bstNode) {
	*q = append(*q, node)
}

func (q *queue) poll() *bstNode {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}
