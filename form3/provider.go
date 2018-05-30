package form3

import (
	"os"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: envDefaultFunc("FORM3_CLIENT_ID"),
				Description: "A form3 client id.",
			},
			"client_secret": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: envDefaultFunc("FORM3_CLIENT_SECRET"),
				Description: "A form3 client secret.",
			},
			"api_host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: envDefaultFuncWithDefault("FORM3_HOST", "api.form3.tech"),
				Description: "A form3 api host i.e. api.form3.tech.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"form3_account_configuration":   resourceForm3AccountConfiguration(),
			"form3_ace":                     resourceForm3Ace(),
			"form3_bank_id":                 resourceForm3BankID(),
			"form3_bic":                     resourceForm3Bic(),
			"form3_credential":              resourceForm3Credential(),
			"form3_organisation":            resourceForm3Organisation(),
			"form3_payport_association":     resourceForm3PayportAssociation(),
			"form3_role":                    resourceForm3Role(),
			"form3_starling_association":    resourceForm3StarlingAssociation(),
			"form3_subscription":            resourceForm3Subscription(),
			"form3_user":                    resourceForm3User(),
			"form3_limit":                   resourceForm3Limit(),
			"form3_bacs_association":        resourceForm3BacsAssociation(),
			"form3_sepainstant_association": resourceForm3SepaInstantAssociation(),
			"form3_payment_defaults":        resourceForm3PaymentDefaults(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func envDefaultFunc(key string) schema.SchemaDefaultFunc {
	return func() (interface{}, error) {
		if v := os.Getenv(key); v != "" {
			if v == "true" {
				return true, nil
			} else if v == "false" {
				return false, nil
			}
			return v, nil
		}
		return nil, nil
	}
}

func envDefaultFuncWithDefault(key string, defaultValue string) schema.SchemaDefaultFunc {
	return func() (interface{}, error) {
		if v := os.Getenv(key); v != "" {
			if v == "true" {
				return true, nil
			} else if v == "false" {
				return false, nil
			}
			return v, nil
		}
		return defaultValue, nil
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		ApiHost:      d.Get("api_host").(string),
		ClientId:     d.Get("client_id").(string),
		ClientSecret: d.Get("client_secret").(string),
	}
	return config.Client()
}

func GetUUIDOK(d *schema.ResourceData, key string) (strfmt.UUID, bool) {
	if attr, ok := d.GetOk(key); ok {
		stringUUID := string(attr.(string))
		return strfmt.UUID(stringUUID), true
	}

	return "", false
}
