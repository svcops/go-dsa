// Package bst
// @Description: 二分搜索树通用定义

package bst

// Action 遍历的Action
type Action func(int, string)

type BST interface {
	// IsEmpty 是否为空
	IsEmpty() bool

	// Size 大小
	Size() int

	// Get 获取值
	Get(int) (string, error)

	//
	// Contains
	//  @Description: 是否包含
	//  @param int
	//  @return bool
	//
	Contains(int) bool

	//
	// Max
	//  @Description: 获取最大值
	//  @return int
	//  @return string
	//  @return error
	//
	Max() (int, string, error)

	//
	// Min
	//  @Description: 获取最小值
	//  @return int
	//  @return string
	//  @return error
	//
	Min() (int, string, error)

	// Add
	//  @Description:  新增 k ,v
	//  @param int
	//  @param string
	//
	Add(int, string)

	// Delete
	//  @Description: 删除值
	//  @param int
	//  @return bool
	//  @return error
	//
	Delete(int) (bool, error)

	// IsBST 是否是二分搜索树
	IsBST() bool

	// Dfs
	//  @Description: 深度遍历
	//  @param Action
	//
	Dfs(Action)

	//
	// Bfs
	//  @Description:
	//  @param Action
	//
	Bfs(Action)

	//
	// PreOrder
	//  @Description: 前序排序
	//  @param Action
	//
	PreOrder(Action)

	//
	// InOrder
	//  @Description: 中序排序
	//  @param Action
	//
	InOrder(Action)

	//
	// PostOrder
	//  @Description: 后序排序
	//  @param Action
	//
	PostOrder(Action)

	//
	// IsBalanced
	//  @Description: 是否是平衡的
	//  @return bool
	//
	IsBalanced() bool

	//
	// SetDebug
	//  @Description: 设置debug打印
	//  @param debug
	//
	SetDebug(debug bool)
}

// TreeError 自定义tree异常
type TreeError string

func (treeError TreeError) Message() string {
	return string(treeError)
}

func (treeError TreeError) Error() string {
	return treeError.Message()
}
