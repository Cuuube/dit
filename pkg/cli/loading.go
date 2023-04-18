package cli

import (
	"sync/atomic"
	"time"
)

// Loading 加载中动画
type Loading interface {
	Play()
	Stop()
	// Pause()
}

const (
	SigNotRunning uint8 = iota
	SigRunning
)

func SimpleLoading(text string) Loading {
	style := NewTextWithDotLoading(text)

	return NewLoadingWithStyle(style)
}

func NewLoadingWithStyle(style LoadingAnimation) Loading {
	sig := atomic.Value{}
	sig.Store(SigNotRunning)

	return &LoadingFrame{
		sig:   sig,
		style: style,
	}
}

var _ (Loading) = (*LoadingFrame)(nil)

// LoadingFrame 加载动画控制框架
type LoadingFrame struct {
	sig   atomic.Value
	style LoadingAnimation
}

func (obj *LoadingFrame) Play() {
	// 原子锁，拒掉多线程执行
	ok := obj.sig.CompareAndSwap(SigNotRunning, SigRunning)
	if !ok {
		return
	}

	obj.style.Print()
	for range time.NewTicker(time.Second).C {
		if obj.sig.Load() != SigRunning { // stoped
			return
		}
		obj.style.Clear()
		obj.style.Print()
	}
}

func (obj *LoadingFrame) Stop() {
	obj.sig.Store(SigNotRunning)
	obj.style.Clear()
	obj.style.Reset()
}

// func (obj *LoadingFrame) Pause() {
// }
