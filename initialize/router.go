package initialize

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/global"
	"server/middleware"
	"server/router"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	// 设置gin模式
	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()
	// 使用日志记录中间件
	Router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	// 使用gin会话路由
	var store = cookie.NewStore([]byte(global.Config.System.SessionsSecret))

	Router.Use(sessions.Sessions("session", store))
	// 将指定目录下的文件提供给客户端
	// "uploads" 是URL路径前缀，http.Dir("uploads")是实际文件系统中存储文件的目录
	//Router.StaticFS(global.Config.Upload.Path, http.Dir(absPath))
	//需要RouterPrefix定义
	Router.StaticFS(global.Config.System.RouterPrefix+"/"+global.Config.Upload.Path, http.Dir("./"+global.Config.Upload.Path))
	// 初始化配置路由
	// 创建路由组
	routerGroup := router.RouterGroupApp
	//公共路由
	publicGroup := Router.Group(global.Config.System.RouterPrefix)
	//用户私人路由
	privateGroup := Router.Group(global.Config.System.RouterPrefix)
	privateGroup.Use(middleware.JWTAuth())
	//管理员路由
	adminGroup := Router.Group(global.Config.System.RouterPrefix)
	adminGroup.Use(middleware.JWTAuth()).Use(middleware.AdminAuth())

	{
		//初始化路由
		routerGroup.InitBaseRouter(publicGroup)
	}
	{
		//用户路由
		routerGroup.InitUserRouter(privateGroup, publicGroup, adminGroup)
		//文章路由
		routerGroup.InitArticleRouter(privateGroup, publicGroup, adminGroup)
		//反馈路由
		routerGroup.InitFeedbackRouter(privateGroup, publicGroup, adminGroup)
		//评论路由
		routerGroup.InitCommentRouter(privateGroup, publicGroup, adminGroup)
	}
	{
		//图片路由
		routerGroup.InitImageRouter(adminGroup)
		//友链路由
		routerGroup.InitFriendLinkRouter(publicGroup, adminGroup)
		//广告路由
		routerGroup.InitAdvertisementRouter(publicGroup, adminGroup)
		//网站路由
		routerGroup.InitWebsiteRouter(publicGroup, adminGroup)
		//配置路由
		routerGroup.InitConfigRouter(adminGroup)
	}
	return Router
}
