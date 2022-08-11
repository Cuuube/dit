package mathletics

import (
	"fmt"
	"testing"
)

func TestMatrix(t *testing.T) {
	data := [][]int{
		{1, 0, 2},
		{-1, 3, 1},
	}
	m := NewMatrix(data)
	fmt.Printf("%+v\n", m)
	fmt.Println(m.Rows(), m.Cols())

	m1 := m.Copy()
	fmt.Printf("%+v\n", m1)
	fmt.Println(m1.Rows(), m1.Cols())

	m2, err := m.Add(&m)
	fmt.Printf("%+v\n", m2)
	fmt.Println(m2, err, m2.Rows(), m2.Cols())

	m3, err := m.MultiNumber(3)
	fmt.Printf("%+v\n", m3)
	fmt.Println(m3, err, m3.Rows(), m3.Cols())

	m4, err := m.Reverse()
	fmt.Printf("%+v\n", m4)
	fmt.Println(m4, err, m4.Rows(), m4.Cols())

}
