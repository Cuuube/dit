package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestMemoryCache(t *testing.T) {
	ctx := context.Background()
	mc := Memory{}
	mc.Init(ctx)

	mc.Set(ctx, "a", 123)
	fmt.Println(mc.Get(ctx, "a"))
	mc.Set(ctx, "a", 456)
	mc.Set(ctx, "a", 789)
	fmt.Println(mc.Get(ctx, "a"))
	mc.SetEx(ctx, "b", 222, 3)
	fmt.Println(mc.Get(ctx, "b"))

	time.Sleep(time.Second * 2)
	fmt.Println(mc.Get(ctx, "b"))
	mc.SetEx(ctx, "b", 333, 5)
	fmt.Println(mc.Get(ctx, "b"))

	time.Sleep(time.Second * 10)
	mc.Set(ctx, "b", "bbb")
	fmt.Println(mc.Get(ctx, "b"))
}

func TestAsyncMemoryCache(t *testing.T) {
	ctx := context.Background()
	mc := Memory{}
	mc.Init(ctx)

	go func() {
		ticker := time.NewTicker(time.Millisecond * 999)
		for range ticker.C {
			mc.Set(ctx, "a", "aaa")
		}
	}()

	go func() {
		ticker := time.NewTicker(time.Millisecond * 777)
		for range ticker.C {
			mc.SetEx(ctx, "b", "bbb", 1)
		}
	}()

	go func() {
		ticker := time.NewTicker(time.Second * 2)
		for range ticker.C {
			mc.SetEx(ctx, "c", "ccc", 1)
		}
	}()

	go func() {
		ticker := time.NewTicker(time.Millisecond * 765)
		for range ticker.C {
			fmt.Println(mc.Get(ctx, "a"))
			fmt.Println(mc.Get(ctx, "b"))
			fmt.Println(mc.Get(ctx, "c"))
		}
	}()

	select {}
}
