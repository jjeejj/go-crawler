package main

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	url := "https://www.thepaper.cn"
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("fetch url error: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("Error status code: %v: ", resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("read content failed: %v", err)
		return
	}
	log.Infof("body: %v", string(body))
}
