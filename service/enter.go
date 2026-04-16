package service

type ServiceGroup struct {
	EsService
	BaseService
	JwtService
	GaodeService
	UserService
	QQService
	ImageService
	FeedbackService
	FriendLinkService
	AdvertisementService
	ArticleService
	CommentService
	ConfigService
	WebsiteService
	HotSearchService
	CalendarService
}

var ServiceGroupApp = new(ServiceGroup)
