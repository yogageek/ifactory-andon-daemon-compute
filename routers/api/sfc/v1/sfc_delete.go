package v1

import (
	"fmt"
	"iii/ifactory/compute/db"
	model "iii/ifactory/compute/model/sfc"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"gopkg.in/mgo.v2/bson"
)

//DELETE
func DeleteWorkOrders(c *gin.Context) {

	id := c.Param("id")

	//use Remove
	// selector := model.WorkOrder{Id: bson.ObjectIdHex(id)} //ok
	// selector := bson.M{"WorkOrderId": "20210126-60102a40959582230048b355"} //ok
	// selector := bson.M{"_id": bson.ObjectIdHex(id)} //ok

	//use RemoveId
	selector := bson.ObjectIdHex(id) //ok

	err := db.MongoDB.UseC(model.C.Workorder).RemoveId(selector)
	if err != nil {
		glog.Error(err)
	}

	/*
		//modify sepc, use workorderId to delete
		workorderId := c.Param("workorderId")
		selector := model.WorkOrder{WorkOrderId: workorderId}
		err := db.MongoDB.UseC(model.C.Workorder).Remove(selector)
		if err != nil {
			glog.Error(err)
		}
	*/

	//delete要用_id還是workorderId--->_id
	//否則workorderId & workorderListId 不同步會搞混
}

//doing here
func DeleteWorkOrderLists(c *gin.Context) {
	// workorderId := c.Param("workorderId")
	id := c.Param("id")
	id2 := c.Param("id2")
	fmt.Println(id, ",", id2)

	//new style
	var updater db.Updater
	selector := updater.GenId(id)
	update := updater.GenPull("WorkOrderList", id2)
	db.Update(model.C.Workorder, selector, update)

	/*
		//old style
		selector := map[string]interface{}{
			"_id": bson.ObjectIdHex(id),
		}
		update := map[string]interface{}{
			"$pull": bson.M{"WorkOrderList": bson.M{"_id": bson.ObjectIdHex(id2)}},
		}
		dbc := db.MongoDB.UseC(model.C.Workorder)
		err := dbc.Update(selector, update)
		if err != nil {
			glog.Error(util.Cerr(err))
		}
	*/

	// unset用法
	// db.Update({"name":"zhang"},{$unset:{"age":1}})
}
