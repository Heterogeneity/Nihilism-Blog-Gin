package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/request"
	"server/model/response"
)

type FriendLinkApi struct {
}

// FriendLinkCreate 友链创建
func (friendLinkApi *FriendLinkApi) FriendLinkCreate(c *gin.Context) {
	var req request.FriendLinkCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = friendLinkService.FriendLinkCreate(req)
	if err != nil {
		global.Log.Error("创建友链失败！", zap.Error(err))
		response.FailWithMessage("创建友链失败！", c)
		return
	}
	response.OkWithMessage("创建友链成功！", c)
}

// FriendLinkDelete 友链删除
func (friendLinkApi *FriendLinkApi) FriendLinkDelete(c *gin.Context) {
	var req request.FriendLinkDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		global.Log.Error(err.Error(), zap.Error(err))
	}
	err = friendLinkService.FriendLinkDelete(req)
	if err != nil {
		global.Log.Error("删除友链失败！", zap.Error(err))
		response.FailWithMessage("删除友链失败！", c)
		return
	}
	response.OkWithMessage("删除友链成功！", c)
}

// FriendLinkUpdate 友链更新
func (friendLinkApi *FriendLinkApi) FriendLinkUpdate(c *gin.Context) {
	var req request.FriendLinkUpdate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = friendLinkService.FriendLinkUpdate(req)
	if err != nil {
		global.Log.Error("更新友链失败！", zap.Error(err))
		response.FailWithMessage("更新友链失败！", c)
		return
	}
	response.OkWithMessage("更新友链成功！", c)

}

// FriendLinkList 友链列表
func (friendLinkApi *FriendLinkApi) FriendLinkList(c *gin.Context) {
	var req request.FriendLinkList
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := friendLinkService.FriendLinkList(req)
	if err != nil {
		global.Log.Error("获取友链后台列表失败！", zap.Error(err))
		response.FailWithMessage("获取友链后台列表失败！", c)
		return
	}
	response.OkWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}

// FriendLinkInfo 友链前台列表
func (friendLinkApi *FriendLinkApi) FriendLinkInfo(c *gin.Context) {
	list, total, err := friendLinkService.FriendLinkInfo()
	if err != nil {
		global.Log.Error("获取友链列表失败！", zap.Error(err))
		response.FailWithMessage("获取友链列表失败！", c)
		return
	}
	response.OkWithData(response.FriendLinkInfo{
		List:  list,
		Total: total,
	}, c)
}
