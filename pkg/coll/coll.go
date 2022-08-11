// coll collection集合包工具
package coll

import "reflect"

// Contains 检查是否在列表中
func Contains[argType any](args []argType, target argType) bool {
	for i := range args {
		if reflect.DeepEqual(args[i], target) {
			return true
		}
	}
	return false
}

// Foreach 高级函数：遍历
func Foreach[argType any](args []argType, exec func(int, argType)) {
	for i := range args {
		exec(i, args[i])
	}
}

// Map 高级函数：映射
func Map[argType any, retType any](args []argType, exec func(int, argType) retType) []retType {
	ret := make([]retType, len(args))
	for i := range args {
		ret[i] = exec(i, args[i])
	}
	return ret
}

// Separate 高级函数：分离
func Separate[argType any](args []argType, exec func(int, argType) bool) (trueList []argType, falseList []argType) {
	for i := range args {
		if exec(i, args[i]) {
			trueList = append(trueList, args[i])
		} else {
			falseList = append(falseList, args[i])
		}
	}
	return trueList, falseList
}

// Some 高级函数：局部满足
func Some[argType any](args []argType, exec func(int, argType) bool) bool {
	for i := range args {
		if exec(i, args[i]) {
			return true
		}
	}
	return false
}

// Every 高级函数：全部满足
func Every[argType any](args []argType, exec func(int, argType) bool) bool {
	for i := range args {
		if !exec(i, args[i]) {
			return false
		}
	}
	return true
}

// Filter 高级函数：条件过滤
func Filter[argType any](args []argType, exec func(int, argType) bool) []argType {
	ret := make([]argType, 0)
	for i := range args {
		if exec(i, args[i]) {
			ret = append(ret, args[i])
		}
	}
	return ret
}

// Reduce 高级函数：累加。列表不可以为空
func Reduce[argType any](args []argType, exec func(sum, next argType) argType) argType {
	if len(args) <= 0 {
		panic("no initial value")
	}
	if len(args) == 1 {
		return args[0]
	}

	var sum argType
	for i := range args {
		if i == 0 {
			sum = args[i]
			continue
		}
		sum = exec(sum, args[i])
	}
	return sum
}

// IgnoreErr 只拿返回值，忽略error
func IgnoreErr[argType any](arg argType, err error) argType {
	return arg
}
