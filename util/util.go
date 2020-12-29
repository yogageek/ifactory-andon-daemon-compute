package util

import (
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

//切斷超過上限的字串
func TruncateString(str string, num int) string {
	bnoden := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		bnoden = str[0:num] + "...(TOO LONG)"
	}
	// fmt.Println(bnoden)
	return bnoden
}

func GetMapKeys(mymap map[int]string) []int {
	keys := make([]int, 0, len(mymap))
	for k := range mymap {
		keys = append(keys, k)
	}
	return keys
}

func ArrayToString(a interface{}, delim string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
	// return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
}

func GetValueByPath(json string, path string) gjson.Result {
	result := gjson.Get(json, path)
	return result
}
