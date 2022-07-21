package system

var _ SystemTool = (*LinuxSystemTool)(nil)

type LinuxSystemTool struct {
	BaseSystemTool
}

// TODO
func (sysTool *LinuxSystemTool) Overview() SystemOverview {
	overview := sysTool.BaseSystemTool.Overview()
	return overview
}
