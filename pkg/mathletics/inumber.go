package mathletics

type INumber[T any] interface {
	Add(INumber[T]) INumber[T]
	Minus(INumber[T]) INumber[T]
	Multi(INumber[T]) INumber[T]
	Divide(INumber[T]) INumber[T]
	Value() T
}
type RawNumber interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}
