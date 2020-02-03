package engine

import (
	"crawler/crawler/model"
	"log"
)

// ConcurrentEngine 并发版
type ConcurrentEngine struct {
	WorkerCount int
	Scheduler   Scheduler
}

// Run 启动engine
func (e ConcurrentEngine) Run(seeds ...*Request) {
	var (
		count = 0
	)
	out := make(chan *ParserResult)

	// run scheduler
	e.Scheduler.Run()

	// create worker
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}

	for _, seed := range seeds {
		e.Scheduler.Submit(seed)
	}
	for {
		result := <-out
		for _, item := range result.Items {
			_, ok := item.(model.Profile)
			if !ok {
				continue
			}
			count++
			log.Printf("got item%d: %v", count, item)
		}

		// 循环等待是因为要等所有request全部submit后，
		// 外面的for循环才会继续，此时才能从out获得result
		for _, req := range result.Requests {
			e.Scheduler.Submit(req)
		}
	}
}
