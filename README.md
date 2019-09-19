# bcplatform

### 服务启动 

```
localhost:bcplatform yekai$ go run *.go -c bc.toml
```

### 开发接口

#### 测试后端服务是否联通

请求如下
```
localhost:test yekai$ curl http://localhost:8080/ping
{"code":"0","msg":"OK","data":null}
```

#### 查看区块链哪个平台处于启动状态

请求如下

```
localhost:test yekai$ curl http://localhost:8080/status
{"code":"0","msg":"OK","data":{"eth":true,"fisco":false,"eos":false}}

```
- false 代表未启动
- true 代表该平台已启动

#### 查看具体某个区块链节点启动情况


请求以太坊情况如下：
```
localhost:test yekai$ curl http://localhost:8080/nodeinfo/eth
{"code":"0","msg":"OK","data":"{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":[{\"enode\":\"enode://16a3884083586ac0e0caefd13d2831714434d913657796e55a08f2665b58a139fa80e422c5be2e59eb0555a794d4345f7413aeac398f10bdefbb1d791f42e4f5@127.0.0.1:30304?discport=0\",\"id\":\"d8d8d732dc69f960eb5fb1db4929e47479ca791c97b7dc040076d366a3228f88\",\"name\":\"Geth/v1.8.22-stable/darwin-amd64/go1.11.5\",\"caps\":[\"eth/63\"],\"network\":{\"localAddress\":\"127.0.0.1:52975\",\"remoteAddress\":\"127.0.0.1:30304\",\"inbound\":false,\"trusted\":false,\"static\":true},\"protocols\":{\"eth\":{\"version\":63,\"difficulty\":131072,\"head\":\"0x5e1fc79cb4ffa4739177b5408045cd5d51c6cf766133f23f7cd72ee1f8d790e0\"}}}]}\n"}
```
请求fisco情况如下：
```
yekai@ubtu-amd64:~$ curl http://localhost:8080/nodeinfo/fisco
{"code":"0","msg":"OK","data":"{\"id\":1,\"jsonrpc\":\"2.0\",\"result\":[{\"Agency\":\"agency\",\"IPAndPort\":\"127.0.0.1:30303\",\"Node\":\"node3\",\"NodeID\":\"b76aa083bebe1be768abf9eb431e128830f97aeb605a13a5029e982e28b01e51cead6dea41f662564ec99c0420fb76a2a404bbbf3e4d2fd1048e33cf111c38ca\",\"Topic\":[]},{\"Agency\":\"agency\",\"IPAndPort\":\"127.0.0.1:30301\",\"Node\":\"node1\",\"NodeID\":\"84ef901cd10078dbdb59d53d3531fdda06e7bc227484dfc4bc1387e3e89eb3b67ea4065cc6e60122ce2c52ec47dda1c1a2b4c46fc1aa37e55e82a5e213549e71\",\"Topic\":[]},{\"Agency\":\"agency\",\"IPAndPort\":\"127.0.0.1:30302\",\"Node\":\"node2\",\"NodeID\":\"3e2cd70fe832a745a17dab5943bc121984409c59bd4f2773e421d038970449de34ea2236fc97f434ae355639d1cca5096c3296182d93631a174e77734a38786e\",\"Topic\":[]}]}\n"}
```

#### 节点测试数据接口

eth
```
localhost:bcplatform yekai$ curl http://localhost:8080/testdata/eth
{"code":"0","msg":"OK","data":{"nodetype":"multi","peers":["127.0.0.1:30303","127.0.0.1:30304","127.0.0.1:30305"]}}
```

fisco 
```
localhost:bcplatform yekai$ curl http://localhost:8080/testdata/fisco
{"code":"0","msg":"OK","data":{"nodetype":"single","peers":["127.0.0.1:30303","127.0.0.1:30304","127.0.0.1:30305","127.0.0.1:30306"]}}
```

eos 
```
localhost:bcplatform yekai$ curl http://localhost:8080/testdata/eos
{"code":"0","msg":"OK","data":{"nodetype":"single","peers":["127.0.0.1:8888"]}}
```

