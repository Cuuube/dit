package numutil

import (
	"fmt"
	"testing"
)

func TestCreateRandByNow(t *testing.T) {
	r1 := CreateRandByNow()
	r2 := CreateRandByNow()

	if r1.Int31() == r2.Int31() {
		t.Fail()
	}
}

func TestRandBool(t *testing.T) {
	randRoll := CreateRandByNow()
	var trueTimes, falseTimes float64 = 0, 0
	for i := 0; i < 1000; i++ {
		if RollBool(randRoll, 100, 1) {
			trueTimes++
		} else {
			falseTimes++
		}
	}
	fmt.Println("trueTimes:", trueTimes, "falseTimes:", falseTimes, "概率大概：", trueTimes/(trueTimes+falseTimes))
}
