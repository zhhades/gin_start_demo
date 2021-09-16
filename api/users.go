package api

import (
	"gin_start_demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ListUsers(c *gin.Context) {
	time.Sleep(time.Second)
	c.JSON(http.StatusOK, gin.H{
		"data": []model.User{
			{
				Name: "zhhades",
				Age:  28,
				Sex:  "male",
			},
		},
	})
}
