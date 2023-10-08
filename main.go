package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"GoMusic/logic"
	"GoMusic/models"
)

func main() {
	e := gin.Default()
	e.POST("/neteasy", func(c *gin.Context) {
		link := c.PostForm("url")
		// 如果 link 不符合网易云规则，直接返回
		if !logic.IsNetEasyDiscover(link) {
			log.Printf("invalid link: %s", link)
			c.JSON(http.StatusBadRequest, &models.Result{Code: -1, Msg: "无效的网易云歌单链接~", Data: nil})
			return
		}
		netEasyDiscover, err := logic.NetEasyDiscover(link)
		if err != nil {
			log.Printf("fail to get net easy discover: %v", err)
			c.JSON(http.StatusBadRequest, &models.Result{Code: -1, Msg: err.Error(), Data: nil})
			return
		}
		c.JSON(200, &models.Result{
			Code: 1,
			Msg:  "success",
			Data: netEasyDiscover,
		})
	})
	e.Run(":8081")
}