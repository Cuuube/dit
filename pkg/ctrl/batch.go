package ctrl

import (
	"fmt"
	"sync"
)

// AsyncForeach 将参数异步执行
func AsyncForeach[T any](worker int, args []T, exec func(int, T)) {
	ch := make(chan struct{}, worker)
	defer close(ch)

	wg := sync.WaitGroup{}
	wg.Add(len(args))

	for i := range args {
		ch <- struct{}{}
		go func(idx int, v T) {
			exec(idx, v)
			wg.Done() // 标记任务完成数
			<-ch      // 释放池
		}(i, args[i])
	}
	// 任务全部结束后终止
	wg.Wait()
}

// ForeachBatch 将参数分批执行。串行执行。func(int: curbatch start index, []T: current batch)
func ForeachBatch[T any](args []T, batchSize int, exec func(int, []T)) {
	if len(args) <= 0 {
		return
	}
	startIdx := 0
	for startIdx < len(args) {
		step := batchSize
		// 最后一批
		if startIdx+batchSize > len(args) {
			step = len(args) - startIdx
		}
		// 默认步进
		exec(startIdx, args[startIdx:startIdx+step])
		startIdx += step
	}
}

// AsyncForeachBatch 分批异步执行
func AsyncForeachBatch[T any](worker int, args []T, batchSize int, exec func(int, []T)) {
	if len(args) <= 0 {
		return
	}

	ch := make(chan struct{}, worker)
	defer close(ch)

	wg := sync.WaitGroup{}

	startIdx := 0
	for startIdx < len(args) {
		step := batchSize
		// 最后一批
		if startIdx+batchSize > len(args) {
			step = len(args) - startIdx
		}
		wg.Add(step)

		ch <- struct{}{}
		go func(startIdx int, v []T) {
			cnt := len(v)
			exec(startIdx, v)
			// 标记任务完成数
			for i := 0; i < cnt; i++ {
				wg.Done()
			}
			<-ch // 释放池
		}(startIdx, args[startIdx:startIdx+step])

		startIdx += step
	}

	// 任务全部结束后终止
	wg.Wait()
}

// AsyncExec 并行执行
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

// AsyncExecWithCancel 将参数分批执行。串行执行
func AsyncExecWithCancel(worker int, execFuncs ...func() error) (chan error, func()) {
	errChan := make(chan error, len(execFuncs))
	if len(execFuncs) <= 0 {
		return errChan, func() {}
	}

	cancel := make(chan struct{})
	cancelFunc := func() {
		cancel <- struct{}{}
		close(cancel)
	}

	ch := make(chan struct{}, worker)
	defer close(ch)

	wg := sync.WaitGroup{}
	wg.Add(len(execFuncs))

	// 并发执行
	for i := range execFuncs {
		select {

		// 如果接收到正常信号，继续执行
		case ch <- struct{}{}:
			go func(i int) {
				exec := execFuncs[i]

				defer func() {
					wg.Done() // 标记任务完成数
					<-ch      // 释放池
				}()

				// 并发接收返回值
				ret := make(chan error)
				defer close(ret)

				// 执行任务
				go func() {
					defer func() {
						recover()
					}()
					ret <- exec()
				}()

				// 执行返回，或者被取消即中断
				select {
				case e := <-ret:
					// 正确返回
					errChan <- e

				case _, ok := <-cancel:
					// 接到取消信号时，强制返回

					if !ok {
						return
					}
				}
			}(i)

			// 如果接收到取消信号，取消掉之后所有执行
		case canceled, ok := <-cancel:
			fmt.Println(canceled, ok)
			break
		}
	}
	// 任务全部结束后终止
	wg.Wait()

	return errChan, cancelFunc
}
