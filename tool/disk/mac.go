package disk

var _ DiskTool = (*MacOSDiskTool)(nil)

type MacOSDiskTool struct {
	BaseDiskTool
}

// // Overview 查看磁盘使用概览
// func (tool *MacOSDiskTool) Overview() {
// 	out, _ := cmdio.Exec("df", "-h")
// 	cmdio.Println(out)
// }
