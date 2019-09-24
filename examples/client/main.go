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
	url, _ := url.Parse("http://hydra.localhost:4445")

	hc := hystrix.NewClient(
		hystrix.WithHTTPTimeout(10*time.Second),
		hystrix.WithCommandName("ory_hydra"),
		hystrix.WithHystrixTimeout(10*time.Second),
		hystrix.WithMaxConcurrentRequests(30),
		hystrix.WithErrorPercentThreshold(20),
	)

	client := hydraclient.New(
		hydraclient.WithHystrixClient(hc),
		hydraclient.WithURL(url),
	)

	_, err := client.IsInstanceAlive(admin.NewIsInstanceAliveParams())
	if err != nil {
		log.Fatalf("instance not ready: %s", err)
	}

	input := admin.NewCreateOAuth2ClientParams().
		WithBody(&models.Client{
			ClientID: "foo-bar-xepoj",
		})

	output, err := client.CreateOAuth2Client(input)
	if err != nil {
		log.Fatalf("unable to create oauth client: %s", err)
	}

	log.Println("Client created:")
	log.Printf("\t ID:     %s", output.Payload.ClientID)
	log.Printf("\t Secret: %s", output.Payload.Secret)
}
