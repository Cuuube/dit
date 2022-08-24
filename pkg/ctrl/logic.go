package ctrl

import "sync"

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

// AsyncExec 将参数分批执行。串行执行
func AsyncExec(worker int, execFuncs ...func() error) []error {
	errs := make([]error, len(execFuncs))
	if len(execFuncs) <= 0 {
		return errs
	}

	ch := make(chan struct{}, worker)
	defer close(ch)

	wg := sync.WaitGroup{}
	wg.Add(len(execFuncs))

	for i := range execFuncs {
		ch <- struct{}{}
		go func(i int) {
			exec := execFuncs[i]
			errs[i] = exec()
			wg.Done() // 标记任务完成数
			<-ch      // 释放池
		}(i)
	}
	// 任务全部结束后终止
	wg.Wait()

	return errs
}
