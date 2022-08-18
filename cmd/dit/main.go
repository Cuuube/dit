package main

import (
	"os"

	"github.com/Cuuube/dit/pkg/cli"
	"github.com/Cuuube/dit/tool/disk"
	"github.com/Cuuube/dit/tool/file"
	"github.com/Cuuube/dit/tool/system"
)

var (
	// moduleArg       *string = flag.String("module", "", "Use -module <module>")
	// simpleModuleArg *string = flag.String("m", "", "Use -m <module>")
	// cmdArg          *string = flag.String("cmd", "", "Use -cmd <cmd>")
	// simpleCmdArg    *string = flag.String("c", "", "Use -c <cmd>")

	// real args
	module    string
	cmd       string
	otherArgs []string
)

func main() {
	// load args
	// loadArgsByFlags()
	loadArgsByArgs()

	switch module {
	case "sys":
		system.RunCmd(cmd, otherArgs...)
	case "disk":
		disk.RunCmd(cmd, otherArgs...)
	case "file":
		file.RunCmd(cmd, otherArgs...)

	// 特殊简化版alias
	case "get":
		file.RunCmd("fetch", os.Args[2:]...)
	case "fetch":
		file.RunCmd("fetch", os.Args[2:]...)
	default:
		cli.Println("暂不支持模块：", module)
	}
}

// // 根据flag包加载参数
// func loadArgsByFlags() {
// 	flag.Parse()

// 	if moduleArg != nil && *moduleArg != "" {
// 		module = *moduleArg
// 	} else if simpleModuleArg != nil && *simpleModuleArg != "" {
// 		module = *simpleModuleArg
// 	}
// 	if cmdArg != nil && *cmdArg != "" {
// 		cmd = *cmdArg
// 	} else if simpleCmdArg != nil && *simpleCmdArg != "" {
// 		cmd = *simpleCmdArg
// 	}
// }

// 根据参数顺序加载参数
func loadArgsByArgs() {
	// cli.Println(os.Args) // [./bin/dit sys overview]
	if module == "" {
		module = os.Args[1]
	}
	if cmd == "" && len(os.Args) > 2 {
		cmd = os.Args[2]
	}
	if len(os.Args) > 3 {
		otherArgs = os.Args[3:]
	}
}
