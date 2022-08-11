package file

import "github.com/Cuuube/dit/pkg/cmdio"

func RunCmd(cmd string, args ...string) {
	tool := NewFileTool()
	switch cmd {
	case "mv":
		fallthrough
	case "move":
		tool.Move(args...)
	default:
		cmdio.Println("暂不支持命令：", cmd)
	}
}
