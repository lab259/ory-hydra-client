package main

import (
	"log"
	"net/url"
	"time"

	"github.com/gojek/heimdall/hystrix"
	hydraclient "github.com/lab259/ory-hydra-client"
	"github.com/ory/hydra/sdk/go/hydra/client/admin"
	"github.com/ory/hydra/sdk/go/hydra/models"
)

func main() {
	url, _ := url.Parse("https://hydra.localhost:4445")

	hc := hystrix.NewClient(
		hystrix.WithHTTPTimeout(10*time.Millisecond),
		hystrix.WithCommandName("ory_hydra"),
		hystrix.WithHystrixTimeout(1000),
		hystrix.WithMaxConcurrentRequests(30),
		hystrix.WithErrorPercentThreshold(20),
	)

	client := hydraclient.New(
		hydraclient.WithHystrixClient(hc),
		hydraclient.WithURL(url),
	)

	response, err := client.Admin.CreateOAuth2Client(admin.NewCreateOAuth2ClientParams().WithBody(&models.Client{
		ClientID: "foo-bar",
	}))
	if err != nil {
		log.Fatalf("unable to create oauth client: %v", err)
	}

	log.Println("Client created:")
	log.Printf("\t ID:     %s", response.Payload.ClientID)
	log.Printf("\t Secret: %s", response.Payload.Secret)
}
