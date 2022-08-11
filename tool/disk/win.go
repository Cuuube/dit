package disk

import "github.com/Cuuube/dit/pkg/cmdio"

var _ DiskTool = (*WinDiskTool)(nil)

type WinDiskTool struct {
	BaseDiskTool
}

// Overview 查看磁盘使用概览
func (tool *WinDiskTool) Overview() {
	out, _ := cmdio.Exec("df", "-h")
	cmdio.Println(out)
}
