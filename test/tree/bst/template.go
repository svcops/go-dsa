// Package bst
// @Description: 二分搜索树测试模板 duckType

package bst

import (
	"go-ads/pkg/tree/bst"
	"math/rand"
	"strconv"
	"testing"
)

// TestBstCreator 测试bst的创建器
type TestBstCreator func() bst.BST

// Prepare 函数式模板
func Prepare(bstCreator TestBstCreator, empty bool) bst.BST {
	// 准备一个空的不平衡的二分搜索树
	if empty {
		return bstCreator()
	} else {
		// 准备一个包含数据的不平衡的二分搜索树
		tree := bstCreator()
		tree.Add(2, "BBB")
		tree.Add(1, "AAA")
		tree.Add(3, "CCC")
		return tree
	}
}

func TestAdd(t *testing.T, creator TestBstCreator) {
	tests := []struct {
		data map[int]string
		size int
	}{
		{
			data: map[int]string{
				2: "BBB",
				1: "AAA",
				3: "CCC",
			},
			size: 3,
		},
		{
			data: map[int]string{
				2: "BBB",
				1: "AAA",
				3: "CCC",
				4: "DDD",
			},
			size: 4,
		},
	}

	for _, test := range tests {
		tree := Prepare(creator, true)

		for k, v := range test.data {
			tree.Add(k, v)
		}

		if tree.Size() != test.size {
			t.Errorf("expect size is (%d)"+
				"and got (%d)", test.size, tree.Size())
		}
	}
}

func TestGet(t *testing.T, creator TestBstCreator) {
	tree := Prepare(creator, true)
	tree.Add(2, "BBB")
	tree.Add(1, "AAA")
	tree.Add(3, "CCC")

	// replace
	tree.Add(2, "B")

	tests := []struct {
		k int
		v string
		e bool
	}{
		{
			k: 1,
			v: "AAA",
			e: false,
		},
		{
			k: 2,
			v: "B",
			e: false,
		},
		{
			k: 3,
			v: "CCC",
			e: false,
		},

		{
			k: 4,
			v: "",
			e: true,
		},
	}

	for _, test := range tests {
		get, err := tree.Get(test.k)
		t.Logf("Get (k = %d, v = %s ,err=%+v)\n", test.k, get, err)
		if test.e == (err == nil) || test.v != get {
			t.Errorf("except get (k=%d,v=%s,isError=%t) "+
				"and got (k=%d,v=%s,isError=%t)", test.k, test.v, test.e, test.k, get, test.e == (err == nil))
		}
	}
}

func TestIsBST(t *testing.T, creator TestBstCreator) {
	if isBST := Prepare(creator, false).IsBST(); isBST {
		t.Logf("是二分搜索树")
	} else {
		t.Errorf("不是二分搜索树")
	}
}

func traverseAction(t *testing.T) func(int, string) {
	return func(i int, s string) {
		t.Logf("k = %d , value = %s\n", i, s)
	}
}

func TestBfs(t *testing.T, creator TestBstCreator) {
	Prepare(creator, false).Bfs(traverseAction(t))
}

func TestDfs(t *testing.T, creator TestBstCreator) {
	Prepare(creator, false).Dfs(traverseAction(t))
}

func TestPreOrder(t *testing.T, creator TestBstCreator) {
	Prepare(creator, false).PreOrder(traverseAction(t))
}

func TestInOrder(t *testing.T, creator TestBstCreator) {
	Prepare(creator, false).InOrder(traverseAction(t))
}

func TestPostOrder(t *testing.T, creator TestBstCreator) {
	Prepare(creator, false).PostOrder(traverseAction(t))
}

func TestDelete(t *testing.T, creator TestBstCreator) {
	tree := Prepare(creator, true)

	keys := []int{2, 1, 4, 3, 5}

	for _, key := range keys {
		// strconv.Itoa(i int) : int -> string
		tree.Add(key, strconv.Itoa(key))
	}

	// tree.Add(2, "2")
	// tree.Add(1, "1")
	// tree.Add(4, "4")
	// tree.Add(3, "3")
	// tree.Add(5, "5")

	tree.Bfs(traverseAction(t))
	/*
	    2
	  /  \
	 1   4
	    / \
	   3   5
	*/

	/*
	    3
	  /  \
	 1   4
	    / \
	   nil   5
	*/
	b, _ := tree.Delete(2)
	t.Logf("Delete result is : %t\n", b)

	t.Logf("After Delete. Size is : %d\n", tree.Size())

	tree.Bfs(traverseAction(t))
}

// TestRandomDelete 测试随机删除
func TestRandomDelete(t *testing.T, creator TestBstCreator, size, max int) {
	tree := creator()

	data := RandomDataInit(tree, size, max)

	beforeSize := tree.Size()
	t.Logf("Before Delete.(data.len = %d) (tree.size = %d) \n", size, beforeSize)

	repeatCount := 0
	for i := 0; i < size; i++ {
		_, err := tree.Delete(data[i])
		if err != nil {
			// t.Logf("Delete Repeated value %d\n", data[i])
			repeatCount++
		}
	}
	t.Logf("After Delete. Size is %d\n", tree.Size())
	if (size - beforeSize - repeatCount) != 0 {
		t.Errorf("Delete Exist Error !\n")
	} else {
		t.Logf("Delete Success! "+
			"(beforeSize = %d,repeateCount = %d)\n",
			beforeSize, repeatCount)
	}
}

// RandomDataInit 随机数据初始化
func RandomDataInit(bst bst.BST, size, max int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		// data = append(data, rand.Intn(size*100))
		data[i] = rand.Intn(max)
	}

	for i := 0; i < size; i++ {
		bst.Add(data[i], strconv.Itoa(data[i]))
	}
	return data
}
