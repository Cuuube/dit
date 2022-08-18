package cli

import "fmt"

// Println std标准打印Println
func Println(a ...any) {
	fmt.Println(a...)
}

// Printf std标准序列化打印
func Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}
