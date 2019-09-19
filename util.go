package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type NodesInfo struct {
	NodeType string   `json:"nodetype"`
	Peers    []string `json:"peers"`
}

func Http_req(url, jsonstr string) ([]byte, error) {
	fmt.Println(url)
	resp, err := http.Post(url, "application/json", bytes.NewBufferString(jsonstr))
	if err != nil {
		fmt.Println("Failed to Post ", url, err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
