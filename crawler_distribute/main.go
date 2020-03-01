package main

import (
	"crawler/crawler/engine"
	"crawler/crawler/parser/zhenai"
	"crawler/crawler/scheduler"
	"crawler/crawler_distribute/config"
	persist "crawler/crawler_distribute/persist/client"
	"crawler/crawler_distribute/rpcsupport"
	worker "crawler/crawler_distribute/worker/client"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"strings"
)

var hosts string

func init() {
	flag.StringVar(&hosts, "hosts", "", "specify hosts")
	flag.Parse()
}

func main() {
	var (
		seeds     []*engine.Request
		itemChan  chan engine.Item
		processor engine.Processor
		e         engine.ConcurrentEngine
		pool      chan *rpc.Client
		err       error
	)

	if hosts == "" {
		panic("must specify hosts")
	}
	hostsArr := strings.Split(hosts, ";")

	if itemChan, err = persist.ItemSaver(
		fmt.Sprintf(":%d", config.ELASTICSEARCHPORT)); err != nil {
		panic(err)
	}
	if pool, err = createClientPool(hostsArr); err != nil {
		panic("create client pool failed.")
	}
	processor = worker.CreateProcessor(pool)
	e = engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueueScheduler{},
		WorkerCount:      10,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	seeds = append(seeds, &engine.Request{
		URL: "http://www.zhenai.com/zhenghun/",
		// Parser: zhenai.CityListParser,
		Parser: zhenai.NewCityListParser(),
	})
	e.Run(seeds...)
}

func createClientPool(hosts []string) (chan *rpc.Client, error) {
	var (
		client     *rpc.Client
		clients    []*rpc.Client
		clientChan chan *rpc.Client
		err        error
	)
	clientChan = make(chan *rpc.Client)

	for _, host := range hosts {
		if client, err = rpcsupport.NewSupport(host); err != nil {
			log.Printf("connect %s failed, err: %v\r\n", host, err)
			continue
		} else {
			log.Printf("connected to %s\r\n", host)
			clients = append(clients, client)
		}
	}

	if len(clients) == 0 {
		return nil, errors.New(
			"no connections available")
	}

	go func() {
		for {
			for _, c := range clients {
				clientChan <- c
			}
		}
	}()

	return clientChan, nil
}
