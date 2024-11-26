package model

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

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

// URLField 定义一个类型，用于兼容单个字符串或数组
type URLField struct {
	Single string      `json:"single"`
	Multi  []URLDetail `json:"multi"`
}

type URLDetail struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type PlayURL struct {
	Code    int      `json:"code"`
	Success int      `json:"success"`
	Type    string   `json:"type"`
	Url     URLField `json:"url"`
	Msg     string   `json:"msg"`
}

// UnmarshalJSON 自定义 JSON 解码器
func (u *URLField) UnmarshalJSON(data []byte) error {
	// 尝试先解析为单个字符串
	var single string
	if err := json.Unmarshal(data, &single); err == nil {
		u.Single = single
		return nil
	}

	// 如果解析为字符串失败，尝试解析为数组
	var multi []URLDetail
	if err := json.Unmarshal(data, &multi); err == nil {
		u.Multi = multi
		return nil
	}

	// 如果都失败，返回错误
	return fmt.Errorf("unsupported format for URLField: %s", string(data))
}
