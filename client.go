package hydraclient

import (
	"net/http"
	"net/url"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/gojek/heimdall/hystrix"
	"github.com/ory/hydra/sdk/go/hydra/client/admin"
)

// Option represents the hydra client options.
type Option func(c *Client)

// Client is the hydra client implementation.
type Client struct {
	admin.Client

	url    url.URL
	client *hystrix.Client
}

// New returns a new instance of hydra Client.
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
			Transport: NewHystrixTransport(c.client),
		},
	)

	c.Client = *admin.New(ht, nil)

	return &c
}

// WithHystrixClient creates an option that will define the `hystrix.Client`
// when creating a new `Client`.
func WithHystrixClient(client *hystrix.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}

// WithURL creates an option that defines a host name for the Hydra server.
func WithURL(u *url.URL) Option {
	return func(c *Client) {
		c.url = *u
	}
}
