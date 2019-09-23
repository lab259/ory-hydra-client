package hydraclient

import (
	"net/url"

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
	var client Client

	for _, opt := range opts {
		opt(&client)
	}

	if client.client == nil {
		client.client = hystrix.NewClient()
	}

	client.OryHydra = *hydra.New(nil, nil)

	return &client
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
