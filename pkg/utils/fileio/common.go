package fileio

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/Cuuube/dit/pkg/cli"
)

// 打印目录列表
func ListDir(path string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(path)
}

// 创建文件夹
func MkDir(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

// 获取文件状态
func GetStat(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

// 读取文件
func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// 写出文件
func WriteFile(path string, bytes []byte) error {
	return ioutil.WriteFile(path, bytes, 0777)
}

// 移动/重命名
func Move(src, dst string) error {
	return os.Rename(src, dst)
}

// 复制文件或文件夹
func Copy(src, dst string) error {
	src = AbsPath(src)
	dst = AbsPath(dst)

	sStat, err := GetStat(src)
	if err != nil {
		return err
	}
	// 文件的处理
	if !sStat.IsDir() {
		sContent, err := os.ReadFile(src)
		if err != nil {
			return err
		}
		return os.WriteFile(dst, sContent, sStat.Mode())
	}

	// 文件夹的处理：/rootpath/abc -> xyz
	// rootpath, dirname := path.Split(src)
	err = MkDir(dst)
	if err != nil {
		return err
	}
	// 递归copy
	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	for _, fileStat := range files {
		subsrc := JoinPath(src, fileStat.Name())
		subdst := JoinPath(dst, fileStat.Name())
		err = Copy(subsrc, subdst)
		if err != nil {
			cli.Printf("文件复制失败:[%s], 原因：%s\n", subsrc, err)
		}
	}
	return nil
}

// 返回当前程序执行目录
func Pwd() string {
	wd, err := os.Getwd()
	if err != nil {
		cli.Println("pwd执行出错：", err.Error())
	}
	return wd
}

// 路径是否存在
func IsExist(src string) bool {
	_, err := GetStat(src)
	return err == nil
}

// 转换为绝对路径
func AbsPath(p string) string {
	if path.IsAbs(p) {
		return p
	}
	return JoinPath(Pwd(), p)
}

// 连接path
func JoinPath(p ...string) string {
	return path.Join(p...)
}
