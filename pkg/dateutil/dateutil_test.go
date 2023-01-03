package dateutil

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	t1 := time.Now()

	fmt.Println(FormatDatetime(t1))
	fmt.Println(FormatDate(t1))
	fmt.Println(FormatDateKey(t1))
	fmt.Println(FormatTime(t1))

	fmt.Println(FormatDatetimeUTC(t1))
	fmt.Println(FormatDateUTC(t1))
	fmt.Println(FormatDateKeyUTC(t1))
	fmt.Println(FormatTimeUTC(t1))
}
