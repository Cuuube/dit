package mathletics

type Number[T RawNumber] struct {
	rawNumber T
}

var _ INumber[int] = Number[int]{rawNumber: 0}

func (num Number[T]) Value() T {
	return num.rawNumber
}

func (num Number[T]) Add(arg INumber[T]) INumber[T] {
	return Number[T]{rawNumber: num.Value() + arg.Value()}

}

func (num Number[T]) Minus(arg INumber[T]) INumber[T] {
	return Number[T]{rawNumber: num.Value() - arg.Value()}
}

func (num Number[T]) Multi(arg INumber[T]) INumber[T] {
	return Number[T]{rawNumber: num.Value() * arg.Value()}
}

func (num Number[T]) Divide(arg INumber[T]) INumber[T] {
	return Number[T]{rawNumber: num.Value() / arg.Value()}
}

// func (num Number[T]) Value() T {
// 	return num.rawNumber
// }

// func (num Number[T]) Add(arg Number[T]) Number[T] {
// 	return Number[T]{rawNumber: num.Value() + arg.Value()}
// }

// func (num Number[T]) Minus(arg Number[T]) Number[T] {
// 	return Number[T]{rawNumber: num.Value() - arg.Value()}
// }

// func (num Number[T]) Multi(arg Number[T]) Number[T] {
// 	return Number[T]{rawNumber: num.Value() * arg.Value()}
// }

// func (num Number[T]) Divide(arg Number[T]) Number[T] {
// 	return Number[T]{rawNumber: num.Value() / arg.Value()}
// }
