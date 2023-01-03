package mathletics

import (
	"fmt"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	m1 := NewMatrix[int](3, 5)
	fmt.Println(m1)
	fmt.Println(m1.Rows(), m1.Cols())

	data := [][]int{
		{1, 0, 2},
		{-1, 3, 1},
	}

	m2 := NewMatrixFrom(data)
	fmt.Println(m2)
	fmt.Println(m2.Rows(), m2.Cols())

	m3 := NewMatrixFromRows(data...)
	fmt.Println(m3)
	fmt.Println(m3.Rows(), m3.Cols())

}
func TestMatrix(t *testing.T) {
	data := [][]int{
		{1, 0, 2},
		{-1, 3, 1},
	}
	m := NewMatrixFrom(data)
	fmt.Printf("%+v\n", m)
	fmt.Println(m.Rows(), m.Cols())

	m1 := m.Copy()
	fmt.Printf("%+v\n", m1)
	fmt.Println(m1.Rows(), m1.Cols())

	m2, err := m.MatrixAdd(&m)
	fmt.Printf("%+v\n", m2)
	fmt.Println(m2, err, m2.Rows(), m2.Cols())

	m3, err := m.MatrixMultiNumber(3)
	fmt.Printf("%+v\n", m3)
	fmt.Println(m3, err, m3.Rows(), m3.Cols())

	m4, err := m.Reverse()
	fmt.Printf("%+v\n", m4)
	fmt.Println(m4, err, m4.Rows(), m4.Cols())

	m5, err := m.MatrixMulti(m4)
	fmt.Printf("%+v\n", m5)
	fmt.Println(m5, err, m5.Rows(), m5.Cols())
}

func TestMatrixMulti(t *testing.T) {
	data1 := [][]int{
		{1, 0, 2},
		{-1, 3, 1},
		{1, 3, 1},
	}

	m := NewMatrixFrom(data1)

	// 乘以常量
	{
		fmt.Println(m.MatrixMultiNumber(2))
	}

	// 统计行和
	{
		m := NewMatrixFrom([][]int{
			{1},
			{1},
			{1},
		})
		fmt.Println(m.MatrixMulti(&m))
	}

	// 不变矩阵
	{
		m := NewMatrixFrom([][]int{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		})
		fmt.Println(m.MatrixMulti(&m))
	}

	// 缩放矩阵
	{
		m := NewMatrixFrom([][]int{
			{2, 0, 0},
			{0, 2, 0},
			{0, 0, 2},
		})
		fmt.Println(m.MatrixMulti(&m))
	}

}

func TestMatrixINumber(t *testing.T) {
	m1 := NewMatrixFrom([][]int{
		{1, 2, 3},
		{4, 5, 6},
	})
	fmt.Printf("m1: [%+v]\n", m1)
	var number1 INumber[IMatrix[int]] = &m1

	m2 := NewMatrixFrom([][]int{
		{2, 2, 2},
		{2, 2, 2},
	})
	fmt.Printf("m2: [%+v]\n", m2)
	var number2 INumber[IMatrix[int]] = &m2

	m3 := NewMatrixFrom([][]int{
		{1},
		{1},
		{1},
	})
	fmt.Printf("m3: [%+v]\n", m3)
	var number3 INumber[IMatrix[int]] = &m3

	fmt.Printf("m1 + m2: [%+v]\n", number1.Add(number2))
	fmt.Printf("m1 - m2: [%+v]\n", number1.Minus(number2))
	fmt.Printf("m1 * m2: [%+v]\n", number1.Multi(number2))
	fmt.Printf("m1 * m3: [%+v]\n", number1.Multi(number3))
	fmt.Printf("m1 / m2: [%+v]\n", number1.Divide(number2))
}
