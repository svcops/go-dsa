// Package avl
// @Description: avl 节点平衡操作 快速记忆：砍子接父算高度

package avl

import "log"

// getHeight 辅助方法，获取节点高度
func getHeight(node *avlNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

// updateHeight 辅助方法，计算节点的高度
func (node *avlNode) updateHeight() {
	if getHeight(node.left) > getHeight(node.right) {
		node.height = getHeight(node.left) + 1
	} else {
		node.height = getHeight(node.right) + 1
	}
}

// getBalanceFactor 获取节点的平衡因子，左子树的高度 - 右子树的高度
func getBalanceFactor(node *avlNode) int {
	if node == nil {
		return 0
	}
	return getHeight(node.left) - getHeight(node.right)
}

// 节点自平衡操作
func (node *avlNode) selfBalance(debug bool) *avlNode {

	nodeBf := getBalanceFactor(node)
	if debug && (nodeBf == 2 || nodeBf == -2) {
		log.Printf("Lost balance : node's key is (%d): node's balanceFactor is %d\n", node.k, nodeBf)
	}

	// avl 树的特点就是在每次添加节点后都保持了平衡
	rtNode := node

	// rightBf := getBalanceFactor(node.right)

	// 左高右低
	if nodeBf == 2 {
		leftBf := getBalanceFactor(node.left)
		if leftBf == 1 || leftBf == 0 {
			rtNode = rtNode.rightRotate()

			if debug {
				log.Println("LL Rotate")
			} else {
				// 不debug打印的话，直接返回，不需要再走其他的判断条件了
				return rtNode
			}
		} else if leftBf == -1 {
			rtNode.left = rtNode.left.leftRotate()
			rtNode = rtNode.rightRotate()

			if debug {
				log.Println("LR Rotate")
			} else {
				return rtNode
			}
		} else {
			// 如果进入了这个循环说明 doBalance写的有问题
			log.Println("Error occurred in selfBalance!!!")
		}

		if debug {
			log.Printf("After selfBalance : node's key is (%d): node's balanceFactor is %d\n", node.k, getBalanceFactor(rtNode))
		}

	}

	// 右高左低
	if nodeBf == -2 {
		rightBf := getBalanceFactor(node.right)
		if rightBf == 0 || rightBf == -1 {
			rtNode = rtNode.leftRotate()

			if debug {
				log.Println("RR Rotate")
			} else {
				// 不debug打印的话，直接返回，不需要再走其他的判断条件了
				return rtNode
			}
		} else if rightBf == 1 {
			rtNode.right = rtNode.right.rightRotate()
			rtNode = rtNode.leftRotate()
			if debug {
				log.Println("RL Rotate")
			} else {
				return rtNode
			}
		} else {
			// 如果进入了这个循环说明 doBalance写的有问题
			log.Println("Error occurred in selfBalance!!!")
		}

		if debug {
			log.Printf("After selfBalance : node's key is (%d): node's balanceFactor is %d\n", node.k, getBalanceFactor(rtNode))
		}
	}

	return rtNode
}

/*
        x
       /\
      y  t4
     /\
    z  t3
   / \
   t1 t2

        y
     /    \
    z      x
   / \    / \
   t1 t2 t3 t4
*/
// 右旋转
func (node *avlNode) rightRotate() *avlNode {
	x := node

	y := x.left

	// 砍子节点
	x.left = y.right
	// 接父节点
	y.right = x
	// 重新计算高度
	x.updateHeight()
	y.updateHeight()
	return y
}

/*
     x
    / \
   t1  y
      / \
    t2   z
        / \
       t3 t4

        y
      /   \
     x     z
    / \   / \
   t1 t2 t3 t4
*/
// 左旋转
func (node *avlNode) leftRotate() *avlNode {
	x := node

	y := x.right

	// 砍子
	x.right = y.left
	// 接父
	y.left = x

	x.updateHeight()
	y.updateHeight()
	return y
}
