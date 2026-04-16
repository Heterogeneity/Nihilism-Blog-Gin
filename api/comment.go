package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/utils"
)

type CommentApi struct {
}

// CommentCreate 创建评论
func (commentApi *CommentApi) CommentCreate(c *gin.Context) {
	var req request.CommentCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	req.UserUID = utils.GetUUID(c)
	err := commentService.CommentCreate(req)
	if err != nil {
		global.Log.Error("failed to create comment:", zap.Error(err))
		response.FailWithMessage("failed to create comment", c)
		return
	}
	response.OkWithMessage("create comment successfully", c)
}

// CommentDelete 删除评论
func (commentApi *CommentApi) CommentDelete(c *gin.Context) {
	var req request.CommentDelete
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = commentService.CommentDelete(c, req)
	if err != nil {
		global.Log.Error("delete comment failed:", zap.Error(err))
		response.FailWithMessage("delete comment failed", c)
		return
	}
	response.OkWithMessage("delete comment successfully", c)
}

// CommentInfo 获取用户评论
func (commentApi *CommentApi) CommentInfo(c *gin.Context) {
	uid := utils.GetUUID(c)
	list, err := commentService.CommentInfo(uid)
	if err != nil {
		global.Log.Error("failed to get comment information", zap.Error(err))
		response.FailWithMessage("failed to get comment information", c)
		return
	}
	response.OkWithData(list, c)
}

// CommentInfoByArticleID 根据文章id获取评论信息
func (commentApi *CommentApi) CommentInfoByArticleID(c *gin.Context) {
	var req request.CommentInfoByArticleID
	err := c.ShouldBindUri(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, err := commentService.CommentInfoByArticleID(req)
	if err != nil {
		global.Log.Error("failed to get comment information", zap.Error(err))
		response.FailWithMessage("failed to get comment information", c)
		return
	}
	response.OkWithData(list, c)
}

// CommentNew 获取最新评论
func (commentApi *CommentApi) CommentNew(c *gin.Context) {
	list, err := commentService.CommentNew()
	if err != nil {
		global.Log.Error("failed to get new comment:", zap.Error(err))
		response.FailWithMessage("fail to get new comment", c)
		return
	}
	response.OkWithData(list, c)
}

// CommentList 获取评论列表
func (commentApi *CommentApi) CommentList(c *gin.Context) {
	var pageInfo request.CommentList
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := commentService.CommentList(pageInfo)
	if err != nil {
		global.Log.Error("failed to get comment list:", zap.Error(err))
		response.FailWithMessage("failed to get comment list", c)
		return
	}
	response.OkWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}
