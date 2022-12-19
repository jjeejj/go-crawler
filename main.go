package main

import (
	"fmt"
	"time"

	"github.com/jjeejj/go-crawler/collect"
	"github.com/jjeejj/go-crawler/engine"
	"github.com/jjeejj/go-crawler/log"
	"github.com/jjeejj/go-crawler/parse/douban"
	"github.com/jjeejj/go-crawler/proxy"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 初始化日志组建
	logPlugin := log.NewStdoutPlugin(zapcore.DebugLevel)
	logger := log.NewLogger(logPlugin)
	logger.Info("log init end")
	proxyURLs := []string{"http://127.0.0.1:7890", "http://127.0.0.1:7890"}
	p, err := proxy.RoundRobinProxy(proxyURLs...)
	if err != nil {
		logger.Error("RoundRobinProxy failed: ", zap.Error(err))
		panic(err)
	}

	var seeds []*collect.Request
	for i := 0; i <= 25; i += 25 {
		url := fmt.Sprintf("https://www.douban.com/group/szsh/discussion?start=%d", i)
		seeds = append(seeds, &collect.Request{
			Url:       url,
			ParseFunc: douban.ParseUrl,
		})
	}
	var f collect.Fetcher = &collect.BrowserFetch{
		Timeout: time.Second * 3,
		Proxy:   p,
	}
	se := engine.NewScheduleEngine(
		engine.WithSeeds(seeds),
		engine.WithFetcher(f),
		engine.WithWorkerCount(5),
		engine.WithLogger(logger),
	)
	se.Run()
	// 广度优先
	// for len(workList) > 0 {
	// 	work := workList[0]
	// 	workList = workList[1:]
	// 	body, err := f.Get(work)
	// 	time.Sleep(time.Second)
	// 	if err != nil {
	// 		logger.Error("read content failed", zap.Error(err))
	// 		continue
	// 	}
	// 	res := work.ParseFunc(body)
	// 	logger.Info("res:", zap.Any("pase res", res))
	// 	for _, item := range res.Items {
	// 		logger.Info("result", zap.String("get urk:", item.(string)))
	// 	}
	// 	workList = append(workList, res.Requests...)
	// }

}
