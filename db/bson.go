package db

import (
	"iii/ifactory/compute/util"
	"iii/ifactory/compute/util/util_business"
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
	ProcessingStatusCode  *int          `json:"ProcessingStatusCode,omitempty" bson:"ProcessingStatusCode,omitempty"`
	ProcessingProgress    string        `json:"ProcessingProgress" bson:"ProcessingProgress,omitempty"`
	ShouldRepairTime      time.Time     `json:"ShouldRepairTime" bson:"ShouldRepairTime,omitempty"`
	PlanRepairTime        time.Time     `json:"PlanRepairTime" bson:"PlanRepairTime,omitempty"`
	// AbnormalCode       string    `json:"AbnormalCode" bson:"AbnormalCode,omitempty"`
	// AbnormalReason     string    `json:"AbnormalReason" bson:"AbnormalReason,omitempty"`
	// AbnormalPosition   string    `json:"AbnormalPosition" bson:"AbnormalPosition,omitempty"`
	// AbnormalSolution   string    `json:"AbnormalSolution" bson:"AbnormalSolution,omitempty"`
	// Timestamp time.Time `json:"Timestamp" bson:"Timestamp,omitempty"`
	Timestamp time.Time `json:"Timestamp,omitempty" bson:"Timestamp,omitempty"`
}

func (o *AbnormalMachineLatest) SetDefaultValue() {
	o.UpdateTime = util.GetNow()
	o.AbnormalStartTime = o.Timestamp
	o.ShouldRepairTime = util_business.GetRepairTime(util.GetNow())
	o.PlanRepairTime = util_business.GetRepairTime(util.GetNow())
	o.Type = "data"
	o.ProcessingProgress = "未指派人員"
	o.ProcessingStatusCode = new(int)
}

//update abnormal lasting second and update time
func (o *AbnormalMachineLatest) UpdateSomeValue() {
	//更新錯誤持續時間 (現在時間-錯誤發生時間)
	o.AbnormalLastingSecond = util.DurationToSecs(util.GetNow().Sub(o.AbnormalStartTime))
	o.UpdateTime = util.GetNow()
}

type AbnormalMachineHist struct {
	Id                    bson.ObjectId `json:"_id,omitempty" bson:"id,omitempty"`
	Type                  string        `json:"Type" bson:"Type,omitempty"`
	GroupID               string        `json:"GroupID" bson:"GroupID,omitempty"`
	GroupName             string        `json:"GroupName" bson:"GroupName,omitempty"`
	MachineID             string        `json:"MachineID" bson:"MachineID,omitempty"`
	MachineName           string        `json:"MachineName" bson:"MachineName,omitempty"`
	AbnormalStartTime     time.Time     `json:"AbnormalStartTime" bson:"AbnormalStartTime,omitempty"`
	AbnormalLastingSecond float64       `json:"AbnormalLastingSecond" bson:"AbnormalLastingSecond,omitempty"`
	ShouldRepairTime      time.Time     `json:"ShouldRepairTime" bson:"ShouldRepairTime,omitempty"`
	PlanRepairTime        time.Time     `json:"PlanRepairTime" bson:"PlanRepairTime,omitempty"`
	RepairReceiveTime     time.Time     `json:"RepairReceiveTime" bson:"RepairReceiveTime,omitempty"`
	// UpdateTime            time.Time     `json:"UpdateTime" bson:"UpdateTime,omitempty"`
}
