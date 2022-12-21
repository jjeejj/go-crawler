package engine

import (
	"github.com/jjeejj/go-crawler/collect"
	"go.uber.org/zap"
)

// DefaultSchedule 默认的调度实现
type DefaultSchedule struct {
	requestCh chan *collect.Request
	workerCh  chan *collect.Request
	Logger    *zap.Logger
	reqQueue  []*collect.Request // 待处理的请求列表
}

func NewDefaultSchedule() *DefaultSchedule {
	s := &DefaultSchedule{
		requestCh: make(chan *collect.Request),
		workerCh:  make(chan *collect.Request),
	}
	return s
}

func (s *DefaultSchedule) Schedule() {
	var req *collect.Request
	var ch chan *collect.Request
	go func() {
		for {
			if len(s.reqQueue) > 0 {
				req = s.reqQueue[0]
				s.reqQueue = s.reqQueue[1:]
				ch = s.workerCh
			}
			select {
			case r := <-s.requestCh:
				s.reqQueue = append(s.reqQueue, r)
			case ch <- req:
			}
		}
	}()
}

func (s *DefaultSchedule) Push(reqs ...*collect.Request) {
	for _, req := range reqs {
		s.requestCh <- req
	}
}

func (s *DefaultSchedule) Pull() *collect.Request {
	return <-s.workerCh
}
