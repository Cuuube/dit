package system

// SystemOverview 系统概览结构体
type SystemOverview struct {
	Goos    string // 操作系统内核
	Arch    string // 系统架构
	SysName string // 系统版本名
	SysVer  string // 系统版本号
}

// SystemTool 系统工具
type SystemTool interface {
	Overview() SystemOverview
}

// NewSystemTool 根据环境创建系统工具
func NewSystemTool() SystemTool {
	var sysTool SystemTool
	if IsLinux() {
		sysTool = &LinuxSystemTool{}
	} else if IsWin() {
		sysTool = &WinSystemTool{}
	} else if IsMacOS() {
		sysTool = &MacOSSystemTool{}
	} else {
		sysTool = &BaseSystemTool{}
	}
	return sysTool
}

var _ SystemTool = (*BaseSystemTool)(nil)

type BaseSystemTool struct{}

func (*BaseSystemTool) Overview() SystemOverview {
	overview := SystemOverview{
		Goos: GetSysGoos(),
		Arch: GetGoArch(),
	}
	return overview
}
