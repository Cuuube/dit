package coll

import (
	"fmt"
	"strings"
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

func TestContains(t *testing.T) {
	newStru := testStruct{"chua", 11}
	fmt.Printf("TestIncludes %+v, res: %v\n", v51, Contains(v51, 5))
	fmt.Printf("TestIncludes %+v, res: %v\n", v55, Contains(v55, v4))
	fmt.Printf("TestIncludes %+v, res: %v\n", v55, Contains(v55, testStruct{"lala", 1}))
	fmt.Printf("TestIncludes %+v, res: %v\n", v55, Contains(v55, newStru))
	fmt.Printf("TestIncludes %+v, res: %v\n", v56, Contains(v56, &v41))
	fmt.Printf("TestIncludes %+v, res: %v\n", v56, Contains(v56, &newStru))
}

func TestSome(t *testing.T) {
	newStru := testStruct{"chua", 11}
	fmt.Printf("TestSome %+v, res: %v\n", v51, Some(v51, func(i int, v int) bool {
		return v <= 1
	}))
	fmt.Printf("TestSome %+v, res: %v\n", v55, Some(v55, func(i int, v testStruct) bool {
		return v.Age < 50
	}))
	fmt.Printf("TestSome %+v, res: %v\n", v55, Some(v55, func(i int, v testStruct) bool {
		return v.Age > 100
	}))
	fmt.Printf("TestSome %+v, res: %v\n", v55, Some(v55, func(i int, v testStruct) bool {
		return v.Name == "Tony"
	}))
	fmt.Printf("TestSome %+v, res: %v\n", v56, Some(v56, func(i int, v *testStruct) bool {
		return v.Age == 11
	}))
	fmt.Printf("TestSome %+v, res: %v\n", v56, Some(v56, func(i int, v *testStruct) bool {
		return *v == newStru
	}))
}

func TestEvery(t *testing.T) {
	newStru := testStruct{"chua", 11}
	fmt.Printf("TestEvery %+v, res: %v\n", v51, Every(v51, func(i int, v int) bool {
		return v <= 1
	}))
	fmt.Printf("TestEvery %+v, res: %v\n", v55, Every(v55, func(i int, v testStruct) bool {
		return v.Age < 100
	}))
	fmt.Printf("TestEvery %+v, res: %v\n", v55, Every(v55, func(i int, v testStruct) bool {
		return v.Age > 0
	}))
	fmt.Printf("TestEvery %+v, res: %v\n", v55, Every(v55, func(i int, v testStruct) bool {
		return v.Name == "Tony"
	}))
	fmt.Printf("TestEvery %+v, res: %v\n", v56, Every(v56, func(i int, v *testStruct) bool {
		return v.Age == 11
	}))
	fmt.Printf("TestEvery %+v, res: %v\n", v56, Every(v56, func(i int, v *testStruct) bool {
		return *v == newStru
	}))
}

func TestFilter(t *testing.T) {
	newStru := testStruct{"chua", 11}
	fmt.Printf("TestFilter %+v, res: %v\n", v51, Filter(v51, func(i int, v int) bool {
		return v <= 1
	}))
	fmt.Printf("TestFilter %+v, res: %v\n", v55, Filter(v55, func(i int, v testStruct) bool {
		return v.Age < 100
	}))
	fmt.Printf("TestFilter %+v, res: %v\n", v55, Filter(v55, func(i int, v testStruct) bool {
		return v.Age > 0
	}))
	fmt.Printf("TestFilter %+v, res: %v\n", v55, Filter(v55, func(i int, v testStruct) bool {
		return v.Name == "Tony"
	}))
	fmt.Printf("TestFilter %+v, res: %v\n", v56, Filter(v56, func(i int, v *testStruct) bool {
		return v.Age == 90
	}))
	fmt.Printf("TestFilter %+v, res: %v\n", v56, Filter(v56, func(i int, v *testStruct) bool {
		return *v == newStru
	}))
}

func TestMap(t *testing.T) {
	fmt.Printf("TestMap %+v, res: %v\n", v52, Map(v52, func(i int, v int) int {
		return v + 1
	}))
	fmt.Printf("TestMap %+v, res: %v\n", v54, Map(v54, func(i int, v string) string {
		return "Hello! " + v
	}))
}

func TestSeparate(t *testing.T) {
	fmt.Printf("TestSeparate %+v, res: ", v53)
	fmt.Println(Separate(v53, func(i int, v int) bool {
		return v > 6
	}))
	fmt.Printf("TestSeparate %+v, res: ", v54)
	fmt.Println(Separate(v54, func(i int, v string) bool {
		return i <= 2
	}))
	fmt.Printf("TestSeparate %+v, res: ", v54)
	fmt.Println(Separate(v54, func(i int, v string) bool {
		return strings.HasPrefix(v, "T")
	}))
}

func TestReduce(t *testing.T) {
	fmt.Printf("TestReduce %+v, res: ", v53)
	fmt.Println(Reduce(v53, func(sum, next int) int {
		return sum + next
	}))
	fmt.Println(Reduce(v53, func(sum, next int) int {
		return sum - next
	}))
	fmt.Printf("TestReduce %+v, res: ", v54)
	fmt.Println(Reduce(v54, func(sum, next string) string {
		return sum + next
	}))
	fmt.Printf("TestReduce %+v, res: ", v55)
	fmt.Println(Reduce(v55, func(sum, next testStruct) testStruct {
		sum.Name += next.Name
		sum.Age += next.Age
		return sum
	}))
}
