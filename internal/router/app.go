package router

import (
	"github.com/gin-gonic/gin"

	"go-jocy/internal/controllers"
)

// InitRouterApp Api
func InitRouterApp(r *gin.RouterGroup) {
	r.POST("/users/login", controllers.UserLogin)
	r.POST("/users/logout", controllers.UserLogout)
	r.GET("/users/info", controllers.UserInfo)
	r.GET("/channel", controllers.Channel)
	r.GET("/video/list", controllers.VideoList)
	r.GET("/banners", controllers.Banners)
	r.GET("/video_update_list/:date", controllers.VideoUpdateList)
	r.GET("/video/detail", controllers.VideoDetail)
	r.GET("/vod_comment/gethitstop", controllers.VodCommentGetHitStop)
	r.GET("/vod_comment/getlist", controllers.VodCommentGetList)
	r.GET("/vod_comment/getsublist", controllers.VodCommentGetSubList)
	r.GET("/video/play", controllers.VideoPlay)
	r.GET("/danmu", controllers.Danmu)
	r.GET("/video/search", controllers.VideoSearch)

	// 获取播放资源
	r.GET("/play/resources", controllers.PlayResources)
}
