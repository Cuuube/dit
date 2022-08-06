package cmdio

import "fmt"

func Scan() string {
	var inp string
	_, err := fmt.Scanln(&inp)
	if err != nil {
		// Println("读取输入出错:", err)
		return ""
	}
	return inp
}
