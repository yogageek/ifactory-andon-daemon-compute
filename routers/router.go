package routers

import (
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
		// apiv1.POST("/iot-2/evt/waconn/fmt/:sourceId", v1.PostOutbound_waconn)
		// apiv1.POST("/iot-2/evt/ifpcfg/fmt/:sourceId", v1.PostOutbound_ifpcfg)
		// apiv1.POST("/iot-2/evt/wacfg/fmt/:sourceId", v1.PostOutbound_wacfg)
	}

	return r
}
