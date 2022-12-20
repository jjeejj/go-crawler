package douban

import (
	"regexp"

	"github.com/jjeejj/go-crawler/collect"
	"go.uber.org/zap"
)

var cityListRegex *regexp.Regexp = regexp.MustCompile("https://www.douban.com/group/topic/[0-9]+/")

func ParseUrl(content []byte, req *collect.Request) collect.ParseResult {
	matches := cityListRegex.FindAllSubmatch(content, -1)
	result := collect.ParseResult{}
	// log.Infof("matches: %v", string(matches))
	for _, m := range matches {
		req.Task.Logger.Info("matches:", zap.ByteStrings("m", m))
		u := string(m[0])
		result.Requests = append(result.Requests, &collect.Request{
			Url:   u,
			Depth: req.Depth + 1,
			ParseFunc: func(c []byte, req *collect.Request) collect.ParseResult {
				return GetContentUrl(c, u)
			},
			Task: req.Task,
		})
	}
	return result
}

var contentRegex *regexp.Regexp = regexp.MustCompile(`<div class="topic-content">[\s\S]*?阳台[\s\S]*?<div`)

func GetContentUrl(content []byte, url string) collect.ParseResult {
	ok := contentRegex.Match(content)
	if !ok {
		return collect.ParseResult{
			Items: []interface{}{},
		}
	}
	return collect.ParseResult{
		Items: []interface{}{url},
	}
}
