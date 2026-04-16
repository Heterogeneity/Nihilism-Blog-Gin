package api

import "server/service"

type ApiGroup struct {
	BaseApi
	UserApi
	ImageApi
	FeedbackApi
	FriendLinkApi
	AdvertisementApi
	ArticleApi
	CommentApi
	ConfigApi
	WebsiteApi
}

var ApiGroupApp = new(ApiGroup)

var baseService = service.ServiceGroupApp.BaseService
var userService = service.ServiceGroupApp.UserService
var qqService = service.ServiceGroupApp.QQService
var jwtService = service.ServiceGroupApp.JwtService
var imageService = service.ServiceGroupApp.ImageService
var feedbackService = service.ServiceGroupApp.FeedbackService
var friendLinkService = service.ServiceGroupApp.FriendLinkService
var advertisementService = service.ServiceGroupApp.AdvertisementService
var articleService = service.ServiceGroupApp.ArticleService
var commentService = service.ServiceGroupApp.CommentService
var configService = service.ServiceGroupApp.ConfigService
var websiteService = service.ServiceGroupApp.WebsiteService
