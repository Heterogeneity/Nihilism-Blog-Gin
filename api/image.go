package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/request"
	"server/model/response"
)

type ImageApi struct {
}

// ImageUpload 图片上传
func (imageApi *ImageApi) ImageUpload(c *gin.Context) {
	_, header, err := c.Request.FormFile("image")
	if err != nil {
		global.Log.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	url, err := imageService.ImageUpload(header)
	if err != nil {
		global.Log.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.ImageUpload{
		Url:     url,
		OssType: global.Config.System.OssType,
	}, "上传成功！", c)
}

// ImageDelete 图片删除
func (imageApi *ImageApi) ImageDelete(c *gin.Context) {
	var req request.ImageDelete
	err := c.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = imageService.ImageDelete(req)
	if err != nil {
		global.Log.Error("删除失败！", zap.Error(err))
		response.FailWithMessage("删除失败！", c)
		return
	}
	response.OkWithMessage("删除图片成功！", c)
}

// ImageList 图片列表
func (imageApi *ImageApi) ImageList(c *gin.Context) {
	var req request.ImageList
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := imageService.ImageList(req)
	if err != nil {
		global.Log.Error("获取图片列表失败！", zap.Error(err))
		response.FailWithMessage("获取图片列表失败！", c)
		return
	}
	response.OkWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}
