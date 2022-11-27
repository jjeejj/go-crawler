package collect

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type BrowserFetch struct{}

func (browser *BrowserFetch) Get(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Errorf("BrowserFetch new request err: %v", err)
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15")
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("BrowserFetch client do  request err: %v", err)
		return nil, err
	}
	return Transfer2Utf8Encoding(resp.Body)
}
