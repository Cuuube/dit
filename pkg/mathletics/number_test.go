package mathletics

import (
	"fmt"
	"testing"
)

func TestNumber(t *testing.T) {
	var i1 int64 = 900000000000000
	var i2 int64 = 2
	n1 := Number[int64]{i1}
	n2 := Number[int64]{i2}
	fmt.Println(n1.Add(n2))
	fmt.Println(n1.Minus(n2))
	fmt.Println(n1.Multi(n2))
	fmt.Println(n1.Divide(n2))

	// var f1 float32 = 3.3
	// n3 := Number[float32]{f1}
	// fmt.Println(n1.Add(n3)) // 格式不同，不能操作
}
