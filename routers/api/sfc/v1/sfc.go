package v1

import (
	"fmt"
	"iii/ifactory/compute/db"
	model "iii/ifactory/compute/model/sfc"
	"net/http"

	"iii/ifactory/compute/util"
	. "iii/ifactory/compute/util/cch/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	. "github.com/logrusorgru/aurora"
	"gopkg.in/mgo.v2/bson"
)

/*
使用 Query 與 DefaultQuery 來取得 request 參數
DefaultQuery 的話如果沒有 firstname 這參數，就會給預設值第二個參數(None)
firstname := c.DefaultQuery("firstname", "None")
lastname := c.Query("lastname")
*/

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

	//delete到底要用哪種id
	//如果都用_id,gin naming會重複 但restful設計沒問題
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

//GET Stats-----------------------------------------------------------
func GetCounts(c *gin.Context) {
	var ci model.CountInfos

	//get all workordersInfo----------------------------
	wos, _ := FindWorkOrdersInfo()
	ci.NewCountInfos(wos)
	r := model.GenGrafanaResponse(ci)
	c.JSON(http.StatusOK, r)
}

func GetTables(c *gin.Context) {
	groupBy := c.Query("groupBy")

	var stsInfo []*model.StatsInfo
	if groupBy == "station" {
		agg := db.Agg{}
		agg.GenUnwind("WorkOrderList")
		agg.GenGroup(
			"WorkOrderList",
			[]string{"WorkOrderId", "StationName"},
			[]string{"WorkOrderId", "StationName", "$Quantity"},
			[]string{"CompletedQty", "NonGoodQty"},
		)
		// util.PrintJson(group)
		err := agg.Aggre(model.C.Workorder, &stsInfo)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		//stats calculate logic
		for _, s := range stsInfo {
			s.CalStats()
		}

		// tutorial
		// var ms []*model.StatsInfo
		// if err := bson.Unmarshal(bsonBytes, &ms); err != nil {
		// 	panic(err)
		// }

		c.JSON(http.StatusOK, stsInfo)
	}
	// util.PrintJson(r)
}

//可一起找出wo, wolist, product, station...
func GetWorkOrders(c *gin.Context) {

	detail := c.Query("detail")

	if detail == "true" {
		//get all workordersInfo----------------------------
		wos, err := FindWorkOrdersInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, wos) //之後須加上by id功能
		return
	} else {
		//get all workorders----------------------------
		wos, err := FindWorkOrders("")
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, wos)
		return
	}
}

func GetWorkOrder(c *gin.Context) {
	workorderId := c.Param("workorderId") //取得URL中参数
	wos, err := FindWorkOrders(workorderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, wos)
}

//POST-------------------

//建立工單 done
func PostWorkOrders(c *gin.Context) {

	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(BrightBlue(string(body)))

	var wo model.WorkOrder
	err := FromJson(string(body), &wo)

	wo.NewWorkOrder()

	if err != nil {
		// c.String(http.StatusOK, `error~~~~`)

		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"code": 1000,
		// 	"msg":  "xxx",
		// 	"data": make(map[string]string),
		// })

		c.JSON(http.StatusBadRequest, err)
		return
	}

	db.Insert(model.C.Workorder, wo)

	// method1: use gjson to get the field you want
	// gjson.GetBytes(req, "factoryId")

	// method2: convert json to struct, 但目前不確定完整格式可能會產生 Bug EOF
	// var ifpcfg model.Ifpcfg
	// if err := c.ShouldBind(&ifpcfg); err != nil {
	// 	fmt.Println("PostOutbound_ifpcfg err:", err)
	// 	// c.String(http.StatusOK, `the body should be formA`)
	// }
	// fmt.Println(BrightCyan(ifpcfg))
}

//建立報工單 done
func PostWorkOrderLists(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(BrightBlue(string(body)))

	workorderId := c.Param("workorderId") //取得URL中参数

	// var v interface{}
	// if err := json.Unmarshal(body, &v); err != nil {
	// 	glog.Error(err)
	// }

	var wols []*model.WorkOrderList

	err := FromJsonNoV(string(body), &wols)
	if err != nil {
		util.Cerr(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	//set initial value
	for _, wol := range wols {
		wol.NewWorkOrderList(workorderId)
	}

	//append wolist到wo底下
	selector := model.WorkOrderList{WorkOrderId: workorderId}

	// debug
	// selector := map[string]interface{}{"WorkOrderID": workorderId}
	// a := reflect.ValueOf(wols).Interface().(interface{})
	// fmt.Println(a)

	// method 1
	// err = db.Pushs(model.C.Workorder, selector, "WorkOrderList", wo.WorkOrderList)

	// method 2
	for _, wol := range wols {
		var jb model.JB
		jb.WorkOrderList = wol
		err := db.Push(model.C.Workorder, selector, jb)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}
}

//PUT-------------------

//更新工單
func PutWorkOrder(c *gin.Context) {
	workorderId := c.Param("workorderId") //取得URL中参数

	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(BrightBlue(string(body)))

	var wo model.WorkOrder
	err := FromJson(string(body), &wo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = UpdateWorkOrder(workorderId, &wo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 4000,
			"msg":  err.Error(),
			"data": make(map[string]string),
		})
		// c.JSON(http.StatusBadRequest, err.Error())
		return
	}
}

// func JustForTest() {
// 	b := GetOutboundSample()
// 	var ifpcfg model.Ifpcfg
// 	err := json.Unmarshal(b, &ifpcfg)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	// fmt.Printf("%+v", ifpcfg)

// 	groupItems := ifpcfg.Group.Items
// 	// machineItems := ifpcfg.Machine
// 	// parameterItems := ifpcfg.Parameter

// 	//method1
// 	//unmarshal裡面寫匿名struct, 在把mapp到struct的值給去 &point
// 	//method2
// 	//tag加上omitempty

// 	for _, i := range groupItems {
// 		fmt.Printf("%+v", i)
// 	}
// }
