package engine

import (
	"github.com/jjeejj/go-crawler/collect"
	"go.uber.org/zap"
)

type ScheduleEngine struct {
	requestCh chan *collect.Request
	workerCh  chan *collect.Request
	out       chan collect.ParseResult
	options
}

func NewScheduleEngine(opts ...Option) *ScheduleEngine {
	options := defaultOptions
	for _, opt := range opts {
		opt(&options)
	}
	se := &ScheduleEngine{}
	se.options = options
	return se
}

func (se *ScheduleEngine) Run() {
	requestCh := make(chan *collect.Request)
	workerCh := make(chan *collect.Request)
	out := make(chan collect.ParseResult)
	se.out = out
	se.requestCh = requestCh
	se.workerCh = workerCh
	go se.Schedule()
	// 创建工作 worker
	for i := 0; i < se.WorkerCount; i++ {
		go se.CreateWorker()
	}
	se.HandleResult()
}

func (se *ScheduleEngine) Schedule() {
	var reqQueue []*collect.Request
	for _, seed := range se.Seeds {
		seed.RootReq.Task = seed
		reqQueue = append(reqQueue, seed.RootReq)
	}
	go func() {
		for {
			var req *collect.Request
			var ch chan *collect.Request
			if len(reqQueue) > 0 {
				req = reqQueue[0]
				reqQueue = reqQueue[1:]
				ch = se.workerCh
			}
			select {
			case r := <-se.requestCh:
				reqQueue = append(reqQueue, r)
			case ch <- req:
			}
		}
	}()
}

func (se *ScheduleEngine) CreateWorker() {
	for {
		r := <-se.workerCh
		// 发送请求之前判断是否超过限制
		if err := r.Check(); err != nil {
			se.Logger.Error("request check failed", zap.Error(err))
			continue
		}
		body, err := se.Fetcher.Get(r)
		if err != nil {
			se.Logger.Error("can't fetch", zap.Error(err))
			continue
		}
		result := r.ParseFunc(body, r)
		se.out <- result

	}
}

func (se *ScheduleEngine) HandleResult() {
	for {
		select {
		case result := <-se.out:
			for _, req := range result.Requests {
				se.requestCh <- req
			}
			for _, item := range result.Items {
				se.Logger.Sugar().Info("get request", item)
			}
		}
	}
}
