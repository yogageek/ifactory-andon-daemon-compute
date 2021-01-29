package model

import (
	"math"
)

type StationInfo struct {
	// stations []Station
	Station
	Calculator
}
type Station struct {
	Name         string
	CompletedQty float64
	GoodQty      float64
	NonGoodQty   float64
}
type Calculator struct {
	ToBeCompletedQty  float64 `json:"ToBeCompletedQty" bson:"ToBeCompletedQty"`
	GoodQtyRate       float64 `json:"GoodQtyRate" bson:"GoodQtyRate"`
	RealCompletedRate float64 `json:"RealCompletedRate" bson:"RealCompletedRate"`
	EstiCompletedRate float64 `json:"EstiCompletedRate" bson:"EstiCompletedRate"`
	Status            float64 `json:"Status" bson:"Status"`
}

func (c *Calculator) NewCalculator(goodQty, completedQty, quantity float64) {
	c.ToBeCompletedQty = c.calToBeCompletedQty(quantity, goodQty)
	c.RealCompletedRate = c.calRealCompletedRate(completedQty, quantity)
	c.GoodQtyRate = c.calGoodQtyRate(goodQty, completedQty)
	c.Status = c.calStatus(completedQty, quantity)
}

func groupStationsByName(stations []Station) (m map[string][]Station) {
	m = make(map[string][]Station)
	for _, s := range stations {
		m[s.Name] = append(m[s.Name], s)
	}
	// util.PrintJson(m)
	return
}

func (s *StationInfo) NewStationInfo(name string, stations []Station, quantity float64) {
	s.Station.Name = name

	for _, station := range stations {
		s.CompletedQty += station.CompletedQty
		s.GoodQty += station.GoodQty
		s.NonGoodQty += station.NonGoodQty
	}

	var c Calculator
	// old style
	// s.ToBeCompletedQty = c.calToBeCompletedQty(quantity, s.GoodQty)
	// s.RealCompletedRate = c.calRealCompletedRate(s.CompletedQty, quantity)
	// s.GoodQtyRate = c.calGoodQtyRate(s.GoodQty, s.CompletedQty)
	// s.Status = s.calStatus(s.CompletedQty, quantity)
	c.NewCalculator(s.GoodQty, s.CompletedQty, quantity)
	s.Calculator = c

}

//移到new struct pointer func名為計算模組
func (c Calculator) calRealCompletedRate(completedQty, quantity float64) float64 {
	if r := (completedQty / quantity) * 100; !math.IsNaN(r) {
		return r
	}
	return 0
}

func (c Calculator) calToBeCompletedQty(quantity, goodQty float64) float64 {
	return quantity - goodQty
}

func (c Calculator) calGoodQtyRate(goodQty, completedQty float64) float64 {
	if r := goodQty / completedQty; !math.IsNaN(r) {
		return r
	}
	return 0
}

func (c Calculator) calStatus(completedQty, quantity float64) float64 {
	switch {
	case completedQty < quantity:
		return 1 //"低於標準"
	case completedQty > quantity:
		return 3 //"高於標準"
	default:
		return 2 //"等於標準"
	}
}

// 脫褲子放屁
// type Stations []struct {
// 	Station
// }

// func (stations Stations) groupStationsByName2() (m map[string][]Station) {
// 	for _, s := range stations {
// 		m[s.Name] = append(m[s.Name], s.Station)
// 	}
// 	return
// }