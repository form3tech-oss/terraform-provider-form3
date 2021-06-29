package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceForm3MandateManagementAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceMandateManagementAssociationCreate,
		Read:   resourceMandateManagementAssociationRead,
		Delete: resourceMandateManagementAssociationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"association_id": {
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

func resourceMandateManagementAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)
	paymentScheme := d.Get("payment_scheme").(string)
	log.Printf("[INFO] Creating mandate management association with payment scheme: %s", paymentScheme)

	mandateManagement, err := createMandateManagementAssociationFromResourceData(d)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] mandate management association create: %#v", mandateManagement)

	createdMandateManagement, err := client.AssociationClient.Associations.PostMandatemanagement(associations.NewPostMandatemanagementParams().
		WithMandateManagementCreationRequest(&models.MandateManagementAssociationCreation{
			Data: mandateManagement,
		}))

	if err != nil {
		return fmt.Errorf("failed to create mandate management association: %s", form3.JsonErrorPrettyPrint(err))
	}
	d.SetId(createdMandateManagement.Payload.Data.ID.String())
	log.Printf("[INFO] mandate management association key: %s", d.Id())

	return resourceMandateManagementAssociationRead(d, meta)

}

func resourceMandateManagementAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	key := d.Id()
	associationID, _ := GetUUIDOK(d, "association_id")

	if associationID == "" {
		associationID = strfmt.UUID(key)
		log.Printf("[INFO] Importing mandate management association id: %s", associationID)
	} else {
		log.Printf("[INFO] Reading mandate management association for id: %s ", key)
	}

	mandateManagement, err := client.AssociationClient.Associations.GetMandatemanagementID(associations.NewGetMandatemanagementIDParams().WithID(associationID))
	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find mandate management association: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	d.Set("association_id", mandateManagement.Payload.Data.ID.String())
	d.Set("payment_scheme", mandateManagement.Payload.Data.Attributes.PaymentScheme)
	d.Set("organisation_id", mandateManagement.Payload.Data.OrganisationID.String())

	return nil
}

func createMandateManagementAssociationFromResourceData(d *schema.ResourceData) (*models.MandateManagementAssociation, error) {
	mandateManagementAssociation := models.MandateManagementAssociation{Attributes: &models.MandateManagementAssociationAttributes{}}
	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		mandateManagementAssociation.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		mandateManagementAssociation.OrganisationID = attr
	}

	if attr, ok := d.GetOk("payment_scheme"); ok {
		mandateManagementAssociation.Attributes.PaymentScheme = models.PaymentScheme(attr.(string))
	}

	return &mandateManagementAssociation, nil

}

func resourceMandateManagementAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	mandateManagementFromResource, err := createMandateManagementAssociationFromResourceDataWithVersion(d, client)
	if err != nil {
		return fmt.Errorf("error deleting mandate management association: %s", form3.JsonErrorPrettyPrint(err))
	}

	log.Printf("[INFO] Deleting mandate management association for id: %s payment scheme: %s", mandateManagementFromResource.ID, mandateManagementFromResource.Attributes.PaymentScheme)

	_, err = client.AssociationClient.Associations.DeleteMandatemanagementID(associations.NewDeleteMandatemanagementIDParams().
		WithID(mandateManagementFromResource.ID).
		WithVersion(*mandateManagementFromResource.Version))

	if err != nil {
		return fmt.Errorf("error deleting mandate management association: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createMandateManagementAssociationFromResourceDataWithVersion(d *schema.ResourceData, client *form3.AuthenticatedClient) (*models.MandateManagementAssociation, error) {
	mandateManagement, err := createMandateManagementAssociationFromResourceData(d)
	if err != nil {
		return nil, err
	}
	version, err := getMandateManagementAssociationVersion(client, mandateManagement.ID)
	if err != nil {
		return nil, err
	}

	mandateManagement.Version = &version

	return mandateManagement, nil
}

func getMandateManagementAssociationVersion(client *form3.AuthenticatedClient, mandateManagementId strfmt.UUID) (int64, error) {
	mandateManagement, err := client.AssociationClient.Associations.GetMandatemanagementID(associations.NewGetMandatemanagementIDParams().WithID(mandateManagementId))
	if err != nil {
		return -1, fmt.Errorf("error reading mandate management association: %s", form3.JsonErrorPrettyPrint(err))
	}

	return *mandateManagement.Payload.Data.Version, nil
}
