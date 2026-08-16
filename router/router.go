package router

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	r := gin.Default()

	// API v1 路由组
	api := r.Group("/api/v1")
	{
		EmailRoutes(api)
	}

	return r
}
