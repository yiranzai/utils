package excel

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

//GetSheet 获取Array Row 的 Sheet数组
func GetSheet(path string, sheet string) [][]string {
	var rows, err = excelize.OpenFile(path)
	if err != nil {
		fmt.Printf("open file " + path + "  faild." + err.Error())
		return nil
	}
	return rows.GetRows(sheet)
}

//GetSheetMap 获取MapRow的Sheet数组
func GetSheetMap(path, sheet string) []map[string]string {
	var rows = GetSheet(path, sheet)
	var result []map[string]string
	var temp map[string]string
	header := make(map[int]string)
	result = make([]map[string]string, 0, len(rows)-1)
	for rowNumber, row := range rows {
		temp = make(map[string]string)
		for colNumber, colCell := range row {
			if rowNumber == 0 {
				header[colNumber] = colCell
			} else {
				temp[header[colNumber]] = colCell
			}
		}
		if rowNumber != 0 {
			result = append(result, temp)
		}
	}
	return result
}
