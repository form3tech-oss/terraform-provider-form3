package form3

import (
	"fmt"
	"github.com/form3tech-oss/go-form3"
	"github.com/form3tech-oss/go-form3/client/associations"
	"github.com/form3tech-oss/go-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3SepaSctAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSepaSctAssociationCreate,
		Read:   resourceSepaSctAssociationRead,
		Delete: resourceSepaSctAssociationDelete,

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
			"bic": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSepaSctAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createSepaSctNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create sepa sct association: %s", err)
	}

	createdAssociation, err := client.AssociationClient.Associations.PostSepasct(associations.NewPostSepasctParams().
		WithCreationRequest(&models.SepaSctAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create sepa sct association: %s", err)
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] sepa sct association key: %s", d.Id())

	return resourceSepaSctAssociationRead(d, meta)
}

func resourceSepaSctAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	sepaSctAssociation, err := client.AssociationClient.Associations.GetSepasctID(associations.NewGetSepasctIDParams().
		WithID(associationId))

	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find sepa sct association: %s", err)
	}

	_ = d.Set("association_id", sepaSctAssociation.Payload.Data.ID.String())
	_ = d.Set("bic", sepaSctAssociation.Payload.Data.Attributes.Bic)
	return nil
}

func resourceSepaSctAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	sepaSctAssociation, err := client.AssociationClient.Associations.GetSepasctID(associations.NewGetSepasctIDParams().
		WithID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting sepa sct association: %s", err)
	}

	_, err = client.AssociationClient.Associations.DeleteSepasctID(associations.NewDeleteSepasctIDParams().
		WithID(sepaSctAssociation.Payload.Data.ID).
		WithVersion(*sepaSctAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting sepa sct association: %s", err)
	}

	return nil
}

func createSepaSctNewAssociationFromResourceData(d *schema.ResourceData) (*models.NewSepaSctAssociation, error) {

	association := models.NewSepaSctAssociation{Attributes: &models.SepaSctAssociationAttributes{}}
	association.Type = "sepasct_associations"
	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("bic"); ok {
		association.Attributes.Bic = attr.(string)
	}
	return &association, nil
}
