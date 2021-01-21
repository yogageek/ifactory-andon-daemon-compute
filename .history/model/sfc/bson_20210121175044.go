package model

import (
	"fmt"
	"iii/ifactory/compute/util"
	"math"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// tutorial
// func (m *StatsInfo) UnmarshalBSON(data []byte) error {
// 	// Unmarshal into a temporary type where the "ends" field is a string.
// 	decoded := new(struct {
// 		ID   primitive.ObjectID `bson:"_id"`
// 		Ends string             `bson:"ends"`
// 	})
// 	if err := bson.Unmarshal(data, decoded); err != nil {
// 		return err
// 	}

// 	return nil
// }

func (c StatsInfo) GetBSON() (interface{}, error) {
	// f := c.Value().Float64()
	// return struct {
	// 	Value        float64 `json:"value" bson:"value"`
	// 	CurrencyCode string  `json:"currencyCode" bson:"currencyCode"`
	// }{
	// 	Value:        f,
	// 	CurrencyCode: c.currencyCode,
	// }, nil
	return nil, nil
}

// SetBSON implements bson.Setter.
func (c *StatsInfo) SetBSON(raw bson.Raw) error {

	// decoded := new(struct { })
	type newStatsInfo StatsInfo
	s := new(newStatsInfo)

	bsonErr := raw.Unmarshal(s)
	if bsonErr != nil {
		return bsonErr
	}
	util.PrintJson(s)

	s.RealCompletedRate = func() float64 {
		if r := (s.CompletedQty / s.Quantity) * 100; !math.IsNaN(r) {
			return r
		}
		return 0
	}()
	s.Status = func() int {
		var status int
		switch s.CompletedQty < s.Quantity {
		case true:
			return
		case true:
		case true:
		}
		return status
	}()

	return nil
}
func (p *StatsInfo) UnmarshalJSON(data []byte) (err error) {

	// var res ProfileJSON
	fmt.Println("x")
	// if err := json.Unmarshal(data, &res); err != nil {
	// 	return err
	// }
	// return nil
	return nil
}

//-------------
type JB struct {
	WorkOrder     *WorkOrder     `json:"WorkOrder,omitempty" bson:"WorkOrder,omitempty"`
	WorkOrderList *WorkOrderList `json:"WorkOrderList,omitempty" bson:"WorkOrderList,omitempty"`
}

type StatsInfo struct {
	WorkOrderId string `json:"WorkOrderId,omitempty" bson:"WorkOrderId"`
	StationName string `json:"StationName,omitempty" bson:"StationName"`

	CompletedQty float64 `json:"CompletedQty,omitempty" bson:"CompletedQty"`
	NonGoodQty   float64 `json:"NonGoodQty,omitempty" bson:"NonGoodQty"`
	GoodQty      float64 `json:"GoodQty,omitempty" bson:"GoodQty"`
	Quantity     float64 `json:"Quantity,omitempty" bson:"Quantity,omitempty"` //預計生產數量

	RealCompletedRate float64 `json:"RealCompletedRate,omitempty" bson:"RealCompletedRate"`

	// future
	// orders list count of the station

	Status string `json:"Status,omitempty" bson:"Status"`
}

// 工單資訊
type WorkOrderInfo struct {
	WorkOrder
	Stations []Station `json:"Stations,omitempty" bson:"Stations"`

	CompletedQty float64 `json:"CompletedQty,omitempty" bson:"CompletedQty"`
	GoodQty      float64 `json:"GoodQty,omitempty" bson:"GoodQty"`
	NonGoodQty   float64 `json:"NonGoodQty,omitempty" bson:"NonGoodQty"`

	GoodProductQty float64 `json:"GoodProductQty,omitempty" bson:"GoodProductQty"`

	GoodQtyRate     float64 `json:"GoodQtyRate,omitempty" bson:"GoodQtyRate"`
	GoodProductRate float64 `json:"GoodProductRate,omitempty" bson:"GoodProductRate"`

	Status string `json:"Status,omitempty" bson:"Status"`
}

//工單
type WorkOrder struct {
	Id          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	WorkOrderId string        `json:"WorkOrderId,omitempty" bson:"WorkOrderId"`

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
	WorkOrderId string        `json:"WorkOrderId,omitempty" bson:"WorkOrderId"`

	CompletedQty float64 `json:"CompletedQty,omitempty" bson:"CompletedQty,omitempty"` //注意!!!不加empty會嚴重影響到push功能而且查不出來!因為c.Update(selector, updated) 系統會自動把selector=interface{}轉成bson,沒帶omitempty會parse成"""
	NonGoodQty   float64 `json:"NonGoodQty,omitempty" bson:"NonGoodQty,omitempty"`
	StationName  string  `json:"StationName,omitempty" bson:"StationName,omitempty" validate:"required"` //注意!!!不加empty會嚴重影響到push功能而且查不出來!
	Reporter     string  `json:"Reporter,omitempty" bson:"Reporter,omitempty" validate:"required"`

	CreateAt *time.Time `json:"CreateAt,omitempty" bson:"CreateAt,omitempty"`
}

//Product Setting
type Product struct {
	Id          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	ProductId   string        `json:"ProductId,omitempty" bson:"ProductId,omitempty"`
	ProductName string        `json:"ProductName,omitempty" bson:"ProductName,omitempty" validate:"required"`
	StationType []string      `json:"StationType,omitempty" bson:"StationType"`
	Unit        string        `json:"Unit,omitempty" bson:"Unit,omitempty"`

	CreateAt time.Time `json:"CreateAt,omitempty" bson:"CreateAt,omitempty"`
}

type MyStations []struct {
	Station
}
type Station struct {
	Name         string
	GoodQty      float64
	NonGoodQty   float64
	CompletedQty float64
}

func (o *WorkOrder) CreateWithDefault() {

	o.CreateAt = func() *time.Time {
		now := util.GetNow()
		return &now
	}()
	//工單產生ID
	o.WorkOrderId = func() string {
		layout := "20060102"
		t := o.CreateAt.Format(layout)
		id := bson.NewObjectId().Hex()
		return t + "-" + id
	}()
	//同時賦予報工單ID
	for _, oo := range o.WorkOrderList {
		oo.CreateWithDefault(o.WorkOrderId)
	}
}

func (o *WorkOrderList) CreateWithDefault(workorderId string) {
	o.Id = bson.NewObjectId()
	o.WorkOrderId = workorderId
	o.CreateAt = func() *time.Time {
		now := util.GetNow()
		return &now
	}()
}

func (woi *WorkOrderInfo) CreateWorkOrderInfo(wo WorkOrder) {

	for _, wol := range wo.WorkOrderList {
		woi.CompletedQty = woi.CompletedQty + wol.CompletedQty
		woi.NonGoodQty = woi.NonGoodQty + wol.NonGoodQty
		woi.GoodQty = woi.CompletedQty - woi.NonGoodQty
	}
	woi.Stations = wo.calStationInfo()

	woi.GoodQtyRate = func() float64 {
		if r := woi.GoodQty / woi.CompletedQty; !math.IsNaN(r) {
			return r
		}
		return 0
	}()

	woi.GoodProductQty = func() (minGood float64) {
		minGood = woi.Stations[0].GoodQty
		for _, s := range woi.Stations {
			if minGood > s.GoodQty {
				minGood = s.GoodQty
			}
		}
		return
	}()

	woi.Status = func() string {
		if wo.Quantity <= woi.GoodProductQty {
			return "done"
		}
		return "available"
	}()

	woi.WorkOrder = wo
}

//----------->

func (o *WorkOrder) calStationInfo() (stations []Station) {
	wols := o.WorkOrderList

	//method1
	mStation := map[string]*Station{}
	for _, wol := range wols {
		s := wol.StationName

		if mStation[s] == nil {
			mStation[s] = &Station{
				Name: s,
			}
		}
		mStation[s].CompletedQty = mStation[s].CompletedQty + wol.CompletedQty
		mStation[s].NonGoodQty = mStation[s].NonGoodQty + wol.NonGoodQty
		mStation[s].GoodQty = mStation[s].CompletedQty - mStation[s].NonGoodQty
	}

	for _, v := range mStation {
		stations = append(stations, *v)
	}
	return

	// //method2
	// for _, stationName := range stationNames {
	// 	stationInfo := &StationInfo{
	// 		Name: stationName,
	// 	}
	// 	for _, wol := range wols {
	// 		s := wol.StationName
	// 		if s == stationInfo.Name {
	// 			stationInfo.CompleteQty = stationInfo.CompleteQty + wol.CompletedQty
	// 			stationInfo.NonGoodQty = stationInfo.NonGoodQty + wol.NonGoodQty
	// 		}
	// 	}
	// 	stationInfos = append(stationInfos, stationInfo)
	// }
	// return stationInfos
}

//----------------------->

//important~!

// myStations := MyStations{struct{ Station }{
// 	Station: Station{},
// }}

// "MyStations": [
//             {
//                 "Name": "A1",
//                 "GoodQty": 100,
//                 "NonGoodQty": 11,
//                 "CompletedQty": 111
//             }
//         ]

// func (ss MyStations) CalSample() {
// 	for _, s := range ss {

// 	}
// }

// func GetDistinctStationNames(wols []*WorkOrderList) (stationNames []string) {
// 	m := map[string]bool{}
// 	for _, wol := range wols {
// 		s := wol.StationName
// 		if !m[s] { //如果stationName不存在map裡面
// 			stationNames = append(stationNames, s)
// 			m[s] = true
// 		}
// 	}
// 	return
// }
