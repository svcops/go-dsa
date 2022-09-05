// Package bst
// @Description: 二分搜索树测试模板 duckType

package bst

import (
	"go-ads/pkg/tree/bst"
	"math/rand"
	"strconv"
	"testing"
)

type prepareFunc func() bst.BST

func testAdd(t *testing.T, prepareEmpty prepareFunc) {
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
		tree := prepareEmpty()

		for k, v := range test.data {
			tree.Add(k, v)
		}

		if tree.Size() != test.size {
			t.Errorf("expect size is (%d)"+
				"and got (%d)", test.size, tree.Size())
		}
	}
}

func testGet(t *testing.T, prepareEmpty prepareFunc) {
	tree := prepareEmpty()
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

func testIsBST(t *testing.T, prepare prepareFunc) {
	if isBST := prepare().IsBST(); isBST {
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

func testBfs(t *testing.T, prepare prepareFunc) {
	prepare().Bfs(traverseAction(t))
}

func testDfs(t *testing.T, prepare prepareFunc) {
	prepare().Dfs(traverseAction(t))
}

func testPreOrder(t *testing.T, prepare prepareFunc) {
	prepare().PreOrder(traverseAction(t))
}

func testInOrder(t *testing.T, prepare prepareFunc) {
	prepare().InOrder(traverseAction(t))
}

func testPostOrder(t *testing.T, prepare prepareFunc) {
	prepare().PostOrder(traverseAction(t))
}

func testDelete(t *testing.T, prepare prepareFunc) {
	tree := prepare()

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

func testDelete2(t *testing.T, prepareEmpty prepareFunc) {
	size := 10000
	data := make([]int, size)
	for i := 0; i < size; i++ {
		// data = append(data, rand.Intn(size*100))
		data[i] = rand.Intn(size * 100)
	}

	tree := prepareEmpty()

	for i := 0; i < size; i++ {
		tree.Add(data[i], strconv.Itoa(data[i]))
	}

	beforeSize := tree.Size()
	t.Logf("Before Delete. Size is %d\n", beforeSize)

	repeatCount := 0
	for i := 0; i < size; i++ {
		_, err := tree.Delete(data[i])
		if err != nil {
			t.Logf("Delete Repeated value %d\n", data[i])
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
