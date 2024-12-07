package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"

	"go-jocy/config"
	"go-jocy/internal/model"
	"go-jocy/utils"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// UserAvatar 随机头像
func UserAvatar(c *gin.Context) {
	// 获取Query参数中的ID
	id := c.Query("id")

	// 判断ID是否存在
	var url string
	var err error
	if id == "" {
		// 未传递ID, 从列表中随机获取一个元素
		url = utils.RandomGetElements(model.AvatarURL)
	} else {
		// 传递ID
		url, err = utils.BindUserToUrl(id, model.AvatarURL)
		if err != nil {
			config.GinLOG.Warn(err.Error())
			url = utils.RandomGetElements(model.AvatarURL)
		}
	}

	// 重定向
	config.GinLOG.Debug(fmt.Sprintf("Redirect: %s", url))
	c.Redirect(http.StatusMovedPermanently, url)
}

// UserCaptcha 用户验证
func UserCaptcha(c *gin.Context) {
	type Captcha struct {
		Type string `json:"type" required:"true"`
	}

	p := new(Captcha)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	client := utils.New("", c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/users/captcha"

	// Struct 转 String
	jsonStr, err := json.Marshal(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 加密参数
	enText, err := utils.EncryptRequests(string(jsonStr))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	resp, err := client.Post(url, enText)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, result)
}

// UserSmsCode 发送验证码
func UserSmsCode(c *gin.Context) {
	type SmsCode struct {
		Phone string `json:"phone" required:"true"`
		Type  string `json:"type" required:"true"`
		UUID  string `json:"uuid" required:"true"`
		Dots  string `json:"dots" required:"true"`
	}

	p := new(SmsCode)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	client := utils.New("", c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/users/smscode"

	// Struct 转 String
	jsonStr, err := json.Marshal(map[string]any{
		"phone": p.Phone,
		"type":  p.Type,
		"enum":  0,
		"uuid":  p.UUID,
		"dots":  p.Dots,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 加密参数
	enText, err := utils.EncryptRequests(string(jsonStr))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	resp, err := client.Post(url, enText)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, result)
}

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	type Register struct {
		Phone    string `json:"phone" required:"true"`
		Password string `json:"password" required:"true"`
		SmsCode  string `json:"sms_code" required:"true"`
		UserName string `json:"user_name" required:"true"`
	}

	p := new(Register)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	client := utils.New("", c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/users/register"

	// Struct 转 String
	jsonStr, err := json.Marshal(map[string]any{
		"phone":     p.Phone,
		"password":  p.Password,
		"smscode":   p.SmsCode,
		"user_name": p.UserName,
		"enum":      0,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 加密参数
	enText, err := utils.EncryptRequests(string(jsonStr))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	resp, err := client.Post(url, enText)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, result)
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	type Login struct {
		Phone    string `json:"phone" required:"true"`
		Password string `json:"password" required:"true"`
	}

	p := new(Login)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	client := utils.New("", c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/users/login"

	// Struct 转 String
	jsonStr, err := json.Marshal(map[string]any{
		"phone":    p.Phone,
		"password": p.Password,
		"enum":     0,
		"symbol":   utils.RandomString(16),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 加密参数
	enText, err := utils.EncryptRequests(string(jsonStr))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	resp, err := client.Post(url, enText)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, result)
}

// UserLogout 用户退出登录
func UserLogout(c *gin.Context) {

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/users/logout"

	resp, err := client.Post(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, result)
}

// UserInfo 用户信息
func UserInfo(c *gin.Context) {

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/users/info"

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, result)
}

// MessageBox 消息通知
func MessageBox(c *gin.Context) {

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/messagebox"

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, result)
}

// MessageBoxType 获取消息通知
func MessageBoxType(c *gin.Context) {

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/messagebox/" + c.Param("type") + "?" + c.Request.URL.RawQuery

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, result)
}

// Channel 频道数据
func Channel(c *gin.Context) {

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/channel?top-level=true"

	kvName := "channel"
	cache, err := config.GinCache.GetIFPresent(kvName)
	if err == nil {
		c.String(http.StatusOK, cache.(string))
		return
	}

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 写入缓存
	if resp.StatusCode() == http.StatusOK {
		_ = config.GinCache.SetWithExpire(kvName, result, time.Hour)
	}

	c.String(http.StatusOK, result)
}

// VideoList 视频列表
func VideoList(c *gin.Context) {

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/video/list?" + c.Request.URL.RawQuery

	kvName := "video:list:" + c.Request.URL.RawQuery
	cache, err := config.GinCache.GetIFPresent(kvName)
	if err == nil {
		c.String(http.StatusOK, cache.(string))
		return
	}

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 写入缓存
	if resp.StatusCode() == http.StatusOK {
		_ = config.GinCache.SetWithExpire(kvName, result, time.Minute*10)
	}

	c.String(http.StatusOK, result)
}

// Banners 横幅数据
func Banners(c *gin.Context) {

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/banners/" + c.Param("id")

	kvName := "banners/" + c.Param("id")
	cache, err := config.GinCache.GetIFPresent(kvName)
	if err == nil {
		c.String(http.StatusOK, cache.(string))
		return
	}

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 写入缓存
	if resp.StatusCode() == http.StatusOK {
		_ = config.GinCache.SetWithExpire(kvName, result, time.Minute*30)
	}

	c.String(http.StatusOK, result)
}

// VideoUpdateList 视频更新列表
func VideoUpdateList(c *gin.Context) {

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/video_update_list/" + c.Param("date") + "?" + c.Request.URL.RawQuery

	kvName := "video_update_list:date:" + c.Param("date") + ":" + c.Request.URL.RawQuery
	cache, err := config.GinCache.GetIFPresent(kvName)
	if err == nil {
		c.String(http.StatusOK, cache.(string))
		return
	}

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 写入缓存
	if resp.StatusCode() == http.StatusOK {
		_ = config.GinCache.SetWithExpire(kvName, result, time.Minute*30)
	}

	c.String(http.StatusOK, result)
}

// VideoDetail 视频详情
func VideoDetail(c *gin.Context) {

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/video/detail?" + c.Request.URL.RawQuery

	kvName := "video:detail:" + c.Request.URL.RawQuery
	cache, err := config.GinCache.GetIFPresent(kvName)
	if err == nil {
		c.String(http.StatusOK, cache.(string))
		return
	}

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 写入缓存
	if resp.StatusCode() == http.StatusOK {
		_ = config.GinCache.SetWithExpire(kvName, result, time.Hour)
	}

	c.String(http.StatusOK, result)
}

// VodCommentGetHitStop 获取视频热评
func VodCommentGetHitStop(c *gin.Context) {

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/vod_comment/gethitstop?" + c.Request.URL.RawQuery

	kvName := "vod_comment:gethitstop:" + c.Request.URL.RawQuery
	cache, err := config.GinCache.GetIFPresent(kvName)
	if err == nil {
		c.String(http.StatusOK, cache.(string))
		return
	}

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 写入缓存
	if resp.StatusCode() == http.StatusOK {
		_ = config.GinCache.SetWithExpire(kvName, result, time.Hour)
	}

	c.String(http.StatusOK, result)
}

// VodCommentGetList 获取视频评论列表
func VodCommentGetList(c *gin.Context) {

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/vod_comment/getlist?" + c.Request.URL.RawQuery

	kvName := "vod_comment:getlist:" + c.Request.URL.RawQuery
	cache, err := config.GinCache.GetIFPresent(kvName)
	if err == nil {
		c.String(http.StatusOK, cache.(string))
		return
	}

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 写入缓存
	if resp.StatusCode() == http.StatusOK {
		_ = config.GinCache.SetWithExpire(kvName, result, time.Minute*10)
	}

	c.String(http.StatusOK, result)
}

// VodCommentGetSubList 获取视频子评论列表
func VodCommentGetSubList(c *gin.Context) {

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/vod_comment/getsublist?" + c.Request.URL.RawQuery

	kvName := "vod_comment:getsublist:" + c.Request.URL.RawQuery
	cache, err := config.GinCache.GetIFPresent(kvName)
	if err == nil {
		c.String(http.StatusOK, cache.(string))
		return
	}

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 写入缓存
	if resp.StatusCode() == http.StatusOK {
		_ = config.GinCache.SetWithExpire(kvName, result, time.Minute*10)
	}

	c.String(http.StatusOK, result)
}

// VideoPlay 获取视频播放线路
func VideoPlay(c *gin.Context) {
	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/video/play?" + c.Request.URL.RawQuery

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 解密数据
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		config.GinLOG.Error(fmt.Sprintf("Failed to deserialize data: %s", result))
		config.GinLOG.Error(err.Error())
		config.GinLOG.Error(fmt.Sprintf("Res String: %s", resp.String()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	config.GinLOG.Debug(fmt.Sprintf("Response: %s", result))

	// 序列化数据
	var res model.Play
	if err = json.Unmarshal([]byte(result), &res); err != nil {
		config.GinLOG.Error(fmt.Sprintf("Failed to deserialize data: %s", result))
		config.GinLOG.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if len(res.Data) == 0 {
		config.GinLOG.Warn(fmt.Sprintf("Failed to fetch data: %s", result))
		c.String(http.StatusInternalServerError, result)
		return
	}

	playURL, err := utils.DecryptPlayUrl(res.Data[0].Url)
	if err != nil {
		config.GinLOG.Error(fmt.Sprintf("Failed to deserialize data: %s", playURL))
		config.GinLOG.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, playURL)
}

// Danmu 弹幕数据
func Danmu(c *gin.Context) {
	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/danmu?" + c.Request.URL.RawQuery

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, resp.String())
}

// VideoSearch 搜索视频
func VideoSearch(c *gin.Context) {

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/video/search?" + c.Request.URL.RawQuery

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, result)
}

// VideoKey 视频预搜索
func VideoKey(c *gin.Context) {

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/video/key?" + c.Request.URL.RawQuery

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, result)
}

// PlayResources 获取播放资源
func PlayResources(c *gin.Context) {
	type Url struct {
		Url string `json:"url" required:"true"`
	}

	p := new(Url)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())

	resp, err := client.Get(p.Url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.Data(http.StatusOK, resp.Header().Get("Content-Type"), resp.Body())
}

// History 历史记录
func History(c *gin.Context) {

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/history?" + c.Request.URL.RawQuery

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, result)
}

// HistoryUpload 上传历史记录
func HistoryUpload(c *gin.Context) {
	type UploadHistory struct {
		Vid       int    `json:"vid" required:"true"`
		Play      string `json:"play" required:"true"`
		Part      string `json:"part" required:"true"`
		TimePoint int64  `json:"time_point" required:"true"`
	}

	p := new(UploadHistory)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/history"

	// Struct 转 String
	jsonStr, err := json.Marshal(map[string]any{
		"vid":        p.Vid,
		"play":       p.Play,
		"part":       p.Part,
		"time_point": p.TimePoint,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 加密参数
	enText, err := utils.EncryptRequests(string(jsonStr))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	resp, err := client.Post(url, enText)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, result)
}

// Collect 我的收藏
func Collect(c *gin.Context) {

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/collect?" + c.Request.URL.RawQuery

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, result)
}

// CollectCreate 创建收藏
func CollectCreate(c *gin.Context) {
	type CreateCollect struct {
		Vid int `json:"vid" required:"true"`
	}

	p := new(CreateCollect)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/collect"

	// Struct 转 String
	jsonStr, err := json.Marshal(map[string]int{
		"vid": p.Vid,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 加密参数
	enText, err := utils.EncryptRequests(string(jsonStr))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	resp, err := client.Post(url, enText)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, result)
}

// CollectDelete 删除收藏
func CollectDelete(c *gin.Context) {
	type CreateCollect struct {
		Vid int `json:"vid" required:"true"`
	}

	p := new(CreateCollect)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	client := utils.New(c.Request.Header.Get("x-token"), c.ClientIP())
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/collect"

	// Struct 转 String
	jsonStr, err := json.Marshal(map[string]int{
		"vid": p.Vid,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 加密参数
	enText, err := utils.EncryptRequests(string(jsonStr))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	resp, err := client.Delete(url, enText)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	result, err := utils.ResponseDecryption(resp.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, result)
}
