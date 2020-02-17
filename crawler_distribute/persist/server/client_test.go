package main

import (
	"crawler/crawler/engine"
	"crawler/crawler/model"
	"crawler/crawler_distribute/persist"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"testing"
	"time"

	"github.com/olivere/elastic/v7"
)

func TestServer(t *testing.T) {
	var (
		conn      net.Conn
		listener  net.Listener
		esClient  *elastic.Client
		rpcClient *rpc.Client
		item      engine.Item
		result    string
		err       error
	)

	if esClient, err = elastic.NewClient(elastic.SetSniff(false),
		elastic.SetURL("http://192.168.94.30:9200")); err != nil {
		t.Fatalf("new elastic search failed, err: %v\r\n", err)
	}

	rpc.Register(&persist.Service{
		Index:  "test",
		Client: esClient,
	})

	if listener, err = net.Listen("tcp", ":9000"); err != nil {
		t.Fatalf("listen failed, err: %v\r\n", err)
	}

	go func() {
		var conn net.Conn
		for {
			conn, err = listener.Accept()
			jsonrpc.ServeConn(conn)
		}
	}()

	time.Sleep(1 * time.Second)

	if conn, err = net.Dial("tcp", ":9000"); err != nil {
		t.Fatalf("dial failed, err: %v\r\n", err)
	}

	rpcClient = jsonrpc.NewClient(conn)

	item = engine.Item{
		ID:      "test0001",
		URL:     "testURL",
		Type:    "test",
		Payload: model.ZhenaiPayLoad{},
	}

	if err = rpcClient.Call("Service.Save", item, &result); err != nil {
		log.Printf("call failed, err: %v\r\n", err)
	} else {
		log.Printf("result: %s", result)
	}
}

// go serveRPC(9000, "test")

// time.Sleep(1 * time.Second)

// conn, err := net.Dial("tcp", ":9000")
// if err != nil {
// 	t.Fatalf("dial failed, err: %v\r\n", err)
// }
// client := rpc.NewClient(conn)

// var result string

// item := engine.Item{
// 	ID:   "test0001",
// 	URL:  "testURL",
// 	Type: "test",
// 	Payload: model.ZhenaiPayLoad{Name: "testName", Gender: "testGender", Age: "testAge",
// 		Height: "testHeight", Weight: "testWeight", Income: "testIncome", Marriage: "testMarriage",
// 		Education: "testEducation"},
// }
// if err = client.Call("Service.Save", item, &result); err != nil {
// 	t.Fatalf("call service failed, err: %v\r\n", err)
// } else {
// 	t.Logf("result: %s", result)
// }
