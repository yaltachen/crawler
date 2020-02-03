package engine

import (
	"crawler/crawler/fetcher"
	"log"
	"time"
)

// doWork 完成fetch + parser
func doWork(req *Request) (result *ParserResult, err error) {
	var (
		content    []byte
		speedLimit <-chan time.Time
	)

	speedLimit = time.Tick(1000 * time.Millisecond)
	<-speedLimit

	if content, err = fetcher.Fetch(req.URL); err != nil {
		log.Printf("fetch %s failed, err: %v\r\n", req.URL, err)
		return nil, err
	}

	if result, err = req.Parser(content); err != nil {
		log.Printf("parser %s failed, err: %v\r\n", req.URL, err)
		return nil, err
	}
	return result, nil
}

// createWorker 创建worker
func createWorker(in chan *Request, out chan *ParserResult, notifier ReadyNotifier) {
	go func() {
		for {
			notifier.WorkerReady(in)
			req := <-in
			result, err := doWork(req)
			if err != nil {
				log.Printf("do work failed. err: %v\r\n", err)
				continue
			}
			out <- result
		}
	}()
}