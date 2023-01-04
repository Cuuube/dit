package regutil

import "regexp"

// IsMatchReg str是否匹配
func IsMatchReg(reg, str string) bool {
	if str == "" {
		return false
	}
	if matched, err := regexp.MatchString(reg, str); err == nil && matched {
		return true
	}
	return false
}
