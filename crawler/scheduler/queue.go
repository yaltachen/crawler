package scheduler

import "crawler/crawler/engine"

// QueueScheduler 队列调度器
type QueueScheduler struct {
	workerChan  chan chan *engine.Request
	requestChan chan *engine.Request
}

// Submit submit request
func (s *QueueScheduler) Submit(req *engine.Request) {
	s.requestChan <- req
}

// Run run queue scheduler
func (s *QueueScheduler) Run() {

	var (
		workerQ  []chan *engine.Request
		requestQ []*engine.Request
	)

	s.workerChan = make(chan chan *engine.Request)
	s.requestChan = make(chan *engine.Request)

	workerQ = make([]chan *engine.Request, 0)
	requestQ = make([]*engine.Request, 0)
	go func() {
		for {
			var activeWorker chan *engine.Request
			var activeRequest *engine.Request
			if len(workerQ) > 0 && len(requestQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case worker := <-s.workerChan:
				workerQ = append(workerQ, worker)
			case request := <-s.requestChan:
				requestQ = append(requestQ, request)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}

// WorkChan return available worker
func (s *QueueScheduler) WorkChan() (req chan *engine.Request) {
	return make(chan *engine.Request)
}

// WorkerReady notify worker is ready
func (s *QueueScheduler) WorkerReady(worker chan *engine.Request) {
	s.workerChan <- worker
}
