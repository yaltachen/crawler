package main

import (
	"crawler/crawler/parser/zhenai"
	"crawler/crawler_distribute/config"
	"crawler/crawler_distribute/worker"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
)

func TestWorker(t *testing.T) {
	var (
		request worker.Request
		result  worker.ParserResult
		client  *rpc.Client
		conn    net.Conn
		err     error
	)
	if conn, err = net.Dial("tcp", ":9001"); err != nil {
		t.Fatal(err)
	}
	client = jsonrpc.NewClient(conn)

	request = worker.Request{
		URL: "https://album.zhenai.com/u/1513404786",
		Parser: worker.SerializedParser{
			Name: config.PERSONPARSER,
			Args: zhenai.PersonParser{
				ID:     "1513404786",
				URL:    "https://album.zhenai.com/u/1513404786",
				Name:   "心语",
				Gender: "女",
				Loc:    "北京",
			},
		},
	}
	if err = client.Call("CrawlerService.Process", request, &result); err != nil {
		t.Logf("call failed, err: %v", err)
	}

	t.Logf("%v", result)
}
