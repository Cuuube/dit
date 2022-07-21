package system

import (
	"runtime"
	"strings"
)

const (
	MacPrefix   = "darwin"
	BSDPrefix   = "freebsd"
	WinPrefix   = "windows"
	LinuxPrefix = "linux"
)

// one of darwin, freebsd, linux, and so on
func GetSysGoos() string {
	return runtime.GOOS
}

// one of 386, amd64, arm, s390x, and so on.
func GetGoArch() string {
	return runtime.GOARCH
}

func IsWin() bool {
	return strings.Contains(GetSysGoos(), WinPrefix)
}

func IsMacOS() bool {
	return strings.Contains(GetSysGoos(), MacPrefix)
}

func IsLinux() bool {
	return strings.Contains(GetSysGoos(), LinuxPrefix)
}

func IsBSD() bool {
	return strings.Contains(GetSysGoos(), BSDPrefix)
}
