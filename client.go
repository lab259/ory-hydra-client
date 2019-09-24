package hydraclient

import (
	"net/http"
	"net/url"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/gojek/heimdall/hystrix"
	hydra "github.com/ory/hydra/sdk/go/hydra/client"
)

type Option func(c *Client)

type Client struct {
	hydra.OryHydra

	url    url.URL
	client *hystrix.Client
}

func New(opts ...Option) *Client {
	var c Client

	for _, opt := range opts {
		opt(&c)
	}

	if c.client == nil {
		c.client = hystrix.NewClient()
	}

	ht := httptransport.NewWithClient(
		c.url.Host,
		c.url.Path,
		[]string{c.url.Scheme},
		&http.Client{
			Transport: &hystrixTransport{c.client},
		},
	)

	c.OryHydra = *hydra.New(ht, nil)

	return &c
}

// WithHystrixClient creates an option that will define the `hystrix.Client`
// when creating a new `Client`.
func WithHystrixClient(client *hystrix.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}

// WithURL creates an option that defines a host name for the Keto server.
func WithURL(u *url.URL) Option {
	return func(c *Client) {
		c.url = *u
	}
}
