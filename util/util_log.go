package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/golang/glog"
	. "github.com/logrusorgru/aurora"
)

var Lg *DebugLogStruct

type DebugLogStruct struct {
	PackageName string
}

// 	debug := os.Getenv("DEBUG")
// 	if debug != "" {
// 		Lg = new(DebugLogStruct)
// 	}

//pretty print any data
func (lg *DebugLogStruct) PrintJson(i interface{}) {
	byt, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		glog.Error(Cerr(err))
	}
	fmt.Println(string(byt))
}

//pretty print []byte
func (lg *DebugLogStruct) PrintByte(byt []byte) {
	var bfer bytes.Buffer
	err := json.Indent(&bfer, byt, "", "  ")
	if err != nil {
		glog.Error(Cerr(err))
	}
	fmt.Println(bfer.String())
}

func DEBUG(formating string, args ...interface{}) {
	LOG("DEBUG", formating, args...)
}

func LOGTEST() {
	p := []byte(`{"id" : 1 , "name" : "Daniel"}`)
	DEBUG("")
	DEBUG("Param p=%s", p)
	DEBUG("Test %s %s", "Hello", "World")
}

//config your print style here
func LOG(level string, formating string, args ...interface{}) {
	filename, line, funcname := "???", 0, "???"
	pc, filename, line, ok := runtime.Caller(2) //Caller(skip int) 是要提升的堆栈帧数，0-当前函数，1-上一层函数，....

	// fmt.Println(reflect.TypeOf(pc), reflect.ValueOf(pc))

	//只取呼叫來源的簡要名稱
	if ok {
		funcname = runtime.FuncForPC(pc).Name()      // main.(*MyStruct).foo
		funcname = filepath.Ext(funcname)            // .foo
		funcname = strings.TrimPrefix(funcname, ".") // foo
		filename = filepath.Base(filename)           // /full/path/basename.go => basename.go
	}

	log.Printf("%s:%d:%s: %s: %s\n", Blue(filename), Blue(line), Blue(funcname), level, fmt.Sprintf(formating, args...))
}

func logLine() string {
	fileName, line, funcName := "???", 0, "???"
	pc, fileName, line, ok := runtime.Caller(2) //Caller(skip int) 是要提升的堆栈帧数，0-当前函数，1-上一层函数，....
	//只取呼叫來源的簡要名稱
	if ok {
		funcName = runtime.FuncForPC(pc).Name()      // main.(*MyStruct).foo
		funcName = filepath.Ext(funcName)            // .foo
		funcName = strings.TrimPrefix(funcName, ".") // foo

		fileName = filepath.Base(fileName) // /full/path/basename.go => basename.go
	}
	return fmt.Sprintf("%s:%d:%s:", fileName, line, funcName)
}

func funcName() string {
	fileName, _, funcName := "???", 0, "???"
	pc, fileName, _, ok := runtime.Caller(2) //Caller(skip int) 是要提升的堆栈帧数，0-当前函数，1-上一层函数，....
	//只取呼叫來源的簡要名稱
	if ok {
		funcName = runtime.FuncForPC(pc).Name()      // main.(*MyStruct).foo
		funcName = filepath.Ext(funcName)            // .foo
		funcName = strings.TrimPrefix(funcName, ".") // foo

		fileName = filepath.Base(fileName) // /full/path/basename.go => basename.go
	}
	return funcName + " "
}

//How to get the name of a function in Go?
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
