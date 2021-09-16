package initialize

import (
	"gin_start_demo/logger"
	"github.com/gin-gonic/gin"

	"gin_start_demo/router"
)

func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	Router := gin.New()
	Router.Use(logger.GinLogger(), logger.GinRecovery(true))
	ApiGroup := Router.Group("/v1")
	router.InitUsersApiRouter(ApiGroup)
	return Router
}
