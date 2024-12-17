package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/victorlui/gopportunities/docs"
	"github.com/victorlui/gopportunities/handler/openings"
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", openings.ShowOpeningHandler)
		v1.POST("/opening", openings.CreateOpeningHandler)
		v1.DELETE("/opening", openings.DeleteOpeningHandler)
		v1.PUT("/opening", openings.UpdateOpeningHandler)
		v1.GET("/openings", openings.ListOpeningsHandler)
	}

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
