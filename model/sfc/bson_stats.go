package model

import "math"

type StatsInfo struct {
	WorkOrderId string `json:"WorkOrderId,omitempty" bson:"WorkOrderId"`
	StationName string `json:"StationName,omitempty" bson:"StationName"`

	GoodQty          float64 `json:"GoodQty" bson:"GoodQty"`
	NonGoodQty       float64 `json:"NonGoodQty" bson:"NonGoodQty"`
	CompletedQty     float64 `json:"CompletedQty" bson:"CompletedQty"`
	ToBeCompletedQty float64 `json:"ToBeCompletedQty" bson:"ToBeCompletedQty"`
	Quantity         float64 `json:"Quantity" bson:"Quantity,omitempty"` //預計生產數量

	RealCompletedRate float64 `json:"RealCompletedRate,omitempty" bson:"RealCompletedRate"`
	EstiCompletedRate float64 `json:"EstiCompletedRate,omitempty" bson:"EstiCompletedRate"`

	//1/28 add
	GoodQtyRate float64 `json:"GoodQtyRate,omitempty" bson:"GoodQtyRate"`

	// orders list count of the station (moved to counts)

	Status float64 `json:"Status,omitempty" bson:"Status"`
}

func (s *StatsInfo) CalStats() {
	var c Cal

	s.GoodQty = s.CompletedQty - s.NonGoodQty
	s.ToBeCompletedQty = c.calToBeCompletedQty(s.Quantity, s.GoodQty)
	s.Status = c.calStatus(s.CompletedQty, s.Quantity)
	s.RealCompletedRate = c.calRealCompletedRate(s.CompletedQty, s.Quantity)
	s.GoodQtyRate = func() float64 {
		if r := s.GoodQty / s.Quantity; !math.IsNaN(r) {
			return r
		}
		return 0
	}()

}
