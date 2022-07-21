package disk

import "github.com/Cuuube/dit/internal/cmdio"

var _ DiskTool = (*WinDiskTool)(nil)

type WinDiskTool struct {
}

// Overview 查看磁盘使用概览
func (sysTool *WinDiskTool) Overview() {
	out, _ := cmdio.Exec("df", "-h")
	cmdio.Println(out)
}
