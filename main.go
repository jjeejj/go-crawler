package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/antchfx/htmlquery"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func main() {
	url := "https://www.thepaper.cn"
	body, err := Fetch(url)
	if err != nil {
		log.Errorf("read content failed err: %v", err)
		return
	}
	// log.Infof("body: %v", string(body))
	// numLinks := strings.Count(string(body), "<a")
	// log.Infof("home page has %d links", numLinks)
	doc, err := htmlquery.Parse(bytes.NewReader(body))
	// log.Infof("%v:", doc.Data)
	if err != nil {
		log.Errorf("htmlquery.Parse failed err: %v", err)
		return
	}
	nodes := htmlquery.Find(doc, `//li[@class="index_wechartcontent__yM1tu"]/span`)
	for _, node := range nodes {
		log.Info(node.FirstChild.Data)
	}

}

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("fetch url error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("Error status code: %v: ", resp.StatusCode)
		return nil, err
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := DetermineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Errorf("r.Peek error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
