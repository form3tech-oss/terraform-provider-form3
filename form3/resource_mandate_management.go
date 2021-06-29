package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/mandates"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceForm3MandateManagement() *schema.Resource {
	return &schema.Resource{
		Create: resourceMandateManagementCreate,
		Read:   resourceMandateManagementRead,
		Delete: resourceMandateManagementDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mandate_management_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"payment_scheme": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceMandateManagementCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)
	paymentScheme := d.Get("payment_scheme").(string)
	log.Printf("[INFO] Creating mandate management with payment scheme: %s", paymentScheme)

	mandateManagement, err := createMandateManagementFromResourceData(d)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] mandate management create: %#v", mandateManagement)

	createdMandateManagement, err := client.AssociationClient.Mandates.PostMandatemanagement(mandates.NewPostMandatemanagementParams().
		WithMandateManagementCreationRequest(&models.MandateManagementCreation{
			Data: mandateManagement,
		}))

	if err != nil {
		return fmt.Errorf("failed to create mandate management: %s", form3.JsonErrorPrettyPrint(err))
	}
	d.SetId(createdMandateManagement.Payload.Data.ID.String())
	log.Printf("[INFO] mandate management key: %s", d.Id())

	return resourceMandateManagementRead(d, meta)

}

func resourceMandateManagementRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	key := d.Id()
	mandateManagementId, _ := GetUUIDOK(d, "mandate_management_id")

	if mandateManagementId == "" {
		mandateManagementId = strfmt.UUID(key)
		log.Printf("[INFO] Importing mandate management id: %s", mandateManagementId)
	} else {
		log.Printf("[INFO] Reading mandate management for id: %s ", key)
	}

	mandateManagement, err := client.AssociationClient.Mandates.GetMandatemanagementID(mandates.NewGetMandatemanagementIDParams().WithID(mandateManagementId))
	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find mandate management: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	d.Set("mandate_management_id", mandateManagement.Payload.Data.ID.String())
	d.Set("payment_scheme", mandateManagement.Payload.Data.Attributes.PaymentScheme)
	d.Set("organisation_id", mandateManagement.Payload.Data.OrganisationID.String())

	return nil
}

func createMandateManagementFromResourceData(d *schema.ResourceData) (*models.MandateManagement, error) {
	mandateManagement := models.MandateManagement{Attributes: &models.MandateManagementAttributes{}}
	if attr, ok := GetUUIDOK(d, "mandate_management_id"); ok {
		mandateManagement.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		mandateManagement.OrganisationID = attr
	}

	if attr, ok := d.GetOk("payment_scheme"); ok {
		mandateManagement.Attributes.PaymentScheme = models.PaymentScheme(attr.(string))
	}

	return &mandateManagement, nil

}

func resourceMandateManagementDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	mandateManagementFromResource, err := createMandateManagementFromResourceDataWithVersion(d, client)
	if err != nil {
		return fmt.Errorf("error deleting mandate management: %s", form3.JsonErrorPrettyPrint(err))
	}

	log.Printf("[INFO] Deleting mandate management for id: %s payment scheme: %s", mandateManagementFromResource.ID, mandateManagementFromResource.Attributes.PaymentScheme)

	_, err = client.AssociationClient.Mandates.DeleteMandatemanagementID(mandates.NewDeleteMandatemanagementIDParams().
		WithID(mandateManagementFromResource.ID).
		WithVersion(*mandateManagementFromResource.Version))

	if err != nil {
		return fmt.Errorf("error deleting mandate management: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createMandateManagementFromResourceDataWithVersion(d *schema.ResourceData, client *form3.AuthenticatedClient) (*models.MandateManagement, error) {
	mandateManagement, err := createMandateManagementFromResourceData(d)
	if err != nil {
		return nil, err
	}
	version, err := getMandateManagementVersion(client, mandateManagement.ID)
	if err != nil {
		return nil, err
	}

	mandateManagement.Version = &version

	return mandateManagement, nil
}

func getMandateManagementVersion(client *form3.AuthenticatedClient, mandateManagementId strfmt.UUID) (int64, error) {
	mandateManagement, err := client.AssociationClient.Mandates.GetMandatemanagementID(mandates.NewGetMandatemanagementIDParams().WithID(mandateManagementId))
	if err != nil {
		return -1, fmt.Errorf("error reading mandate management: %s", form3.JsonErrorPrettyPrint(err))
	}

	return *mandateManagement.Payload.Data.Version, nil
}
