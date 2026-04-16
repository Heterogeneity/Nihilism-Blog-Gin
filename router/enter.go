package router

type RouterGroup struct {
	BaseRouter
	UserRouter
	ImageRouter
	FeedbackRouter
	FriendLinkRouter
	AdvertisementRouter
	ArticleRouter
	CommentRouter
	ConfigRouter
	WebsiteRouter
}

var RouterGroupApp = new(RouterGroup)
