package form3

import (
	"fmt"
	"github.com/ewilde/go-form3"
	"github.com/ewilde/go-form3/client/associations"
	"github.com/ewilde/go-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3SepaInstantAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSepaInstantAssociationCreate,
		Read:   resourceSepaInstantAssociationRead,
		Delete: resourceSepaInstantAssociationDelete,

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
			"business_user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSepaInstantAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createSepaInstantNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create sepa instant association: %s", err)
	}

	createdAssociation, err := client.AssociationClient.Associations.PostSepainstant(associations.NewPostSepainstantParams().
		WithCreationRequest(&models.SepaInstantAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create sepa instant association: %s", err)
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] sepa instant association key: %s", d.Id())

	return resourceSepaInstantAssociationRead(d, meta)
}

func resourceSepaInstantAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	sepaInstantAssociation, err := client.AssociationClient.Associations.GetSepainstantID(associations.NewGetSepainstantIDParams().
		WithID(associationId))

	if err != nil {
		apiError := err.(*runtime.APIError)
		if apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find sepa instant association: %s", err)
	}

	d.Set("association_id", sepaInstantAssociation.Payload.Data.ID.String())
	d.Set("business_user_id", sepaInstantAssociation.Payload.Data.Attributes.BusinessUserID)

	return nil
}

func resourceSepaInstantAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	sepaInstantAssociation, err := client.AssociationClient.Associations.GetSepainstantID(associations.NewGetSepainstantIDParams().
		WithID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting sepa instant association: %s", err)
	}

	log.Printf("[INFO] Deleting sepa instant association for id: %s business user id: %s", sepaInstantAssociation.Payload.Data.ID, sepaInstantAssociation.Payload.Data.Attributes.BusinessUserID)

	_, err = client.AssociationClient.Associations.DeleteSepainstantID(associations.NewDeleteSepainstantIDParams().
		WithID(sepaInstantAssociation.Payload.Data.ID).
		WithVersion(*sepaInstantAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting sepa instant association: %s", err)
	}

	return nil
}

func createSepaInstantNewAssociationFromResourceData(d *schema.ResourceData) (*models.NewSepaInstantAssociation, error) {

	association := models.NewSepaInstantAssociation{Attributes: &models.SepaInstantAssociationAttributes{}}

	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("business_user_id"); ok {
		association.Attributes.BusinessUserID = attr.(string)
	}

	return &association, nil
}
