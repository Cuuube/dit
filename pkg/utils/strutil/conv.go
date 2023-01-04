package strutil

import (
	"strconv"
	"strings"
)

func StringToBool(str string) bool {
	str = strings.Trim(str, " ")
	if str == "true" {
		return true
	}
	return false
}

// StringToUint
func StringToUint(str string) uint {
	ui, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0
	}
	return uint(ui)
}

// UintToString
func UintToString(num uint) string {
	str := strconv.FormatUint(uint64(num), 10)

	return str
}

func StringToLong(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func LongToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

func IntToString(num int) string {
	return strconv.Itoa(num)
}
