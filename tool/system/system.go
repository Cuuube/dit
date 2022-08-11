package system

import "github.com/Cuuube/dit/pkg/cmdio"

func RunCmd(cmd string, args ...string) {
	sysTool := NewSystemTool()
	switch cmd {
	case "":
		fallthrough
	case "overview":
		cmdio.PrintStruct(sysTool.Overview())
	default:
		cmdio.Println("暂不支持命令：", cmd)
	}
}
