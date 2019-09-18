package main

import (
	"github.com/gin-gonic/gin"
)

type RespMsg struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ping(c *gin.Context) {
	resp := &RespMsg{
		"0",
		"OK",
		nil,
	}
	c.JSON(200, resp)
}

//监控各个区块链平台状态
func blockstatus(c *gin.Context) {
	resp := &RespMsg{
		"0",
		"OK",
		nil,
	}

	c.JSON(200, resp)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("ping", ping)
	r.Run(":8080")
}
