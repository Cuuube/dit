package cli

import "testing"

func TestPrintDict(t *testing.T) {
	data1 := map[string]any{
		"goos":       "darwin",
		"arch":       "arm64",
		"longstring": "aaaaaaaaaaaaaaa aaaaaaa",
	}
	PrintDict(data1, "KEY", "VALUE")
}

func TestPrintTable(t *testing.T) {
	data1 := [][]string{
		{"goos", "darwin"},
		{"arch", "arm64"},
		{"longstring", "aaaaaaaaaaaaaaa aaaaaaa"},
	}
	PrintTable(data1)
	PrintTableWithHeader([]string{"hello", "world"}, data1)
}

func TestPrintStruct(t *testing.T) {
	type stru1 struct {
		TestStr string
		TestInt int
	}
	PrintStruct(stru1{TestStr: "haha", TestInt: 999})
}
