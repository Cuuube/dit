package mathletics

import "fmt"

var _ IMatrix[int] = (*Matrix[int])(nil)

// 根据二维数据创建矩阵
func NewMatrixFrom[T RawNumber](data [][]T) Matrix[T] {
	return Matrix[T]{
		data: data,
		cols: len(data[0]),
		rows: len(data),
	}
}

// 根据多行数据创建矩阵
func NewMatrixFromRows[T RawNumber](data ...[]T) Matrix[T] {
	return Matrix[T]{
		data: data,
		cols: len(data[0]),
		rows: len(data),
	}
}

// 根据行列初始化矩阵
func NewMatrix[T RawNumber](rows, cols int) Matrix[T] {
	mat := make([][]T, rows)
	for i := 0; i < rows; i++ {
		mat[i] = make([]T, cols)
	}
	return Matrix[T]{
		data: mat,
		cols: cols,
		rows: rows,
	}
}

// 实数矩阵
type Matrix[T RawNumber] struct {
	data [][]T // 二维矩阵，是行的数组
	rows int   // 行数
	cols int   // 列数
}

func (m *Matrix[T]) Rows() int {
	return m.rows
}

func (m *Matrix[T]) Cols() int {
	return m.cols
}

func (m *Matrix[T]) Get(rowIdx, colIdx int) T {
	return m.data[rowIdx][colIdx]
}

func (m *Matrix[T]) Set(rowIdx, colIdx int, val T) {
	m.data[rowIdx][colIdx] = val
}

func (m *Matrix[T]) Copy() Matrix[T] {
	res := Matrix[T]{
		rows: m.Rows(),
		cols: m.Cols(),
	}
	lines := make([][]T, res.rows)
	for i := 0; i < m.Rows(); i++ {
		newRow := make([]T, res.Cols())
		copy(newRow, m.data[i])
		lines[i] = newRow
	}
	res.data = lines
	return res
}

// 矩阵+矩阵
func (m *Matrix[T]) MatrixAdd(target IMatrix[T]) (IMatrix[T], error) {
	if m.Cols() != target.Cols() || m.Rows() != target.Rows() {
		return nil, ErrMatrixFormatNotEqual
	}
	result := m.Copy()

	for rIdx := 0; rIdx < m.Rows(); rIdx++ {
		for cIdx := 0; cIdx < m.Cols(); cIdx++ {
			result.Set(rIdx, cIdx, m.Get(rIdx, cIdx)+target.Get(rIdx, cIdx))
		}
	}
	return &result, nil
}

// 矩阵-矩阵
func (m *Matrix[T]) MatrixMinus(target IMatrix[T]) (IMatrix[T], error) {
	if m.Cols() != target.Cols() || m.Rows() != target.Rows() {
		return nil, ErrMatrixFormatNotEqual
	}
	result := m.Copy()

	for rIdx := 0; rIdx < m.Rows(); rIdx++ {
		for cIdx := 0; cIdx < m.Cols(); cIdx++ {
			result.Set(rIdx, cIdx, m.Get(rIdx, cIdx)-target.Get(rIdx, cIdx))
		}
	}
	return &result, nil
}

// 矩阵*数字
func (m *Matrix[T]) MatrixMultiNumber(target T) (IMatrix[T], error) {
	result := m.Copy()

	for rIdx := 0; rIdx < m.Rows(); rIdx++ {
		for cIdx := 0; cIdx < m.Cols(); cIdx++ {
			result.Set(rIdx, cIdx, m.Get(rIdx, cIdx)*target)
		}
	}
	return &result, nil
}

// 矩阵*矩阵;  https://baike.baidu.com/item/%E7%9F%A9%E9%98%B5%E4%B9%98%E6%B3%95
func (m *Matrix[T]) MatrixMulti(target IMatrix[T]) (IMatrix[T], error) {
	// 矩阵A的列数必须等于矩阵B的行数 A.Cols() === B.Rows()
	if m.Cols() != target.Rows() {
		return m, ErrMatrixFormatNotEqual
	}

	// ret的行数为A的行数，列数为B的列数
	retRows := m.Rows()
	retCols := target.Cols()
	ret := NewMatrix[T](retRows, retCols)

	/*
		针对每格(r, c)，运算如下：
			Ret(r, c) = sum[i,j:0->A.Cols()]{A(r,i)*B(j, c)}
	*/

	// 迭代目标matrix
	for r := 0; r < m.Rows(); r++ {
		for c := 0; c < target.Cols(); c++ {
			var sum T
			for i := 0; i < m.Cols(); i++ {
				mx := m.Get(r, i)
				tx := target.Get(i, c)
				// sum += m.Get(r, i) * target.Get(j, c)
				sum += mx * tx
			}
			// 赋值
			ret.Set(r, c, sum)
		}
	}

	return &ret, nil
}

func (m *Matrix[T]) Reverse() (IMatrix[T], error) {
	result := Matrix[T]{
		cols: m.Rows(),
		rows: m.Cols(),
	}
	lines := make([][]T, m.Cols())
	for i := range lines {
		lines[i] = make([]T, m.Rows())
	}

	for i := 0; i < m.Rows(); i++ {
		for j := 0; j < m.Cols(); j++ {
			lines[j][i] = m.Get(i, j)
		}
	}
	result.data = lines

	return &result, nil
}

func (m *Matrix[T]) Add(args INumber[IMatrix[T]]) INumber[IMatrix[T]] {
	ret, err := m.MatrixAdd(args.Value())
	if err != nil {
		fmt.Println(err)
	}
	return ret
}
func (m *Matrix[T]) Minus(args INumber[IMatrix[T]]) INumber[IMatrix[T]] {
	ret, err := m.MatrixMinus(args.Value())
	if err != nil {
		fmt.Println(err)
	}
	return ret
}
func (m *Matrix[T]) Multi(args INumber[IMatrix[T]]) INumber[IMatrix[T]] {
	ret, err := m.MatrixMulti(args.Value())
	if err != nil {
		fmt.Println(err)
	}
	return ret
}
func (m *Matrix[T]) Divide(args INumber[IMatrix[T]]) INumber[IMatrix[T]] {
	fmt.Println("矩阵不支持除法!")
	return m
}

func (m *Matrix[T]) Value() IMatrix[T] {
	return m
}

// func MatrixAdd[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](m1, m2 IMatrix[T]) (IMatrix[T], error) {
// 	if m1.Cols() != m2.Cols() || m1.Rows() != m2.Rows() {
// 		return nil, ErrMatrixFormatNotEqual
// 	}
// }

// func MatrixMulti[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](m1, m2 IMatrix[T]) (IMatrix[T], error) {

// }
