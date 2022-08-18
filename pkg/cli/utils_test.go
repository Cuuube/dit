package cli

import (
	"fmt"
	"testing"
)

func TestSplitToDict(t *testing.T) {
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

func TestSplitToTable(t *testing.T) {
	example1 := `Filesystem	Size	Used	Avail	Use%	Mounted on
udev	32G	0	32G	0%	/dev
tmpfs	6.3G	755M	5.6G	12%	/run
/dev/sda1	86G	69G	13G	85%	/`
	result := SplitToTable(example1, "\t")
	fmt.Printf("%+v\n", result)

	if len(result) != 4 {
		t.Fail()
	}
	if result[0][0] != "Filesystem" {
		t.Fail()
	}
	if result[3][4] != "85%" {
		t.Fail()
	}
}
