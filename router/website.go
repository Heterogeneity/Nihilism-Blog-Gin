package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type WebsiteRouter struct {
}

func (w *WebsiteRouter) InitWebsiteRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	websiteRouter := Router.Group("website")
	websiteAdminRouter := PublicRouter.Group("website")
	websiteApi := api.ApiGroupApp.WebsiteApi
	{
		websiteAdminRouter.POST("addCarousel", websiteApi.WebsiteAddCarousel)
		websiteAdminRouter.PUT("cancelCarousel", websiteApi.WebsiteCancelCarousel)
		websiteAdminRouter.POST("createFooterLink", websiteApi.WebsiteCreateFooterLink)
		websiteAdminRouter.DELETE("deleteFooterLink", websiteApi.WebsiteDeleteFooterLink)
	}
	{
		websiteRouter.GET("logo", websiteApi.WebsiteLogo)
		websiteRouter.GET("title", websiteApi.WebsiteTitle)
		websiteRouter.GET("info", websiteApi.WebsiteInfo)
		websiteRouter.GET("carousel", websiteApi.WebsiteCarousel)
		websiteRouter.GET("news", websiteApi.WebsiteNews)
		websiteRouter.GET("calendar", websiteApi.WebsiteCalendar)
		websiteRouter.GET("footerLink", websiteApi.WebsiteFooterLink)
	}
}
