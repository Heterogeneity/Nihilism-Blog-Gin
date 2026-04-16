package elasticsearch

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

// Article 文章表
type Article struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

	Cover    string   `json:"cover"`
	Title    string   `json:"title"`
	Keyword  string   `json:"keyword"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
	Abstract string   `json:"abstract"`
	Content  string   `json:"content"`

	Views    int `json:"views"`
	Comments int `json:"comments"`
	Likes    int `json:"likes"`
}

// ArticleIndex 文章 ES 索引
func ArticleIndex() string {
	return "article_index"
}

func ArticleMapping() *types.TypeMapping {
	return &types.TypeMapping{
		Properties: map[string]types.Property{
			"created_at": types.DateProperty{NullValue: nil, Format: func(s string) *string { return &s }("yyyy-MM-dd HH:mm:ss")},
			"updated_at": types.DateProperty{NullValue: nil, Format: func(s string) *string { return &s }("yyyy-MM-dd HH:mm:ss")},
			"cover":      types.TextProperty{},
			"title":      types.TextProperty{},
			"keyword":    types.KeywordProperty{},
			"category":   types.KeywordProperty{},
			"tags":       []types.KeywordProperty{},
			"abstract":   types.TextProperty{},
			"content":    types.TextProperty{},
			"views":      types.IntegerNumberProperty{},
			"comments":   types.IntegerNumberProperty{},
			"likes":      types.IntegerNumberProperty{},
		},
	}
}
