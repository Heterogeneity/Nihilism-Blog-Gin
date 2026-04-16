package other

type HotItem struct {
	Index       int    `json:"index"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Popularity  string `json:"popularity"`
	URL         string `json:"url"`
}

type HotSearchData struct {
	Source     string    `json:"source"`      //
	UpdateTime string    `json:"update_time"` //
	HotList    []HotItem `json:"hot_list"`    //
}
