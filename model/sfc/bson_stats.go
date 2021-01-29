package model

type StatsInfo struct {
	WorkOrderId string `json:"WorkOrderId,omitempty" bson:"WorkOrderId"`
	StationName string `json:"StationName,omitempty" bson:"StationName"`

	CompletedQty float64 `json:"CompletedQty" bson:"CompletedQty"`
	GoodQty      float64 `json:"GoodQty" bson:"GoodQty"`
	NonGoodQty   float64 `json:"NonGoodQty" bson:"NonGoodQty"`
	Quantity     float64 `json:"Quantity" bson:"Quantity,omitempty"` //預計生產數量

	Calculator
	// ToBeCompletedQty  float64 `json:"ToBeCompletedQty" bson:"ToBeCompletedQty"`
	// RealCompletedRate float64 `json:"RealCompletedRate,omitempty" bson:"RealCompletedRate"`
	// EstiCompletedRate float64 `json:"EstiCompletedRate,omitempty" bson:"EstiCompletedRate"`
	// GoodQtyRate       float64 `json:"GoodQtyRate,omitempty" bson:"GoodQtyRate"`
	// Status            float64 `json:"Status,omitempty" bson:"Status"`

	// orders list count of the station (moved to counts)

}

func (s *StatsInfo) CalStats() {

	s.GoodQty = s.CompletedQty - s.NonGoodQty //待檢查是否自動有

	var c Calculator
	c.NewCalculator(s.GoodQty, s.CompletedQty, s.Quantity)
	s.Calculator = c

	// old style
	// s.ToBeCompletedQty = c.calToBeCompletedQty(s.Quantity, s.GoodQty)
	// s.RealCompletedRate = c.calRealCompletedRate(s.CompletedQty, s.Quantity)
	// s.GoodQtyRate = c.calGoodQtyRate(s.GoodQty, s.CompletedQty)
	// s.Status = c.calStatus(s.CompletedQty, s.Quantity)
}
