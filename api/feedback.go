package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/utils"
)

type FeedbackApi struct {
}

// FeedbackCreate 创建反馈
func (feedbackApi *FeedbackApi) FeedbackCreate(c *gin.Context) {
	var req request.FeedBackCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	req.UID = utils.GetUUID(c)
	err = feedbackService.FeedbackCreate(req)
	if err != nil {
		global.Log.Error("提交反馈失败！", zap.Error(err))
		response.FailWithMessage("提交反馈失败！", c)
		return
	}
	response.OkWithMessage("提交反馈成功！", c)
}

// FeedbackInfo 获取用户反馈信息
func (feedbackApi *FeedbackApi) FeedbackInfo(c *gin.Context) {
	uid := utils.GetUUID(c)
	list, err := feedbackService.FeedbackInfo(uid)
	if err != nil {
		global.Log.Error("获取反馈信息失败！", zap.Error(err))
		response.FailWithMessage("获取反馈信息失败！", c)
		return
	}
	response.OkWithData(list, c)
}

// FeedbackNew 获取最新反馈
func (feedbackApi *FeedbackApi) FeedbackNew(c *gin.Context) {
	list, err := feedbackService.FeedbackNew()
	if err != nil {
		global.Log.Error("获取新反馈失败！", zap.Error(err))
		response.FailWithMessage("获取新反馈失败！", c)
		return
	}
	response.OkWithData(list, c)
}

// FeedbackDelete 删除反馈
func (feedbackApi *FeedbackApi) FeedbackDelete(c *gin.Context) {
	var req request.FeedbackDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = feedbackService.FeedbackDelete(req)
	if err != nil {
		global.Log.Error("删除反馈失败！", zap.Error(err))
		response.FailWithMessage("删除反馈失败！", c)
		return
	}
	response.OkWithMessage("删除反馈成功！", c)
}

// FeedbackReply 回复反馈
func (feedbackApi *FeedbackApi) FeedbackReply(c *gin.Context) {
	var req request.FeedbackReply
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = feedbackService.FeedbackReply(req)
	if err != nil {
		global.Log.Error("回复反馈失败！", zap.Error(err))
		response.FailWithMessage("回复反馈失败！", c)
		return
	}
	response.OkWithMessage("回复反馈成功！", c)
}

// FeedbackList 获取反馈列表
func (feedbackApi *FeedbackApi) FeedbackList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := feedbackService.FeedbackList(pageInfo)
	if err != nil {
		global.Log.Error("获取反馈列表失败！", zap.Error(err))
		response.FailWithMessage("获取反馈列表失败", c)
		return
	}
	response.OkWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}
