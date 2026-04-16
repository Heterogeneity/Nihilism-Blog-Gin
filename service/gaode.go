package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/global"
	"server/model/other"
	"server/utils"
)

// GaodeService 提供与高德相关的服务
type GaodeService struct {
}

// GetLocationByIP 根据IP地址获取地理位置信息
func (gaodeService *GaodeService) GetLocationByIP(ip string) (other.IPResponse, error) {
	data := other.IPResponse{}
	key := global.Config.Gaode.Key
	urlStr := "https://restapi.amap.com/v3/ip"
	method := "GET"
	params := map[string]string{
		"ip":  ip,
		"key": key,
	}
	res, err := utils.HttpRequest(urlStr, method, nil, params, nil)
	if err != nil {
		return data, err
	}
	//延迟函数return后执行 如果一个函数中有多个defer语句，它们会以LIFO（后进先出）的顺序执行。
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return data, fmt.Errorf("请求失败代码: %d", res.StatusCode)
	}
	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetWeatherByAdcode 根据城市编码获取实时天气信息
func (gaodeService *GaodeService) GetWeatherByAdcode(adcode string) (other.Live, error) {
	data := other.WeatherResponse{}
	key := global.Config.Gaode.Key
	urlStr := "https://restapi.amap.com/v3/weather/weatherInfo"
	method := "GET"
	params := map[string]string{
		"city": adcode,
		"key":  key,
	}
	res, err := utils.HttpRequest(urlStr, method, nil, params, nil)
	if err != nil {
		return other.Live{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return other.Live{}, fmt.Errorf("请求错误代码:%d", res.StatusCode)
	}
	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		return other.Live{}, err
	}
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return other.Live{}, err
	}

	if len(data.Lives) == 0 {
		return other.Live{}, fmt.Errorf("天气数据返回错误")
	}
	return data.Lives[0], nil
}
