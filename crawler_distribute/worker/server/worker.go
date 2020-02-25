package main

import (
	"crawler/crawler_distribute/rpcsupport"
	"crawler/crawler_distribute/worker"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServeRPC(":9001", worker.CrawlerService{}))
}
