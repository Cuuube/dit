package mathletics

var _ IMatrix[int] = (*Matrix[int])(nil)

func NewMatrix[T RawNumber](data [][]T) Matrix[T] {
	return Matrix[T]{
		data: data,
		cols: len(data[0]),
		rows: len(data),
	}
}

func NewMatrixFromRows[T RawNumber](data ...[]T) Matrix[T] {
	return Matrix[T]{
		data: data,
		cols: len(data[0]),
		rows: len(data),
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

func (m *Matrix[T]) Add(target IMatrix[T]) (IMatrix[T], error) {
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

func (m *Matrix[T]) Minus(target IMatrix[T]) (IMatrix[T], error) {
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

func (m *Matrix[T]) MultiNumber(target T) (IMatrix[T], error) {
	result := m.Copy()

	for rIdx := 0; rIdx < m.Rows(); rIdx++ {
		for cIdx := 0; cIdx < m.Cols(); cIdx++ {
			result.Set(rIdx, cIdx, m.Get(rIdx, cIdx)*target)
		}
	}
	return &result, nil
}

// TODO https://baike.baidu.com/item/%E7%9F%A9%E9%98%B5%E4%B9%98%E6%B3%95
func (m *Matrix[T]) Multi(target IMatrix[T]) (IMatrix[T], error) {
	return m, ErrMatrixFormatNotEqual
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

// func MatrixAdd[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](m1, m2 IMatrix[T]) (IMatrix[T], error) {
// 	if m1.Cols() != m2.Cols() || m1.Rows() != m2.Rows() {
// 		return nil, ErrMatrixFormatNotEqual
// 	}
// }

// func MatrixMulti[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](m1, m2 IMatrix[T]) (IMatrix[T], error) {

// }
