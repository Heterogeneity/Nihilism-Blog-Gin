package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/request"
	"server/model/response"
)

type AdvertisementApi struct {
}

// AdvertisementList 广告列表
func (advertisementApi *AdvertisementApi) AdvertisementList(c *gin.Context) {
	var pageInfo request.AdvertisementList
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := advertisementService.AdvertisementList(pageInfo)
	if err != nil {
		global.Log.Error("获取广告后台列表失败！", zap.Error(err))
		response.FailWithMessage("获取广告后台列表失败！", c)
		return
	}
	response.OkWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}

// AdvertisementCreate 广告创建
func (advertisementApi *AdvertisementApi) AdvertisementCreate(c *gin.Context) {
	var req request.AdvertisementCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = advertisementService.AdvertisementCreate(req)
	if err != nil {
		global.Log.Error("创建广告失败！", zap.Error(err))
		response.FailWithMessage("创建广告失败！", c)
		return
	}
	response.OkWithData("创建广告成功！", c)
}

// AdvertisementDelete 广告删除
func (advertisementApi *AdvertisementApi) AdvertisementDelete(c *gin.Context) {
	var req request.AdvertisementDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = advertisementService.AdvertisementDelete(req)
	if err != nil {
		global.Log.Error("删除广告失败！", zap.Error(err))
		response.FailWithMessage("删除广告失败！", c)
		return
	}
	response.OkWithMessage("删除广告成功！", c)
}

// AdvertisementUpdate 广告更新
func (advertisementApi *AdvertisementApi) AdvertisementUpdate(c *gin.Context) {
	var req request.AdvertisementUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = advertisementService.AdvertisementUpdate(req)
	if err != nil {
		global.Log.Error("更新广告失败！", zap.Error(err))
		response.FailWithMessage("更新广告失败！", c)
		return
	}
	response.OkWithMessage("更新广告成功！", c)
}

// AdvertisementInfo 广告前台列表
func (advertisementApi *AdvertisementApi) AdvertisementInfo(c *gin.Context) {
	list, total, err := advertisementService.AdvertisementInfo()
	if err != nil {
		global.Log.Error("获取广告列表失败！", zap.Error(err))
		response.FailWithMessage("获取广告列表失败！", c)
		return
	}
	response.OkWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}
