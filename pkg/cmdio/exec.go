package cmdio

import (
	"os/exec"
	"strings"
)

// ExecCmd 执行命令，并输出命令执行信息
func Exec(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// ExecIgnoreErr 执行命令，并输出命令执行信息，忽略错误
func ExecIgnoreErr(name string, arg ...string) string {
	out, _ := Exec(name, arg...)
	return strings.Trim(out, "\t \n")
}
