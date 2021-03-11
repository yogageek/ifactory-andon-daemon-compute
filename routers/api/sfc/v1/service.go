package v1

import (
	"iii/ifactory/compute/db"
	model "iii/ifactory/compute/model/sfc"
	"iii/ifactory/compute/util"

	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
)

//#暫時解掉finishAt時間
//#M新增工單完成時間 FinishAt
//似乎也可以前端put做? (當需要 手動結掉 或是 數量滿足時要結掉)
func CallFinishAtService() {
	wos, err := FindWorkOrders("")
	if err != nil {
		glog.Info("CallFinishAtService err", err)
	}

	for _, wo := range wos {
		var woInfo model.WorkOrderInfo
		woInfo.NewWorkOrderInfo(wo)
		if woInfo.Status == "2" {

			// db.getCollection('iii.sfc.workorder').update(
			// 	{WorkOrderId:  "WO-2021",
			// 	 FinishAt: {$exists: true}},
			// 	{$set: {FinishAt: new Date()}},
			// )

			selector := model.WorkOrder{
				WorkOrderId: woInfo.WorkOrderId,
			}

			// now := util.GetNow()
			// selectorTest := WorkOrderInfo{
			// FinishAt: &now,
			// }

			value := bson.M{"FinishAt": util.GetNow()}
			db.UpdateIfOneFieldExistOrNot(model.C.Workorder, selector, value, "FinishAt", false)
		}
	}
}
