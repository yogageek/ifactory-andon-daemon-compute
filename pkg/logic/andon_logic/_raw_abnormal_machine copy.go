package andon_logic

import (
	"fmt"
	"iii/ifactory/compute/db"
	"iii/ifactory/compute/util"

	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
)

//注意from/to collection分別操縱的data

var (
	fromCollection = db.AMRawdata
	toCollection   = db.AMLatest
	//異常機器對應值  --->拉去enum
	theAbnormalColumn = "StatusLay1Value"
	theAbnormalValue  = 2000
)

func Do() {
	Processing()
	Updating()
}

//select abnormal data from raw data by option key and value
func selectAbnormalDataFromRaw() []interface{} {
	search1 := bson.M{"$match": bson.M{theAbnormalColumn: theAbnormalValue}}
	r := db.MatchBySearchOptionsI(fromCollection, search1)
	return r
}

//select abnormal data
func selectAbnormalData() []interface{} {
	return db.QueryAll(toCollection)
}

//dynamic query 抽出

func bsonToByte(bDatas interface{}) []byte {
	//[]bson to []byte

	//bson.Marshal // 會亂碼
	abmDatas_byte, err := bson.MarshalJSON(bDatas)
	if err != nil {
		glog.Error(err)
	}
	// fmt.Println(string(abmDatas_byte))
	return abmDatas_byte
}

//將raw data依照指定條件拷貝到abnormal data
func Processing() {

	// bson一定要先轉byte才能在轉struct

	//get []bson in collection
	bsonData := selectAbnormalDataFromRaw()
	// util.PrintJson(bsonData)

	// bson to byte
	//撈出來後要加一些欄位(之後可能常修改)
	bsonByte := bsonToByte(bsonData)
	// util.PrintReq(bsonByte)

	// bson.UnmarshalJSON to unmarshal rawdata to latestdata type
	var abnormalMachineLatestData []db.AbnormalMachineLatest
	err := bson.UnmarshalJSON(bsonByte, &abnormalMachineLatestData)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
	// util.PrintJson(abnormalMachineLatestData)

	//loop selected data
	for _, data := range abnormalMachineLatestData {
		// fmt.Printf("abnormalMachineLatestData: %+v\n", data)

		//logic 1. business
		data.SetDefaultValue()

		// 以MachineID當作primaryKey當作query條件
		query := bson.M{
			db.PrimaryKey: data.MachineID, //upsert條件, 如果machineId一樣就更新，不一樣就新增
		}

		//logic 2. insert data into collection base on query(if MachineID not exist)
		func() {
			//set on insert
			valueAndOption := bson.M{
				"$setOnInsert": data,
			}
			db.UpsertByOption(toCollection, query, valueAndOption)
			// 同上寫法
			// i, err := db.MongoDB.UseC(toCollection).Upsert(query, updateAndoption)
			// util.PrintJson(i)
			// if err != nil {
			// 	util.Cerr(err)
			// }
		}()
	}
}

//定時更新abnormal collection
func Updating() {
	bsonData := selectAbnormalData()
	// for _, v := range bsonData {
	// 	fmt.Println(v)
	// }
	// bson to byte
	// 撈出來後要加一些欄位
	bsonByte := bsonToByte(bsonData)

	// bson.UnmarshalJSON to unmarshal rawdata to latestdata type
	abnormalMachineLatestData := []db.AbnormalMachineLatest{}
	err := bson.UnmarshalJSON(bsonByte, &abnormalMachineLatestData)
	if err != nil {
		glog.Error(util.Cerr(err))
	}
	// util.PrintJson(abnormalMachineLatestData)

	// abnormalMachineLatestData2 := []db.AbnormalMachineLatest{}
	// db.FindAllI(toCollection, nil, nil, &abnormalMachineLatestData2)
	// fmt.Printf("abnormalMachineLatestData2: %+v\n", abnormalMachineLatestData2)
	// return

	//loop selected data
	for _, data := range abnormalMachineLatestData {
		// fmt.Printf("abnormalMachineLatestData: %+v\n", data)

		//logic 2. update abnormal lasting second and update time
		func() {
			//更新錯誤持續時間 (現在時間-錯誤發生時間)
			now := util.GetNow()
			abnormalLastingSecond := now.Sub(data.AbnormalStartTime) //#attention! can's use time.since coz it use utc+0
			fmt.Println("機器名稱:", data.MachineName)
			fmt.Println("錯誤發生時間:", data.AbnormalStartTime)
			fmt.Println("現在時間:", util.GetNow())
			fmt.Println("錯誤持續時間:", abnormalLastingSecond)

			valueAndOption := bson.M{
				"$set": db.AbnormalMachineLatest{ //bson.go 的tag bson要帶omitempty, 否則會用default空值寫入
					UpdateTime:            now,
					AbnormalLastingSecond: util.DurationToSecs(abnormalLastingSecond),
				},
			}
			// util.PrintJson(valueAndOption)
			db.UpdateByOption(toCollection, nil, valueAndOption)
		}()
	}
	fmt.Println("done")
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
