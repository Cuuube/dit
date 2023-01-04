package strutil

import "strings"

func SubStr(raw string, maxlen int) string {
	if maxlen >= len(raw) {
		return raw
	}
	return raw[0:maxlen]
}

func PadPrefix(raw string, padChar rune, targetlen int) string {
	var sb strings.Builder
	for len(raw)+sb.Len() < targetlen {
		sb.WriteRune(padChar)
	}
	sb.WriteString(raw)
	return sb.String()
}
func PadSuffix(raw string, padChar rune, targetlen int) string {
	var sb strings.Builder
	sb.WriteString(raw)

	for sb.Len() < targetlen {
		sb.WriteRune(padChar)
	}
	return sb.String()
}
