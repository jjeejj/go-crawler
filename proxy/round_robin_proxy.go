package proxy

import (
	"errors"
	"net/http"
	"net/url"
	"sync/atomic"
)

type ProxyFunc func(*http.Request) (*url.URL, error)
type roundRobinProxy struct {
	proxyUrls []*url.URL
	index     uint32
}

func RoundRobinProxy(ProxyUrls ...string) (ProxyFunc, error) {
	if len(ProxyUrls) == 0 {
		return nil, errors.New("Proxy url list is empty")
	}
	urls := make([]*url.URL, len(ProxyUrls))
	for i, u := range ProxyUrls {
		u, err := url.Parse(u)
		if err != nil {
			return nil, err
		}
		urls[i] = u
	}
	return (&roundRobinProxy{urls, 0}).GetProxy, nil
}

func (rr *roundRobinProxy) GetProxy(pr *http.Request) (*url.URL, error) {
	index := atomic.AddUint32(&rr.index, 1) - 1
	u := rr.proxyUrls[index%uint32(len(rr.proxyUrls))]
	return u, nil
}
