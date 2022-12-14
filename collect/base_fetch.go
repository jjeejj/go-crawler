package collect

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type BaseFetch struct {
}

func (base *BaseFetch) Get(urlReq *Request) ([]byte, error) {
	resp, err := http.Get(urlReq.Url)
	if err != nil {
		log.Errorf("BaseFetch Get url error: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Errorf("Error status code: %v", resp.StatusCode)
	}
	return Transfer2Utf8Encoding(resp.Body)
}
