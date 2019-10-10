package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RespMsg struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PlatformRun struct {
	Eth   bool `json:"eth"`
	Fisco bool `json:"fisco"`
	Eos   bool `json:"eos"`
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
func nodeinfo(c *gin.Context) {
	resp := &RespMsg{
		"0",
		"OK",
		nil,
	}
	defer c.JSON(200, resp)
	name := c.Param("name")
	if name != "eth" && name != "fisco" && name != "eos" {
		resp.Code = "100"
		resp.Msg = "param err"
		return
	}
	var peers NodesInfo
	switch name {
	case "eth":
		eth := NewEthReq()
		bdata, err := eth.Http_req()
		if err != nil {
			fmt.Println("以太坊平台并未启动")
		} else {
			peers.NodeType = "multi"
			peers.Peers = append(peers.Peers, "127.0.0.1:30303")
			peers.Peers = append(peers.Peers, "127.0.0.1:30304")
			peers.Peers = append(peers.Peers, "127.0.0.1:30305")
			//resp.Data = peers
			resp.Data = string(bdata)
			// fmt.Println(resp.Data)
			return
		}
	case "eos":
	case "fisco":
		//fisco
		fisco := NewFiscoReq()
		bdata, err := fisco.Http_req()
		if err != nil {
			fmt.Println("fisco平台并未启动")
		} else {
			resp.Data = string(bdata)
		}
	}
}

//监控各个区块链平台状态
func peerinfo(c *gin.Context) {
	resp := &RespMsg{
		"0",
		"OK",
		nil,
	}
	defer c.JSON(200, resp)
	name := c.Param("name")

	var peers NodesInfo
	if name == "eth" {
		peers.NodeType = "multi"
		peers.Peers = append(peers.Peers, "127.0.0.1:30303")
		peers.Peers = append(peers.Peers, "127.0.0.1:30304")
		peers.Peers = append(peers.Peers, "127.0.0.1:30305")
		resp.Data = peers
	} else if name == "fisco" {
		peers.NodeType = "single"
		peers.Peers = append(peers.Peers, "127.0.0.1:30303")
		peers.Peers = append(peers.Peers, "127.0.0.1:30304")
		peers.Peers = append(peers.Peers, "127.0.0.1:30305")
		peers.Peers = append(peers.Peers, "127.0.0.1:30306")
		resp.Data = peers
	} else if name == "eos" {
		peers.NodeType = "single"
		peers.Peers = append(peers.Peers, "127.0.0.1:8888")
		resp.Data = peers
	} else {
		resp.Code = "100"
		resp.Msg = "Param err"
	}

}

//获得节点信息
func platformRun(c *gin.Context) {
	var rdata PlatformRun
	resp := &RespMsg{
		"0",
		"OK",
		&rdata,
	}
	defer c.JSON(200, resp)

	//一般认为只有一个平台处于运行状态
	{
		//eth
		eth := NewEthReq()
		_, err := eth.Http_req()
		if err != nil {
			fmt.Println("以太坊平台并未启动")
		} else {
			rdata.Eth = true
			return
		}
	}
	{
		//fisco
		fisco := NewFiscoReq()
		_, err := fisco.Http_req()
		if err != nil {
			fmt.Println("fisco平台并未启动")
		} else {
			rdata.Fisco = true
		}
	}

}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}

		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.StaticFile("/", "html/index.html")
	r.StaticFS("/html", http.Dir("html"))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(Cors())
	r.GET("ping", ping)
	r.GET("status", platformRun)
	r.GET("nodeinfo/:name", nodeinfo)
	r.GET("testdata/:name", peerinfo)
	r.Run(":8080")
}
