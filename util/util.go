package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/golang/glog"
	"github.com/thoas/go-funk"
	"github.com/tidwall/gjson"
)

func CheckAndPrettyJson(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

//抓倒數第一個字符並返回字符"以後"的值
func SubStringAfter(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}

//抓倒數第一個字符並返回字符"以前"的值
func SubStringBefore(value string, a string) string {
	// Get substring before a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

func GetMapValues(mymap map[string]int) []int {
	//形態要對 不然會err
	r := funk.Map(mymap, func(k string, v int) int {
		return v
	})

	data := r.([]int)
	return data
}

//json轉interface{}再轉map[string]interface{}
func JsonToMapStringIface(myjson []byte) map[string]interface{} {

	//Marshal the json to a map
	var iface interface{}
	err := json.Unmarshal(myjson, &iface)
	if err != nil {
		glog.Error(err)
	}

	var msiface map[string]interface{}

	switch iface.(type) {
	case interface{}:
		if msiface, ok := iface.(map[string]interface{}); ok {
			return msiface
		}
		// if msiface, ok := iface.([]map[string]interface{}); ok {
		// 	return []msiface
		// }
	}

	//print the map
	// fmt.Println(m)

	//unmarshal the map to json
	//result, _ := json.Marshal(m)

	//print the json
	// os.Stdout.Write(result)

	return msiface
}

//將 arrayJson 轉為 []map[string]interface
func JsonAryToMap(myjson []byte) []map[string]interface{} {

	//Marshal the json to a map
	var aryMFace []map[string]interface{}
	err := json.Unmarshal(myjson, &aryMFace)
	if err != nil {
		glog.Error(err)
	}
	return aryMFace

	// a := mmsiface["dtInstance"].(map[string]interface{})["feature"].(map[string]interface{})["monitor"]
	// fmt.Println(a)

}

//interface{}轉map[string]interface{}
func IfaceToMapStringIface(iface interface{}) map[string]interface{} {
	msiface := iface.(map[string]interface{})
	return msiface
}

func AddKeyValueToJson(key string, v interface{}, data []byte) []byte {
	//add a new key-value to json
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		glog.Error(err)
	}

	m[key] = v

	newData, err := json.Marshal(m)
	if err != nil {
		glog.Error(err)
	}
	return newData
}

func MapStringIfaceToJson(m map[string]interface{}) []byte {
	jsonStr, err := json.Marshal(m)
	if err != nil {
		glog.Error(err)
	}
	return jsonStr
}

func IfaceToJson(m interface{}) []byte {
	jsonStr, err := json.Marshal(m)
	if err != nil {
		glog.Error(err)
	}
	return jsonStr
}

func SliceInIfaceToMapStrIface(t interface{}) []map[string]interface{} {
	mStrIfaces := []map[string]interface{}{}
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)
		for i := 0; i < s.Len(); i++ {
			rv := s.Index(i)
			fmt.Println(rv)
			iface := rv.Interface() // convert reflect value to interface
			mStrIface := iface.(map[string]interface{})
			mStrIfaces = append(mStrIfaces, mStrIface)
		}
	}
	return mStrIfaces
}

func ReadSliceInIface(t interface{}) {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)
		for i := 0; i < s.Len(); i++ {
			rv := s.Index(i)
			fmt.Println(rv)
		}
	}
}

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
