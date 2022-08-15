// diff两个值的差异
package diff

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Cuuube/dit/pkg/coll"
	"github.com/Cuuube/dit/pkg/ctrl"
)

// DiffJSON 两个json，要求是dict格式的json
func DiffJSONMap(jsonLeft, jsonRight []byte) (dictl, dictr map[string]any, leftOnly, rightOnly map[string]any, err error) {
	dictl, dictr = make(map[string]any), make(map[string]any)
	leftOnly, rightOnly = make(map[string]any), make(map[string]any)

	err = json.Unmarshal(jsonLeft, &dictl)
	if err != nil {
		return dictl, dictr, leftOnly, rightOnly, fmt.Errorf("json1解析失败：%s", err)
	}
	err = json.Unmarshal(jsonRight, &dictr)
	if err != nil {
		return dictl, dictr, leftOnly, rightOnly, fmt.Errorf("json2解析失败：%s", err)
	}

	leftOnly, rightOnly = DiffDict(dictl, dictr, "")
	return dictl, dictr, leftOnly, rightOnly, err
}

// DiffDict diff两个dict
func DiffDict(dictl, dictr map[string]any, path string) (leftOnly, rightOnly map[string]any) {
	leftOnly, rightOnly = make(map[string]any), make(map[string]any)
	for k := range dictl {
		curPath := fmt.Sprintf("%s%s%s", path, ctrl.Ternary(path == "", "", "."), k)

		// 一方无元素的判断
		vl := dictl[k]
		vr, found := dictr[k]
		if !found {
			leftOnly[curPath] = dictl[k]
			continue
		}

		// 都存在，执行value diff
		resl, resr := DiffValue(vl, vr, curPath)
		leftOnly = merge(leftOnly, resl)
		rightOnly = merge(rightOnly, resr)
	}
	for k := range dictr {
		curPath := fmt.Sprintf("%s%s%s", path, ctrl.Ternary(path == "", "", "."), k)
		// 另一方无元素的判断
		_, found := dictl[k]
		if found {
			continue
		}
		rightOnly[curPath] = dictr[k]
	}
	return leftOnly, rightOnly
}

// DiffSlice diff两个slice
func DiffSlice(slicel, slicer []any, path string) (leftOnly, rightOnly map[string]any) {
	leftOnly, rightOnly = make(map[string]any), make(map[string]any)
	lenl, lenr := len(slicel), len(slicer)

	for i := 0; i < lenl || i < lenr; i++ {
		curPath := fmt.Sprintf("%s[%d]", path, i)

		// 一方无元素的判断
		if i >= lenl {
			rightOnly[curPath] = slicer[i]
			continue
		} else if i >= lenr {
			leftOnly[curPath] = slicel[i]
			continue
		}
		vl, vr := slicel[i], slicer[i]

		resl, resr := DiffValue(vl, vr, curPath)
		leftOnly = merge(leftOnly, resl)
		rightOnly = merge(rightOnly, resr)
	}

	return leftOnly, rightOnly
}

// DiffValue diff任意值
func DiffValue(vl, vr any, curPath string) (leftOnly, rightOnly map[string]any) {
	leftOnly, rightOnly = make(map[string]any), make(map[string]any)

	if vl == nil || vr == nil {
		leftOnly[curPath] = vl
		rightOnly[curPath] = vr
		return
	}

	ktl, ktr := reflect.TypeOf(vl).Kind(), reflect.TypeOf(vr).Kind()
	// 类型不同
	if ktl != ktr {
		leftOnly[curPath] = vl
		rightOnly[curPath] = vr
		return
	}
	// 类型相同
	// 如果是数组
	if ktl == reflect.Slice || ktl == reflect.Array {
		resl, resr := DiffSlice(vl.([]any), vr.([]any), curPath)
		leftOnly = merge(leftOnly, resl)
		rightOnly = merge(rightOnly, resr)
		return
	}
	// 如果是map
	if ktl == reflect.Map {
		resl, resr := DiffDict(vl.(map[string]any), vr.(map[string]any), curPath)
		leftOnly = merge(leftOnly, resl)
		rightOnly = merge(rightOnly, resr)
		return
	}
	// 其他的直接判断值
	if !reflect.DeepEqual(vl, vr) {
		leftOnly[curPath] = vl
		rightOnly[curPath] = vr
		return
	}
	return
}

// merge 合并两个dict
func merge[keyType coll.DictKeyable](mapl, mapr map[keyType]any) map[keyType]any {
	for k := range mapr {
		if _, found := mapl[k]; !found {
			mapl[k] = mapr[k]
		}
	}
	return mapl
}
