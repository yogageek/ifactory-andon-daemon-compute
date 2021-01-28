package model

import (
	"encoding/json"
	"fmt"
	"iii/ifactory/compute/util"

	"github.com/golang/glog"
)

func te(s string) {

}

func RealRecursive(m map[string]interface{}) {
	for k, v := range m {
		_ = k
		// fmt.Println("k:", k)
		// fmt.Println("v:", v)
		if v, ok := v.(map[string]interface{}); ok {
			RealRecursive(v)
		}
	}
}

func MaunalRecursive(m map[string]interface{}) ([]string, []interface{}) {
	var ss []string
	var vv []interface{}
	for k, v := range m { //station,workorder
		header := k
		if m, ok := v.(map[string]interface{}); ok {
			for k, v := range m { //count,typecount,eachtypecount
				if m, ok := v.(map[string]interface{}); ok {
					for k, v := range m {
						if m, ok := v.(map[string]interface{}); ok {
							_ = m
						} else {
							ss = append(ss, header+k)
							vv = append(vv, v)
						}
					}
				} else {
					ss = append(ss, header+k)
					vv = append(vv, v)
				}
			}
		} else {
			ss = append(ss, header+k)
			vv = append(vv, v)
		}
	}
	return ss, vv
}

func GenGrafanaResponse(columns []string, values []interface{}, v interface{}) map[string]interface{} {

	util.PrintJson(v)

	b, err := json.Marshal(&v)
	if err != nil {
		glog.Error(err)
	}

	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		glog.Error(err)
	}

	columns, values = MaunalRecursive(m)
	fmt.Println(columns)
	fmt.Println(values)

	Columns := []map[string]interface{}{}
	for i := 0; i < len(columns); i++ {
		m := map[string]interface{}{
			"text": columns[i],
			"type": "string",
		}
		Columns = append(Columns, m)
	}

	Rows := []interface{}{}
	Rows = append(Rows, values)

	r := map[string]interface{}{
		"columns": Columns,
		"rows":    Rows,
		"type":    "table",
	}

	util.PrintJson(r)

	return r
}
