package engine

import (
	"sync"

	"github.com/jjeejj/go-crawler/collect"
	"go.uber.org/zap"
)

// Crawler 爬虫引擎实例
type Crawler struct {
	out            chan collect.ParseResult // 结果
	options                                 // 引擎实例的配置项
	VisitedMap     map[string]bool          // 已经爬取过的地址map
	VisitedMapLock sync.Mutex
}

// Scheduler 调度的接口定义
type Scheduler interface {
	Schedule()                // 调度方法
	Push(...*collect.Request) // 添加爬虫的任务请求
	Pull() *collect.Request   // 取出爬虫的任务请求
}

func NewEngine(opts ...Option) *Crawler {
	options := defaultOptions
	for _, opt := range opts {
		opt(&options)
	}
	engine := &Crawler{}
	engine.options = options
	engine.out = make(chan collect.ParseResult)
	engine.VisitedMap = make(map[string]bool, 100)
	return engine
}

func (crawler *Crawler) Run() {
	go crawler.Schedule()
	// 创建工作 worker
	for i := 0; i < crawler.WorkerCount; i++ {
		go crawler.CreateWorker()
	}
	crawler.HandleResult()
}

func (crawler *Crawler) Schedule() {
	var reqs []*collect.Request
	for _, seed := range crawler.Seeds {
		seed.RootReq.Task = seed
		reqs = append(reqs, seed.RootReq)
	}
	go crawler.scheduler.Schedule()
	go crawler.scheduler.Push(reqs...)
}

func (crawler *Crawler) CreateWorker() {
	for {
		r := crawler.scheduler.Pull()
		// 发送请求之前判断是否超过限制
		if err := r.Check(); err != nil {
			crawler.Logger.Error("request check failed", zap.Error(err))
			continue
		}
		// 判断之前没有访问过
		if crawler.HasVisited(r) {
			crawler.Logger.Debug("request has visit", zap.String("url:", r.Url))
			continue
		}
		body, err := crawler.Fetcher.Get(r)
		if err != nil {
			crawler.Logger.Error("can't fetch", zap.Error(err))
			continue
		}
		result := r.ParseFunc(body, r)
		if len(result.Requests) > 0 {
			crawler.scheduler.Push(result.Requests...)
		}
		crawler.out <- result

		crawler.StoreVisited(r)
	}
}

func (crawler *Crawler) HandleResult() {
	for {
		select {
		case result := <-crawler.out:
			for _, item := range result.Items {
				crawler.Logger.Sugar().Info("get request", item)
			}
		}
	}
}

// HasVisited 判断请求是否处理过
func (crawler *Crawler) HasVisited(req *collect.Request) bool {
	crawler.VisitedMapLock.Lock()
	defer crawler.VisitedMapLock.Unlock()
	unique := req.Unique()
	return crawler.VisitedMap[unique]
}

func (crawler *Crawler) StoreVisited(reqs ...*collect.Request) {
	crawler.VisitedMapLock.Lock()
	defer crawler.VisitedMapLock.Unlock()
	for _, req := range reqs {
		unique := req.Unique()
		crawler.VisitedMap[unique] = true
	}
}
