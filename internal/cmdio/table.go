package cmdio

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// 打印带有表头的二维表格
func PrintTableWithHeader(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

// 打印二维表格
func PrintTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

// 打印字典，作为二维表格输出
func PrintDict[T any](data map[string]T, keyColName, valColName string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{keyColName, valColName})

	for k, v := range data {
		table.Append([]string{k, fmt.Sprintf("%v", v)})
	}
	table.Render() // Send output
}

// 打印结构体的key、value
func PrintStruct(stru any) {
	kvDict := make(map[string]any)

	// typ := reflect.TypeOf(stru)
	// vals := reflect.ValueOf(typ)
	// if typ.Kind() == reflect.Ptr {
	// 	typ = typ.Elem()
	// 	vals = vals.Elem()
	// }

	// // 否则正常解析：
	// for i := 0; i < typ.NumField(); i++ {
	// 	f := typ.Field(i)

	// 	kvDict[f.Name] = vals.Field(i).Interface()
	// }
	byt, _ := json.Marshal(stru)
	json.Unmarshal(byt, &kvDict)
	PrintDict(kvDict, "key", "val")
}
