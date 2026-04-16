package other

// AccessTokenResponse 表示通过授权码获取的AccessToken返回结构
type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
}

// UserInfoResponse 表示获取用户消息的返回结构
type UserInfoResponse struct {
	Ret          int    `json:"ret"`
	Msg          string `json:"msg"`
	IsLost       int    `json:"is_lost"`
	Nickname     string `json:"nickname"`
	Figureurl    string `json:"figureurl"`
	Figureurl1   string `json:"figureurl_1"`
	Figureurl2   string `json:"figureurl_2"`
	FigureurlQQ1 string `json:"figureurl_qq_1"`
	FigureurlQQ2 string `json:"figureurl_qq_2"`
}
