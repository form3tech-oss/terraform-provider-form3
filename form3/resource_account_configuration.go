package form3

import (
	"bytes"
	"fmt"
	"github.com/form3tech-oss/go-form3"
	"github.com/form3tech-oss/go-form3/client/accounts"
	"github.com/form3tech-oss/go-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3AccountConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountConfigurationCreate,
		Read:   resourceAccountConfigurationRead,
		Delete: resourceAccountConfigurationDelete,
		Update: resourceAccountConfigurationUpdate,
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
			"account_generation_configuration": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"country": {
							Type:     schema.TypeString,
							Required: true,
						},
						"valid_account_ranges": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"minimum": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"maximum": {
										Type:     schema.TypeInt,
										Required: true,
									},
								},
							},
						},
					},
				},
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
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find account configuration: %s", err)
	}

	d.Set("organisation_id", configuration.Payload.Data.OrganisationID.String())
	d.Set("account_configuration_id", configuration.Payload.Data.ID.String())
	d.Set("account_generation_enabled", configuration.Payload.Data.Attributes.AccountGenerationEnabled)

	accountGenerationConfigurations :=
		make([]interface{}, 0, len(configuration.Payload.Data.Attributes.AccountGenerationConfiguration))

	for _, element := range configuration.Payload.Data.Attributes.AccountGenerationConfiguration {
		accountGenerationConfigurations = append(accountGenerationConfigurations, map[string]interface{}{
			"country":              element.Country,
			"valid_account_ranges": flattenValidAccountRanges(element.ValidAccountRanges),
		})
	}

	if err := d.Set("account_generation_configuration", accountGenerationConfigurations); err != nil {
		return err
	}

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

func resourceAccountConfigurationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)
	configuration, err := createAccountConfigurationFromResourceDataWithVersion(d, client)

	if err != nil {
		return fmt.Errorf("failed to update account configuration: %s", err)
	}

	log.Printf("[INFO] Updating account configuration with id: %s", configuration.ID)
	_, err = client.AccountClient.Accounts.PatchAccountconfigurationsID(accounts.NewPatchAccountconfigurationsIDParams().
		WithID(configuration.ID).
		WithConfigAmendRequest(&models.ConfigurationAmendment{
			Data: configuration,
		}))

	if err != nil {
		return fmt.Errorf("failed to update account configuration: %s", err)
	}

	return nil
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

	if attr, ok := d.GetOk("account_generation_configuration"); ok {
		accountConfigurationArray := attr.([]interface{})

		var accountConfigs []*models.AccountGenerationConfiguration

		for _, accountConfigElement := range accountConfigurationArray {
			country := accountConfigElement.(map[string]interface{})["country"].(string)

			validAccountRangesSet := accountConfigElement.(map[string]interface{})["valid_account_ranges"].(*schema.Set).List()

			var validAccountRanges []*models.AccountGenerationConfigurationValidAccountRangesItems

			for _, accountRangeElement := range validAccountRangesSet {
				validAccountRange := models.AccountGenerationConfigurationValidAccountRangesItems{
					Minimum: int64(accountRangeElement.(map[string]interface{})["minimum"].(int)),
					Maximum: int64(accountRangeElement.(map[string]interface{})["maximum"].(int)),
				}
				validAccountRanges = append(validAccountRanges, &validAccountRange)
			}

			accountConfig := models.AccountGenerationConfiguration{
				Country:            country,
				ValidAccountRanges: validAccountRanges,
			}

			accountConfigs = append(accountConfigs, &accountConfig)
		}

		configuration.Attributes.AccountGenerationConfiguration = accountConfigs
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

func flattenValidAccountRanges(validAccountRanges []*models.AccountGenerationConfigurationValidAccountRangesItems) *schema.Set {
	validAccountRangesSet := schema.NewSet(validAccountRangeHash, []interface{}{})

	if validAccountRanges == nil {
		return validAccountRangesSet
	}

	for _, value := range validAccountRanges {
		validAccountRangesSet.Add(flattenValidAccountRange(value))
	}
	return validAccountRangesSet
}

func validAccountRangeHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]int64)
	buf.WriteString(fmt.Sprintf("%d", m["minimum"]))
	buf.WriteString(fmt.Sprintf("%d", m["maximum"]))
	return hashcode.String(buf.String())
}

func flattenValidAccountRange(value *models.AccountGenerationConfigurationValidAccountRangesItems) map[string]int64 {
	m := map[string]int64{}
	m["minimum"] = value.Minimum
	m["maximum"] = value.Maximum

	return m
}
