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
	. "github.com/logrusorgru/aurora"
)

/*
使用 Query 與 DefaultQuery 來取得 request 參數
DefaultQuery 的話如果沒有 firstname 這參數，就會給預設值第二個參數(None)
firstname := c.DefaultQuery("firstname", "None")
lastname := c.Query("lastname")
*/

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
