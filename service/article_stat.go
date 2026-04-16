package service

import (
	"server/global"
	"strconv"
)

func (articleService *ArticleService) NewArticleView() CountDB {
	return CountDB{
		Index: "article_views",
	}
}

type CountDB struct {
	Index string
}

// Set 在原有基础上加一
func (c CountDB) Set(id string) error {
	return global.Redis.HIncrBy(c.Index, id, 1).Err()
}

// GetInfo 取出数据
func (c CountDB) GetInfo() map[string]int {
	var Info = map[string]int{}
	maps := global.Redis.HGetAll(c.Index).Val()

	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		Info[id] = num
	}
	return Info
}

// Clear 清除数据
func (c CountDB) Clear() {
	global.Redis.Del(c.Index)
}
