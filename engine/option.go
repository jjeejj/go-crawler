package engine

import (
	"github.com/jjeejj/go-crawler/collect"
	"go.uber.org/zap"
)

type options struct {
	WorkerCount int
	Logger      *zap.Logger
	Seeds       []*collect.Task
	Fetcher     collect.Fetcher
	scheduler   Scheduler
}

type Option func(opt *options)

var defaultOptions = options{
	Logger: zap.NewNop(),
}

func WithLogger(logger *zap.Logger) Option {
	return func(opt *options) {
		opt.Logger = logger
	}
}

func WithFetcher(fetcher collect.Fetcher) Option {
	return func(opt *options) {
		opt.Fetcher = fetcher
	}
}

func WithWorkerCount(count int) Option {
	return func(opt *options) {
		opt.WorkerCount = count
	}
}

func WithScheduler(scheduler Scheduler) Option {
	return func(opt *options) {
		opt.scheduler = scheduler
	}
}

func WithSeeds(seeds []*collect.Task) Option {
	return func(opt *options) {
		opt.Seeds = seeds
	}
}
