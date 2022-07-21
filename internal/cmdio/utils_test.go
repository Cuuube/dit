package cmdio

import (
	"fmt"
	"testing"
)

func TestParseStrToDict(t *testing.T) {
	example1 := `aaaaaa
	ProductName:	macOS
	ProductVersion:	12.4`
	result := SplitToDict(example1, ":")
	fmt.Printf("%+v\n", result)

	if result["ProductName"] != "macOS" {
		t.Fail()
	}
	if result["ProductVersion"] != "12.4" {
		t.Fail()
	}
	if _, found := result["aaaaaa"]; found {
		t.Fail()
	}

	example2 := `PRETTY_NAME="Debian GNU/Linux 9 (stretch)"
	NAME="Debian GNU/Linux"`
	result2 := SplitToDict(example2, "=")
	fmt.Printf("%+v\n", result2)
	if result2["PRETTY_NAME"] != "Debian GNU/Linux 9 (stretch)" {
		t.Fail()
	}
}
