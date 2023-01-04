package dateutil

import "regexp"

// 校验日期时间格式合法性
func CheckDatekeyValid(datekey string) bool {
	reg, _ := regexp.Compile(`^(\d{4}-\d{2}-\d{2})?T?\s?(\d{2}:\d{2}:\d{2})?$`)
	return reg.MatchString(datekey)
}
