package model

import "fmt"

type CountInfos struct {
	Station struct {
		Count           float64
		TypeCount       float64
		EachTypeCount   map[string]float64
		EachStatusCount map[string]float64
	}
	WorkOrder struct {
		Count         float64
		TypeCount     float64
		EachTypeCount map[string]float64
	}
}

func (s *CountInfos) NewCountInfos(wois []WorkOrderInfo) {
	s.WorkOrder.EachTypeCount = map[string]float64{}
	s.Station.EachTypeCount = map[string]float64{}
	s.Station.EachStatusCount = map[string]float64{}

	//bug 為了解決沒有值 json就出不來的情況 所以要把每個選項都先放入
	s.Station.EachStatusCount["1"] = 0
	s.Station.EachStatusCount["2"] = 0
	s.Station.EachStatusCount["3"] = 0

	for _, woi := range wois {
		s.WorkOrder.Count++
		s.WorkOrder.EachTypeCount[fmt.Sprintf("%v", woi.Status)]++
		s.WorkOrder.TypeCount = float64(len(s.WorkOrder.EachTypeCount))

		for _, sta := range woi.StationDetails {
			s.Station.Count++
			s.Station.EachTypeCount[fmt.Sprintf("%v", sta.Name)]++
			s.Station.TypeCount = float64(len(s.Station.EachTypeCount))
			s.Station.EachStatusCount[fmt.Sprintf("%v", sta.StationCalInfo.Status)]++
		}
	}

	// check same in array
	// for _, a := range s {
	//     if a == e {
	//         return true
	//     }
	// }
	// return false

	// check same in map
	// stations := map[string]interface{}{}
	// for _, o := range ss {
	// 	if _, ok := stations[o.StationName]; !ok {
	// 		s.StationCount++
	// 	}
	// }
}
