package coll

import (
	"fmt"
	"strings"
	"testing"
)

func TestStream(t *testing.T) {
	raw := []string{"1", "2"}
	mapExec := func(arg string) string {
		return "Hello" + arg
	}
	filterExec := func(arg string) bool {
		return strings.HasSuffix(arg, "2")
	}

	ret := make([]string, 0)
	NewStream(raw).Map(mapExec).Filter(filterExec).Export(&ret)
	fmt.Println(raw, ret)

	raw2 := []int{1, 5, 2, 5, 8, 5, 4, 8, 0}
	ret2 := NewStream(raw2).Map(func(i int) int { return i * 10 }).
		Filter(func(i int) bool { return i >= 50 }).
		View(func(i int) { fmt.Println(i) }).
		Map(func(i int) int { return i / 10 }).Execute()
	fmt.Println(ret2)

}
