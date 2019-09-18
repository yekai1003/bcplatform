package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

//fisco-bcos
const jsonstr = `{"jsonrpc":"2.0","method":"getPeers","params":[1],"id":1}`

func main() {
	fmt.Println("hello world")
	resp, _ := http.Post("http://localhost:8080", "application/x-www-form-urlencoded", bytes.NewBufferString(jsonstr))
	// cli, err := jsonrpc.Dial("tcp", "http://localhost:8080")
	// cli.Call()
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
