package main

import (
	"crawler/crawler/engine"
	"crawler/crawler/parser/zhenai"
	"crawler/crawler/scheduler"
	"crawler/crawler_distribute/config"
	persist "crawler/crawler_distribute/persist/client"
	worker "crawler/crawler_distribute/worker/client"
	"fmt"
)

func main() {
	var (
		seeds     []*engine.Request
		itemChan  chan engine.Item
		processor engine.Processor
		e         engine.ConcurrentEngine
		err       error
	)
	if itemChan, err = persist.ItemSaver(
		fmt.Sprintf(":%d", config.ELASTICSEARCHPORT)); err != nil {
		panic(err)
	}

	if processor, err = worker.CreateProcessor(); err != nil {
		panic(err)
	}
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
