package form3

import (
	"fmt"
	"github.com/ewilde/go-form3"
	"github.com/ewilde/go-form3/client/accounts"
	"github.com/ewilde/go-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3AccountConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountConfigurationCreate,
		Read:   resourceAccountConfigurationRead,
		Delete: resourceAccountConfigurationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"account_configuration_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_generation_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAccountConfigurationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	id := d.Get("account_configuration_id").(string)
	log.Printf("[INFO] Creating account configuration with id: %s", id)

	configuration, err := createAccountConfigurationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create account configuration: %s", err)
	}

	createdConfiguration, err := client.AccountClient.Accounts.PostAccountconfigurations(accounts.NewPostAccountconfigurationsParams().
		WithAccountConfigurationCreationRequest(&models.AccountConfigurationCreation{
			Data: configuration,
		}))

	if err != nil {
		return fmt.Errorf("failed to create account configuration: %s", err)
	}

	d.SetId(createdConfiguration.Payload.Data.ID.String())
	log.Printf("[INFO] configuration key: %s", d.Id())

	return resourceAccountConfigurationRead(d, meta)
}

func resourceAccountConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	key := d.Id()
	configurationId, _ := GetUUIDOK(d, "account_configuration_id")
	log.Printf("[INFO] Reading account configuration for id: %s", key)

	if configurationId == "" {
		configurationId = strfmt.UUID(key)
		log.Printf("[INFO] Importing account configuration id: %s ", key)
	} else {
		log.Printf("[INFO] Reading account configuration for id: %s", configurationId)
	}

	configuration, err := client.AccountClient.Accounts.GetAccountconfigurationsID(accounts.NewGetAccountconfigurationsIDParams().
		WithID(configurationId))

	if err != nil {
		apiError := err.(*runtime.APIError)
		if apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find account configuration: %s", err)
	}

	d.Set("organisation_id", configuration.Payload.Data.OrganisationID.String())
	d.Set("account_configuration_id", configuration.Payload.Data.ID.String())
	d.Set("account_generation_enabled", configuration.Payload.Data.Attributes.AccountGenerationEnabled)
	return nil
}

func resourceAccountConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	configurationFromResource, err := createAccountConfigurationFromResourceDataWithVersion(d, client)
	if err != nil {
		return fmt.Errorf("error deleting account configuration: %s", err)
	}

	log.Printf("[INFO] Deleting account configuration for id: %s", configurationFromResource.ID)

	_, err = client.AccountClient.Accounts.DeleteAccountconfigurationsID(accounts.NewDeleteAccountconfigurationsIDParams().
		WithID(configurationFromResource.ID).
		WithVersion(*configurationFromResource.Version))

	if err != nil {
		return fmt.Errorf("error deleting account configuration: %s", err)
	}

	return nil
}

func resourceAccountConfigurationUpdate() {

}

func createAccountConfigurationFromResourceDataWithVersion(d *schema.ResourceData, client *form3.AuthenticatedClient) (*models.AccountConfiguration, error) {
	configuration, err := createAccountConfigurationFromResourceData(d)
	version, err := getAccountConfigurationVersion(client, configuration.ID)
	if err != nil {
		return nil, err
	}

	configuration.Version = &version

	return configuration, nil
}

func createAccountConfigurationFromResourceData(d *schema.ResourceData) (*models.AccountConfiguration, error) {

	configuration := models.AccountConfiguration{Attributes: &models.AccountConfigurationAttributes{}}
	configuration.Type = "account_configurations"
	if attr, ok := GetUUIDOK(d, "account_configuration_id"); ok {
		configuration.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		configuration.OrganisationID = attr
	}

	if attr, ok := d.GetOk("account_generation_enabled"); ok {
		configuration.Attributes.AccountGenerationEnabled = attr.(bool)
	}
	return &configuration, nil
}

func getAccountConfigurationVersion(client *form3.AuthenticatedClient, configurationId strfmt.UUID) (int64, error) {
	configuration, err := client.AccountClient.Accounts.GetAccountconfigurationsID(accounts.NewGetAccountconfigurationsIDParams().WithID(configurationId))
	if err != nil {
		if err != nil {
			return -1, fmt.Errorf("error reading account configuration: %s", err)
		}
	}

	return *configuration.Payload.Data.Version, nil
}
