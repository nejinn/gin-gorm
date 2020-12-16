package router

import (
	"github.com/gin-gonic/gin"
	v1 "note/api/v1"
)

func InitNavRouers(Router *gin.RouterGroup) {
	NavRouter := Router.Group("nav")
	{
		NavRouter.POST("create", v1.CreateNav)
		NavRouter.GET("list", v1.GetNavList)
	}

}
