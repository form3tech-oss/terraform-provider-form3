package form3

import (
	"github.com/ewilde/go-form3"
	"github.com/ewilde/go-form3/client"
	"log"
)

// Config contains form3 provider settings
type Config struct {
	ClientId       string
	ClientSecret   string
	ApiHost        string
	OrganisationId string
}

func (c *Config) Client() (*form3.AuthenticatedClient, error) {

	config := client.DefaultTransportConfig()
	config.WithHost(c.ApiHost)

	client := form3.NewAuthenticatedClient(config)

	log.Printf("[INFO] form3 client configured for server %s", c.ApiHost)

	client.Authenticate(c.ClientId, c.ClientSecret)

	return client, nil
}
