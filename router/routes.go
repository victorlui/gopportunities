package router

import (
	"github.com/gin-gonic/gin"
	"github.com/victorlui/gopportunities/handler/openings"
)

func initializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", openings.ShowOpeningHandler)
		v1.POST("/opening", openings.CreateOpeningHandler)
		v1.DELETE("/opening", openings.DeleteOpeningHandler)
		v1.PUT("/opening", openings.UpdateOpeningHandler)
		v1.GET("/openings", openings.ListOpeningsHandler)
	}
}
