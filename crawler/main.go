package main

import (
	"crawler/crawler/engine"
	"crawler/crawler/parser/zhenai"
	"crawler/crawler/persist"
	"crawler/crawler/scheduler"
	"log"
)

func main() {
	var (
		seeds []*engine.Request
	)

	seeds = append(seeds, &engine.Request{
		URL:    "http://www.zhenai.com/zhenghun/",
		Parser: zhenai.CityListParser,
	})

	// engine.SimpleEngine{}.Run(seeds...)

	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		log.Printf("start item saver failed, err: %v\r\n", err)
		return
	}

	engine.ConcurrentEngine{
		WorkerCount: 10,
		// Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler: &scheduler.QueueScheduler{},
		ItemChan:  itemChan,
	}.Run(seeds...)
}
