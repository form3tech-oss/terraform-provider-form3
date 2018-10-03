package form3

import (
	"fmt"
	"github.com/form3tech-oss/go-form3"
	"github.com/form3tech-oss/go-form3/client/payment_defaults"
	"github.com/form3tech-oss/go-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3PaymentDefaults() *schema.Resource {
	return &schema.Resource{
		Create: resourcePaymentDefaultsCreate,
		Read:   resourcePaymentDefaultsRead,
		Delete: resourcePaymentDefaultsDelete,

		Schema: map[string]*schema.Schema{
			"payment_defaults_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"default_payment_scheme": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePaymentDefaultsCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	defaultPaymentScheme := d.Get("default_payment_scheme").(string)
	log.Printf("[INFO] Creating Payment Defaults with payment scheme: %s", defaultPaymentScheme)

	paymentDefaults, err := createPaymentDefaultsFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create Payment Defaults: %s", err)
	}

	postPaymentdefaultsParams := payment_defaults.NewPostPaymentdefaultsParams().WithDefaultConfiguration(&models.PaymentDefaultsCreate{
		Data: paymentDefaults,
	})

	createdPaymentDefaults, err := client.PaymentdefaultsClient.PaymentDefaults.PostPaymentdefaults(postPaymentdefaultsParams)

	if err != nil {
		return fmt.Errorf("failed to create Payment Defaults: %s", err)
	}

	d.SetId(createdPaymentDefaults.Payload.Data.ID.String())
	log.Printf("[INFO] Payment Defaults key: %s", d.Id())

	return resourcePaymentDefaultsRead(d, meta)
}

func resourcePaymentDefaultsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	key := d.Id()
	paymentDefaultsId, _ := GetUUIDOK(d, "payment_defaults_id")
	defaultPaymentScheme := d.Get("default_payment_scheme").(string)
	log.Printf("[INFO] Reading Payment Defaults for id: %s default payment scheme: %s", key, defaultPaymentScheme)

	paymentdefaultsIDOK, err := client.PaymentdefaultsClient.PaymentDefaults.GetPaymentdefaultsID(payment_defaults.NewGetPaymentdefaultsIDParams().
		WithID(paymentDefaultsId))

	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find Payment Defaults: %s", err)
	}

	d.Set("payment_defaults_id", paymentdefaultsIDOK.Payload.Data.ID.String())
	d.Set("default_payment_scheme", paymentdefaultsIDOK.Payload.Data.Attributes.DefaultPaymentScheme)
	return nil
}

func resourcePaymentDefaultsDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	paymentDefaultsFromResource, err := createPaymentDefaultsFromResourceDataWithVersion(d, client)
	if err != nil {
		return fmt.Errorf("error deleting Payment Defaults: %s", err)
	}

	log.Printf("[INFO] Deleting Payment Defaults for id: %s default payment scheme: %s", paymentDefaultsFromResource.ID, paymentDefaultsFromResource.Attributes.DefaultPaymentScheme)

	_, err = client.PaymentdefaultsClient.PaymentDefaults.DeletePaymentdefaultsID(payment_defaults.NewDeletePaymentdefaultsIDParams().
		WithID(paymentDefaultsFromResource.ID).
		WithVersion(*paymentDefaultsFromResource.Version))

	if err != nil {
		return fmt.Errorf("error deleting Payment Defaults: %s", err)
	}

	return nil
}

func createPaymentDefaultsFromResourceDataWithVersion(d *schema.ResourceData, client *form3.AuthenticatedClient) (*models.PaymentDefaults, error) {
	paymentDefaults, err := createPaymentDefaultsFromResourceData(d)
	version, err := getPaymentDefaults(client, paymentDefaults.ID)
	if err != nil {
		return nil, err
	}

	paymentDefaults.Version = &version

	return paymentDefaults, nil
}

func createPaymentDefaultsFromResourceData(d *schema.ResourceData) (*models.PaymentDefaults, error) {

	paymentDefaults := models.PaymentDefaults{Attributes: &models.PaymentDefaultsAttributes{}}

	if attr, ok := GetUUIDOK(d, "payment_defaults_id"); ok {
		paymentDefaults.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		paymentDefaults.OrganisationID = attr
	}

	if attr, ok := d.GetOk("default_payment_scheme"); ok {
		paymentDefaults.Attributes.DefaultPaymentScheme = attr.(string)
	}

	return &paymentDefaults, nil
}

func getPaymentDefaults(client *form3.AuthenticatedClient, paymentDefaultsId strfmt.UUID) (int64, error) {
	paymentDefaults, err := client.PaymentdefaultsClient.PaymentDefaults.GetPaymentdefaultsID(payment_defaults.NewGetPaymentdefaultsIDParams().
		WithID(paymentDefaultsId))
	if err != nil {
		if err != nil {
			return -1, fmt.Errorf("error reading Payment Defaults: %s", err)
		}
	}

	return *paymentDefaults.Payload.Data.Version, nil
}
