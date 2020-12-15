package router

import (
	"github.com/gin-gonic/gin"
	v1 "note/api/v1"
)

func InitNavRouers(Router *gin.RouterGroup) {
	UserRouter := Router.Group("nav")
	{
		UserRouter.POST("create", v1.CreateUserNav)

	}

}
