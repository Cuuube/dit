package system

import (
	"fmt"
	"testing"
)

func TestSysSymbol(t *testing.T) {
	fmt.Println("GetSysGoos: ", GetSysGoos())
	fmt.Println("GetGoArch: ", GetGoArch())
	fmt.Println("IsWin: ", IsWin())
	fmt.Println("IsMacOS: ", IsMacOS())
	fmt.Println("IsLinux: ", IsLinux())
	fmt.Println("IsBSD: ", IsBSD())
}
