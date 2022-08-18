package file

import "github.com/Cuuube/dit/pkg/cli"

func RunCmd(cmd string, args ...string) {
	tool := NewFileTool()
	switch cmd {
	case "mv":
		fallthrough
	case "move":
		tool.Move(args...)
	case "cp":
		fallthrough
	case "copy":
		tool.Copy(args...)
	case "get":
		fallthrough
	case "fetch":
		tool.Fetch(args...)
	default:
		cli.Println("暂不支持命令：", cmd)
	}
}
