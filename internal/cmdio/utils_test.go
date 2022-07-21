package cmdio

import (
	"testing"
)

func TestParseStrToDict(t *testing.T) {
	examples := `aaaaaa
	ProductName:	macOS
	ProductVersion:	12.4`
	result := ParseStrToDict(examples)
	// cmdio.Printf("%+v", result)

	if result["ProductName"] != "macOS" {
		t.Fail()
	}
	if result["ProductVersion"] != "12.4" {
		t.Fail()
	}
	if _, found := result["aaaaaa"]; found {
		t.Fail()
	}
}
