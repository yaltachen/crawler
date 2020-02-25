package engine

import (
	"crawler/crawler/fetcher"
	"log"
	"time"
)

// DoWork 完成fetch + parser
func DoWork(req *Request) (result *ParserResult, err error) {
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

	if result, err = req.Parser.Parse(content); err != nil {
		log.Printf("parser %s failed, err: %v\r\n", req.URL, err)
		return nil, err
	}
	return result, nil
}

// createWorker 创建worker
func (e *ConcurrentEngine) createWorker(in chan *Request, out chan *ParserResult, notifier ReadyNotifier) {
	go func() {
		for {
			notifier.WorkerReady(in)
			req := <-in
			// TODO: call rpc
			// result, err := DoWork(req)
			result, err := e.RequestProcessor(req)
			if err != nil {
				log.Printf("Do work failed. err: %v\r\n", err)
				continue
			}
			out <- result
		}
	}()
}
