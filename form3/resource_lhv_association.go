package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3LhvAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceLhvAssociationCreate,
		Read:   resourceLhvAssociationRead,
		Delete: resourceLhvAssociationDelete,

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

func resourceLhvAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createLhvNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create lhv sct association: %s", err)
	}

	createdAssociation, err := client.AssociationClient.Associations.PostLhv(associations.NewPostLhvParams().
		WithCreationRequest(&models.LhvAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create lhv sct association: %s", err)
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] lhv sct association key: %s", d.Id())

	return resourceLhvAssociationRead(d, meta)
}

func resourceLhvAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	lhvAssociation, err := client.AssociationClient.Associations.GetLhvID(associations.NewGetLhvIDParams().
		WithID(associationId))

	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find lhv sct association: %s", err)
	}

	_ = d.Set("association_id", lhvAssociation.Payload.Data.ID.String())
	_ = d.Set("bic", lhvAssociation.Payload.Data.Attributes.Bic)
	return nil
}

func resourceLhvAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	lhvAssociation, err := client.AssociationClient.Associations.GetLhvID(associations.NewGetLhvIDParams().
		WithID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting lhv sct association: %s", err)
	}

	_, err = client.AssociationClient.Associations.DeleteLhvID(associations.NewDeleteLhvIDParams().
		WithID(lhvAssociation.Payload.Data.ID).
		WithVersion(*lhvAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting lhv sct association: %s", err)
	}

	return nil
}

func createLhvNewAssociationFromResourceData(d *schema.ResourceData) (*models.NewLhvAssociation, error) {

	association := models.NewLhvAssociation{Attributes: &models.LhvAssociationAttributes{}}
	association.Type = "lhv_associations"
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
