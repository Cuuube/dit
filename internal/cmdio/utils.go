package cmdio

import "strings"

// 将文本格式解析为kv。例如`Name: val`
func SplitToDict(raw string, KVSplitChar string) map[string]string {
	ret := map[string]string{}

	lines := strings.Split(raw, "\n")
	for _, line := range lines {
		clips := strings.SplitN(line, KVSplitChar, 2)
		if len(clips) > 1 {
			k := strings.Trim(clips[0], "\t \n\"'")
			v := strings.Trim(clips[1], "\t \n\"'")
			ret[k] = v
		}
	}
	return ret
}
