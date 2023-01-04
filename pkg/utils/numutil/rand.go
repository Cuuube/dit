package numutil

import (
	"math/rand"
	"time"
)

// 使用当前时间戳生成随机种子
func CreateRandByNow() *rand.Rand {
	// 1649608992 1649608992218 1649608992218914 1649608992218914000
	// fmt.Println(now.Unix(), now.UnixMilli(), now.UnixMicro(), now.UnixNano())
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// 根据随机数roll概率
func RollBool(randRoll *rand.Rand, total float64, hits float64) bool {
	return randRoll.Float64() < (hits / total)
}
