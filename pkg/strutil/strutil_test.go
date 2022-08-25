package strutil

import (
	"fmt"
	"testing"

	"github.com/Cuuube/dit/pkg/tst"
)

func TestSubStr(t *testing.T) {
	fmt.Println(SubStr("1234567890", 5))
	fmt.Println(SubStr("1234567890", 1))
	fmt.Println(SubStr("1234567890", 9))
	fmt.Println(SubStr("1234567890", 10))
	fmt.Println(SubStr("1234567890", 11))
	fmt.Println(SubStr("1234567890", 0))
	fmt.Println(SubStr("1234567890", 44))
}

func TestPad(t *testing.T) {
	tu := tst.New(t)
	fmt.Println(PadPrefix("11", '0', 5))
	fmt.Println(PadSuffix("11", '0', 5))
	fmt.Println(PadPrefix("1111111", '0', 5))
	fmt.Println(PadSuffix("1111111", '0', 5))

	tu.MustEqual(PadPrefix("11", '0', 5), "00011")
	tu.MustEqual(PadPrefix("11", '0', 4), "00011")
}
