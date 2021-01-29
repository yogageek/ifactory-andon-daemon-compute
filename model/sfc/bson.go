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

func calRealCompletedRate(completedQty, quantity float64) float64 {
	if r := (completedQty / quantity) * 100; !math.IsNaN(r) {
		return r
	}
	return 0
}

func calStatus(completedQty, quantity float64) float64 {
	switch {
	case completedQty < quantity:
		return 1 //"低於標準"
	case completedQty > quantity:
		return 3 //"高於標準"
	default:
		return 2 //"等於標準"
	}
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

// 工單資訊
type WorkOrderInfo struct {
	WorkOrder
	Stations []Station `json:"Stations,omitempty" bson:"Stations"`

	MixedCompletedQty float64 `json:"MixedCompletedQty,omitempty" bson:"MixedCompletedQty"`
	MixedGoodQty      float64 `json:"MixedGoodQty,omitempty" bson:"MixedGoodQty"`
	MixedNonGoodQty   float64 `json:"MixedNonGoodQty,omitempty" bson:"MixedNonGoodQty"`

	GoodProductQty float64 `json:"GoodProductQty" bson:"GoodProductQty"`

	MixedGoodQtyRate   float64 `json:"MixedGoodQtyRate,omitempty" bson:"MixedGoodQtyRate"`
	GoodProductQtyRate float64 `json:"GoodProductQtyRate,omitempty" bson:"GoodProductQtyRate"`

	Status string `json:"Status,omitempty" bson:"Status"`
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

type MyStations []struct {
	Station
}

// // 改成
// type StationDetail struct {
// 	Station
// 	CalInfo
// }
type Station struct {
	Name         string
	CompletedQty float64
	GoodQty      float64
	NonGoodQty   float64

	CalInfo
}

type CalInfo struct {
	ToBeCompletedQty  float64 `json:"-"`
	RealCompletedRate float64 `json:"-"`
	EstiCompletedRate float64 `json:"-"`
	Status            float64 `json:"Status"`
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

func (woi *WorkOrderInfo) NewWorkOrderInfo(wo WorkOrder) {

	for _, wol := range wo.WorkOrderList {
		woi.MixedCompletedQty = woi.MixedCompletedQty + wol.CompletedQty
		woi.MixedNonGoodQty = woi.MixedNonGoodQty + wol.NonGoodQty
		woi.MixedGoodQty = woi.MixedCompletedQty - woi.MixedNonGoodQty
	}
	woi.Stations = wo.calStationInfo()

	woi.MixedGoodQtyRate = func() float64 {
		if r := woi.MixedGoodQty / woi.MixedCompletedQty; !math.IsNaN(r) {
			return r
		}
		return 0
	}()

	woi.GoodProductQty = func() (minGood float64) {
		//當有工單沒報工會error
		// minGood = woi.Stations[0].GoodQty
		// for _, s := range woi.Stations {
		// 	if minGood > s.GoodQty {
		// 		minGood = s.GoodQty
		// 	}
		// }

		//1/29 fix
		for i := 0; i < len(woi.Stations); i++ {
			minGood = woi.Stations[0].GoodQty
			if minGood > woi.Stations[i].GoodQty {
				minGood = woi.Stations[i].GoodQty
			}
		}

		return
	}()

	woi.Status = func() string {
		if wo.Quantity <= woi.GoodProductQty {
			return "2"
		}
		return "1"
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

		mStation[s].CalInfo.ToBeCompletedQty = o.Quantity - mStation[s].GoodQty
		mStation[s].CalInfo.Status = calStatus(mStation[s].CompletedQty, o.Quantity)
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
