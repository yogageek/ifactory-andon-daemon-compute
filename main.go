package main

import (
	"flag"
	"fmt"
	"iii/ifactory/compute/db"
	"iii/ifactory/compute/pkg/logic"

	"iii/ifactory/compute/routers"
	"iii/ifactory/compute/setenv"
	"net/http"
	"os"
)

//其他包的init會先執行，因为在编译的时候会先去检查导入的包，首先发现其他包里面的init，然后才会到main包里面的init
func init() {
	fmt.Println("DEBUG:", os.Getenv("DEBUG"))

	setFlag()

	setenv.SetEnv()

	db.StartMongo()
}

var port = ":8080"

func main() {
	// glog.Infoln("VERSION:", "v1.0.0 ", port) //改從build args拿
	// go thesso.RunTokenLoop()
	// time.Sleep(2 * time.Second)
	doSth()

	//REFACTOR LATER!!!!!! 參考go project layout分到另一個main------------------------->
	refacLater()

}

func doSth() {
	// logic.RunDaemonLoop()

	// r := db.FindAll(db.AMRawdata)
	// util.PrintJson(r)
	go logic.RunDaemonLoop()
}

func setFlag() {
	flag.Usage = usage

	//after setting here, no more need to put in .vscode args
	flag.Set("alsologtostderr", "true")
	flag.Set("v", "5")        //Info,Error...這種的不用set v就印的出來, glog.V(3)...這種的要set v才印的出來
	flag.Set("log_dir", "gg") //put glog log data into "gg" folder, (注意要先把資料夾創出來!)
	//stderrthreshold确保了只有大于或者等于该级别的日志才会被输出到stderr中，也就是标准错误输出中，默认为ERROR。当设置为FATAL时候，不会再有任何error信息的输出。

	flag.Parse() //解析上面的set。 after parse(), so that your flag.set start effected
}

func usage() {
	flag.PrintDefaults()
	os.Exit(2)
}

func refacLater() {

	router := routers.InitRouter()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: router,
		// ReadTimeout:    ReadTimeout,
		// WriteTimeout:   WriteTimeout,
		// MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
