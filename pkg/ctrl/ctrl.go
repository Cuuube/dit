package ctrl

// Ternary 三元运算符 根据控制逻辑，返回值
func Ternary[retType any](condition bool, trueVal, falseVal retType) retType {
	if condition {
		return trueVal
	}
	return falseVal
}

// IfElse 控制逻辑
func IfElse(condition bool, trueFunc, elseFunc func()) {
	if condition {
		trueFunc()
		return
	}
	elseFunc()
}

// IgnoreErr 只拿返回值，忽略error
func IgnoreErr[argType any](arg argType, err error) argType {
	return arg
}
