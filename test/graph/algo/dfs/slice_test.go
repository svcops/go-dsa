// Package dfs
// @Description: 传值 不能 修改结构体的内容 ！！！
// @Description:  传指针 可以 修改结构体的内容 ！！！
package dfs

import "testing"

// TestAppendFailure
//
//	@Description:
//	@param t
func TestAppendFailure(t *testing.T) {
	ta := testAppend{
		flag:   false,
		values: []string{"hello"},
	}
	appendFailure(ta)
	t.Log("after append flag is  :", ta.flag)
	t.Log("after append values is  :", ta.values)

	// hello
}

// TestAppendSuccess
//
//	@Description:
//	@param t
func TestAppendSuccess(t *testing.T) {
	ta := testAppend{
		values: []string{"hello"},
	}
	appendSuccess(&ta)
	t.Log("after append flag is  :", ta.flag)
	t.Log("after append values is  :", ta.values)
	// hello world
}

// appendFailure
//
//	@Description: 传值不能修改结构体的内容
//	@param ta
func appendFailure(ta testAppend) {
	ta.flag = true
	ta.values = append(ta.values, "world")
}

// appendSuccess
//
//	@Description: 传指针可以修改结构体的内容
//	@param ta
func appendSuccess(ta *testAppend) {
	ta.flag = true
	ta.values = append(ta.values, "world")
}

type testAppend struct {
	flag   bool
	values []string
}
