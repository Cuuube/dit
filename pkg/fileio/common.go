package fileio

import (
	"io/ioutil"
	"os"

	"github.com/Cuuube/dit/pkg/cmdio"
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

// 返回当前程序执行目录
func Pwd() string {
	wd, err := os.Getwd()
	if err != nil {
		cmdio.Println("pwd执行出错：", err.Error())
	}
	return wd
}

// 路径是否存在
func IsExist(src string) bool {
	_, err := GetStat(src)
	if err != nil {
		return false
	}
	return true
}
