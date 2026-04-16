package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type FriendLinkRouter struct {
}

func (f *FriendLinkRouter) InitFriendLinkRouter(Router *gin.RouterGroup, AdminRouter *gin.RouterGroup) {
	friendLinkPublicRouter := Router.Group("friendLink")
	friendLinkRouter := AdminRouter.Group("friendLink")
	friendLinkApi := api.ApiGroupApp.FriendLinkApi
	{
		friendLinkRouter.POST("create", friendLinkApi.FriendLinkCreate)
		friendLinkRouter.DELETE("delete", friendLinkApi.FriendLinkDelete)
		friendLinkRouter.PUT("update", friendLinkApi.FriendLinkUpdate)
		friendLinkRouter.GET("list", friendLinkApi.FriendLinkList)
	}
	{
		friendLinkPublicRouter.GET("info", friendLinkApi.FriendLinkInfo)
	}
}
