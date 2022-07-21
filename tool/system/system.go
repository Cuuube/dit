package system

import "github.com/Cuuube/dit/internal/cmdio"

func RunCmd(cmd string) {
	sysTool := NewSystemTool()
	switch cmd {
	case "overview":
		cmdio.PrintStruct(sysTool.Overview())
	default:
		cmdio.Println("暂不支持命令：", cmd)
	}
}
