package model

import (
	"iii/ifactory/compute/util"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// 工單資訊
type WorkOrderInfo struct {
	Sum        float64 `json:"Sum,omitempty" bson:"Sum"`
	SumGood    float64 `json:"SumGood,omitempty" bson:"SumGood"`
	SumNonGood float64 `json:"SumNonGood,omitempty" bson:"SumNonGood"`
	GoodRate   float64 `json:"GoodRate,omitempty" bson:"GoodRate"`
	Status     string  `json:"Status,omitempty" bson:"GoodRate"`
	WorkOrder
}

type JB struct {
	WorkOrder     *WorkOrder     `json:"WorkOrder,omitempty" bson:"WorkOrder,omitempty"`
	WorkOrderList *WorkOrderList `json:"WorkOrderList,omitempty" bson:"WorkOrderList,omitempty"`
}

//工單
type WorkOrder struct {
	Id          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	WorkOrderID string        `json:"WorkOrderID,omitempty" bson:"WorkOrderID"`

	Quantity      float64    `json:"Quantity,omitempty" bson:"Quantity,omitempty" validate:"required"` //預計生產數量
	PlanStartDate *time.Time `json:"PlanStartDate,omitempty" bson:"PlanStartDate,omitempty"`
	DeliverAt     *time.Time `json:"DeliverAt,omitempty" bson:"DeliverAt,omitempty"`

	Product       *Product         `json:"Product,omitempty" bson:"Product,omitempty" validate:"required"`
	WorkOrderList []*WorkOrderList `json:"WorkOrderList,omitempty" bson:"WorkOrderList,omitempty"`

	CreateAt *time.Time `json:"CreateAt,omitempty" bson:"CreateAt,omitempty"`
}

//報工單
type WorkOrderList struct {
	Id          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	WorkOrderID string        `json:"WorkOrderID,omitempty" bson:"WorkOrderID"`

	CompletedQty float64 `json:"CompletedQty,omitempty" bson:"CompletedQty,omitempty"` //注意!!!不加empty會嚴重影響到push功能而且查不出來!因為c.Update(selector, updated) 系統會自動把selector=interface{}轉成bson,沒帶omitempty會parse成"""
	NonGoodQty   float64 `json:"NonGoodQty,omitempty" bson:"NonGoodQty,omitempty"`
	StationName  string  `json:"StationName,omitempty" bson:"StationName,omitempty" validate:"required"` //注意!!!不加empty會嚴重影響到push功能而且查不出來!
	Reporter     string  `json:"Reporter,omitempty" bson:"Reporter,omitempty" validate:"required"`

	CreateAt *time.Time `json:"CreateAt,omitempty" bson:"CreateAt,omitempty"`
}

//Product Setting
type Product struct {
	Id          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	ProductID   string        `json:"ProductID,omitempty" bson:"ProductID,omitempty"`
	ProductName string        `json:"ProductName,omitempty" bson:"ProductName,omitempty" validate:"required"`
	StationType []string      `json:"StationType,omitempty" bson:"StationType"`
	Unit        string        `json:"Unit,omitempty" bson:"Unit,omitempty"`

	CreateAt time.Time `json:"CreateAt,omitempty" bson:"CreateAt,omitempty"`
}

func (o *WorkOrder) CreateWithDefault() {

	o.CreateAt = func() *time.Time {
		now := util.GetNow()
		return &now
	}()
	//工單產生ID
	o.WorkOrderID = func() string {
		layout := "20060102"
		t := o.CreateAt.Format(layout)
		id := bson.NewObjectId().Hex()
		return t + "-" + id
	}()
	//同時賦予報工單ID
	for _, oo := range o.WorkOrderList {
		oo.CreateWithDefault(o.WorkOrderID)
	}
}

func (o *WorkOrderList) CreateWithDefault(workorderId string) {
	o.Id = bson.NewObjectId()
	o.WorkOrderID = workorderId
	o.CreateAt = func() *time.Time {
		now := util.GetNow()
		return &now
	}()
}
