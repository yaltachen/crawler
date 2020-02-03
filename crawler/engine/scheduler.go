package engine

// Scheduler 调度器
type Scheduler interface {
	Submit(req *Request)
	Run()
	WorkChan() (req chan *Request)
	ReadyNotifier
}

// ReadyNotifier tell Scheduler worker is ready
type ReadyNotifier interface {
	WorkerReady(chan *Request)
}
