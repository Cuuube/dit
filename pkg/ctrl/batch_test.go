package ctrl

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestIgnoreErr(t *testing.T) {
	testFunc := func() (int, error) {
		return 111, errors.New("err!")
	}
	fmt.Println(testFunc())
	fmt.Println(IgnoreErr(testFunc()))
}

func TestAsyncForeach(t *testing.T) {
	args := make([]string, 100)
	for i := 0; i < 100; i++ {
		args[i] = strconv.Itoa(i)
	}
	AsyncForeach(17, args, func(idx int, v string) {
		time.Sleep(time.Second)
		fmt.Println(idx, v)
	})
}

func TestForeachBatch(t *testing.T) {
	size := 100
	args := make([]string, size)
	for i := 0; i < size; i++ {
		args[i] = strconv.Itoa(i)
	}
	ForeachBatch(args, 17, func(idx int, v []string) {
		// time.Sleep(time.Second)
		fmt.Println(idx, v[0], "~", v[len(v)-1])
	})
}

func TestBatchAsyncForeach(t *testing.T) {
	size := 1000
	args := make([]string, size)
	for i := 0; i < size; i++ {
		args[i] = strconv.Itoa(i)
	}
	AsyncForeachBatch(5, args, 142, func(idx int, v []string) {
		// time.Sleep(time.Second)
		fmt.Println(idx, v[0], "~", v[len(v)-1])
		time.Sleep(time.Second)
	})
}

func TestAsyncExec(t *testing.T) {
	// 循环打印goroutine数
	fmt.Println(runtime.NumGoroutine())
	go func() {
		t := time.NewTicker(time.Second)
		for range t.C {
			fmt.Println(runtime.NumGoroutine())
		}
	}()
	fmt.Println(runtime.NumGoroutine())

	execs := make([]func() error, 0)
	for i := 0; i < 100; i++ {
		idx := i
		execs = append(execs, func() error {
			time.Sleep(time.Millisecond * 100 * time.Duration(idx))
			return fmt.Errorf("err: %d", idx)
		})
	}
	// AsyncExec(7, func() error {
	// 	time.Sleep(time.Millisecond * 100)
	// })
	errs := AsyncExec(17, execs...)
	fmt.Println(errs)
}

func TestAsyncExecWithCancel(t *testing.T) {
	execs := make([]func() error, 100)
	for i := 0; i <= 100; i++ {
		idx := i
		execs = append(execs, func() error {
			fmt.Println("start", idx)
			time.Sleep(time.Second * time.Duration(i))
			fmt.Println("end", idx)
			return errors.New(fmt.Sprint(i))
		})
	}

	erChan, cancel := AsyncExecWithCancel(2, execs...)
	defer close(erChan)
	for {
		select {
		case er, ok := <-erChan:
			fmt.Println(er, ok)

			if !ok {
				return
			}
		}
		go func() {
			time.Sleep(3)
			cancel()
		}()
	}

}
