package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParseDuration 解析持续时间字符串为 time.Duration。
// 持续时间字符串应由数字值和时间单位组成，单位可以是 "d" 表示天，"h" 表示小时，"m" 表示分钟，"s" 表示秒。
// 例如，"1d2h30m" 会被解析为 1 天、2 小时和 30 分钟。
// 如果字符串为空或格式无效，则返回错误。
func ParseDuration(d string) (time.Duration, error) {
	//去除空格
	d = strings.TrimSpace(d)
	if len(d) == 0 {
		return 0, fmt.Errorf("空字符串！")
	}
	//定义每个单位持续时间值
	unitPattern := map[string]time.Duration{
		"d": time.Hour * 24,
		"h": time.Hour,
		"m": time.Minute,
		"s": time.Second,
	}

	var totalDuration time.Duration
	//遍历单位
	for _, unit := range []string{"d", "h", "m", "s"} {
		for strings.Contains(d, unit) {
			//找到单位的位置
			unitIndex := strings.Index(d, unit)
			//提取单位前的部分
			part := d[:unitIndex]
			if part == "" {
				part = "0"
			}
			//转为整数
			val, err := strconv.Atoi(part)
			if err != nil {
				return 0, fmt.Errorf("无效的参数：%v", err)
			}
			//累积时间叠加
			totalDuration += time.Duration(val) * unitPattern[unit]
			//移除已处理的部分
			d = d[unitIndex+len(unit):]
		}
	}

	if len(d) > 0 {
		return 0, fmt.Errorf("时间格式未知。")
	}

	return totalDuration, nil
}
