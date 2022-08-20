package cache

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrNotFound = errors.New("not found")

type memoryValue struct {
	key         string
	value       interface{}
	expiredtime int64 // 过期时间点的时间戳

	pre  *memoryValue
	next *memoryValue
}

type Memory struct {
	data       map[string]*memoryValue
	dataRWLock sync.RWMutex

	// quickExpiredLinkedList *memoryValue
	// slowExpiredLinkedList  *memoryValue
	expiredLinkedList *memoryValue
	expiredLock       sync.RWMutex
}

func (cache *Memory) startExpiredCheck(ctx context.Context) {
	fmt.Println("启动过期检查！")
	i := 1
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		fmt.Println("检测次数：", i)
		i++

		node := cache.expiredLinkedList.next
		for node != nil {
			cur := node
			node = node.next

			now := time.Now().UnixMilli()
			if now > cur.expiredtime {
				fmt.Printf("监测到过期key:%s, 过期时间：%d, cur time：%d\n", cur.key, cur.expiredtime, now)
				cur.expiredtime = -1
				cache.simpleDel(ctx, cur)
			}
		}
	}
}

func (cache *Memory) Init(ctx context.Context) error {
	cache.data = map[string]*memoryValue{}
	// cache.quickExpiredLinkedList = &memoryValue{expiredtime: 999999999999}
	// cache.slowExpiredLinkedList = &memoryValue{expiredtime: 999999999999}
	cache.expiredLinkedList = &memoryValue{expiredtime: -1}

	go cache.startExpiredCheck(ctx)

	return nil
}

func (cache *Memory) Get(ctx context.Context, key string) (interface{}, error) {
	v, found := cache.simpleGet(ctx, key)
	if !found {
		return nil, ErrNotFound
	}

	// 鉴别是否过期
	now := time.Now().UnixMilli()
	if v.expiredtime > 0 && now > v.expiredtime {
		fmt.Println("已经过期：", key, v.expiredtime, now)
		// 如果过期，执行回收
		cache.simpleDel(ctx, v)
		return nil, ErrNotFound
	}
	return v.value, nil
}

func (cache *Memory) Set(ctx context.Context, key string, value interface{}) error {
	v, found := cache.simpleGet(ctx, key)
	if !found {
		v = &memoryValue{key: key}
	}
	v.value = value
	v.expiredtime = -1
	cache.simpleSet(ctx, key, v)
	cache.updateExpiredLinkedList(v)
	return nil
}

func (cache *Memory) SetEx(ctx context.Context, key string, value interface{}, expiredSeconds int) error {
	expiredTime := time.Now().Add(time.Second * time.Duration(expiredSeconds)).UnixMilli()

	if expiredSeconds < 0 {
		expiredTime = -1
	}

	v, found := cache.simpleGet(ctx, key)
	if !found {
		v = &memoryValue{key: key}
	}
	v.value = value
	v.expiredtime = expiredTime
	cache.simpleSet(ctx, key, v)
	cache.updateExpiredLinkedList(v)
	return nil
}

func (cache *Memory) Del(ctx context.Context, key string) error {
	v, found := cache.simpleGet(ctx, key)
	if !found {
		return nil
	}
	// 删除map
	cache.simpleDel(ctx, v)
	return nil
}

func (cache *Memory) simpleGet(ctx context.Context, key string) (*memoryValue, bool) {
	cache.dataRWLock.RLock()
	defer cache.dataRWLock.RUnlock()

	v, found := cache.data[key]
	return v, found
}

func (cache *Memory) simpleSet(ctx context.Context, key string, v *memoryValue) {
	cache.dataRWLock.Lock()
	defer cache.dataRWLock.Unlock()

	cache.data[key] = v
}

func (cache *Memory) simpleDel(ctx context.Context, v *memoryValue) error {
	// 加锁
	cache.dataRWLock.Lock()
	defer cache.dataRWLock.Unlock()

	// 删除map
	delete(cache.data, v.key)
	// 链表中移除
	cache.updateExpiredLinkedList(v)
	return nil
}

func (cache *Memory) updateExpiredLinkedList(v *memoryValue) {
	cache.expiredLock.Lock()
	defer cache.expiredLock.Unlock()

	// 先移除val原来所在的链表位置
	if v.pre != nil {
		v.pre.next = v.next
	}
	// 如果不失效，不加入链表
	if v.expiredtime <= 0 {
		return
	}
	// 分快慢分别加入链表
	// if isQuick(v.expiredtime) {
	// 	appendToLinkedList(v, cache.quickExpiredLinkedList)
	// } else {
	// 	appendToLinkedList(v, cache.slowExpiredLinkedList)
	// }
	appendToLinkedList(v, cache.expiredLinkedList)
}

func appendToLinkedList(v *memoryValue, linkedListHead *memoryValue) {
	temp := linkedListHead.next
	linkedListHead.next = v
	v.pre = linkedListHead
	if temp != nil {
		temp.pre = v
		v.next = temp
	}
}
