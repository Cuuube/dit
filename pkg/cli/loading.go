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

func NewSimpleLoading(text string, confFn ...LoadingConfFn) Loading {
	animation := NewTextWithDotLoading(text)

	return NewLoadingWith(animation, confFn...)
}

func NewLoadingWith(animation LoadingAnimation, confFn ...LoadingConfFn) Loading {
	signal := atomic.Value{}
	signal.Store(SigNotRunning)

	conf := DefaultLoadingConf()
	for _, fn := range confFn {
		fn(conf)
	}

	return &LoadingFrame{
		signal:    signal,
		animation: animation,

		conf: conf,
	}
}

var _ (Loading) = (*LoadingFrame)(nil)

// LoadingFrame 加载动画控制框架
type LoadingFrame struct {
	signal    atomic.Value
	animation LoadingAnimation

	conf *LoadingConf
}

func (obj *LoadingFrame) Play() {
	// 原子锁，拒掉多线程执行
	ok := obj.signal.CompareAndSwap(SigNotRunning, SigRunning)
	if !ok {
		return
	}

	obj.animation.Print()
	for range time.NewTicker(obj.conf.frameDuration).C {
		if obj.signal.Load() != SigRunning { // stoped
			return
		}
		obj.animation.Clear()
		obj.animation.Print()
	}
}

func (obj *LoadingFrame) Stop() {
	obj.signal.Store(SigNotRunning)
	if obj.conf.clearAfterStop {
		obj.animation.Clear()
	}
	obj.animation.Reset()
}

// func (obj *LoadingFrame) Pause() {
// }

// --- loading conf ---

type LoadingConfFn func(*LoadingConf)

type LoadingConf struct {
	clearAfterStop bool
	frameDuration  time.Duration
}

func DefaultLoadingConf() *LoadingConf {
	return &LoadingConf{
		clearAfterStop: true,
		frameDuration:  time.Second,
	}
}

func WithClearAfterStop(inp bool) LoadingConfFn {
	return func(lc *LoadingConf) {
		lc.clearAfterStop = inp
	}
}

func WithFrameDuration(inp time.Duration) LoadingConfFn {
	return func(lc *LoadingConf) {
		lc.frameDuration = inp
	}
}
