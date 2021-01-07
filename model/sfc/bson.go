package model

import (
	"iii/ifactory/compute/util"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type WorkOrder struct {
	Id            string    `json:"id,omitempty" bson:"_id,omitempty"`
	WorkOrderId   string    `json:"WorkOrderId,omitempty" bson:"WorkOrderId"`
	Station       string    `json:"Station,omitempty" bson:"Station" validate:"required"`
	Machine       string    `json:"Machine,omitempty" bson:"Machine" validate:"required"`
	Status        string    `json:"Status,omitempty" bson:"Status" validate:"required"`
	Good          float64   `json:"Good,omitempty" bson:"Good"`
	NonGood       float64   `json:"NonGood,omitempty" bson:"NonGood"`
	Quantity      float64   `json:"Quantity,omitempty" bson:"Quantity" validate:"required"`
	Reporter      string    `json:"Reporter,omitempty" bson:"Reporter" validate:"required"`
	PlanStartDate time.Time `json:"PlanStartDate,omitempty" bson:"PlanStartDate" validate:"required"`
	CreateAt      time.Time `json:"CreateAt,omitempty" bson:"CreateAt,omitempty"`
}

func (o *WorkOrder) CreateWithDefault() {
	o.CreateAt = util.GetNow()
	o.WorkOrderId = func() string {
		layout := "20060102"
		t := o.CreateAt.Format(layout)
		id := bson.NewObjectId().Hex()
		return t + id
	}()
}
