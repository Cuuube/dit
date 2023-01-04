package hashutil

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// HashUrl 将字符串hash为十六进制MD5字符串
func HashUrl(url string) string {
	h := md5.New()
	h.Write([]byte(strings.ToLower(url)))
	return hex.EncodeToString(h.Sum(nil))
}
