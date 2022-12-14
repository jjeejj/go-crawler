package collect

// Request 网站的请求信息
type Request struct {
	Url       string `json:"url"`
	Cookie    string `json:"cookie"`
	ParseFunc func([]byte) ParseResult
}

// ParseResult 请求返回的结果
type ParseResult struct {
	Requests []*Request
	Items    []interface{}
}
