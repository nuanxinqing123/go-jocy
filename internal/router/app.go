package router

import (
	"github.com/gin-gonic/gin"

	"go-jocy/internal/controllers"
)

// InitRouterApp Api
func InitRouterApp(r *gin.RouterGroup) {
	// 验证码
	r.POST("/users/captcha", controllers.UserCaptcha)
	// 发送验证码
	r.POST("/users/smscode", controllers.UserSmsCode)
	// 注册
	r.POST("/users/register", controllers.UserRegister)
	// 登录
	r.POST("/users/login", controllers.UserLogin)
	// 退出
	r.POST("/users/logout", controllers.UserLogout)
	// 用户信息
	r.GET("/users/info", controllers.UserInfo)
	// 消息通知
	r.GET("/messagebox", controllers.MessageBox)
	// 获取消息通知
	r.GET("/messagebox/:type", controllers.MessageBoxType)
	// 频道
	r.GET("/channel", controllers.Channel)
	// 视频
	r.GET("/video/list", controllers.VideoList)
	// 横幅
	r.GET("/banners", controllers.Banners)
	// 视频更新
	r.GET("/video_update_list/:date", controllers.VideoUpdateList)
	// 视频详情
	r.GET("/video/detail", controllers.VideoDetail)
	// 视频热评
	r.GET("/vod_comment/gethitstop", controllers.VodCommentGetHitStop)
	// 视频评论
	r.GET("/vod_comment/getlist", controllers.VodCommentGetList)
	// 视频子评论
	r.GET("/vod_comment/getsublist", controllers.VodCommentGetSubList)
	// 视频播放
	r.GET("/video/play", controllers.VideoPlay)
	// 弹幕
	r.GET("/danmu", controllers.Danmu)
	// 搜索
	r.GET("/video/search", controllers.VideoSearch)
	// 预搜索
	r.GET("/video/key", controllers.VideoKey)
	// 获取播放资源
	r.POST("/play/resources", controllers.PlayResources)
	// 观看历史
	r.GET("/history", controllers.History)
	// 上传历史
	r.POST("/history", controllers.HistoryUpload)
	// 我的收藏
	r.GET("/collect", controllers.Collect)
	// 我的收藏-创建
	r.POST("/collect", controllers.CollectCreate)
	// 我的收藏-删除
	r.DELETE("/collect", controllers.CollectDelete)
	// todo 发送弹幕
	// todo 发送评论
}
