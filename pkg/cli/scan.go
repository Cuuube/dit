package cli

import (
	"bufio"
	"fmt"
	"os"
)

// func Scan() string {
// 	var inp string
// 	_, err := fmt.Scanln(&inp)
// 	if err != nil {
// 		// Println("读取输入出错:", err)
// 		return ""
// 	}
// 	return inp
// }

func Scan() string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		if len(text) != 0 {
			return text
		} else {
			// exit if user entered an empty string
			break
		}

	}

	// handle error
	if scanner.Err() != nil {
		fmt.Println("read input error: ", scanner.Err())
	}
	return ""
}

func ReadInput() string {
	return Scan()
}
