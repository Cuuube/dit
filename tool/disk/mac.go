package disk

import (
	"github.com/Cuuube/dit/internal/cmdio"
)

var _ DiskTool = (*MacOSDiskTool)(nil)

type MacOSDiskTool struct {
}

// Overview 查看磁盘使用概览
func (sysTool *MacOSDiskTool) Overview() {
	out, _ := cmdio.Exec("df", "-h")
	cmdio.Println(out)
}
