package hydraclient

import (
	"net/http"

	"github.com/gojek/heimdall/hystrix"
)

// NewHystrixTransport returns a http.RoundTripper implementation that uses
// a hydrix.Client.
func NewHystrixTransport(client *hystrix.Client) http.RoundTripper {
	return &hystrixTransport{client: client}
}

type hystrixTransport struct {
	client *hystrix.Client
}

func (ht *hystrixTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return ht.client.Do(req)
}
