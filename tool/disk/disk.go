package disk

import "github.com/Cuuube/dit/pkg/cmdio"

func RunCmd(cmd string, args ...string) {
	diskTool := NewDiskTool()
	switch cmd {
	case "":
		fallthrough
	case "overview":
		diskTool.Overview()
	default:
		cmdio.Println("暂不支持命令：", cmd)
	}
}
