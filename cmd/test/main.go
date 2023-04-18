package main

import (
	"time"

	"github.com/Cuuube/dit/pkg/cli"
)

func main() {
	// loading := cli.NewSimpleLoading("加载中", cli.WithClearAfterStop(true), cli.WithFrameDuration(time.Second/10))
	// go loading.Play()
	// time.Sleep(time.Second * 5)
	// loading.Stop()

	// ani := cli.NewTextWithProgressLoading("加载中……")
	// ani.Progress.Total = 100
	// // loading.
	// loading := cli.NewLoadingWith(ani, cli.WithClearAfterStop(false), cli.WithFrameDuration(time.Second/2))
	// go func() {
	// 	i := 1
	// 	for range time.NewTicker(time.Second).C {
	// 		ani.Progress.Current = i * 10
	// 		i++
	// 		if i >= 100 {
	// 			return
	// 		}
	// 	}
	// }()

	strs := []string{`-`, `\`, `|`, `/`}
	// strs := []string{`我`, `爱`, `北`, `京`, `天`, `安`, `门`}
	// strs := []string{`(◕‿◕)`, `(-‿-)`}
	// strs := []string{`◴`, `◷`, `◶`, `◵`}
	// strs := []string{`⠁`, `⠂`, `⠄`, `⠠`, `⠐`, `⠈`}
	// strs := []string{`⛤`, `⛧`}

	ani := cli.NewTextWithStringsLoopLoading("请等待 ", strs)

	loading := cli.NewLoadingWith(ani, cli.WithFrameDuration(time.Second/2))

	go loading.Play()
	time.Sleep(time.Second * 12)
	loading.Stop()
}
