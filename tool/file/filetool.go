package file

import (
	"regexp"

	"github.com/Cuuube/dit/internal/cmdio"
	"github.com/Cuuube/dit/internal/fileutil"
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

	if fileutil.IsExist(src) {
		// 更改名称为dst
		fileutil.Move(src, dst)
		return
	}

	// 将src视为正则
	files, err := fileutil.ListDir(fileutil.Pwd())
	if err != nil {
		cmdio.Println("读取当前文件失败:", err)
		return
	}

	failedFlg := true
	reg := regexp.MustCompile(src)
	for _, file := range files {
		fname := file.Name()
		matched := reg.MatchString(fname)
		if !matched {
			continue
		}
		newName := reg.ReplaceAllString(fname, dst)
		err := fileutil.Move(fname, newName)
		if err != nil {
			cmdio.Printf("将【%s】修改为【%s】失败！\n", fname, newName)
			continue
		}
		cmdio.Printf("将【%s】修改为【%s】\n", fname, newName)
		failedFlg = false
	}

	if failedFlg {
		cmdio.Println("匹配不到任何文件:", src)
	}
}
