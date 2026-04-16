package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"server/global"
	"server/model/appTypes"
	"server/model/database"
	"server/model/other"
	"server/model/request"
	"server/utils"
)

type CommentService struct {
}

func (commentService *CommentService) CommentCreate(req request.CommentCreate) error {
	return global.DB.Create(&database.Comment{
		ArticleID: req.ArticleID,
		PID:       req.PID,
		UserUID:   req.UserUID,
		Content:   req.Content,
	}).Error
}
func (commentService *CommentService) CommentDelete(c *gin.Context, req request.CommentDelete) error {
	if len(req.IDs) == 0 {
		return nil
	}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		for _, id := range req.IDs {
			var comment database.Comment
			if err := global.DB.Take(&comment, id).Error; err != nil {
				return err
			}
			userUid := utils.GetUUID(c)
			userRoleId := utils.GetRoleID(c)
			if userUid != comment.UserUID && userRoleId != appTypes.Admin {
				return errors.New("you do not have permission to delete this comment")
			}
			if err := commentService.DeleteCommentAndChildren(tx, id); err != nil {
				return err
			}
		}
		return nil
	})

}
func (commentService *CommentService) CommentInfo(uid uuid.UUID) ([]database.Comment, error) {
	var rawComments []database.Comment
	err := global.DB.Order("id desc").Where("user_uid=?", uid).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("uid,username,avatar,address,signature")
	}).Find(&rawComments).Error
	if err != nil {
		return nil, err
	}
	for i := range rawComments {
		if err := commentService.LoadChildren(&rawComments[i]); err != nil {
			return nil, err
		}
	}
	var comments []database.Comment
	idMap := commentService.FindChildCommentsIDByRootCommentUserUUID(rawComments)
	for i := range rawComments {
		if _, exists := idMap[rawComments[i].ID]; !exists {
			comments = append(comments, rawComments[i])
		}
	}
	return rawComments, nil
}
func (commentService *CommentService) CommentInfoByArticleID(req request.CommentInfoByArticleID) ([]database.Comment, error) {
	var comment []database.Comment
	if err := global.DB.Where("article_id = ? AND p_id is NULL", req.ArticleID).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("uid,username,avatar,address,signature")
	}).Find(&comment).Error; err != nil {
		return nil, err
	}
	for i := range comment {
		if err := commentService.LoadChildren(&comment[i]); err != nil {
			return nil, err
		}
	}
	return comment, nil
}
func (commentService *CommentService) CommentNew() ([]database.Comment, error) {
	var comment []database.Comment
	err := global.DB.Order("id desc").Limit(5).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("uid,username,avatar,address,signature")
	}).Find(&comment).Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}
func (commentService *CommentService) CommentList(info request.CommentList) (interface{}, int64, error) {
	db := global.DB
	if info.ArticleID != nil {
		db = db.Where("article_id = ?", *info.ArticleID)
	}
	if info.UserUID != nil {
		db = db.Where("user_uid = ?", *info.UserUID)
	}
	if info.Content != nil {
		db = db.Where("content = ?", "%"+*info.Content+"%")
	}
	option := other.MySQLOption{
		PageInfo: info.PageInfo,
		Where:    db,
	}
	return utils.MySQLPagination(&database.Comment{}, option)
}
