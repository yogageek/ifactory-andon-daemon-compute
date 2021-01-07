package compute_logic

import (
	"fmt"

	"iii/ifactory/compute/db"
	"iii/ifactory/compute/model"
	"iii/ifactory/compute/util"

	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
)

type AmHist struct {
	sourceCollection string
	targetCollection string
}

//注意from/to collection分別操縱的data
var (
	amHist = AmHist{
		sourceCollection: model.AMLatest,
		targetCollection: model.AMHist,
	}
)

const (
	//抓指定CODE  --->拉去enum
	ProcessStatusCodeColumn = "ProcessingStatusCode"
	ProcessStatusCodeValue  = 5
)

func Daemon_AbnormalMachineHist() {
	amHist.Inserting()
	amHist.Deleteing()
	fmt.Println("------Daemon_AbnormalMachineHist--------")
}

//select abnormal data from raw data by option key and value
func (o AmHist) FindHistDataFromLatest(column string, value int) (datas []model.AbnormalMachineHist) {
	search := []bson.M{
		bson.M{
			"$match": bson.M{column: value},
		},
	}
	err := db.Match(o.sourceCollection, search, &datas)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
	return datas
}

//將latest data依照指定條件拷貝到hist data
func (o AmHist) Inserting() {

	//select
	datas := o.FindHistDataFromLatest(ProcessStatusCodeColumn, ProcessStatusCodeValue)
	// fmt.Printf("AbnormalMachineHistData: %+v\n", datas)

	//logic. insert data into collection base on query(if MachineID not exist)
	func() {
		for _, data := range datas {
			// fmt.Printf("AbnormalMachineHistData: %+v\n", data)

			// 以ID當作primaryKey當作query條件
			// query := bson.M{
			// 	db.QueryKey.ID: data.Id, //upsert條件, 如果Id一樣就更新，不一樣就新增
			// }

			// //set on insert
			// valueAndOption := bson.M{
			// 	"$setOnInsert": data, //如果不存在就insert
			// }
			db.Insert(o.targetCollection, data)
		}
	}()
}

func (o AmHist) Deleteing() {
	//delete latest collection documents by options
	selector := bson.M{ProcessStatusCodeColumn: ProcessStatusCodeValue}
	db.Remove(o.sourceCollection, selector)
}
