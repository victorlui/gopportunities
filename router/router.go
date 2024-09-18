package router

import "github.com/gin-gonic/gin"

func Initialize() {
	g := gin.Default()
	initializeRoutes(g)
	g.Run(":8080")
}
