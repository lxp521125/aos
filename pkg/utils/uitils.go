package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var reLineBreak = regexp.MustCompile("\n")

// PathExists ...
func PathExists(path string) (bool, error) {
	dir, _ := os.Getwd() //当前的目录
	path = dir + path
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// MakeDir ...
func MakeDir(path string) error {
	dir, _ := os.Getwd() //当前的目录
	err := os.Mkdir(dir+path, os.ModePerm)
	return err
}

// ChangeStruct2OtherStruct ...
func ChangeStruct2OtherStruct(item interface{}, toItem interface{}) {

	j, _ := json.Marshal(item)
	json.Unmarshal(j, &toItem)
}

// ChangeRedis2OtherStruct ...
func ChangeRedis2OtherStruct(item interface{}) []uint8 {
	// var toItem []uint8
	// j, _ := json.Marshal(item)
	// json.Unmarshal(j, &toItem)
	return item.([]uint8)
}

// ChangeUint82OtherStruct ...
func ChangeUint82OtherStruct(item interface{}, toItem interface{}) {
	j, _ := json.Marshal(item)
	json.Unmarshal(j, &toItem)
}

// ChangeByteStruct2OtherStruct ...
func ChangeByteStruct2OtherStruct(str []byte, toItem interface{}) error {

	if erri := json.Unmarshal([]byte(str), &toItem); erri != nil && toItem == nil {
		return erri
	}
	return nil
}

// ChangeArrayString2Int ...
func ChangeArrayString2Int(before []string) []int {
	after := make([]int, len(before))
	for k, item := range before {
		after[k], _ = strconv.Atoi(item)
	}
	return after
}

// ChangeArrayString2Int64 ...
func ChangeArrayString2Int64(before []string) []int64 {
	after := make([]int64, len(before))
	for k, item := range before {
		after[k], _ = strconv.ParseInt(item, 10, 64)
	}
	return after
}

// JoinInt64Array2String 拆分int64的数组为string
func JoinInt64Array2String(before []int64, sep string) string {
	after := ""
	for _, item := range before {
		after = after + strconv.Itoa(int(item)) + sep
	}
	return strings.Trim(after, sep)
}

// String2Int ...
func String2Int(before string) int {
	after, _ := strconv.Atoi(before)
	return after
}

// String2Int64 ...
func String2Int64(before string) int64 {
	after, _ := strconv.ParseInt(before, 10, 64)
	return after
}

// String2Float64 ...
func String2Float64(before string) float64 {
	after, _ := strconv.ParseFloat(before, 64)
	return after
}

// String2Int32 ...
func String2Int32(before string) int32 {
	after, _ := strconv.Atoi(before)
	return int32(after)
}

// String2Int8 ...
func String2Int8(before string) int8 {
	after, _ := strconv.Atoi(before)
	return int8(after)
}

// Slise2Map ....
func Slise2Map(originData []string) map[string]int {
	retData := make(map[string]int, 0)
	if originData == nil {
		return retData
	}
	// count := len(originData)
	for i, v := range originData {
		retData[v] = i
	}
	// for i := 0; i < count; i++ {
	// 	// if mapContains(retData, originData[i]) { //前后取前

	// 	// }
	// }
	return retData
}

//  InArrayInt64 是否存在数组中 int64
func InArrayInt64(arr []int64, val int64) bool {
	flag := false
	for _, item := range arr {
		if item == val {
			flag = true
			break
		}
	}
	return flag
}

//  InArrayInt  是否存在数组中 int
func InArrayInt(arr []int, val int) bool {
	flag := false
	for _, item := range arr {
		if item == val {
			flag = true
			break
		}
	}
	return flag
}

// InArrayString 是否存在数组中 字符串
func InArrayString(arr []string, val string) bool {
	flag := false
	for _, item := range arr {
		if item == val {
			flag = true
			break
		}
	}
	return flag
}

//判断key是否存在
func MapContains(needMap map[string]int, key string) bool {
	if _, ok := needMap[key]; ok {
		return true
	}
	return false
}

// FormartDate2Time
func FormartDate2Time(dataTimeStr, ms string) int64 {
	dataTime, _ := time.Parse(ms, dataTimeStr)
	return dataTime.Unix()
}

// ABCToRune ...
func ABCToRune(abc string) rune {
	abcrune := []rune(abc)
	return abcrune[0]
}

// TODO 读取文件可能是别的形式 用factory进行构造读函数
func ReadExcel(sheet [][]string, obj interface{}, startRow, startCol int) ([]interface{}, error) {
	objT := reflect.TypeOf(obj)
	if !(objT.Kind() == reflect.Ptr && objT.Elem().Kind() == reflect.Struct) {
		return nil, errors.New("readSheet must pass a struct type")
	}
	if len(sheet) <= startRow || len(sheet[0]) <= startCol {
		return nil, errors.New("empty sheet ")
	}

	type FieldInfo struct {
		Index   int
		Field   *reflect.StructField
		Group   string
		ColName string
	}
	objT = objT.Elem()
	var colMap = make(map[int]*FieldInfo)
	var maxColumnIndex = 0 //it's the column index of first invalid column
	columnFound := make(map[string]bool)
	for i, cell := range sheet[startRow-1] {
		if i < startCol-1 {
			continue
		} else if cell == "" || len(strings.TrimSpace(cell)) == 0 { // break when meet first empty column
			break
		}
		maxColumnIndex = i
		cellValue := strings.TrimSpace(cell)
		for j := 0; j < objT.NumField(); j++ {
			field := objT.Field(j)
			if field.Tag.Get("col") == cellValue {
				colMap[i] = &FieldInfo{Index: j, Field: &field, Group: field.Tag.Get("group"), ColName: cellValue}
				columnFound[cellValue] = true
			}
		}
	}
	if len(colMap) == 0 {
		return nil, errors.New("no column found for sheet ")
	}
	//检查是否缺少column配置
	for j := 0; j < objT.NumField(); j++ {
		field := objT.Field(j)
		colInStruct := field.Tag.Get("col")
		if len(colInStruct) < 1 {
			continue
		}
		if !columnFound[colInStruct] {
			//fmt.Printf("no column found for sheet %s, column %s \n", sheet.Name, colInStruct)
			return nil, errors.New(fmt.Sprintf("表格 ,中没有列%s,更新checkconfig.exe再试试", colInStruct))
		}
	}

	errFunc := func(elem reflect.Type, fieldIndex, i, j int, err error) error {
		return fmt.Errorf("field %s at %c%d error for sheet %s: %s", elem.Field(fieldIndex).Name, 'A'+j%26, i+1, err.Error())
	}
	var result = make([]interface{}, 0)
	for i, row := range sheet {
		if i < startRow {
			continue
		} else if row == nil || len(row) == 0 {
			break
		}
		objInstance := reflect.New(objT).Interface()
		objV := reflect.ValueOf(objInstance).Elem()
		for j, cell := range row {
			if j < startCol-1 {
				continue
			}
			fieldInfo := colMap[j]
			if fieldInfo == nil {
				continue
			}
			cellString := cell
			cellString = strings.TrimSpace(cellString)
			if j == startCol-1 && i >= startRow && (cell == "" || len(cellString) == 0) {
				goto exit //finish when meet first empty row (the first column of this row is empty)
			}
			if j > maxColumnIndex {
				break
			}
			fieldV := objV.Field(fieldInfo.Index)
			if !fieldV.CanSet() {
				return nil, fmt.Errorf("field %s can not be set for sheet %s", objT.Field(fieldInfo.Index).Name)
			}
			if len(cellString) == 0 {
				continue
			}
			switch objT.Field(fieldInfo.Index).Type.Kind() {
			case reflect.Bool:
				if cellString == "1" {
					fieldV.SetBool(true)
				} else if cellString == "0" {
					fieldV.SetBool(false)
				} else {
					b, err := strconv.ParseBool(cellString)
					if err != nil {
						return nil, errFunc(objT, fieldInfo.Index, i, j, err)
					}
					fieldV.SetBool(b)
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				x, err := strconv.ParseInt(strings.Split(cellString, ".")[0], 10, 64) //需防止自动计算字段为float类型
				if err != nil {
					return nil, errFunc(objT, fieldInfo.Index, i, j, err)
				}
				fieldV.SetInt(x)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				x, err := strconv.ParseUint(cellString, 10, 64)
				if err != nil {
					return nil, errFunc(objT, fieldInfo.Index, i, j, err)
				}
				fieldV.SetUint(x)
			case reflect.Float32, reflect.Float64:
				x, err := strconv.ParseFloat(cellString, 64)
				if err != nil {
					return nil, errFunc(objT, fieldInfo.Index, i, j, err)
				}
				fieldV.SetFloat(x)
			case reflect.String:
				s1 := reLineBreak.ReplaceAllString(cellString, "")
				// s2 := strings.Replace(s1, `\`, `\\`, -1)
				fieldV.SetString(strings.Replace(s1, `"`, `\"`, -1))
				// fieldV.SetString(reLineBreak.ReplaceAllString(cellString, ""))
			}
		}
		result = append(result, objInstance)
	}
exit:
	return result, nil
}
