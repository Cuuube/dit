package disk

var _ DiskTool = (*LinuxDiskTool)(nil)

type LinuxDiskTool struct {
	BaseDiskTool
}

// // Overview 查看磁盘使用概览
// func (tool *LinuxDiskTool) Overview() {
// 	out, _ := cmdio.Exec("df", "-h")
// 	cmdio.Println(out)
// }
