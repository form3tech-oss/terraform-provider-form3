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

func resourceForm3Bic() *schema.Resource {
	return &schema.Resource{
		Create: resourceBicCreate,
		Read:   resourceBicRead,
		Delete: resourceBicDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"bic_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bic": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBicCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*form3.AuthenticatedClient)

	bicID := d.Get("bic_id").(string)
	log.Printf("[INFO] Creating bic with id: %s", bicID)

	bicResource, err := createBicFromResourceData(d)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] bic create: %#v", bicResource)

	createdBic, err := client.AccountClient.Accounts.PostBics(accounts.NewPostBicsParams().
		WithBicCreationRequest(&models.BicCreation{
			Data: bicResource,
		}))

	if err != nil {
		return fmt.Errorf("failed to create bic: %s", err)
	}

	d.SetId(createdBic.Payload.Data.ID.String())
	log.Printf("[INFO] bic key: %s", d.Id())

	return resourceBicRead(d, meta)
}

func resourceBicRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	bicID, _ := GetUUIDOK(d, "bic_id")

	if bicID == "" {
		bicID = strfmt.UUID(d.Id())
		log.Printf("[INFO] Importing bic with resource id: %s.", bicID)
	} else {
		log.Printf("[INFO] Reading bic with resource id: %s.", bicID)
	}

	bicResponse, err := client.AccountClient.Accounts.GetBicsID(accounts.NewGetBicsIDParams().
		WithID(bicID))

	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find bic: %s", err)
	}

	d.Set("bic_id", bicResponse.Payload.Data.ID.String())
	d.Set("organisation_id", bicResponse.Payload.Data.OrganisationID.String())
	d.Set("bic", bicResponse.Payload.Data.Attributes.Bic)
	return nil
}

func resourceBicDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	bicFromResource, err := createBicFromResourceDataWithVersion(d, client)
	if err != nil {
		return fmt.Errorf("error deleting bic: %s", err)
	}

	log.Printf("[INFO] Deleting bic with id: %s bic: %s", bicFromResource.ID, bicFromResource.Attributes.Bic)

	_, err = client.AccountClient.Accounts.DeleteBicsID(accounts.NewDeleteBicsIDParams().
		WithID(bicFromResource.ID).
		WithVersion(*bicFromResource.Version))

	if err != nil {
		return fmt.Errorf("error deleting bic: %s", err)
	}

	return nil
}

func createBicFromResourceData(d *schema.ResourceData) (*models.Bic, error) {
	bic := models.Bic{Attributes: &models.BicAttributes{}}
	bic.Type = "bics"
	if attr, ok := GetUUIDOK(d, "bic_id"); ok {
		bic.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		bic.OrganisationID = attr
	}

	if attr, ok := d.GetOk("bic"); ok {
		bic.Attributes.Bic = attr.(string)
	}

	return &bic, nil
}

func createBicFromResourceDataWithVersion(d *schema.ResourceData, client *form3.AuthenticatedClient) (*models.Bic, error) {
	bic, err := createBicFromResourceData(d)
	version, err := getBicVersion(client, bic.ID)
	if err != nil {
		return nil, err
	}

	bic.Version = &version

	return bic, nil
}

func getBicVersion(client *form3.AuthenticatedClient, id strfmt.UUID) (int64, error) {
	bic, err := client.AccountClient.Accounts.GetBicsID(accounts.NewGetBicsIDParams().WithID(id))
	if err != nil {
		if err != nil {
			return -1, fmt.Errorf("error reading bic id: %s", err)
		}
	}

	return *bic.Payload.Data.Version, nil
}
