package scheduler

import "crawler/crawler/engine"

// SimpleScheduler 简单的调度器
type SimpleScheduler struct {
	reqChan chan *engine.Request
}

// Submit 提交worker
func (s SimpleScheduler) Submit(request *engine.Request) {
	go func() { s.reqChan <- request }()
	// s.reqChan <- request
}

// Run 运行
func (s *SimpleScheduler) Run() {
	s.reqChan = make(chan *engine.Request)
}

// WorkChan return available req chan
func (s SimpleScheduler) WorkChan() (req chan *engine.Request) {
	return s.reqChan
}

// WorkerReady SimpleScheduler 不需要实现
func (s SimpleScheduler) WorkerReady(chan *engine.Request) {
	// do nothing
}
