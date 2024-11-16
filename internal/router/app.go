package router

import (
	"github.com/gin-gonic/gin"

	"go-jocy/internal/controllers"
)

// InitRouterApp Api
func InitRouterApp(r *gin.RouterGroup) {
	r.GET("/users/info", controllers.UserInfo)
}
