package api

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"server/global"
	"server/model/database"
	"server/model/request"
	"server/model/response"
	"server/utils"
	"time"
)

type UserApi struct {
}

// Logout 退出登录
func (userApi *UserApi) Logout(c *gin.Context) {
	userService.Logout(c)
	response.OkWithMessage("退出成功！", c)
}

// UserResetPassword 密码重置
func (userApi *UserApi) UserResetPassword(c *gin.Context) {
	var req request.UserResetPassword
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	req.UserID = utils.GetUserID(c)
	err = userService.UserResetPassword(req)
	if err != nil {
		global.Log.Error("密码修改失败!", zap.Error(err))
		response.FailWithMessage("密码修改失败!", c)
		return
	}
	response.OkWithData("密码修改成功！", c)
	userApi.Logout(c)

}

// UserInfo 用户信息
func (userApi *UserApi) UserInfo(c *gin.Context) {
	userID := utils.GetUserID(c)
	user, err := userService.UserInfo(userID)
	if err != nil {
		global.Log.Error("获取用户资料失败！", zap.Error(err))
		response.FailWithMessage("获取用户资料失败！", c)
		return
	}
	response.OkWithData(user, c)
}

// UserChangeInfo 用户信息更改
func (userApi *UserApi) UserChangeInfo(c *gin.Context) {
	var req request.UserChangeInfo
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	req.UserID = utils.GetUserID(c)
	err = userService.UserChangeInfo(req)
	if err != nil {
		global.Log.Error("修改用户信息失败！", zap.Error(err))
		response.FailWithMessage("修改用户信息失败！", c)
		return
	}
	response.OkWithMessage("修改用户信息成功！", c)
}

// UserWeather 用户当地天气API
func (userApi *UserApi) UserWeather(c *gin.Context) {
	ip := c.ClientIP()
	if ip == "127.0.0.1" {
		ip = "112.49.249.228"
	}
	weather, err := userService.UserWeather(ip)
	if err != nil {
		global.Log.Error("获取天气失败：", zap.Error(err))
	}
	response.OkWithData(weather, c)
}

// UserChart 用户登录数据图表
func (userApi *UserApi) UserChart(c *gin.Context) {
	var req request.UserChart
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := userService.UserChart(req)
	if err != nil {
		global.Log.Error("获取图表数据失败", zap.Error(err))
		response.FailWithMessage("获取图表数据失败", c)
		return
	}
	response.OkWithData(data, c)
}

// ForgotPassword 忘记密码
func (userApi *UserApi) ForgotPassword(c *gin.Context) {
	var req request.ForgotPassword
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	session := sessions.Default(c)
	savedEmail := session.Get("email")
	if savedEmail == nil || savedEmail.(string) != req.Email {
		response.FailWithMessage("匹配邮箱无法验证！", c)
		return
	}
	savedCode := session.Get("verification_code")
	if savedCode == nil || savedCode.(string) != req.VerificationCode {
		response.FailWithMessage("无效的验证码！", c)
		return
	}
	savedTime := session.Get("expire_time")
	if savedTime.(int64) < time.Now().Unix() {
		response.FailWithMessage("验证码过期！", c)
		return
	}
	err := userService.ForgotPassword(req)
	if err != nil {
		global.Log.Error("取回密码失败！", zap.Error(err))
		response.FailWithMessage("取回密码失败！", c)
		return
	}
	response.OkWithMessage("取回密码成功！", c)
}

// UserCard 用户卡片
func (userApi *UserApi) UserCard(c *gin.Context) {
	var req request.UserCard
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userCard, err := userService.UserCard(req)
	if err != nil {
		global.Log.Error("获取卡片失败！", zap.Error(err))
		response.FailWithMessage("获取卡片失败！", c)
		return
	}
	response.OkWithData(userCard, c)
}

// Register 注册
func (userApi *UserApi) Register(c *gin.Context) {
	var req request.Register
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	session := sessions.Default(c)
	savedEmail := session.Get("email")
	if savedEmail == nil || savedEmail.(string) != req.Email {
		response.FailWithMessage("邮箱与验证邮箱不匹配", c)
		return
	}
	savedCode := session.Get("verification_code")
	if savedCode == nil || savedCode.(string) != req.VerificationCode {
		response.FailWithMessage("无效的验证码", c)
		return
	}

	savedTime := session.Get("expire_time")
	if savedTime.(int64) < time.Now().Unix() {
		response.FailWithMessage("验证码过期", c)
		return
	}

	u := database.User{Username: req.Username, Password: req.Password, Email: req.Email}
	user, err := userService.Register(u)
	if err != nil {
		global.Log.Error("注册失败", zap.Error(err))
		response.FailWithMessage("注册失败", c)
		return
	}
	userApi.TokenNext(c, user)
}

// Login 登录
func (userApi *UserApi) Login(c *gin.Context) {
	switch c.Query("flag") {
	case "email":
		userApi.EmailLogin(c)
	case "qq":
		userApi.QQLogin(c)
	default:
		userApi.EmailLogin(c)
	}
}

// EmailLogin 邮箱登录
func (userApi *UserApi) EmailLogin(c *gin.Context) {
	var req request.Login
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if store.Verify(req.CaptchaID, req.Captcha, true) {
		u := database.User{Email: req.Email, Password: req.Password}
		user, err := userService.EmailLogin(u)
		if err != nil {
			global.Log.Error("登陆失败：", zap.Error(err))
			response.FailWithMessage("登录失败!", c)
			return
		}
		userApi.TokenNext(c, user)
		return
	}
	response.FailWithMessage("验证码错误!", c)
}

// QQLogin QQ登录
func (userApi *UserApi) QQLogin(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		response.FailWithMessage("验证码已使用", c)
		return
	}
	accessTokenResponse, err := qqService.GetAccessTokenByCode(code)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user, err := userService.QQLogin(accessTokenResponse)
	if err != nil {
		global.Log.Error("登录失败：", zap.Error(err))
		response.FailWithMessage("登录失败！", c)
	}
	userApi.TokenNext(c, user)
}

// TokenNext token验证
func (userApi *UserApi) TokenNext(c *gin.Context, user database.User) {
	if user.Freeze {
		response.FailWithMessage("用户被冻结，请联系管理员！", c)
		return
	}
	baseClaims := request.BaseClaims{
		UserID: user.ID,
		UID:    user.UID,
		RoleID: user.RoleID,
	}
	j := utils.NewJWT()
	// 创建访问令牌
	accessClaims := j.CreateAccessClaims(baseClaims)
	accessToken, err := j.CreateAccessToken(accessClaims)
	if err != nil {
		global.Log.Error("获取accessToken失败！", zap.Error(err))
		response.FailWithMessage("获取accessToken失败！", c)
		return
	}
	// 创建刷新令牌
	refreshClaim := j.CreateRefreshClaims(baseClaims)
	refreshToken, err := j.CreateRefreshToken(refreshClaim)
	if err != nil {
		global.Log.Error("获取refreshToken失败！", zap.Error(err))
		response.FailWithMessage("获取refreshToken失败！", c)
		return
	}
	// 是否开启了多地点登录拦截
	if !global.Config.System.UseMultipoint {
		// 设置刷新令牌并返回
		utils.SetRefreshToken(c, refreshToken, int(refreshClaim.ExpiresAt.Unix()-time.Now().Unix()))
		c.Set("user_id", user.ID)
		response.OkWithDetailed(response.Login{
			User:                user,
			AccessToken:         accessToken,
			AccessTokenExpireAt: accessClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功！", c)
		return
	}
	// 检查 Redis 中是否已存在该用户的 JWT
	if jwtStr, err := jwtService.GetRedisJWT(user.UID); errors.Is(err, redis.Nil) {
		if err := jwtService.SetRedisJWT(refreshToken, user.UID); err != nil {
			global.Log.Error("设置登录状态失败！", zap.Error(err))
			response.FailWithMessage("设置登录状态失败！", c)
			return
		}
		utils.SetRefreshToken(c, refreshToken, int(refreshClaim.ExpiresAt.Unix()-time.Now().Unix()))
		c.Set("user_id", user.ID)
		response.OkWithDetailed(response.Login{
			User:                user,
			AccessToken:         accessToken,
			AccessTokenExpireAt: accessClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功！", c)
	} else if err != nil {
		// 出现错误处理
		global.Log.Error("设置登录状态失败！", zap.Error(err))
		response.FailWithMessage("设置登录状态失败！", c)
	} else {
		// Redis 中已存在该用户的 JWT，将旧的 JWT 加入黑名单，并设置新的 token
		var blacklist database.JwtBlacklist
		blacklist.Jwt = jwtStr
		if err := jwtService.JoinInBlacklist(blacklist); err != nil {
			global.Log.Error("设置登录状态失败！", zap.Error(err))
			response.FailWithMessage("设置登录状态失败！", c)
			return
		}
		// 设置新的 JWT 到 Redis
		if err := jwtService.SetRedisJWT(refreshToken, user.UID); err != nil {
			global.Log.Error("设置登录状态失败！", zap.Error(err))
			response.FailWithMessage("设置登录状态失败！", c)
			return
		}
		// 设置刷新令牌并返回
		utils.SetRefreshToken(c, refreshToken, int(refreshClaim.ExpiresAt.Unix()-time.Now().Unix()))
		c.Set("user_id", user.ID)
		response.OkWithDetailed(response.Login{
			User:                user,
			AccessToken:         accessToken,
			AccessTokenExpireAt: accessClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功!", c)
	}
}

// UserList 用户列表
func (userApi *UserApi) UserList(c *gin.Context) {
	var pageInfo request.UserList
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userService.UserList(pageInfo)
	if err != nil {
		global.Log.Error("获取用户列表失败", zap.Error(err))
		response.FailWithMessage("获取用户列表失败", c)
	}
	response.OkWithData(
		response.PageResult{
			List:  list,
			Total: total,
		}, c)
}

// UserFreeze 用户冻结
func (userApi *UserApi) UserFreeze(c *gin.Context) {
	var req request.UserOperation
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.UserFreeze(req)
	if err != nil {
		global.Log.Error("冻结失败！", zap.Error(err))
		response.FailWithMessage("冻结失败！", c)
		return
	}
	response.FailWithMessage("冻结成功！", c)
}

// UserUnfreeze 用户解冻
func (userApi *UserApi) UserUnfreeze(c *gin.Context) {
	var req request.UserOperation
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.UserUnfreeze(req)
	if err != nil {
		global.Log.Error("解冻失败！", zap.Error(err))
		response.FailWithMessage("解冻失败！", c)
		return
	}
	response.FailWithMessage("解冻成功！", c)
}

// UserLoginList 用户登录列表
func (userApi *UserApi) UserLoginList(c *gin.Context) {
	var pageInfo request.UserLoginList
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userService.UserLoginList(pageInfo)
	if err != nil {
		global.Log.Error("获取登录列表失败！", zap.Error(err))
		response.FailWithMessage("获取登录列表失败!", c)
		return
	}
	response.OkWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}
