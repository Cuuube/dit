package coll

import (
	"sync"
)

// AsyncExec 将参数异步执行
func AsyncExec[T any](worker int, args []T, exec func(T)) {
	ch := make(chan struct{}, worker)
	defer close(ch)

	wg := sync.WaitGroup{}
	wg.Add(len(args))

	for i := range args {
		ch <- struct{}{}
		go func(v T) {
			exec(v)
			wg.Done() // 标记任务完成数
			<-ch      // 释放池
		}(args[i])
	}
	// 任务全部结束后终止
	wg.Wait()
}

// BatchExec 将参数分批执行。串行执行
func BatchExec[T any](args []T, batchSize int, exec func([]T)) {
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
		exec(args[startIdx : startIdx+step])
		startIdx += step
	}
}

// AsyncBatchExec 分批异步执行
func AsyncBatchExec[T any](worker int, args []T, batchSize int, exec func([]T)) {
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
		go func(v []T) {
			cnt := len(v)
			exec(v)
			// 标记任务完成数
			for i := 0; i < cnt; i++ {
				wg.Done()
			}
			<-ch // 释放池
		}(args[startIdx : startIdx+step])

		startIdx += step
	}

	// 任务全部结束后终止
	wg.Wait()
}
