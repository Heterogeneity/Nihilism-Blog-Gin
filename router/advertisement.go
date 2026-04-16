package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type AdvertisementRouter struct {
}

func (a *AdvertisementRouter) InitAdvertisementRouter(Router *gin.RouterGroup, AdminRouter *gin.RouterGroup) {
	publicRouter := Router.Group("advertisement")
	adminRouter := AdminRouter.Group("advertisement")

	advertisementApi := api.ApiGroupApp.AdvertisementApi
	{
		adminRouter.GET("list", advertisementApi.AdvertisementList)
		adminRouter.POST("create", advertisementApi.AdvertisementCreate)
		adminRouter.DELETE("delete", advertisementApi.AdvertisementDelete)
		adminRouter.PUT("update", advertisementApi.AdvertisementUpdate)
	}
	{
		publicRouter.GET("info", advertisementApi.AdvertisementInfo)
	}
}
