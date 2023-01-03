package file

import (
	"regexp"
	"strings"

	"github.com/Cuuube/dit/pkg/cli"
	"github.com/Cuuube/dit/pkg/fileio"
	"github.com/Cuuube/dit/pkg/httpio"
)

// FileTool 系统工具
type FileTool interface {
	Move(args ...string)
	Copy(args ...string)
	Fetch(args ...string)
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
		cli.Println("未找到src和dst!", args)
		return
	}
	src := args[0]
	dst := args[1]

	if fileio.IsExist(src) {
		// 更改名称为dst
		err := fileio.Move(src, dst)
		if err != nil {
			cli.Println("move失败: ", err)
		}
		return
	}

	// 将src视为正则
	files, err := fileio.ListDir(fileio.Pwd())
	if err != nil {
		cli.Println("读取当前文件失败:", err)
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
		cli.Println("匹配不到任何文件:", src)
		return
	}

	cli.PrintDict(nameWillChangeMap, "源文件", "目标文件")
	cli.Printf("需要更改%d个文件，是否继续？y/N\n", len(nameWillChangeMap))
	if strings.ToLower(cli.ReadInput()) != "y" {
		cli.Println("操作终止")
		return
	}

	// 执行
	for fname, newName := range nameWillChangeMap {
		err := fileio.Move(fname, newName)
		if err != nil {
			cli.Printf("将【%s】修改为【%s】失败！\n", fname, newName)
			continue
		}
		cli.Printf("成功将【%s】修改为【%s】\n", fname, newName)
	}
	cli.Printf("move成功！\nsrc: %s\ndst: %s\n", src, dst)
}

func (tool *BaseFileTool) Copy(args ...string) {
	if len(args) < 2 {
		cli.Println("未找到src和dst!", args)
		return
	}
	src := args[0]
	dst := args[1]

	if !fileio.IsExist(src) {
		// 更改名称为dst
		cli.Println("路径不存在: ", src)
		return
	}
	err := fileio.Move(src, dst)
	if err != nil {
		cli.Println("copy失败: ", err)
	}
	cli.Printf("copy成功！\nsrc: %s\ndst: %s\n", src, dst)
}

func (tool *BaseFileTool) Fetch(args ...string) {
	if len(args) < 1 {
		cli.Println("参数错误!")
		return
	}
	src := args[0]
	dst := fileio.JoinPath(fileio.Pwd(), fileio.GetFileName(src))
	if len(args) >= 2 {
		dst = args[1]
	}

	err := httpio.Fetch(src, dst)
	if err != nil {
		cli.Printf("获取网络资源出错：%v\nsrc: %s\ndst: %s\n", err, src, dst)
	}
	cli.Printf("获取网络资源成功！\nsrc: %s\ndst: %s\n", src, dst)
}
