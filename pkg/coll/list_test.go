package coll

import (
	"fmt"
	"sort"
	"testing"
)

func TestList(t *testing.T) {
	// 常规元素测试
	var li1 List[string] = []string{}
	fmt.Println("raw:", li1)
	li1.Append("a")
	fmt.Println("append a:", li1)
	li1.Append("b")
	fmt.Println("append b:", li1)
	li1.Prepend("+++")
	fmt.Println("prepend +++:", li1)
	fmt.Println("get 1:", li1.Get(1))
	fmt.Println("len:", li1.Len())
	li1.Pop()
	fmt.Println("pop:", li1)
	fmt.Println("len:", li1.Len())
	li1.Shift()
	fmt.Println("shift:", li1)
	li1.Set(0, "5")
	fmt.Println("set 0 5:", li1)
	// li1.Set(1, "y") // error
	// fmt.Println("set 1 y:", li1)
	fmt.Println("len:", li1.Len())
}

func TestListSort(t *testing.T) {
	// 常规元素排序
	var li1 List[string] = []string{}
	li1.Append("1")
	li1.Append("7")
	li1.Append("9")
	li1.Append("y")
	li1.Append("2")
	li1.Append("+")
	li1.Prepend("-")
	li1.Prepend("1")
	li1.Prepend(".")
	li1.Prepend("7")
	li1.Prepend("a")
	li1.Prepend("3")
	fmt.Println("len:", li1.Len(), "li1:", li1)
	sort.Sort(&li1)
	fmt.Println("after sort: len:", li1.Len(), "li1:", li1)

	// 指针元素排序
	var li2 List[*int] = []*int{}
	var (
		one   = 1
		two   = 2
		three = 3
	)
	li2.Append(&two)
	li2.Append(&one)
	li2.Append(&three)
	fmt.Println("len:", li2.Len(), "li2:", li2)
	sort.Sort(&li2)
	fmt.Println("after sort: len:", li2.Len(), "li2:", *li2.Get(0), *li2.Get(1), *li2.Get(2))

	// 多层指针元素排序排序
	var li3 List[**int] = []**int{}
	var (
		oneP   = &one
		twoP   = &two
		threeP = &three
	)
	li3.Append(&twoP)
	li3.Append(&oneP)
	li3.Append(&threeP)
	fmt.Println("len:", li3.Len(), "li3:", li3)
	sort.Sort(&li3)
	fmt.Println("after sort: len:", li3.Len(), "li3:", **li3.Get(0), **li3.Get(1), **li3.Get(2))
}
