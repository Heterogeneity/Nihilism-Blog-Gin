package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"server/global"
	"server/model/database"
	"server/model/request"
	"server/model/response"
	"time"
)

type WebsiteApi struct {
}

// WebsiteAddCarousel 添加首页背景
func (websiteApi *WebsiteApi) WebsiteAddCarousel(c *gin.Context) {
	var req request.WebsiteCarouselOperation
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = websiteService.WebsiteAddCarousel(req)
	if err != nil {
		global.Log.Error("failed to create carousel", zap.Error(err))
		response.FailWithMessage("failed to create carousel", c)
		return
	}
	response.OkWithMessage("successfully created carousel", c)
}

// WebsiteCancelCarousel 移除首页背景
func (websiteApi *WebsiteApi) WebsiteCancelCarousel(c *gin.Context) {
	var req request.WebsiteCarouselOperation
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = websiteService.WebsiteCancelCarousel(req)
	if err != nil {
		global.Log.Error("failed to cancel carousel", zap.Error(err))
		response.FailWithMessage("failed to cancel carousel", c)
		return
	}
	response.OkWithMessage("carousel successfully cancelled", c)
}

// WebsiteCreateFooterLink 创建页脚链接
func (websiteApi *WebsiteApi) WebsiteCreateFooterLink(c *gin.Context) {
	var req database.FooterLink
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = websiteService.WebsiteCreateFooterLink(req)
	if err != nil {
		global.Log.Error("failed to get footer link", zap.Error(err))
		response.FailWithMessage("failed to get footer link", c)
		return
	}
	response.OkWithMessage("successfully created footer link", c)
}

// WebsiteDeleteFooterLink 删除页脚链接
func (websiteApi *WebsiteApi) WebsiteDeleteFooterLink(c *gin.Context) {
	var req database.FooterLink
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = websiteService.WebsiteDeleteFooterLink(req)
	if err != nil {
		global.Log.Error("failed to delete footer link", zap.Error(err))
		response.FailWithMessage("failed to delete footer link", c)
		return
	}
	response.OkWithMessage("successfully deleted footer link", c)
}

// WebsiteLogo 网站 Logo 链接
func (websiteApi *WebsiteApi) WebsiteLogo(c *gin.Context) {
	if global.Config.Website.Logo != "" {
		c.Redirect(http.StatusMovedPermanently, global.Config.Website.Logo)
	} else {
		c.Redirect(http.StatusMovedPermanently, "images/logo.svg")
	}
}

// WebsiteTitle 网站标题栏
func (websiteApi *WebsiteApi) WebsiteTitle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"title": global.Config.Website.Title,
	})
}

// WebsiteInfo 获取网站信息
func (websiteApi *WebsiteApi) WebsiteInfo(c *gin.Context) {
	response.OkWithData(global.Config.Website, c)
}

// WebsiteCarousel 获取首页背景
func (websiteApi *WebsiteApi) WebsiteCarousel(c *gin.Context) {
	urls := websiteService.WebsiteCarousel()
	response.OkWithData(urls, c)
}

// WebsiteNews 获取新闻
func (websiteApi *WebsiteApi) WebsiteNews(c *gin.Context) {
	sourceStr := c.Query("source")
	hotSearchData, err := websiteService.WebsiteNews(sourceStr)
	if err != nil {
		global.Log.Error("Failed to get news:", zap.Error(err))
		response.FailWithMessage("Failed to get news", c)
		return
	}
	response.OkWithData(hotSearchData, c)
}

// WebsiteCalendar 获取日历
func (websiteApi *WebsiteApi) WebsiteCalendar(c *gin.Context) {
	dateStr := time.Now().Format("2006/0102")
	calendar, err := websiteService.WebsiteCalendar(dateStr)
	if err != nil {
		global.Log.Error("failed to get calendar", zap.Error(err))
		response.FailWithMessage("failed to get calendar", c)
		return
	}
	response.OkWithData(calendar, c)
}

// WebsiteFooterLink 获取页脚链接
func (websiteApi *WebsiteApi) WebsiteFooterLink(c *gin.Context) {
	footerLinks := websiteService.WebsiteFooterLink()
	response.OkWithData(footerLinks, c)
}
