package request

import "github.com/gofrs/uuid"

type CommentInfoByArticleID struct {
	ArticleID string `json:"article_id" form:"article_id" uri:"article_id" binding:"required"`
}

type CommentCreate struct {
	UserUID   uuid.UUID `json:"-"`
	ArticleID string    `json:"article_id" binding:"required"`
	PID       *uint     `json:"p_id"`
	Content   string    `json:"content" binding:"required,max=320"`
}

type CommentDelete struct {
	IDs []uint `json:"ids"`
}

type CommentList struct {
	ArticleID *string `json:"article_id" form:"article_id"`
	UserUID   *string `json:"user_uid" form:"user_uid"`
	Content   *string `json:"content" form:"content"`
	PageInfo
}
