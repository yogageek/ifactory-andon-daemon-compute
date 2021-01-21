package routers

import (
	v1 "iii/ifactory/compute/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	// gin.SetMode(setting.RunMode)

	//----------------->
	apiv1 := r.Group("/")
	{
		apiv1.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "OK",
			})
		})

		// apiv1.GET("/workorders/info", v1.GetWorkOrdersInfo) //列出all wo(包含底下list,product等資訊)

		//get all workorders(include list)
		apiv1.GET("/workorders", v1.GetWorkOrders)
		//get one workorders(include list)
		apiv1.GET("/workorders/:workorderId", v1.GetWorkOrderLists)
		//get all list of one workorders
		apiv1.GET("/workorders/:workorderId/workorder_list")

		//新增wo (如果有帶list則同時新增list)
		apiv1.POST("/workorders", v1.PostWorkOrders)
		//新增list到指定wo底下
		apiv1.POST("/workorders/:workorderId/workorderLists", v1.PostWorkOrderLists)

		// apiv1.PUT("/workorders/:workorderId", v1.PutWorkOrder)

	}

	return r
}
