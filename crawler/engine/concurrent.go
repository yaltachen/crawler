package engine

import (
	"crawler/crawler/duplicate"
)

// ConcurrentEngine 并发版
type ConcurrentEngine struct {
	WorkerCount      int
	Scheduler        Scheduler
	ItemChan         chan Item
	RequestProcessor Processor
}

// Processor 处理任务
type Processor func(*Request) (*ParserResult, error)

// Run 启动engine
func (e ConcurrentEngine) Run(seeds ...*Request) {

	out := make(chan *ParserResult)

	// run scheduler
	e.Scheduler.Run()

	// create worker
	for i := 0; i < e.WorkerCount; i++ {
		// createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
		e.createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}

	for _, seed := range seeds {
		e.Scheduler.Submit(seed)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			// TODO: save to elasticsearch
			go func(item Item) {
				e.ItemChan <- item
			}(item)
		}

		// 循环等待是因为要等所有request全部submit后，
		// 外面的for循环才会继续，此时才能从out获得result
		for _, req := range result.Requests {
			if duplicate.IsDuplicate(req.URL) {
				// 重复的url
				// log.Printf("dulicate url: %s\r\n", req.URL)
				continue
			}
			e.Scheduler.Submit(req)
		}
	}
}
