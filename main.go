package main

import (
	"github.com/jjeejj/go-crawler/collect"
	"github.com/jjeejj/go-crawler/proxy"
	log "github.com/sirupsen/logrus"
)

func main() {
	url := "https://www.google.com/?client=safari"
	proxyURLs := []string{"http://127.0.0.1:7890", "http://127.0.0.1:7890"}
	p, err := proxy.RoundRobinProxy(proxyURLs...)
	var f collect.Fetcher = &collect.BrowserFetch{
		Proxy: p,
	}
	body, err := f.Get(url)
	if err != nil {
		log.Errorf("read content failed err: %v", err)
		return
	}
	log.Infof("body: %v", string(body))
	// numLinks := strings.Count(string(body), "<a")
	// log.Infof("home page has %d links", numLinks)
	// doc, err := htmlquery.Parse(bytes.NewReader(body))
	// log.Infof("%v:", doc.Data)
	// if err != nil {
	// 	log.Errorf("htmlquery.Parse failed err: %v", err)
	// 	return
	// }
	// nodes := htmlquery.Find(doc, `//li[@class="index_wechartcontent__yM1tu"]/span`)
	// for _, node := range nodes {
	// 	log.Info(node.FirstChild.Data)
	// }

}
