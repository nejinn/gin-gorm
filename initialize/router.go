package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"
	"net/http"
	_ "note/docs"
	"note/global"
	"note/middleware"
	"note/router"
)

// 初始化总路由

func Routers() *gin.Engine {
	global.NLY_LOG.Info("开始初始化路由设置")
	var Router = gin.New()
	Router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	global.NLY_LOG.Info("使用 Zap 自定义 Logger 和 Recovery ")
	Router.StaticFS(global.NLY_CONFIG.Static.Path, http.Dir(global.NLY_CONFIG.Static.Path))
	global.NLY_LOG.Info("注册静态文件目录:", zap.String("dir", global.NLY_CONFIG.Static.Path))
	// 跨域
	Router.Use(middleware.Cors())
	global.NLY_LOG.Info("注册跨域中间件")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.NLY_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	global.NLY_LOG.Info("开始注册路由")
	PublicApiGroup := Router.Group("api/v1/")
	{
		router.InitUserRouers(PublicApiGroup) // 注册用户路由
	}
	PrivateApiGroup := Router.Group("api/v1/")
	PrivateApiGroup.Use(middleware.JWTAuthMiddleware())
	{
		router.InitNavRouers(PrivateApiGroup)
	}
	global.NLY_LOG.Info("全部路由注册成功")
	return Router
}
