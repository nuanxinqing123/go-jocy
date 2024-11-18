package model

type Play struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Play      string `json:"play"`
		PlayZh    string `json:"play_zh"`
		Part      string `json:"part"`
		AutoPlay  int    `json:"auto_play"`
		Url       string `json:"url"`
		Parse     string `json:"parse"`
		Ps        int    `json:"ps"`
		Extension struct {
			Player string `json:"player"`
			Url    string `json:"url"`
		} `json:"extension"`
		Completeness  int    `json:"completeness"`
		TotalDuration int    `json:"total_duration"`
		Resolution    string `json:"resolution"`
		LuaHeader     struct {
			UserAgent string `json:"User-Agent"`
		} `json:"lua_header"`
		LuaType string `json:"lua_type"`
		VipType int    `json:"vip_type"`
	} `json:"data"`
}

type PlayURL struct {
	Code    int    `json:"code"`
	Success int    `json:"success"`
	Type    string `json:"type"`
	Url     string `json:"url"`
	Msg     string `json:"msg"`
}
