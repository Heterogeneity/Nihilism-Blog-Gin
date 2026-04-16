package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type ConfigRouter struct {
}

func (c *ConfigRouter) InitConfigRouter(AdminRouter *gin.RouterGroup) {

	configAdminRouter := AdminRouter.Group("config")
	configApi := api.ApiGroupApp.ConfigApi
	{
		configAdminRouter.GET("website", configApi.GetWebsite)
		configAdminRouter.PUT("website", configApi.UpdateWebsite)

		configAdminRouter.GET("system", configApi.GetSystem)
		configAdminRouter.PUT("system", configApi.UpdateSystem)

		configAdminRouter.GET("email", configApi.GetEmail)
		configAdminRouter.PUT("email", configApi.UpdateEmail)

		configAdminRouter.GET("qq", configApi.GetQQ)
		configAdminRouter.PUT("qq", configApi.UpdateQQ)

		configAdminRouter.GET("qiniu", configApi.GetQiniu)
		configAdminRouter.PUT("qiniu", configApi.UpdateQiniu)

		configAdminRouter.GET("jwt", configApi.GetJwt)
		configAdminRouter.PUT("jwt", configApi.UpdateJwt)

		configAdminRouter.GET("gaode", configApi.GetGaode)
		configAdminRouter.PUT("gaode", configApi.UpdateGaode)
	}

}
