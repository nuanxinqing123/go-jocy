package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"

	"go-jocy/config"
	"go-jocy/internal/model"
	"go-jocy/utils"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// UserInfo 用户信息
func UserInfo(c *gin.Context) {
	client := utils.New(c.Request.Header.Get("x-token"))
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

// Channel 频道数据
func Channel(c *gin.Context) {
	client := utils.New(c.Request.Header.Get("x-token"))
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/channel?top-level=true"

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

// VideoList 视频列表
func VideoList(c *gin.Context) {
	client := utils.New(c.Request.Header.Get("x-token"))
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/video/list?" + c.Request.URL.RawQuery

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

// Banners 横幅数据
func Banners(c *gin.Context) {
	client := utils.New(c.Request.Header.Get("x-token"))
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/banners/0"

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

// VideoUpdateList 视频更新列表
func VideoUpdateList(c *gin.Context) {
	client := utils.New(c.Request.Header.Get("x-token"))
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/video_update_list/" + c.Param("date") + "?" + c.Request.URL.RawQuery

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

// VideoDetail 视频详情
func VideoDetail(c *gin.Context) {
	client := utils.New(c.Request.Header.Get("x-token"))
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/video/detail?" + c.Request.URL.RawQuery

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

// VodCommentGetHitStop 获取视频热评
func VodCommentGetHitStop(c *gin.Context) {
	client := utils.New(c.Request.Header.Get("x-token"))
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/vod_comment/gethitstop?" + c.Request.URL.RawQuery

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

// VodCommentGetList 获取视频评论列表
func VodCommentGetList(c *gin.Context) {
	client := utils.New(c.Request.Header.Get("x-token"))
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/vod_comment/getlist?" + c.Request.URL.RawQuery

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

// VodCommentGetSubList 获取视频子评论列表
func VodCommentGetSubList(c *gin.Context) {
	client := utils.New(c.Request.Header.Get("x-token"))
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/vod_comment/getsublist?" + c.Request.URL.RawQuery

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

// VideoPlay 获取视频播放线路
func VideoPlay(c *gin.Context) {
	client := utils.New(c.Request.Header.Get("x-token"))
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/video/play?" + c.Request.URL.RawQuery

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

	// 序列化数据
	var res model.Play
	if err = json.Unmarshal([]byte(result), &res); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if len(res.Data) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Failed to fetch data",
		})
		return
	}

	playURL, err := utils.DecryptPlayUrl(res.Data[0].Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, playURL)
}

// Danmu 弹幕数据
func Danmu(c *gin.Context) {
	client := utils.New(c.Request.Header.Get("x-token"))
	url := utils.RandomChoice(config.GinConfig.App.BaseURL) + "/app/danmu?" + c.Request.URL.RawQuery

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

// VideoSearch 搜索视频
func VideoSearch(c *gin.Context) {
	client := utils.New(c.Request.Header.Get("x-token"))
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

// PlayResources 获取播放资源
func PlayResources(c *gin.Context) {
	url := c.Query("url")
	client := utils.New(c.Request.Header.Get("x-token"))

	resp, err := client.Get(url, nil)
	config.GinLOG.Debug(fmt.Sprintf("StatusCode: %d", resp.StatusCode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.Data(http.StatusOK, resp.Header().Get("Content-Type"), resp.Body())
}
