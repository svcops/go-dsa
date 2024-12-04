package graph

import (
	"go-dsa/pkg/graph"
	"reflect"
	"testing"
)

func Test_Set(t *testing.T) {

	set := graph.CreateSet()

	set.Add("hello")

	t.Log("auxset's size is :", set.Size())
	t.Log("auxset is empty :", set.IsEmpty())
	t.Log("contains hello :", set.Contains("hello"))
	t.Log("contains world :", set.Contains("world"))

	set.Delete("hello")

	t.Log("auxset's size is :", set.Size())
	t.Log("auxset is empty :", set.IsEmpty())
	t.Log("contains hello :", set.Contains("hello"))
	t.Log("contains world :", set.Contains("world"))

}

func Test_SetReflect(t *testing.T) {
	set := graph.CreateSet()
	set.Add("hello")

	t.Log("reflect.TypeOf(set)", reflect.TypeOf(set))
	t.Log("reflect.TypeOf(set).Kind() ", reflect.TypeOf(set).Kind())
	if reflect.TypeOf(set).Kind() == reflect.Map {
	}

}

type mapValue struct {
	flag bool
}

func Test_Map(t *testing.T) {
	m := make(map[string]map[string]float64)

	t.Log("m[hello] is nil", m["hello"] == nil)
	t.Log("m[hello] =", m["hello"])

	m2 := make(map[string]mapValue)
	hello := m2["hello"]
	hello.flag = true
	t.Log("m2[hello] is nil ? ", !hello.flag)
	t.Log("m2[hello] =", m2["hello"])

}
