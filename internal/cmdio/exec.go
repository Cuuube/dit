package cmdio

import (
	"os/exec"
)

// ExecCmd 执行命令，并输出命令执行信息
func Exec(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}
