package util

import "fmt"

const (
	BYTE_LENGTH uint = 8
	MAX_UINT    uint = 4294967295
)

type BitMap struct {
	data []byte // 是否可以用指针，避免启动时吃大量内存？但使用时相应效率很低
	// length int
}

func (bm *BitMap) Init(maxLength uint) {
	if maxLength > MAX_UINT {
		panic(fmt.Errorf("不可超过最大值%d", MAX_UINT))
	}
	maxLeng := maxLength / BYTE_LENGTH
	if maxLength%BYTE_LENGTH != 0 {
		maxLeng++
	}
	bm.data = make([]byte, maxLeng)
	bm.logMemery()
}

func (bm *BitMap) Add(value uint) {
	index, mod := bm.getIndex(value)

	var newByt byte = 1
	newByt = newByt << mod
	bm.data[index] = bm.data[index] | newByt
}

func (bm *BitMap) Remove(value uint) {
	index, mod := bm.getIndex(value)

	var newByt byte = ^(1 << mod)
	bm.data[index] = bm.data[index] & newByt
}

func (bm *BitMap) Has(value uint) bool {
	index, mod := bm.getIndex(value)

	var newByt byte = bm.data[index] | (1 << mod)
	if bm.data[index] == newByt {
		return true
	}
	return false
}

func (bm *BitMap) getIndex(value uint) (uint, byte) {
	index := value / BYTE_LENGTH
	mod := value % BYTE_LENGTH
	return index, byte(mod)
}

func (bm *BitMap) logMemery() {
	length := len(bm.data)
	var num int = length
	var unit string = "byte"

	for num > 1024 {
		num = num / 1024
		switch unit {
		case "byte":
			unit = "KB"
		case "KB":
			unit = "MB"
		case "MB":
			unit = "GB"
		case "GB":
			unit = "TB"
		default:
			unit = ""
		}
	}
	// fmt.Println(fmt.Sprintf("bytes长度:%d，消耗内存%d%s", length, num, unit))
}
