package other

// IPResponse 用于表示 IP 定位查询的响应结果
type IPResponse struct {
	Status    string `json:"status"`
	Info      string `json:"info"`
	InfoCode  string `json:"infocode"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Adcode    string `json:"adcode"`
	Rectangle string `json:"rectangle"`
}

// Cast 表示天气预报中的每日数据
type Cast struct {
	Date         string `json:"date"`
	Week         string `json:"week"`
	DayWeather   string `json:"dayweather"`
	NightWeather string `json:"nightweather"`
	DayTemp      string `json:"daytemp"`
	NightTemp    string `json:"nighttemp"`
	DayWind      string `json:"daywind"`
	NightWind    string `json:"nightwind"`
	DayPower     string `json:"daypower"`
	NightPower   string `json:"nightpower"`
}

// Live 表示实况天气数据
type Live struct {
	Province         string `json:"province"`
	City             string `json:"city"`
	Adcode           string `json:"adcode"`
	Weather          string `json:"weather"`
	Temperature      string `json:"temperature"`
	WindDirection    string `json:"winddirection"`
	WindPower        string `json:"windpower"`
	Humidity         string `json:"humidity"`
	ReportTime       string `json:"reporttime"`
	TemperatureFloat string `json:"temperaturefloat"`
	HumidityFloat    string `json:"humidityfloat"`
}

// Forecast 表示天气预报信息
type Forecast struct {
	City       string `json:"city"`
	Adcode     string `json:"adcode"`
	Province   string `json:"province"`
	ReportTime string `json:"reporttime"`
	Casts      []Cast `json:"casts"`
}

// WeatherResponse 用于表示天气查询的响应结果
type WeatherResponse struct {
	Status   string   `json:"status"`
	Count    string   `json:"count"`
	Info     string   `json:"info"`
	InfoCode string   `json:"infocode"`
	Lives    []Live   `json:"lives"`
	Forecast Forecast `json:"forecast"`
}
