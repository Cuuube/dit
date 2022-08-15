package mathletics

type INumber interface {
	Add(INumber) INumber
	Minus(INumber) INumber
	Multi(INumber) INumber
	Divide(INumber) INumber
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}
