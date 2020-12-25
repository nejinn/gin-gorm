package router

import (
	"github.com/gin-gonic/gin"
	v1 "note/api/v1"
)

func InitDashRouers(Router *gin.RouterGroup) {
	DashRouter := Router.Group("dashboard")
	{
		DashRouter.GET("", v1.InfoBox)
	}

}

