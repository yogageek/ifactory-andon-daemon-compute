package util

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/imroc/req"
	"github.com/logrusorgru/aurora"
)

type MyError struct {
	ErrDesc string
	ErrMsg  string
	Code    int
}

//只要實現了 error接口, 就可以當成是 error 型態的物件
//可以使用 return &MyError{err, 0, ""}
func (e *MyError) Error() string {
	// return fmt.Sprintf("radius %0.2f: %s", e.Code, e.Message)
	return e.ErrMsg
}

func DoAPI(apiType string, url string, param interface{}) ([]byte, error) {

	// #fix for krunshan
	req.EnableInsecureTLS(true)

	switch param.(type) {
	case nil:
		// fmt.Println(aurora.Yellow("no param"))
	case req.Param:
		// fmt.Println(aurora.Yellow("this is req.Param"))
	// case *req.bodyJson:
	// 	yellow("this is *req.bodyJson")
	default:
		// fmt.Println(aurora.Yellow("this is *req.bodyJson"))
		// yellow(reflect.TypeOf(param))
	}

	fmt.Println(aurora.Blue(fmt.Sprintf("%v %v ", apiType, url)))

	header := req.Header{
		"Accept": "application/json",
		// "Authorization": Token,
	}

	// green := color.New(color.FgGreen).SprintFunc()
	// glog.Infof("param: %+v", green(param))

	var (
		r   *req.Resp
		err error
	)

	if apiType == "GET" {
		r, err = req.Get(url, header, param)
	} else if apiType == "POST" {
		r, err = req.Post(url, header, param)
	} else if apiType == "PATCH" {
		r, err = req.Patch(url, header, param)
	} else if apiType == "PUT" {
		r, err = req.Put(url, header, param)
	} else if apiType == "DELETE" {
		r, err = req.Delete(url, header, param)
	} else {
		panic("apiType invalid")
	}

	if err != nil {
		apiErr := fmt.Errorf("API Err: %s", err.Error())
		glog.Error(Cerr(apiErr))
	}

	rCode := r.Response().StatusCode
	if rCode != 200 && rCode != 201 {
		respErr := fmt.Errorf(string(r.Bytes())) //對方api返回的錯誤訊息
		fmt.Println(aurora.Red(fmt.Errorf("%s FAIL! code=%d, resp=%s ", apiType, rCode, respErr)))
		//# New Design
		myErr := &MyError{
			ErrMsg: respErr.Error(),
			Code:   rCode,
		}
		return nil, myErr
		//how to panic and recover
	}

	// method1
	// r.ToJSON(&foo)       // response => struct/map
	// log.Printf("%+v", r) // print info (try it, you may surprise)
	// method2
	res, err := r.ToBytes()
	if err != nil {
		glog.Error(Cerr(err))
	}

	resStr := string(res)
	resStr = TruncateString(resStr, 1000)
	fmt.Println(aurora.Green(fmt.Sprintf("%s SUCCESS! code=%d, resp=%s ", apiType, rCode, resStr)))
	return res, nil
}
