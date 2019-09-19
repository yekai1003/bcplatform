package main

import (
	"fmt"

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

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("ping", ping)
	r.GET("status", platformRun)
	r.GET("nodeinfo/:name", nodeinfo)
	r.GET("testdata/:name", peerinfo)
	r.Run(":8080")
}
