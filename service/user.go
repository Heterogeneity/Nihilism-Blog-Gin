package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"server/global"
	"server/model/appTypes"
	"server/model/database"
	"server/model/other"
	"server/model/request"
	"server/model/response"
	"server/utils"
	"time"
)

type UserService struct {
}

func (userService *UserService) Logout(c *gin.Context) {
	uid := utils.GetUUID(c)
	jwtStr := utils.GetRefreshToken(c)
	utils.ClearRefreshToken(c)
	global.Redis.Del(uid.String())
	_ = ServiceGroupApp.JwtService.JoinInBlacklist(database.JwtBlacklist{Jwt: jwtStr})
}

func (userService *UserService) UserResetPassword(req request.UserResetPassword) error {
	var user database.User
	if err := global.DB.Take(&user, req.UserID).Error; err != nil {
		return err
	}
	if ok := utils.BcryptCheck(req.Password, user.Password); !ok {
		return errors.New("原密码错误！")
	}
	user.Password = utils.BcryptHash(req.NewPassword)
	return global.DB.Save(&user).Error
}

func (userService *UserService) UserInfo(userID uint) (database.User, error) {
	var user database.User
	if err := global.DB.Take(&user, userID).Error; err != nil {
		return database.User{}, err
	}
	return user, nil
}

func (userService *UserService) UserChangeInfo(req request.UserChangeInfo) error {
	var user database.User
	if err := global.DB.Take(&user, req.UserID).Error; err != nil {
		return err
	}
	return global.DB.Model(&user).Updates(req).Error
}

func (userService *UserService) UserWeather(ip string) (string, error) {
	result, err := global.Redis.Get("weather:" + ip).Result()
	if err != nil {
		ipResponse, err := ServiceGroupApp.GaodeService.GetLocationByIP(ip)
		if err != nil {
			return "", err
		}
		live, err := ServiceGroupApp.GaodeService.GetWeatherByAdcode(ipResponse.Adcode)
		if err != nil {
			return "", err
		}
		weather := "地区：" + live.Province + "-" + live.City + "天气：" + live.Weather + "温度：" + live.Temperature + "°C" + "风向：" + live.WindDirection + "风级：" + live.WindPower + "湿度：" + live.Humidity + "%"
		if err := global.Redis.Set("weather"+ip, weather, time.Hour*1).Err(); err != nil {
			return "", err
		}
		return weather, nil
	}
	return result, nil
}

func (userService *UserService) UserChart(req request.UserChart) (response.UserChart, error) {
	where := global.DB.Where(fmt.Sprintf("date_sub(curdate(),interval %d day) <= created_at", req.Date))
	var res response.UserChart
	startDate := time.Now().AddDate(0, 0, -req.Date)
	for i := 1; i <= req.Date; i++ {
		res.DateList = append(res.DateList, startDate.AddDate(0, 0, i).Format("2006-01-02"))
	}
	loginCounts := utils.FetchDateCounts(global.DB.Model(&database.Login{}), where)
	registerCounts := utils.FetchDateCounts(global.DB.Model(&database.User{}), where)
	for _, date := range res.DateList {
		loginCount := loginCounts[date]
		registerCount := registerCounts[date]
		res.LoginData = append(res.LoginData, loginCount)
		res.RegisterData = append(res.RegisterData, registerCount)
	}
	return res, nil
}

func (userService *UserService) ForgotPassword(req request.ForgotPassword) error {
	var user database.User
	if err := global.DB.Where("email=?", req.Email).First(&user).Error; err != nil {
		return err
	}
	user.Password = utils.BcryptHash(req.NewPassword)
	return global.DB.Save(&user).Error
}

func (userService *UserService) UserCard(req request.UserCard) (response.UserCard, error) {
	var user database.User
	if err := global.DB.Where("uid = ?", req.UID).Select("uid", "username", "avatar", "address", "signature").First(&user).Error; err != nil {
		return response.UserCard{}, err
	}
	return response.UserCard{
		UID:       user.UID,
		Username:  user.Username,
		Avatar:    user.Avatar,
		Address:   user.Address,
		Signature: user.Signature,
	}, nil
}

func (userService *UserService) Register(u database.User) (database.User, error) {
	if !errors.Is(global.DB.Where("email=?", u.Email).First(&database.User{}).Error, gorm.ErrRecordNotFound) {
		return database.User{}, errors.New("邮箱已注册！")
	}
	u.Password = utils.BcryptHash(u.Password)
	u.UID = uuid.Must(uuid.NewV4())
	u.Avatar = "/image/avatar.jpg"
	u.RoleID = appTypes.User
	u.Register = appTypes.Email
	if err := global.DB.Create(&u).Error; err != nil {
		return database.User{}, err
	}
	return u, nil
}

func (userService *UserService) EmailLogin(u database.User) (database.User, error) {
	var user database.User
	err := global.DB.Where("email=?", u.Email).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return database.User{}, errors.New("邮箱或密码错误！")
		}
		return user, nil
	}
	return database.User{}, err
}

func (userService *UserService) QQLogin(accessTokenResponse other.AccessTokenResponse) (database.User, error) {
	var user database.User
	err := global.DB.Where("openid=?", accessTokenResponse.Openid).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return database.User{}, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		userInfoResponse, err := ServiceGroupApp.QQService.GetUserInfoByAccessTokenAndOpenid(accessTokenResponse.AccessToken, accessTokenResponse.Openid)
		if err != nil {
			return database.User{}, err
		}
		user.UID = uuid.Must(uuid.NewV4())
		user.Username = userInfoResponse.Nickname
		user.Openid = accessTokenResponse.Openid
		user.Avatar = userInfoResponse.FigureurlQQ2
		user.RoleID = appTypes.User
		user.Register = appTypes.QQ
		if err := global.DB.Create(&user).Error; err != nil {
			return database.User{}, err
		}
	}
	return user, nil
}

func (userService *UserService) UserList(info request.UserList) (interface{}, int64, error) {
	db := global.DB
	if info.UID != nil {
		db = db.Where("uid=?", info.UID)
	}
	if info.Register != nil {
		db = db.Where("register = ?", info.Register)
	}
	option := other.MySQLOption{
		PageInfo: info.PageInfo,
		Where:    db,
	}
	return utils.MySQLPagination(&database.User{}, option)
}

func (userService *UserService) UserFreeze(req request.UserOperation) error {
	var user database.User
	if err := global.DB.Take(&user, req.ID).Update("freeze", true).Error; err != nil {
		return err
	}
	jwtStr, _ := ServiceGroupApp.JwtService.GetRedisJWT(user.UID)
	if jwtStr != "" {
		_ = ServiceGroupApp.JwtService.JoinInBlacklist(database.JwtBlacklist{Jwt: jwtStr})
	}
	return nil
}

func (userService *UserService) UserUnfreeze(req request.UserOperation) error {
	return global.DB.Take(&database.User{}, req.ID).Update("freeze", false).Error
}

func (userService *UserService) UserLoginList(info request.UserLoginList) (interface{}, int64, error) {
	db := global.DB
	if info.UID != nil {
		var userID uint
		if err := global.DB.Model(&database.User{}).Where("uid = ?", info.UID).Pluck("id", &userID).Error; err != nil {
			return nil, 0, nil
		}
		db.Where("user_id = ?", userID)
	}
	option := other.MySQLOption{
		PageInfo: info.PageInfo,
		Where:    db,
		Preload:  []string{"User"},
	}
	return utils.MySQLPagination(&database.Login{}, option)
}
