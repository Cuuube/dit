package hashutil

import (
	"hash/crc32"
)

// HashStringToUint 用crc32将字符串hash为uint
func HashStringToUint(str string) uint {
	return uint(crc32.ChecksumIEEE([]byte(str)))
}
