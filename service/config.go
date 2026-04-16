package service

import (
	"gorm.io/gorm"
	"server/config"
	"server/global"
	"server/model/appTypes"
	"server/utils"
)

type ConfigService struct {
}

func (configService *ConfigService) UpdateWebsite(info config.Website) error {
	oldArray := []string{
		global.Config.Website.Logo,
		global.Config.Website.FullLogo,
		global.Config.Website.QQImage,
		global.Config.Website.WechatImage,
		global.Config.Website.Skill,
	}
	newArray := []string{
		info.Logo,
		info.FullLogo,
		info.QQImage,
		info.WechatImage,
		info.Skill,
	}

	added, removed := utils.DiffArrays(oldArray, newArray)
	return global.DB.Transaction(func(tx *gorm.DB) error {
		if err := utils.InitImagesCategory(global.DB, removed); err != nil {
			return err
		}
		if err := utils.ChangeImagesCategory(global.DB, added, appTypes.System); err != nil {
			return err
		}
		global.Config.Website = info
		if err := utils.SaveYAML(); err != nil {
			return err
		}
		return nil
	})
}
func (configService *ConfigService) UpdateSystem(info config.System) error {
	global.Config.System.UseMultipoint = info.UseMultipoint
	global.Config.System.SessionsSecret = info.SessionsSecret
	global.Config.System.OssType = info.OssType
	return utils.SaveYAML()
}
func (configService *ConfigService) UpdateEmail(info config.Email) error {
	global.Config.Email = info
	return utils.SaveYAML()
}
func (configService *ConfigService) UpdateQQ(info config.QQ) error {
	global.Config.QQ = info
	return utils.SaveYAML()
}
func (configService *ConfigService) UpdateQiniu(info config.Qiniu) error {
	global.Config.Qiniu = info
	return utils.SaveYAML()
}
func (configService *ConfigService) UpdateJwt(info config.Jwt) error {
	global.Config.Jwt = info
	return utils.SaveYAML()
}
func (configService *ConfigService) UpdateGaode(info config.Gaode) error {
	global.Config.Gaode = info
	return utils.SaveYAML()
}
