package main

import (
	"time"

	"github.com/Cuuube/dit/pkg/cli"
)

func main() {
	loading := cli.SimpleLoading("加载中")
	go loading.Play()
	time.Sleep(time.Second * 5)
	loading.Stop()

	go loading.Play()
	time.Sleep(time.Second * 3)
	loading.Stop()
}
