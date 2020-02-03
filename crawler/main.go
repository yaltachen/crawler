package main

import (
	"crawler/crawler/engine"
	"crawler/crawler/parser/zhenai"
	"crawler/crawler/scheduler"
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

	engine.ConcurrentEngine{
		WorkerCount: 10,
		// Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler: &scheduler.QueueScheduler{},
	}.Run(seeds...)
}
