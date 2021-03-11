package routers

import (
	v1 "iii/ifactory/compute/routers/api/sfc/v1"

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

		//get all workorders(include list)
		apiv1.GET("/workorders", v1.GetWorkOrders)
		//get one workorders(include list)
		apiv1.GET("/workorders/:workorderId", v1.GetWorkOrder)
		//get all list of one workorders(上面那隻可以達到)
		// apiv1.GET("/workorders/:workorderId/workorder_list")

		//add wo/list (如果有帶list則同時新增list)
		apiv1.POST("/workorders", v1.PostWorkOrders)
		//add list (到指定wo底下)
		apiv1.POST("/workorders/:workorderId/workorderLists", v1.PostWorkOrderLists)

		apiv1.DELETE("/workorders/:id", v1.DeleteWorkOrders)
		apiv1.DELETE("/workorders/:id/workorderLists/:id2", v1.DeleteWorkOrderLists)
		apiv1.PUT("/workorders/:workorderId", v1.PutWorkOrder)

		apiv1.GET("/grafana/tables", v1.GetTables) //列出all wo(包含底下list,product等資訊)
		apiv1.GET("/grafana/counts", v1.GetCounts) //列出all wo(包含底下list,product等資訊)

		//命名規則不太確定 有點不正確
		apiv1.GET("/grafana/switchingPanel/workorders/id", v1.GetListsOfWorkOrderId)
	}

	return r
}
