package andon_logic

import (
	"encoding/json"
	"fmt"
	"iii/ifactory/compute/db"

	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
)

var (
	fromCollection = db.AMRawdata
	toCollection   = db.AMLatest
	//異常機器對應值
	theAbnormalColumn = "StatusLay1Value"
	theAbnormalValue  = 2000
)

//select abnormal data from raw data by option key and value
func selectAbnormalData() []bson.M {
	search1 := bson.M{"$match": bson.M{theAbnormalColumn: theAbnormalValue}}
	r := db.MatchBySearchOptions(fromCollection, search1)
	return r
}

//dynamic query 抽出

func bsonToByte(bDatas interface{}) []byte {
	//[]bson to []byte
	abmDatas_byte, err := bson.MarshalJSON(bDatas)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(abmDatas_byte))
	return abmDatas_byte
}

// func byteToStruct(byt []byte) interface{} {
// 	//[]byte to []struct
// 	var machineRawDatas []db.MachineRawData
// 	err = json.Unmarshal(abmDatas_byte, &machineRawDatas)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

//bson raw testing

//insert selected data into collection
func Processing() {
	//get []bson
	bDatas := selectAbnormalData()
	// fmt.Println(bDatas)

	//bDatas撈出來後要加一些欄位(重要規劃 之後可能常修改)
	byt := bsonToByte(bDatas)

	//clone rawdata to latestdata type
	var theStruct []db.AbnormalMachineLatest
	err := json.Unmarshal(byt, &theStruct)
	if err != nil {
		glog.Error(err)
	}

	fmt.Printf("the struct: %+v\n", theStruct)
	// util.PrintJson(theStruct)
	//the important logic
	for _, s := range theStruct {
		s.ProcessingStatusCode = 0
		//bson一定要先轉byte才能在轉struct
		KEY := "MachineID" //upsert條件, 如果machineId一樣就更新，不一樣就新增
		query := bson.M{KEY: s.MachineID}

		//sturct to bson(key變小寫了)
		b, _ := structToBson(s)
		fmt.Println(b)

		db.UpsertByOption(toCollection, query, b)
	}
	return
	//用bson迴圈
	//來不及的話就先給peter這段------------->>>>>
	fmt.Println("------------bson style")
	for _, b := range bDatas {
		//bson一定要先轉byte才能在轉struct
		KEY := "MachineID" //upsert條件, 如果machineId一樣就更新，不一樣就新增
		query := bson.M{KEY: b[KEY]}
		db.UpsertByOption(toCollection, query, b)
	}
	//<<<<來不及的話就先給peter這段-------------
	return

	// -----//data=byte
	// var value interface{} = bson.M{"some": "value"}
	// data, err := bson.Marshal(value)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// -----// byte to raw
	// var doc bson.Raw
	// err = bson.Unmarshal(data, &doc)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//----------

	//[]bson to []byte
	abmDatas_byte, err := bson.MarshalJSON(bDatas)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(abmDatas_byte))

	//[]byte to []struct
	var machineRawDatas []db.MachineRawData
	err = json.Unmarshal(abmDatas_byte, &theStruct)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", machineRawDatas)

	fmt.Println("------------struct style")
	for _, data := range theStruct {
		dynamicQuery := db.MachineRawData{
			MachineID: data.MachineID,
		}
		doc, _ := structToBson(dynamicQuery)
		fmt.Println("doc(query bson):", doc)

		// struct to []byte
		// dynamicQueryJson, _ := json.Marshal(dynamicQuery)
		// fmt.Println(string(dynamicQueryJson))

		//[]byte to bson
		// var dynamicQueryBson bson.M
		// err := bson.UnmarshalJSON(dynamicQueryJson, &dynamicQueryBson)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(dynamicQueryBson) //沒辦法把Id:ObjectIdHex("5fe2f22a414c2f0b5149aa8b") 轉bson 會漏

		//debugging
		// data.StatusMapValue = 3000
		fmt.Println("data:", data)

		// query := bson.M{db.MachineRawData.Id: data.MachineId}
		db.UpsertByOption(toCollection, *doc, data)
	}

}

//struct to bson
func structToBson(v interface{}) (doc *bson.M, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

// //bson to struct
// func toStruct(v interface{}) (doc *bson.M, err error) {
// 	data, err := bson.Marshal(v)
// 	if err != nil {
// 		return
// 	}

// 	err = bson.Unmarshal(data, &doc)
// 	return
// }
