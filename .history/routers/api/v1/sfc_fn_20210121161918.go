package v1

import (
	"iii/ifactory/compute/db"
	model "iii/ifactory/compute/model/sfc"
	"math"

	"github.com/golang/glog"
)

func FindWorkOrders(workorderId *string, i *interface{}) {
	var wo []model.WorkOrder
	err := db.FindAll(model.C.Workorder, nil, nil, &wo)
	if err == nil {
		glog.Info(err)
	}
	return wo
}

func FindWorkOrdersById() {

}

func UpdateWorkOrder(workorderId string, wo *model.WorkOrder) error {

	//temprary wo (為了不誤修改其他值)
	func() {
		// wo = model.WorkOrder{
		// 	WorkOrderId: "",
		// }
		wo.WorkOrderID = ""
	}()

	selector := model.WorkOrder{
		WorkOrderID: workorderId,
	}

	//注意bson有無帶omitempty, 因為mgo會將interface轉bson傳入, 如果是日期格時會默認起始值
	err := db.UpdateOne(model.C.Workorder, selector, wo)
	if err != nil {
		return err
	}
	return nil
}

func FindWorkOrderInfo() []model.WorkOrderInfo {
	var wo []model.WorkOrder
	err := db.Lookup(model.C.Workorder, model.C.Workorder_list, "WorkOrderId", &wo)
	if err == nil {
		glog.Info(err)
	}

	// var rs []map[string]interface{}
	var wois []model.WorkOrderInfo
	for _, o := range wo {

		var sumGood float64
		var sumNonGood float64

		//sub collection
		for _, l := range o.WorkOrderList {
			sumGood = sumGood + (l.CompletedQty - l.NonGoodQty)
			sumNonGood = sumNonGood + l.NonGoodQty
		}

		// oInfo := map[string]interface{}{
		// 	"WorkOrderId": o.WorkOrderId,
		// 	"Good":        sumGood,
		// 	"NonGood":     sumNonGood,
		// }
		woi := model.WorkOrderInfo{
			WorkOrder:  o,
			SumGood:    sumGood,
			SumNonGood: sumNonGood,
			Sum:        sumGood + sumNonGood,
			GoodRate: func() float64 {
				if r := sumGood / (sumGood + sumNonGood); !math.IsNaN(r) {
					return r
				}
				return 0
			}(),
			Status: func() string {
				if o.Quantity <= sumGood {
					return "closed"
				}
				return "open"
			}(),
		}

		//如何把map string i 併入 struct 物件中, 扁平化
		// rs = append(rs, oInfo)
		wois = append(wois, woi)
	}
	return wois
}
