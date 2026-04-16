package service

import (
	"gorm.io/gorm"
	"server/global"
	"server/model/appTypes"
	"server/model/database"
	"server/model/other"
	"server/model/request"
	"server/utils"
)

type FriendLinkService struct {
}

// FriendLinkCreate 友链创建
func (friendLinkService *FriendLinkService) FriendLinkCreate(req request.FriendLinkCreate) error {
	linkToCreate := database.FriendLink{
		Logo:        req.Logo,
		Link:        req.Link,
		Name:        req.Name,
		Description: req.Description,
	}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := utils.ChangeImagesCategory(tx, []string{linkToCreate.Logo}, appTypes.Logo); err != nil {
			return err
		}
		return tx.Create(&linkToCreate).Error
	})

}

// FriendLinkDelete 友链删除
func (friendLinkService *FriendLinkService) FriendLinkDelete(req request.FriendLinkDelete) error {
	if len(req.IDs) == 0 {
		return nil
	}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		for _, id := range req.IDs {
			var friendLink database.FriendLink
			if err := tx.Take(&friendLink, id).Error; err != nil {
				return err
			}
			if err := utils.InitImagesCategory(tx, []string{friendLink.Logo}); err != nil {
				return err
			}
			if err := tx.Delete(&friendLink).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// FriendLinkUpdate 友链更新
func (friendLinkService *FriendLinkService) FriendLinkUpdate(req request.FriendLinkUpdate) error {
	update := struct {
		Link        string `json:"link"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}{
		Link:        req.Link,
		Name:        req.Name,
		Description: req.Description,
	}
	return global.DB.Take(&database.FriendLink{}, req.ID).Updates(update).Error
}

// FriendLinkList 友链列表
func (friendLinkService *FriendLinkService) FriendLinkList(req request.FriendLinkList) (list interface{}, total int64, err error) {
	db := global.DB
	if req.Name != nil {
		db = db.Where("name LIKE ?", "%"+*req.Name+"%")
	}
	if req.Description != nil {
		db = db.Where("description LIKE ?", "%"+*req.Description+"%")
	}
	option := other.MySQLOption{
		PageInfo: req.PageInfo,
		Where:    db,
	}
	return utils.MySQLPagination(&database.FriendLink{}, option)
}

// FriendLinkInfo 友链前台列表
func (friendLinkService *FriendLinkService) FriendLinkInfo() (links []database.FriendLink, total int64, err error) {
	err = global.DB.Model(&database.FriendLink{}).Count(&total).Find(&links).Error
	if err != nil {
		return nil, 0, err
	}
	return links, total, nil
}
