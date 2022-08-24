package mathletics

type Number[T RawNumber] struct {
	rawNumber T
}

func (num Number[T]) Value() T {
	return num.rawNumber
}

func (num Number[T]) Add(arg Number[T]) Number[T] {
	return Number[T]{rawNumber: num.rawNumber + arg.Value()}
}

func (num Number[T]) Minus(arg Number[T]) Number[T] {
	return Number[T]{rawNumber: num.rawNumber - arg.Value()}
}

func (num Number[T]) Multi(arg Number[T]) Number[T] {
	return Number[T]{rawNumber: num.rawNumber * arg.Value()}
}

func (num Number[T]) Divide(arg Number[T]) Number[T] {
	return Number[T]{rawNumber: num.rawNumber / arg.Value()}
}
