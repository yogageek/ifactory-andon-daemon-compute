package v1

import (
	"iii/ifactory/compute/db"
	model "iii/ifactory/compute/model/sfc"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

//Switching Panel-----------------------------------------------------------
func GetListsOfWorkOrderId(c *gin.Context) {
	//get all workordersInfo----------------------------
	wos, _ := FindWorkOrdersInfo()
	mm := []map[string]interface{}{}
	for _, wo := range wos {
		id := wo.WorkOrderId
		m := map[string]interface{}{
			"text":  id,
			"value": id,
		}
		mm = append(mm, m)
	}
	c.JSON(http.StatusOK, mm)
}
