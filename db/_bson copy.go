package db

import (
	"fmt"
	"iii/ifactory/compute/util"
	"iii/ifactory/compute/util/util_business"
	"reflect"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//called by other pkg
var (
	PrimaryKey = "MachineID"
)

type MachineRawData struct {
	// Id bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	// ID              string `json:"_id"`
	Id              bson.ObjectId `json:"_id,omitempty" bson:"id,omitempty"`
	GroupID         string        `json:"GroupID"`
	GroupName       string        `json:"GroupName"`
	MachineID       string        `json:"MachineID"`
	MachineName     string        `json:"MachineName"`
	StatusLay1Value int           `json:"StatusLay1Value"`
	StatusMapValue  int           `json:"StatusMapValue"`
	StatusRawValue  int           `json:"StatusRawValue"` //不用也可以拿到
	Timestamp       time.Time     `json:"Timestamp" bson:"Timestamp"`
}

// ps:
// StatusMapValue int           `json:"statusMapValue,omitempty" bson:"statusMapValue,omitempty"` //轉json會轉成這樣
//研究ID型態差異

type AbnormalMachineLatest struct {
	// ID string `json:"_id"`
	// Id                    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Id                    bson.ObjectId `json:"_id,omitempty" bson:"id,omitempty"`
	UpdateTime            time.Time     `json:"UpdateTime" bson:"UpdateTime,omitempty"`
	Type                  string        `json:"Type" bson:"Type,omitempty"`
	GroupID               string        `json:"GroupID" bson:"GroupID,omitempty"`
	GroupName             string        `json:"GroupName" bson:"GroupName,omitempty"`
	MachineID             string        `json:"MachineID" bson:"MachineID,omitempty"`
	MachineName           string        `json:"MachineName" bson:"MachineName,omitempty"`
	AbnormalStartTime     time.Time     `json:"AbnormalStartTime" bson:"AbnormalStartTime,omitempty"`
	AbnormalLastingSecond float64       `json:"AbnormalLastingSecond" bson:"AbnormalLastingSecond,omitempty"`
	ProcessingStatusCode  int           `json:"ProcessingStatusCode" bson:"ProcessingStatusCode"`
	ProcessingProgress    string        `json:"ProcessingProgress" bson:"ProcessingProgress,omitempty"`
	ShouldRepairTime      time.Time     `json:"ShouldRepairTime" bson:"ShouldRepairTime,omitempty"`
	PlanRepairTime        time.Time     `json:"PlanRepairTime" bson:"PlanRepairTime,omitempty"`
	// AbnormalCode       string    `json:"AbnormalCode" bson:"AbnormalCode,omitempty"`
	// AbnormalReason     string    `json:"AbnormalReason" bson:"AbnormalReason,omitempty"`
	// AbnormalPosition   string    `json:"AbnormalPosition" bson:"AbnormalPosition,omitempty"`
	// AbnormalSolution   string    `json:"AbnormalSolution" bson:"AbnormalSolution,omitempty"`
	// Timestamp time.Time `json:"Timestamp" bson:"Timestamp,omitempty"`

}

func (o *AbnormalMachineLatest) SetDefaultValue() {
	o.UpdateTime = util.GetNow()
	o.AbnormalStartTime = util.GetNow()
	o.ShouldRepairTime = util_business.GetRepairTime(util.GetNow())
	o.PlanRepairTime = util_business.GetRepairTime(util.GetNow())
	o.Type = "data"
	o.ProcessingProgress = "未指派人員"
	// o.ProcessingStatusCode = 0
}

//update abnormal lasting second and update time
func (o *AbnormalMachineLatest) UpdateSomeValue() {
	//更新錯誤持續時間 (現在時間-錯誤發生時間)
	o.AbnormalLastingSecond = util.DurationToSecs(util.GetNow().Sub(o.AbnormalStartTime))
	o.UpdateTime = util.GetNow()
}

// //db.FindAllI(toCollection, nil, nil, &abnormalMachineLatestData2)會進來這
// func (s *AbnormalMachineLatest) SetBSON(raw bson.Raw) error {
// 	fmt.Println(s.AbnormalStartTime)
// 	return raw.Unmarshal(s)
// }

// type AbnormalMachineLatest struct {
// 	Id             bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
// 	StatusMapValue int           `json:"statusMapValue,omitempty" bson:"statusMapValue,omitempty"` //轉json會轉成這樣
// 	MachineId      string        `bson:"MachineID,omitempty"`                                      //不用也可以拿到
// }

// 會影響到bson時區轉換 原本應utc+8 轉出來變成utc+0
//initial default value logic-----------------------
// func (o *AbnormalMachineLatest) UnmarshalJSON(r []byte) error {
// 	// Create a new type from the target type to avoid recursion.
// 	type ob AbnormalMachineLatest

// 	// Unmarshal into an instance of the new type.
// 	var obStruct ob

// 	err := bson.UnmarshalJSON(r, &obStruct) // attention!
// 	if err != nil {
// 		glog.Error(err)
// 	}

// 	// //parse time 錯誤 試試ensaas寫法
// 	// type abnormalmachinelatest AbnormalMachineLatest
// 	// s := &abnormalmachinelatest{
// 	// 	UpdateTime:           util.GetNow(),
// 	// 	AbnormalStartTime:    util.GetNow(),
// 	// 	ShouldRepairTime:     util_business.GetRepairTime(util.GetNow()),
// 	// 	PlanRepairTime:       util_business.GetRepairTime(util.GetNow()),
// 	// 	Type:                 "Data",
// 	// 	ProcessingStatusCode: 0,
// 	// 	ProcessingProgress:   "未指派人員",
// 	// }

// 	// if obStruct.AbnormalLastingSecond==0{
// 	// 	ob.
// 	// }

// 	fmt.Printf("UnmarshalJSON struct: %+v\n", obStruct)

// 	// fmt.Println(o)
// 	// fmt.Println(s)
// 	*o = AbnormalMachineLatest(obStruct)
// 	// o===s
// 	// util.PrintJson(o)
// 	// util.PrintJson(s)
// 	return nil
// }

// type MyString string

//自定義tag用法-----------------------
// declaring a person struct
type Person struct {

	// setting the default value
	// of name to "geek"
	name string `default:"geek"`
}

func default_tag(p Person) string {

	// TypeOf returns type of
	// interface value passed to it
	typ := reflect.TypeOf(p)

	// checking if null string
	if p.name == "" {

		// returns the struct field
		// with the given parameter "name"
		f, _ := typ.FieldByName("name")

		// returns the value associated
		// with key in the tag string
		// and returns empty string if
		// no such key in tag
		p.name = f.Tag.Get("default")
	}

	return fmt.Sprintf("%s", p.name)
}

// main function
func main() {

	// prints out the default name
	fmt.Println("Default name is:", default_tag(Person{}))
}
