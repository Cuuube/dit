package system

import (
	"strings"

	"github.com/Cuuube/dit/internal/cmdio"
)

var _ SystemTool = (*MacOSSystemTool)(nil)

type MacOSSystemTool struct {
	BaseSystemTool
}

// 查看系统概览
func (sysTool *MacOSSystemTool) Overview() SystemOverview {
	overview := sysTool.BaseSystemTool.Overview()

	sysTool.fillSysInfo(&overview)

	return overview
}

// 查看系统详细版本
func (sysTool *MacOSSystemTool) fillSysInfo(info *SystemOverview) {
	out, err := cmdio.Exec("sw_vers")
	if err != nil {
		return
	}

	// 小写化
	out = strings.ToLower(out)

	kvs := cmdio.ParseStrToDict(out)
	info.SysName = kvs["productname"]
	info.SysVer = kvs["productversion"]
}
