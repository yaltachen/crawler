package main

import (
	"crawler/crawler/model"
	"crawler/crawler_distribute/persist"
	"crawler/crawler_distribute/rpcsupport"
	"encoding/gob"
	"flag"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 0, "specify service port")
	flag.Parse()
}

func main() {
	var err error
	if port == 0 {
		log.Fatalf("must specify service port")
	}
	if err = serveRPC(port, "dating_profile"); err != nil {
		log.Fatalf("serve persist service failed, err: %v\r\n", err)
	}
}

func serveRPC(port int, index string) error {
	var (
		client *elastic.Client
		err    error
	)
	gob.Register(model.ZhenaiPayLoad{})
	if client, err =
		elastic.NewClient(elastic.SetSniff(false),
			elastic.SetURL("http://192.168.94.30:9200")); err != nil {

	}
	return rpcsupport.ServeRPC(fmt.Sprintf(":%d", port), &persist.Service{
		Client: client,
		Index:  index,
	})
}
