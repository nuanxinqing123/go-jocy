package router

import (
	"github.com/gin-gonic/gin"

	"go-jocy/internal/controllers"
	"go-jocy/internal/middleware"
)

// InitRouterApp Api
func InitRouterApp(r *gin.RouterGroup) {
	// 随机头像
	r.GET("/users/avatar", controllers.UserAvatar)

	auth := r.Group("")
	auth.Use(middleware.Auth())
	// 验证码
	auth.POST("/users/captcha", controllers.UserCaptcha)
	// 发送验证码
	auth.POST("/users/smscode", controllers.UserSmsCode)
	// 注册
	auth.POST("/users/register", controllers.UserRegister)
	// 登录
	auth.POST("/users/login", controllers.UserLogin)
	// 退出
	auth.POST("/users/logout", controllers.UserLogout)
	// 用户信息
	auth.GET("/users/info", controllers.UserInfo)
	// 消息通知
	auth.GET("/messagebox", controllers.MessageBox)
	// 获取消息通知
	auth.GET("/messagebox/:type", controllers.MessageBoxType)
	// 频道
	auth.GET("/channel", controllers.Channel)
	// 视频
	auth.GET("/video/list", controllers.VideoList)
	// 横幅
	auth.GET("/banners/:id", controllers.Banners)
	// 视频更新
	auth.GET("/video_update_list/:date", controllers.VideoUpdateList)
	// 视频详情
	auth.GET("/video/detail", controllers.VideoDetail)
	// 视频热评
	auth.GET("/vod_comment/gethitstop", controllers.VodCommentGetHitStop)
	// 视频评论
	auth.GET("/vod_comment/getlist", controllers.VodCommentGetList)
	// 视频子评论
	auth.GET("/vod_comment/getsublist", controllers.VodCommentGetSubList)
	// 视频播放
	auth.GET("/video/play", controllers.VideoPlay)
	// 弹幕
	auth.GET("/danmu", controllers.Danmu)
	// 搜索
	auth.GET("/video/search", controllers.VideoSearch)
	// 预搜索
	auth.GET("/video/key", controllers.VideoKey)
	// 获取播放资源
	auth.POST("/play/resources", controllers.PlayResources)
	// 观看历史
	auth.GET("/history", controllers.History)
	// 上传历史
	auth.POST("/history", controllers.HistoryUpload)
	// 我的收藏
	auth.GET("/collect", controllers.Collect)
	// 我的收藏-创建
	auth.POST("/collect", controllers.CollectCreate)
	// 我的收藏-删除
	auth.DELETE("/collect", controllers.CollectDelete)
	// todo 发送弹幕
	// todo 发送评论
}
