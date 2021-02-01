package v1

import (
	"fmt"
	model "iii/ifactory/compute/model/sfc"
	"io/ioutil"
	"net/http"

	. "iii/ifactory/compute/util/cch/json"

	"github.com/gin-gonic/gin"
	. "github.com/logrusorgru/aurora"
)

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
