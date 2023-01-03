package mark

import "time"

func TodayDatekey() string {
	return time.Now().Format("20060102")
}

func CurrentMinuteDatekey() string {
	return time.Now().Format("200601021504")
}

func CurrentSecondDatekey() string {
	return time.Now().Format("20060102150405")
}
