package hotSearch

import "server/model/other"

type Source interface {
	GetHotSearchData(maxNum int) (HotSearchData other.HotSearchData, err error)
}

func NewSource(sourceStr string) Source {
	switch sourceStr {
	case "baidu":
		return &Baidu{}
	case "toutiao":
		return &Toutiao{}
	case "zhihu":
		return &Zhihu{}
	case "kuaishou":
		return &Kuaishou{}
	default:
		return nil
	}
}
