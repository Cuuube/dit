package coll

import (
	"reflect"
	"sort"
)

var _ Collection[any] = (*List[any])(nil)

// List 动态列表包装
type List[T any] []T

// 长度
func (li *List[T]) Len() int {
	return len(*li)
}

// 尾部插入
func (li *List[T]) Append(v T) {
	*li = append(*li, v)
}

// 尾部弹出
func (li *List[T]) Pop() T {
	v := (*li)[li.Len()-1]
	(*li) = (*li)[0 : li.Len()-1]
	return v
}

// 获取元素下标
func (li *List[T]) IndexOf(v T) int {
	for i := 0; i < li.Len(); i++ {
		if reflect.DeepEqual(li.Get(i), v) {
			return i
		}
	}
	return -1
}

// 是否包含某元素
func (li *List[T]) Contains(v T) bool {
	return li.IndexOf(v) >= 0
}

// 删除固定元素
func (li *List[T]) Remove(v T) {
	idx := li.IndexOf(v)
	if idx >= 0 {
		l := *li
		(*li) = ConcatList(l[0:idx], l[idx+1:])
	}
}

// 头部弹出
func (li *List[T]) Shift() T {
	v := (*li)[0]
	(*li) = (*li)[1:li.Len()]
	return v
}

// 头部插入
func (li *List[T]) Prepend(v T) {
	li.Append(v)
	for i := li.Len() - 1; i >= 1; i-- {
		li.Swap(i-1, i)
	}
}

// 获取元素位置
func (li *List[T]) Get(i int) T {
	return (*li)[i]
}

// 设置元素固定位置
func (li *List[T]) Set(i int, v T) {
	(*li)[i] = v
}

// 交换
func (li *List[T]) Swap(i, j int) {
	(*li)[i], (*li)[j] = (*li)[j], (*li)[i]
}

func (li *List[T]) SortBy(lessfunc func(i, j int) bool) {
	sort.Slice((*li), lessfunc)
}

// 是否小于
func (li *List[T]) Less(i, j int) bool {
	len := li.Len()
	if i >= len || j >= len {
		return false
	}

	iItem := (*li)[i]
	jItem := (*li)[j]
	iItemType := reflect.TypeOf(iItem)
	iItemValue := reflect.ValueOf(iItem)
	jItemValue := reflect.ValueOf(jItem)
	// ptrFlg := false // 指针允许寻址一次

	for {
		switch iItemType.Kind() {

		// 数字按照值比大小
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return iItemValue.Int() < jItemValue.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return iItemValue.Uint() < jItemValue.Uint()
		case reflect.Float32, reflect.Float64:
			return iItemValue.Float() < jItemValue.Float()

		// 字符串按照本身大小
		case reflect.String:
			return iItemValue.String() < jItemValue.String()

		// 指针按照内部的值
		case reflect.Pointer:
			iItemType = iItemType.Elem()
			iItemValue = iItemValue.Elem()
			jItemValue = jItemValue.Elem()

		default:
			// 其他的结构体、数组、函数等都不执行默认排序
			return false
		}
	}
}
