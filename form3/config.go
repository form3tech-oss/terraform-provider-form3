package form3

import (
	"log"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client"
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

	err := client.Authenticate(c.ClientId, c.ClientSecret)

	return client, err
}
