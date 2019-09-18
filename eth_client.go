package main

import (
	"fmt"
)

type EthReq struct {
	connstr string
	jsonstr string
}

const eth_jsonstr = `{"jsonrpc":"2.0","method":"admin_peers","params":[],"id":1}`

func NewEthReq() *EthReq {
	return &EthReq{Config.Fisco.Connstr, eth_jsonstr}
}

func (r *EthReq) Http_req() ([]byte, error) {
	fmt.Println(r)
	return Http_req(r.connstr, r.jsonstr)
}
