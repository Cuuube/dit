package system

import (
	"strings"

	"github.com/Cuuube/dit/internal/cmdio"
)

var _ SystemTool = (*LinuxSystemTool)(nil)

type LinuxSystemTool struct {
	BaseSystemTool
}

// Overview 查看系统概览
func (sysTool *LinuxSystemTool) Overview() SystemOverview {
	overview := sysTool.BaseSystemTool.Overview()

	sysTool.fillSysInfo(&overview)
	sysTool.fillCompInfo(&overview)

	return overview
}

// fillSysInfo 增加系统信息
func (sysTool *LinuxSystemTool) fillSysInfo(info *SystemOverview) {
	out, err := cmdio.Exec("lsb_release", "-a")
	if err != nil {
		cmdio.Println(err.Error())
		return
	}
	kvs := cmdio.SplitToDict(out, ":")
	info.Sys = kvs["Description"]
	info.Kernel = strings.Split(cmdio.ExecIgnoreErr("cat", "/proc/version"), "(")[0]
}

// fillCompInfo 增加电脑信息
func (sysTool *LinuxSystemTool) fillCompInfo(info *SystemOverview) {
	info.Hostname = cmdio.ExecIgnoreErr("hostname")
	info.User = cmdio.ExecIgnoreErr("whoami")
}

// func (sysTool *LinuxSystemTool) fillSysInfo2(info *SystemOverview) {
// 	out, err := cmdio.Exec("cat", "/etc/os-release")
// 	/*
// 		PRETTY_NAME="Debian GNU/Linux 9 (stretch)"
// 		NAME="Debian GNU/Linux"
// 		VERSION_ID="9"
// 		VERSION="9 (stretch)"
// 		VERSION_CODENAME=stretch
// 		ID=debian
// 		HOME_URL="https://www.debian.org/"
// 		SUPPORT_URL="https://www.debian.org/support"
// 		BUG_REPORT_URL="https://bugs.debian.org/"
// 	*/

// 	if err != nil {
// 		return
// 	}
// 	kvs := cmdio.SplitToDict(out, "=")
// 	info.SysName = kvs["NAME"]
// 	info.SysVer = kvs["VERSION"]
// }
