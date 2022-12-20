package collect

import (
	"errors"
	"time"

	"go.uber.org/zap"
)

// Task 爬虫的具体任务
// 一个任务下面多个请求
type Task struct {
	Url      string        `json:"url"`
	Cookie   string        `json:"cookie"`
	WaitTime time.Duration // 每个请求任务之间等待的时间
	MaxDepth int           `json:"max_depth"` // 当前任务的最大深度
	Logger   *zap.Logger
	RootReq  *Request // 任务的第一个请求
}

// Request 网站的请求信息
type Request struct {
	Url       string `json:"url"`
	Task      *Task  `json:"task"`
	Depth     int    `json:"depth"` // 当前任务的当前深度
	ParseFunc func([]byte, *Request) ParseResult
}

// ParseResult 请求返回的结果
type ParseResult struct {
	Requests []*Request
	Items    []interface{}
}

func (r *Request) Check() error {
	if r.Depth > r.Task.MaxDepth {
		return errors.New("max depth exceeded limit for request")
	}
	return nil
}
