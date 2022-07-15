package util

import (
	"encoding/json"
	"strconv"
)

func IntToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func StringToInt(i string) int {
	j, _ := strconv.Atoi(i)
	return j
}
func StringToInt64(i string) int64 {
	j, _ := strconv.ParseInt(i, 10, 64)
	return j
}

// 判断字符串是否在字符串列表中
func IsContain(target string, List []string) bool {

	for _, element := range List {

		if target == element {
			return true
		}
	}
	return false

}

func InterfaceArrayToStringArray(data []interface{}) (i []string) {
	for _, param := range data {
		i = append(i, param.(string))
	}
	return i
}

func StructToJsonString(param interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

// 传入参数必须是指针
func JsonStringToStruct(s string, args interface{}) error {
	err := json.Unmarshal([]byte(s), args)
	return err
}

func int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
