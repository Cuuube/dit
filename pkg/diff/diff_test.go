package diff

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDiffDict(t *testing.T) {
	d1 := map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": []any{1, 4, 6},
		"e": []any{9, 9, 9},
		"g": map[string]any{
			"11": "abc",
			"33": "bbb",
			"d":  []any{1, 4, 6},
		},
	}
	d2 := map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": "ccc",
		"d": []any{1, 4, 7, 9},
		"f": []any{1, 4, 6},
		"g": map[string]any{
			"11": "abc",
			"22": "aaa",
			"d":  []any{1, 4, 7, 9},
		},
	}

	// dl, dr, ol, or, err
	ol, or := DiffDict(d1, d2, "")
	fmt.Printf("leftOnly: %+v\nrightOnly: %+v\nerr:%v\n", ol, or, nil)

	// fmt.Println("=======test2:=======")
	// ol, or = DiffDict(map[int]any{1: "123", 2: map[int]any{1: "123", 2: "789"}, 3: 4}, map[int]any{1: "123", 2: map[int]any{1: "123", 2: "222"}, 3: 3}, "")
	// fmt.Printf("leftOnly: %+v\nrightOnly: %+v\nerr:%v\n", ol, or, nil)
}

func TestDiffJSONMap(t *testing.T) {
	d1 := map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": []any{1, 4, 6},
		"e": []any{9, 9, 9},
		"g": map[string]any{
			"11": "abc",
			"33": "bbb",
			"d":  []any{1, 4, 6},
		},
	}
	d2 := map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": "ccc",
		"d": []any{1, 4, 7, 9},
		"f": []any{1, 4, 6},
		"g": map[string]any{
			"11": "abc",
			"22": "aaa",
			"d":  []any{1, 4, 7, 9},
		},
	}
	j1, _ := json.Marshal(d1)
	j2, _ := json.Marshal(d2)
	fmt.Printf("jsonLeft: %s\njsonRight:%s\n", string(j1), string(j2))

	dl, dr, ol, or, err := DiffJSONMap(j1, j2)
	fmt.Printf("jsonValLeft: %+v\njsonValRight: %+v\nleftOnly: %+v\nrightOnly: %+v\nerr:%v\n", dl, dr, ol, or, err)

	fmt.Println("=======test2:=======")
	dl, dr, ol, or, err = DiffJSONMap([]byte("[1,2,3]"), []byte("<div>567</div>"))
	fmt.Printf("jsonValLeft: %+v\njsonValRight: %+v\nleftOnly: %+v\nrightOnly: %+v\nerr:%v\n", dl, dr, ol, or, err)
}

func TestDiffValue(t *testing.T) {
	ol, or := DiffValue(1, "true", "")
	fmt.Printf("leftOnly: %+v\nrightOnly: %+v\nerr:%v\n", ol, or, nil)

	fmt.Println("=======test2:=======")
	ol, or = DiffValue(map[string]any{"a": 123, "b": 456, "c": 7}, map[string]any{"a": 456, "b": 456}, "")
	fmt.Printf("leftOnly: %+v\nrightOnly: %+v\nerr:%v\n", ol, or, nil)

	fmt.Println("=======test3:=======")
	ol, or = DiffValue(map[string]any{"a": 123, "b": 456, "c": nil}, map[string]any{"a": 456, "b": 456}, "")
	fmt.Printf("leftOnly: %+v\nrightOnly: %+v\nerr:%v\n", ol, or, nil)

	fmt.Println("=======test4:=======")
	ol, or = DiffValue(nil, map[string]any{"a": 456, "b": 456}, "")
	fmt.Printf("leftOnly: %+v\nrightOnly: %+v\nerr:%v\n", ol, or, nil)
}
