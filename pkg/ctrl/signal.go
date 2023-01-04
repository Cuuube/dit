package ctrl

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 终止前执行代码，并且可以设定延时关闭
func DosthBeforeQuitWithDelay(delaySeconds uint, do func()) {
	// 线程退出前先清理
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				AwaitDelaySeconds(delaySeconds)
				do()
				fallthrough
			default:
				os.Exit(0)
			}
		}
	}()
}

// 执行倒计时，倒计时结束前会阻塞当前goroutine
func AwaitDelaySeconds(delaySeconds uint) {
	if delaySeconds > 0 {
		fmt.Println("倒计时", delaySeconds, "秒")
		delaySeconds -= 1
	} else {
		return
	}

	for range time.NewTicker(time.Second).C {
		if delaySeconds > 0 {
			fmt.Println("倒计时", delaySeconds, "秒")
			delaySeconds--
		} else {
			return
		}
	}
}
