package ctrl

import (
	"fmt"
	"testing"
)

type testStruct struct {
	Name string
	Age  int
}

var (
	v1  = 1
	v11 = 2
	v2  = "str1"
	v21 = "str2"
	v3  = 'a'
	v31 = 'b'
	v4  = testStruct{"chua", 11}
	v41 = testStruct{"ha", 90}

	v51 = []int{1, 2, 3, 4, 5}
	v52 = []int{1, 2, 3, 5, 7}
	v53 = []int{1, 2, 3, 5, 8, 13, 21}
	v54 = []string{"Tom", "Anna", "Jim", "Zhang", "Tim"}
	v55 = []testStruct{v4, v41}
	v56 = []*testStruct{&v4, &v41}
)

func TestTernary(t *testing.T) {
	fmt.Println(Ternary(true, v1, v11))
	fmt.Println(Ternary(false, v1, v11))
	fmt.Println(Ternary(true, v2, v21))
	fmt.Println(Ternary(false, v2, v21))
	fmt.Println(Ternary(true, v3, v31))
	fmt.Println(Ternary(false, v3, v31))
	fmt.Println(Ternary(true, v4, v41))
	fmt.Println(Ternary(false, v4, v41))
}

func TestIfElse(t *testing.T) {
	IfElse(1 < 2, func() { fmt.Println("1<2") }, func() { fmt.Println("1 not <2") })
	IfElse(1 > 2, func() { fmt.Println("1>2") }, func() { fmt.Println("1 not >2") })
}
