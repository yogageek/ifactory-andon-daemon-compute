package model

import (
	"fmt"
	"iii/ifactory/compute/util"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//-------------
type JB struct {
	WorkOrder     *WorkOrder     `json:"WorkOrder,omitempty" bson:"WorkOrder,omitempty"`
	WorkOrderList *WorkOrderList `json:"WorkOrderList,omitempty" bson:"WorkOrderList,omitempty"`
}

// 工單資訊
type WorkOrderInfo struct {
	WorkOrder
	StationInfos []StationInfo `json:"StationInfos,omitempty" bson:"StationInfos"`

	GoodProductQty float64 `json:"GoodProductQty" bson:"GoodProductQty"`

	MixedCompletedQty float64 `json:"MixedCompletedQty" bson:"MixedCompletedQty"`
	MixedGoodQty      float64 `json:"MixedGoodQty" bson:"MixedGoodQty"`
	MixedNonGoodQty   float64 `json:"MixedNonGoodQty" bson:"MixedNonGoodQty"`

	MixedGoodQtyRate float64 `json:"MixedGoodQtyRate" bson:"MixedGoodQtyRate"`

	// GoodProductQtyRate float64 `json:"GoodProductQtyRate,omitempty" bson:"GoodProductQtyRate"`

	Status string `json:"Status" bson:"Status"`
}

//工單
type WorkOrder struct {
	Id          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	WorkOrderId string        `json:"WorkOrderId,omitempty" bson:"WorkOrderId,omitempty"`

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

	//#是否把station模組放到這取代以下四欄?(但是會動到很多地方)
	StationName  string  `json:"StationName,omitempty" bson:"StationName,omitempty" validate:"required"` //注意!!!不加empty會嚴重影響到push功能而且查不出來!
	CompletedQty float64 `json:"CompletedQty,omitempty" bson:"CompletedQty,omitempty"`                   //注意!!!不加empty會嚴重影響到push功能而且查不出來!因為c.Update(selector, updated) 系統會自動把selector=interface{}轉成bson,沒帶omitempty會parse成"""
	GoodQty      float64 `json:"GoodQty,omitempty" bson:"GoodQty,omitempty"`
	NonGoodQty   float64 `json:"NonGoodQty,omitempty" bson:"NonGoodQty,omitempty"`

	Reporter string     `json:"Reporter,omitempty" bson:"Reporter,omitempty" validate:"required"`
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

func (o *WorkOrder) NewWorkOrder() {

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
		oo.NewWorkOrderList(o.WorkOrderId)
	}
}

func (o *WorkOrderList) NewWorkOrderList(workorderId string) {
	o.Id = bson.NewObjectId()
	o.WorkOrderId = workorderId
	//1/29 add
	o.GoodQty = o.CompletedQty - o.NonGoodQty
	o.CreateAt = func() *time.Time {
		now := util.GetNow()
		return &now
	}()
}

func (o *WorkOrderList) GenStation() (s Station) {
	s.CompletedQty = o.CompletedQty
	s.GoodQty = o.GoodQty
	s.NonGoodQty = o.NonGoodQty
	s.Name = o.StationName
	return
}

func (woInfo *WorkOrderInfo) NewWorkOrderInfo(wo WorkOrder) {

	// try
	// var Stations Stations
	var stations []Station

	for _, wolist := range wo.WorkOrderList {
		// old style
		// w.MixedCompletedQty = w.MixedCompletedQty + l.CompletedQty
		// w.MixedNonGoodQty = w.MixedNonGoodQty + l.NonGoodQty
		// w.MixedGoodQty = w.MixedCompletedQty - w.MixedNonGoodQty
		// new style
		woInfo.MixedCompletedQty += wolist.CompletedQty
		woInfo.MixedNonGoodQty += wolist.NonGoodQty
		woInfo.MixedGoodQty += wolist.GoodQty

		// try
		// for _, v := range Stations {
		// 	v.Station = wolist.GenStation()
		// }

		//gen stations by wolists
		stations = append(stations, wolist.GenStation())
	}

	//stations logic
	mStations := groupStationsByName(stations)
	for name, stations := range mStations {
		var si StationInfo
		si.NewStationInfo(name, stations, wo.Quantity)
		woInfo.StationInfos = append(woInfo.StationInfos, si)
	}

	//暫用
	var c Calculator
	woInfo.MixedGoodQtyRate = c.calGoodQtyRate(woInfo.MixedGoodQty, woInfo.MixedCompletedQty)

	woInfo.GoodProductQty = func() (minGood float64) {
		//1/29 fix
		for i := 0; i < len(woInfo.StationInfos); i++ {
			minGood = woInfo.StationInfos[0].GoodQty
			if minGood > woInfo.StationInfos[i].GoodQty {
				minGood = woInfo.StationInfos[i].GoodQty
			}
		}
		return
	}()

	woInfo.Status = func() string {
		if wo.Quantity <= woInfo.GoodProductQty {
			return "2"
		}
		return "1"
	}()

	woInfo.WorkOrder = wo
}

// // SetBSON implements bson.Setter.
// func (c *StatsInfo) SetBSON(raw bson.Raw) error {

// 	// decoded := new(struct { })
// 	type newStatsInfo StatsInfo
// 	s := new(newStatsInfo)

// 	bsonErr := raw.Unmarshal(s)
// 	if bsonErr != nil {
// 		return bsonErr
// 	}
// 	// util.PrintJson(s)

// 	s.RealCompletedRate = func() float64 {
// 		if r := (s.CompletedQty / s.Quantity) * 100; !math.IsNaN(r) {
// 			return r
// 		}
// 		return 0
// 	}()
// 	s.Status = func() float64 {
// 		switch {
// 		case s.CompletedQty < s.Quantity:
// 			return -1 //"低於標準"
// 		case s.CompletedQty > s.Quantity:
// 			return 1 //"高於標準"
// 		default:
// 			return 0 //"等於標準"
// 		}
// 	}()
// 	s.NonGoodQty = s.CompletedQty - s.NonGoodQty

// 	c.WorkOrderId = s.WorkOrderId
// 	c.CompletedQty = s.CompletedQty
// 	c.GoodQty = s.GoodQty
// 	c.Quantity = s.Quantity
// 	c.StationName = s.StationName
// 	c.Status = s.Status
// 	c.RealCompletedRate = s.RealCompletedRate
// 	c.NonGoodQty = s.NonGoodQty

// 	// util.PrintJson(c)

// 	return nil
// }

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

func (p *StatsInfo) UnmarshalJSON(data []byte) (err error) {

	// var res ProfileJSON
	fmt.Println("x")
	// if err := json.Unmarshal(data, &res); err != nil {
	// 	return err
	// }
	// return nil
	return nil
}
