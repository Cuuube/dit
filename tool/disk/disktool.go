package disk

import (
	"regexp"

	"github.com/Cuuube/dit/internal/cmdio"
	"github.com/Cuuube/dit/internal/fileutil"
	"github.com/Cuuube/dit/tool/system"
)

// DiskTool 系统工具
type DiskTool interface {
	Overview()
	Rename(args ...string)
}

// NewDiskTool 根据环境创建系统工具
func NewDiskTool() DiskTool {
	var sysTool DiskTool
	if system.IsLinux() {
		sysTool = &LinuxDiskTool{}
	} else if system.IsWin() {
		sysTool = &WinDiskTool{}
	} else if system.IsMacOS() {
		sysTool = &MacOSDiskTool{}
	} else {
		sysTool = &BaseDiskTool{}
	}
	return sysTool
}

var _ DiskTool = (*BaseDiskTool)(nil)

type BaseDiskTool struct{}

// Overview 查看磁盘使用概览
func (tool *BaseDiskTool) Overview() {
	out, _ := cmdio.Exec("df", "-h")
	cmdio.Println(out)
}

// Rename 文件重命名
func (tool *BaseDiskTool) Rename(args ...string) {
	if len(args) < 2 {
		cmdio.Println("参数错误!")
		return
	}
	src := args[0]
	dst := args[1]

	if fileutil.IsExist(src) {
		// 更改名称为dst
		fileutil.Rename(src, dst)
		return
	}

	// 将src视为正则
	files, err := fileutil.ListDir(fileutil.Pwd())
	if err != nil {
		cmdio.Println("读取当前文件失败:", err)
		return
	}

	reg := regexp.MustCompile(src)
	for _, file := range files {
		fname := file.Name()
		matched := reg.MatchString(fname)
		if !matched {
			continue
		}
		newName := reg.ReplaceAllString(fname, dst)
		err := fileutil.Rename(fname, newName)
		if err != nil {
			cmdio.Printf("将【%s】修改为【%s】失败！\n", fname, newName)
			continue
		}
		cmdio.Printf("将【%s】修改为【%s】\n", fname, newName)
	}

	cmdio.Println("匹配不到任何文件:", src)
}
