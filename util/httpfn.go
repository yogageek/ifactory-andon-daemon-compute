package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang/glog"
	"github.com/logrusorgru/aurora"
)

var (
	ParaFail error = errors.New("illegal parameter.")
)

func WErr(w http.ResponseWriter, err interface{}, code int) {
	w.WriteHeader(code)

	//# New Design
	var desc string

	// for narmal err
	func() {
		if err, ok := err.(error); ok {
			ChooseDescByError(err)
			//從err interface轉為MyError取出其中ErrDesc
			desc = func(interface{}) string {
				if err, ok := err.(*MyError); ok {
					return err.ErrDesc
				}
				return ""
			}(err)
		}
	}()

	// for array err
	func() {
		if errstr, ok := err.([]string); ok {
			errMsg := strings.Join(errstr, "") //把[]string轉string
			err := &MyError{"", errMsg, code}
			ChooseDescByError(err)
			desc = err.ErrDesc
		}
	}()

	switch v := err.(type) {
	case error:
		err = v.(error).Error() //須轉為string 否則前端顯示不出來
	}

	errfmt := map[string]interface{}{
		"description": desc,
		"error":       err,
	}

	json.NewEncoder(w).Encode(errfmt)
}

func ChooseDescByError(err error) {
	const a = "projectId is duplicated"
	const aa = "group is duplicated"

	const b = "parent id is not exist"
	const bb = "is your Id correct?"

	const c = "userName is duplicated"
	const cc = "userName is duplicated"

	if strings.Contains(err.Error(), a) {
		SetDesc(err, aa)
	}
	if strings.Contains(err.Error(), b) {
		SetDesc(err, bb)
	}
	if strings.Contains(err.Error(), c) {
		SetDesc(err, cc)
	}

	//無法用switch(Execution order)
}

func SetDesc(err error, desc string) {
	if err, ok := err.(*MyError); ok {
		err.ErrDesc = desc
		return
	}
	//增加來源行數
	glog.Error(aurora.Magenta("error type is not MyError"))
}

func WErrNew(w http.ResponseWriter, errmap map[string]interface{}, code int) {
	w.WriteHeader(code)

	newerrmap := map[string]interface{}{}
	for k, v := range errmap {
		var str string
		switch vt := v.(type) {
		case error:
			str = vt.(error).Error() //須轉為string 否則前端顯示不出來
		}
		newerrmap[k] = str
	}

	errfmt := map[string]interface{}{
		"error": newerrmap,
	}

	json.NewEncoder(w).Encode(errfmt)
	// {
	// 	"error": {
	// 		"dashb": "{\"message\":\"Folder not found\",\"status\":\"not-found\"}",
	// 		"dh": "{\"message\":\"srp-frame not found \"}"
	// 	}
	// }
}

// func switchIface(i interface{}) {
// 	switch v := i.(type) {

// 	case string:
// 		fmt.Printf("iface is string")
// 		return i.(string)
// 	case []string:
// 		fmt.Printf("iface is []string")
// 	case error:
// 		fmt.Printf("iface is error")
// 		return i.(error)
// 	default:
// 		fmt.Printf("I don't know about type %T!\n", v)
// 	}
// }

func CheckRequestParameter(iAry ...interface{}) bool {
	for _, i := range iAry {
		switch iv := i.(type) {
		case string:
			if iv == "" {
				fmt.Println("string is empty")
				return false
			}
		case []string:
			if len(iv) == 0 {
				fmt.Println("[]string is empty")
				return false
			}
			for _, v := range iv {
				if v == "" {
					fmt.Println("string is empty")
					return false
				}
			}
		case int:
			if iv == 0 {
				fmt.Println("int is empty")
				return false
			}
		default:
			fmt.Printf("I don't know about type %T!\n", iv)
			return false
		}
	}
	return true
}
