package util

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	. "github.com/logrusorgru/aurora"
)

func PrintJson(s interface{}) {
	out, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%s\n%s\n", logLine(), string(out))
}

func PrintReq(s []byte) {
	fmt.Println(BrightBlue("PrintReq..."))
	pstring := Sprintf("-> request body is %s", string(s))
	fmt.Println(BrightBlue(pstring))
}

func Pdo(s interface{}) {
	ps := Sprintf("%s...", s)
	fmt.Println(Blue(ps))
	// a := Blue(ps)
	// return a
}

func Pdone(s interface{}) {
	ps := Sprintf("%s Success", s)
	fmt.Println(Blue(ps))
}

func Cerr(s interface{}) interface{} {
	ps := Sprintf("err: %s", s)
	return (Red(funcName() + ps))
}

func CerrLine(s interface{}) {
	ps := Sprintf("err: %s", s)
	l := logLine()
	fmt.Println(Red(l + ps))
}

// defer util.Elapsed("DDDDDDD1") not work
func Elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}
