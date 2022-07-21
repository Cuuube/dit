package system

var _ SystemTool = (*WinSystemTool)(nil)

type WinSystemTool struct {
	BaseSystemTool
}

// TODO
func (sysTool *WinSystemTool) Overview() SystemOverview {
	overview := sysTool.BaseSystemTool.Overview()
	return overview
}
