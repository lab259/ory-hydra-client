package hydraclient

import (
	"net/http"

	"github.com/gojek/heimdall/hystrix"
)

type hystrixTransport struct {
	client *hystrix.Client
}

func (ht *hystrixTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return ht.client.Do(req)
}
