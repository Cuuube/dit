package mathletics

import "errors"

var (
	ErrMatrixFormatNotEqual = errors.New("matrix type is not same!")
)

// 矩阵 https://baike.baidu.com/item/%E7%9F%A9%E9%98%B5/18069
type IMatrix[T RawNumber] interface {
	Cols() int       // 列数
	Rows() int       // 行数
	Get(int, int) T  // 获取值
	Set(int, int, T) // 设置值

	MatrixAdd(IMatrix[T]) (IMatrix[T], error)   // 矩阵加法
	MatrixMinus(IMatrix[T]) (IMatrix[T], error) // 矩阵减法
	MatrixMultiNumber(T) (IMatrix[T], error)    // 矩阵乘法
	MatrixMulti(IMatrix[T]) (IMatrix[T], error) // 矩阵乘法
	Reverse() (IMatrix[T], error)               // 转置

	// impliments INumber
	Add(INumber[IMatrix[T]]) INumber[IMatrix[T]]
	Minus(INumber[IMatrix[T]]) INumber[IMatrix[T]]
	Multi(INumber[IMatrix[T]]) INumber[IMatrix[T]]
	Divide(INumber[IMatrix[T]]) INumber[IMatrix[T]]
	Value() IMatrix[T]
}
