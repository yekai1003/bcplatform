package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

//fisco-bcos
//const jsonstr = `{"jsonrpc":"2.0","method":"getPeers","params":[1],"id":1}`
const jsonstr = `{"jsonrpc":"2.0","method":"admin_nodeInfo","params":[],"id":1}`

func main() {
	fmt.Println("hello world")
	resp, err := http.Post("http://localhost:8546", "application/json", bytes.NewBufferString(jsonstr))
	// cli, err := jsonrpc.Dial("tcp", "http://localhost:8080")
	// cli.Call()
	if err != nil {
		fmt.Println("Failed to Post ", err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
