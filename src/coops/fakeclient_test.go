package coops_test

import (
	"net/http"

	"github.com/webercoder/gocean/src/lib"
)

type FakeCoopsClient struct {
	Err      error
	JsonData string
}

func (fsc *FakeCoopsClient) Get(url string) (resp *http.Response, err error) {
	if fsc.Err != nil {
		return nil, fsc.Err
	}

	return &http.Response{
		Body: lib.NewStringReadCloser(fsc.JsonData),
	}, nil
}
