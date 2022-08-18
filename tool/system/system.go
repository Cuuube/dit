package system

import "github.com/Cuuube/dit/pkg/cli"

func RunCmd(cmd string, args ...string) {
	sysTool := NewSystemTool()
	switch cmd {
	case "":
		fallthrough
	case "overview":
		cli.PrintStruct(sysTool.Overview())
	default:
		cli.Println("暂不支持命令：", cmd)
	}
}
