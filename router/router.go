package router

import (
	"gin_start_demo/api"
	"github.com/gin-gonic/gin"
)

func InitUsersApiRouter(Router *gin.RouterGroup) {
	userApiRouter := Router.Group("/user")
	userApiRouter.GET("/list", api.ListUsers)
}
