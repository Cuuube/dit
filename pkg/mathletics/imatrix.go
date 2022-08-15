package mathletics

import "errors"

var (
	ErrMatrixFormatNotEqual = errors.New("matrix type is not same!")
)

// 矩阵 https://baike.baidu.com/item/%E7%9F%A9%E9%98%B5/18069
type IMatrix[T Number] interface {
	Cols() int                            // 列数
	Rows() int                            // 行数
	Get(int, int) T                       // 获取值
	Set(int, int, T)                      // 设置值
	Add(IMatrix[T]) (IMatrix[T], error)   // 矩阵加法
	Minus(IMatrix[T]) (IMatrix[T], error) // 矩阵减法
	MultiNumber(T) (IMatrix[T], error)    // 矩阵乘法
	Multi(IMatrix[T]) (IMatrix[T], error) // 矩阵乘法
	Reverse() (IMatrix[T], error)         // 转置
}
