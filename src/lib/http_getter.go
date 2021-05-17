package lib

import "net/http"

// HTTPGetter for mocking the HTTP client in tests.
type HTTPGetter interface {
	Get(url string) (resp *http.Response, err error)
}
