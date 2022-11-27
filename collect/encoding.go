package collect

import (
	"bufio"
	"io"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Errorf("r.Peek error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func Transfer2Utf8Encoding(rd io.Reader) ([]byte, error) {
	bodyReader := bufio.NewReader(rd)
	e := DetermineEncoding(bodyReader)
	ut8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(ut8Reader)
}
