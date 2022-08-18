package disk

import "github.com/Cuuube/dit/pkg/cli"

var _ DiskTool = (*WinDiskTool)(nil)

type WinDiskTool struct {
	BaseDiskTool
}

// Overview 查看磁盘使用概览
func (tool *WinDiskTool) Overview() {
	out, _ := cli.Exec("df", "-h")
	cli.Println(out)
}
