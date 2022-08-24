package cache

import (
	"context"
	"errors"
	"sync"
	"time"
)

const (
	NeverExpiredTime = -1
)

var ErrNotFound = errors.New("not found")

type memoryValue struct {
	key         string
	value       interface{}
	expiredtime int64 // 过期时间点的时间戳

	pre  *memoryValue
	next *memoryValue
}

var _ Cache = (*Memory)(nil)

// 创建内存缓存
func NewMemoryCache(ctx context.Context) *Memory {
	cache := Memory{}
	cache.Init(ctx)
	return &cache
}

// 内存缓存
type Memory struct {
	data       map[string]*memoryValue
	dataRWLock sync.RWMutex

	expiredLinkedList *memoryValue
	expiredLock       sync.Mutex
}

// 开启过期循环检查
func (cache *Memory) startExpiredCheckLoop(ctx context.Context) {
	// fmt.Println("启动过期检查！")
	// i := 1
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		// fmt.Println("检测次数：", i)
		// i++

		node := cache.expiredLinkedList.next
		for node != nil {
			cur := node
			node = node.next

			now := time.Now().UnixMilli()
			if now > cur.expiredtime {
				// fmt.Printf("监测到过期key:%s, 过期时间：%d, cur time：%d\n", cur.key, cur.expiredtime, now)
				cur.expiredtime = NeverExpiredTime
				cache.simpleDel(ctx, cur)
			}
		}
	}
}

// 初始化
func (cache *Memory) Init(ctx context.Context) error {
	cache.data = map[string]*memoryValue{}
	cache.expiredLinkedList = &memoryValue{expiredtime: NeverExpiredTime}

	go cache.startExpiredCheckLoop(ctx)

	return nil
}

// 获取耽搁元素
func (cache *Memory) Get(ctx context.Context, key string) (interface{}, error) {
	v, found := cache.simpleGet(ctx, key)
	if !found {
		return nil, ErrNotFound
	}

	// 鉴别是否过期
	now := time.Now().UnixMilli()
	if v.expiredtime > 0 && now > v.expiredtime {
		// fmt.Println("已经过期：", key, v.expiredtime, now)
		// 如果过期，执行回收
		cache.simpleDel(ctx, v)
		return nil, ErrNotFound
	}
	return v.value, nil
}

// 设置单个元素
func (cache *Memory) Set(ctx context.Context, key string, value interface{}) error {
	v, found := cache.simpleGet(ctx, key)
	if !found {
		v = &memoryValue{key: key}
	}
	v.value = value
	v.expiredtime = NeverExpiredTime
	cache.simpleSet(ctx, key, v)
	cache.updateExpiredLinkedList(v)
	return nil
}

// 设置元素，并且设置过期秒数
func (cache *Memory) SetEx(ctx context.Context, key string, value interface{}, expiredSeconds int) error {
	expiredTime := time.Now().Add(time.Second * time.Duration(expiredSeconds)).UnixMilli()

	if expiredSeconds < 0 {
		expiredTime = NeverExpiredTime
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

// 移除单个元素
func (cache *Memory) Del(ctx context.Context, key string) error {
	v, found := cache.simpleGet(ctx, key)
	if !found {
		return nil
	}
	// 删除map
	cache.simpleDel(ctx, v)
	return nil
}

// map里读取单个元素
func (cache *Memory) simpleGet(ctx context.Context, key string) (*memoryValue, bool) {
	cache.dataRWLock.RLock()
	defer cache.dataRWLock.RUnlock()

	v, found := cache.data[key]
	return v, found
}

// map里设置单个元素
func (cache *Memory) simpleSet(ctx context.Context, key string, v *memoryValue) {
	cache.dataRWLock.Lock()
	defer cache.dataRWLock.Unlock()

	cache.data[key] = v
}

// map里删除单个元素
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

// 更新当前元素在过期链表里的状态
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

// 插入链表颈部
func appendToLinkedList(v *memoryValue, linkedListHead *memoryValue) {
	temp := linkedListHead.next
	linkedListHead.next = v
	v.pre = linkedListHead
	if temp != nil {
		temp.pre = v
		v.next = temp
	}
}
