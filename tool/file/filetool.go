package file

import (
	"regexp"
	"strings"

	"github.com/Cuuube/dit/pkg/cmdio"
	"github.com/Cuuube/dit/pkg/fileio"
)

// FileTool 系统工具
type FileTool interface {
	Move(args ...string)
}

// NewFileTool 根据环境创建系统工具
func NewFileTool() FileTool {
	var tool FileTool
	// if system.IsLinux() {
	// 	tool = &LinuxFileTool{}
	// } else if system.IsWin() {
	// 	tool = &WinFileTool{}
	// } else if system.IsMacOS() {
	// 	tool = &MacOSFileTool{}
	// } else {
	tool = &BaseFileTool{}
	// }
	return tool
}

var _ FileTool = (*BaseFileTool)(nil)

type BaseFileTool struct{}

// Move 文件移动/重命名
func (tool *BaseFileTool) Move(args ...string) {
	if len(args) < 2 {
		cmdio.Println("参数错误!")
		return
	}
	src := args[0]
	dst := args[1]

	if fileio.IsExist(src) {
		// 更改名称为dst
		fileio.Move(src, dst)
		return
	}

	// 将src视为正则
	files, err := fileio.ListDir(fileio.Pwd())
	if err != nil {
		cmdio.Println("读取当前文件失败:", err)
		return
	}

	nameWillChangeMap := map[string]string{}
	reg := regexp.MustCompile(src)
	for _, file := range files {
		fname := file.Name()
		matched := reg.MatchString(fname)
		if !matched {
			continue
		}
		newName := reg.ReplaceAllString(fname, dst)
		nameWillChangeMap[fname] = newName
	}

	if len(nameWillChangeMap) <= 0 {
		cmdio.Println("匹配不到任何文件:", src)
		return
	}

	cmdio.PrintDict(nameWillChangeMap, "源文件", "目标文件")
	cmdio.Printf("需要更改%d个文件，是否继续？y/N\n", len(nameWillChangeMap))
	if strings.ToLower(cmdio.Scan()) != "y" {
		cmdio.Println("操作终止")
		return
	}

	// 执行
	for fname, newName := range nameWillChangeMap {
		err := fileio.Move(fname, newName)
		if err != nil {
			cmdio.Printf("将【%s】修改为【%s】失败！\n", fname, newName)
			continue
		}
		cmdio.Printf("成功将【%s】修改为【%s】\n", fname, newName)
	}
	cmdio.Println("操作完成")
}
