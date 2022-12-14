package douban

import (
	"regexp"

	"github.com/jjeejj/go-crawler/collect"
	log "github.com/sirupsen/logrus"
)

var cityListRegex *regexp.Regexp = regexp.MustCompile("https://www.douban.com/group/topic/[0-9]+/")

func ParseUrl(content []byte) collect.ParseResult {
	matches := cityListRegex.FindAllSubmatch(content, -1)
	result := collect.ParseResult{}
	log.Infof("matches: %v", matches)
	// for _, m := range matches {

	// }
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
