package router

import (
	"github.com/gin-gonic/gin"
	"note/api/v1"
)

func InitUserRouers(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		//UserRouter.GET("user_list", v1.GetUserList)
		UserRouter.POST("register", v1.Register)
		//UserRouter.POST("uploadImg", v1.UploadImg)
		UserRouter.POST("login", v1.LoginAdmin)
	}

}

func InitProviderUserRouers(Router *gin.RouterGroup) {
	ProviderUserRouter := Router.Group("user")
	{
		ProviderUserRouter.GET("", v1.GetUserInfo)
		ProviderUserRouter.GET("list",v1.GetUserList)
	}

}
