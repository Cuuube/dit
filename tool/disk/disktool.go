package disk

import (
	"github.com/Cuuube/dit/pkg/cli"
	"github.com/Cuuube/dit/tool/system"
)

// DiskTool 系统工具
type DiskTool interface {
	Overview()
}

// NewDiskTool 根据环境创建系统工具
func NewDiskTool() DiskTool {
	var sysTool DiskTool
	if system.IsLinux() {
		sysTool = &LinuxDiskTool{}
	} else if system.IsWin() {
		sysTool = &WinDiskTool{}
	} else if system.IsMacOS() {
		sysTool = &MacOSDiskTool{}
	} else {
		sysTool = &BaseDiskTool{}
	}
	return sysTool
}

var _ DiskTool = (*BaseDiskTool)(nil)

type BaseDiskTool struct{}

// Overview 查看磁盘使用概览
func (tool *BaseDiskTool) Overview() {
	out, _ := cli.Exec("df", "-h")
	cli.Println(out)
}
