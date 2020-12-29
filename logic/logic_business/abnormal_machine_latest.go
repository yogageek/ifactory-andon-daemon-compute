package logic_business

import (
	"fmt"
	"iii/ifactory/compute/db"
	"iii/ifactory/compute/util"

	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
)

type AmLatest struct {
	sourceCollection string
	targetCollection string
}

//注意from/to collection分別操縱的data
var (
	amLatest = AmLatest{
		sourceCollection: db.AMRawdata,
		targetCollection: db.AMLatest,
	}
)

const (
	//異常機器對應值  --->拉去enum
	AbnormalColumn = "StatusLay1Value"
	AbnormalValue  = 2000

	//異常機器對應值(平板)  --->拉去enum
	AbnormalColumnForPanel = "TabletStatusValue"
	AbnormalValueForPanel  = 1
)

// {"$or", []interface{}{
// 	bson.D{{"key2", 2}},
// 	bson.D{{"key3", 2}},
// }},

func Daemon_AbnormalMachineLatest() {
	amLatest.Inserting()
	amLatest.Updating()
	fmt.Println("------Daemon_AbnormalMachineLatest------")
}

//select abnormal data from raw data by option key and value
func (o AmLatest) FindAbnormalDataFromRaw(abnormalColumn string, abnormalValue int) (datas []db.AbnormalMachineLatest) {
	search := []bson.M{
		bson.M{
			"$match": bson.M{abnormalColumn: abnormalValue},
		},
	}
	err := db.Match(o.sourceCollection, search, &datas)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
	return datas
}

//select abnormal data
func (o AmLatest) FindAbnormalData() (datas []db.AbnormalMachineLatest) {
	err := db.FindAll(o.targetCollection, nil, nil, &datas)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
	return datas
}

//將raw data依照指定條件拷貝到abnormal data
func (o AmLatest) Inserting() {

	//select (將AMRawdata的bsonData轉存到AbnormalMachineLatest物件)
	datas := o.FindAbnormalDataFromRaw(AbnormalColumn, AbnormalValue)
	// fmt.Printf("abnormalMachineLatestData: %+v\n", datas)

	//logic. insert data into collection base on query(if MachineID not exist)
	func() {
		for _, data := range datas {
			// fmt.Printf("abnormalMachineLatestData: %+v\n", data)
			data.SetDefaultValue()

			// 以MachineID當作primaryKey當作query條件
			query := bson.M{
				db.PrimaryKey: data.MachineID, //upsert條件, 如果machineId一樣就更新，不一樣就新增
			}

			//set on insert
			valueAndOption := bson.M{
				"$setOnInsert": data, //如果不存在就insert
			}
			db.Upsert(o.targetCollection, query, valueAndOption)
		}
	}()

	datas = o.FindAbnormalDataFromRaw(AbnormalColumnForPanel, AbnormalValueForPanel)
	//logic. insert data into collection base on query(if MachineID not exist)
	func() {
		for _, data := range datas {
			// fmt.Printf("abnormalMachineLatestData: %+v\n", data)
			data.SetDefaultValue()

			// 以MachineID當作primaryKey當作query條件
			query := bson.M{
				db.PrimaryKey: data.MachineID, //upsert條件, 如果machineId一樣就更新，不一樣就新增
			}

			//set on insert
			valueAndOption := bson.M{
				"$setOnInsert": data, //如果不存在就insert
			}
			db.Upsert(o.targetCollection, query, valueAndOption)
		}
	}()
}

//定時更新collection
func (o AmLatest) Updating() {

	//select
	datas := o.FindAbnormalData()
	// fmt.Printf("abnormalMachineLatestData2: %+v\n", abnormalMachineLatestData)

	//logic
	func() {
		for _, data := range datas {
			// fmt.Printf("abnormalMachineLatestData: %+v\n", data)
			data.UpdateSomeValue()
			fmt.Println("機器名稱:", data.MachineName)
			fmt.Println("錯誤發生時間:", data.AbnormalStartTime)
			fmt.Println("現在時間:", util.GetNow())
			fmt.Println("錯誤持續時間:", data.AbnormalLastingSecond)
			valueAndOption := bson.M{
				"$set": db.AbnormalMachineLatest{ //bson.go 的tag bson要帶omitempty, 否則會用default空值寫入
					UpdateTime:            data.UpdateTime,
					AbnormalLastingSecond: data.AbnormalLastingSecond,
				},
			}
			// util.PrintJson(valueAndOption)
			db.Update(o.targetCollection, nil, valueAndOption)
		}
	}()
}

//sturct to bson(key變小寫了)
// b, _ := structToBson(s)
// fmt.Println(b)

//struct to bson
func structToBson(v interface{}) (doc *bson.M, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &doc)
	return
}

// func byteToStruct(byt []byte) interface{} {
// 	//[]byte to []struct
// 	var machineRawDatas []db.MachineRawData
// 	err = json.Unmarshal(abmDatas_byte, &machineRawDatas)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
