package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/accounts"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceForm3BankID() *schema.Resource {
	return &schema.Resource{
		Create: resourceBankIDCreate,
		Read:   resourceBankIDRead,
		Delete: resourceBankIDDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"bank_resource_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bank_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bank_id_code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBankIDCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*form3.AuthenticatedClient)

	bankResourceID := d.Get("bank_resource_id").(string)
	log.Printf("[INFO] Creating bank id with id: %s", bankResourceID)

	bankIDResource, err := createBankIDFromResourceData(d)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] bank id create: %#v", bankIDResource)

	createdBankID, err := client.AccountClient.Accounts.PostBankids(accounts.NewPostBankidsParams().
		WithBankIDCreationRequest(&models.BankIDCreation{
			Data: bankIDResource,
		}))

	if err != nil {
		return fmt.Errorf("failed to create bank id: %s", err)
	}

	d.SetId(createdBankID.Payload.Data.ID.String())
	log.Printf("[INFO] bankId key: %s", d.Id())

	return resourceBankIDRead(d, meta)
}

func resourceBankIDRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	bankResourceID, _ := GetUUIDOK(d, "bank_resource_id")

	if bankResourceID == "" {
		bankResourceID = strfmt.UUID(d.Id())
		log.Printf("[INFO] Importing bank id with resource id: %s.", bankResourceID)
	} else {
		log.Printf("[INFO] Reading bank id with resource id: %s.", bankResourceID)
	}

	bankIDResponse, err := client.AccountClient.Accounts.GetBankidsID(accounts.NewGetBankidsIDParams().
		WithID(bankResourceID))

	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find bank id: %s", err)
	}

	d.Set("bank_resource_id", bankIDResponse.Payload.Data.ID.String())
	d.Set("organisation_id", bankIDResponse.Payload.Data.OrganisationID.String())
	d.Set("bank_id", bankIDResponse.Payload.Data.Attributes.BankID)
	d.Set("bank_id_code", bankIDResponse.Payload.Data.Attributes.BankIDCode)
	d.Set("country", bankIDResponse.Payload.Data.Attributes.Country)
	return nil
}

func resourceBankIDDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	bankIdFromResource, err := createbBankIDFromResourceDataWithVersion(d, client)
	if err != nil {
		return fmt.Errorf("error deleting bankId: %s", err)
	}

	log.Printf("[INFO] Deleting bank id with resource id: %s bank id: %s", bankIdFromResource.ID, bankIdFromResource.Attributes.BankID)

	_, err = client.AccountClient.Accounts.DeleteBankidsID(accounts.NewDeleteBankidsIDParams().
		WithID(bankIdFromResource.ID).
		WithVersion(*bankIdFromResource.Version))

	if err != nil {
		return fmt.Errorf("error deleting bankId: %s", err)
	}

	return nil
}

func createBankIDFromResourceData(d *schema.ResourceData) (*models.BankID, error) {

	bankId := models.BankID{Attributes: &models.BankIDAttributes{}}
	bankId.Type = "bankids"
	if attr, ok := GetUUIDOK(d, "bank_resource_id"); ok {
		bankId.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		bankId.OrganisationID = attr
	}

	if attr, ok := d.GetOk("bank_id"); ok {
		bankId.Attributes.BankID = attr.(string)
	}

	if attr, ok := d.GetOk("bank_id_code"); ok {
		bankId.Attributes.BankIDCode = attr.(string)
	}

	if attr, ok := d.GetOk("country"); ok {
		bankId.Attributes.Country = attr.(string)
	}

	return &bankId, nil
}

func createbBankIDFromResourceDataWithVersion(d *schema.ResourceData, client *form3.AuthenticatedClient) (*models.BankID, error) {
	bankID, err := createBankIDFromResourceData(d)
	version, err := getBankVersion(client, bankID.ID)
	if err != nil {
		return nil, err
	}

	bankID.Version = &version

	return bankID, nil
}

func getBankVersion(client *form3.AuthenticatedClient, id strfmt.UUID) (int64, error) {
	bankID, err := client.AccountClient.Accounts.GetBankidsID(accounts.NewGetBankidsIDParams().WithID(id))
	if err != nil {
		if err != nil {
			return -1, fmt.Errorf("error reading bankd id: %s", err)
		}
	}

	return *bankID.Payload.Data.Version, nil
}
