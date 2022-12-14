package collect

type Fetcher interface {
	Get(url *Request) ([]byte, error)
}
