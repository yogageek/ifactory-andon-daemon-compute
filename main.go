package main

import (
	"flag"
	"iii/ifactory/compute/db"
	"iii/ifactory/compute/logic"
	"iii/ifactory/compute/setenv"
	"os"
)

//其他包的init會先執行，因为在编译的时候会先去检查导入的包，首先发现其他包里面的init，然后才会到main包里面的init
func init() {
	setFlag()
	os.Setenv("DEBUG", "true")
	setenv.SetEnv()

	db.StartMongo()
}

var port = ":8080"

func main() {
	// glog.Infoln("VERSION:", "v1.0.0 ", port) //改從build args拿
	// go thesso.RunTokenLoop()
	// time.Sleep(2 * time.Second)
	doSth()

	// r := middleware.Router()
	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowCredentials: true,
	// 	AllowedMethods:   []string{"GET", "OPTIONS"},
	// })

	// handler := c.Handler(r)

	// log.Fatal(http.ListenAndServe(port, handler))
}

func doSth() {
	// logic.RunDaemonLoop()

	// r := db.FindAll(db.AMRawdata)
	// util.PrintJson(r)
	logic.RunDaemonLoop()
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
