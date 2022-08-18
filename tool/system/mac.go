package system

import (
	"github.com/Cuuube/dit/pkg/cli"
)

var _ SystemTool = (*MacOSSystemTool)(nil)

type MacOSSystemTool struct {
	BaseSystemTool
}

// 查看系统概览
func (sysTool *MacOSSystemTool) Overview() SystemOverview {
	overview := sysTool.BaseSystemTool.Overview()

	sysTool.fillSysInfo(&overview)
	sysTool.fillCompInfo(&overview)

	return overview
}

// fillSysInfo 增加系统信息
func (sysTool *MacOSSystemTool) fillSysInfo(info *SystemOverview) {
	// out, err := cli.Exec("sw_vers")
	out, err := cli.Exec("system_profiler", "SPSoftwareDataType")

	if err != nil {
		return
	}

	kvs := cli.SplitToDict(out, ":")
	info.Sys = kvs["System Version"]
	info.Kernel = kvs["Kernel Version"]
}

// fillCompInfo 增加电脑信息
func (sysTool *MacOSSystemTool) fillCompInfo(info *SystemOverview) {
	info.Hostname = cli.ExecIgnoreErr("hostname")
	info.User = cli.ExecIgnoreErr("whoami")
}
