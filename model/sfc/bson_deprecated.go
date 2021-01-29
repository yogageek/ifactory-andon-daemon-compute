package model

//----------->

// //deprecated
// func (o *WorkOrder) calStationInfo() (stations []Station) {
// 	wols := o.WorkOrderList

// 	//method1
// 	mStation := map[string]*Station{}
// 	for _, wol := range wols {
// 		s := wol.StationName

// 		if mStation[s] == nil {
// 			mStation[s] = &Station{
// 				Name: s,
// 			}
// 		}
// 		mStation[s].CompletedQty = mStation[s].CompletedQty + wol.CompletedQty
// 		mStation[s].NonGoodQty = mStation[s].NonGoodQty + wol.NonGoodQty
// 		mStation[s].GoodQty = mStation[s].CompletedQty - mStation[s].NonGoodQty

// 		// mStation[s].StationCalInfo.ToBeCompletedQty = o.Quantity - mStation[s].GoodQty
// 		// mStation[s].StationCalInfo.Status = calStatus(mStation[s].CompletedQty, o.Quantity)
// 	}

// 	for _, v := range mStation {
// 		stations = append(stations, *v)
// 	}
// 	return

// 	// //method2
// 	// for _, stationName := range stationNames {
// 	// 	stationInfo := &StationInfo{
// 	// 		Name: stationName,
// 	// 	}
// 	// 	for _, wol := range wols {
// 	// 		s := wol.StationName
// 	// 		if s == stationInfo.Name {
// 	// 			stationInfo.CompleteQty = stationInfo.CompleteQty + wol.CompletedQty
// 	// 			stationInfo.NonGoodQty = stationInfo.NonGoodQty + wol.NonGoodQty
// 	// 		}
// 	// 	}
// 	// 	stationInfos = append(stationInfos, stationInfo)
// 	// }
// 	// return stationInfos
// }

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
