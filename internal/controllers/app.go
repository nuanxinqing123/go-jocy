package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-jocy/config"
	"go-jocy/utils"
)

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
	fmt.Println(resp.String())

	c.String(http.StatusOK, resp.String())
}
