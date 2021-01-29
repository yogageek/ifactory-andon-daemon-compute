package model

type StatsInfo struct {
	WorkOrderId string `json:"WorkOrderId,omitempty" bson:"WorkOrderId"`
	StationName string `json:"StationName,omitempty" bson:"StationName"`

	CompletedQty float64 `json:"CompletedQty" bson:"CompletedQty"`
	GoodQty      float64 `json:"GoodQty" bson:"GoodQty"`
	NonGoodQty   float64 `json:"NonGoodQty" bson:"NonGoodQty"`
	Quantity     float64 `json:"Quantity" bson:"Quantity,omitempty"` //預計生產數量

	Calculator

	// orders list count of the station (moved to counts)
}

func (s *StatsInfo) CalStats() {

	s.GoodQty = s.CompletedQty - s.NonGoodQty //待檢查是否自動有

	var c Calculator
	c.NewCalculator(s.GoodQty, s.CompletedQty, s.Quantity)
	s.Calculator = c
}
