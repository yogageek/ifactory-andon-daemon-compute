package model

import (
	"encoding/json"

	// . "iii/ifactory/compute/model/sfc"

	"github.com/golang/glog"
)

type JsonParser interface {
	ParseToObject(body []byte) WorkOrder
	ParseToInterface(body []byte) interface{}
}

type Interfacer struct {
}

// type WorkOrder2 struct {
// }

//只是為了都要實作
func (i Interfacer) ParseToObject(body []byte) {
	return
}

func (i Interfacer) ParseToInterface(body []byte) interface{} {
	var v interface{}
	if err := json.Unmarshal(body, &v); err != nil {
		glog.Error(err)
	}
	return v
}

func (o WorkOrder) ParseToObject(body []byte) WorkOrder {
	var wo WorkOrder
	if err := json.Unmarshal(body, &wo); err != nil {
		glog.Error(err)
	}
	return wo
}

//只是為了都要實作
func (o WorkOrder) ParseToInterface(body []byte) interface{} {
	var v interface{}
	return v
}

func t() {
	var jp JsonParser
	jp = new(WorkOrder) //WorkOrder也要實作parse to interface才能使用

	// o.ParseToObject()
	// j.ParseToObject()
}

// 實作做其他類型的parse to object
// func (o WorkOrder) ParseToObject(body []byte) WorkOrder {
// 	var wo WorkOrder
// 	if err := json.Unmarshal(body, &wo); err != nil {
// 		glog.Error(err)
// 	}
// 	return wo
// }
