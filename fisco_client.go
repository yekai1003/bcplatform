package main

//fisco-bcos
const fisco_jsonstr = `{"jsonrpc":"2.0","method":"getPeers","params":[1],"id":1}`

type FiscoReq struct {
	connstr string
	jsonstr string
}

func NewFiscoReq() *FiscoReq {
	return &FiscoReq{Config.Fisco.Connstr, fisco_jsonstr}
}

func (r *FiscoReq) Http_req() ([]byte, error) {
	return Http_req(r.connstr, r.jsonstr)
}
