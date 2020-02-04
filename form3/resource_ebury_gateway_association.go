package form3

import (
	"fmt"
	"log"

	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/runtime"

	"github.com/form3tech-oss/terraform-provider-form3/client/associations"

	"github.com/form3tech-oss/terraform-provider-form3/models"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceForm3EburyAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceEburyAssociationCreate,
		Read:   resourceEburyAssociationRead,
		Delete: resourceEburyAssociationDelete,
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
		},
		SchemaVersion: 0,
	}
}

func resourceEburyAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createEburyAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create ebury association: %s", err)
	}

	createdAssociation, err := client.AssociationClient.Associations.PostEbury(
		associations.NewPostEburyParams().WithCreationRequest(&models.EburyAssociationCreation{Data: association}))
	if err != nil {
		return fmt.Errorf("failed to create ebury association: %s", err)
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] ebury association key: %s", d.Id())

	return resourceEburyAssociationRead(d, meta)
}

func resourceEburyAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	eburyAssociation, err := client.AssociationClient.Associations.GetEburyID(
		associations.NewGetEburyIDParams().WithID(associationId))
	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("couldn't find ebury gateway association: %s", err)
	}

	_ = d.Set("association_id", eburyAssociation.Payload.Data.ID.String())
	return nil
}

func resourceEburyAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	eburyAssociation, err := client.AssociationClient.Associations.GetEburyID(associations.NewGetEburyIDParams().
		WithID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting ebury gateway association: %s", err)
	}

	_, err = client.AssociationClient.Associations.DeleteEburyID(associations.NewDeleteEburyIDParams().
		WithID(eburyAssociation.Payload.Data.ID).
		WithVersion(*eburyAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting ebury gateway association: %s", err)
	}

	return nil
}

func createEburyAssociationFromResourceData(d *schema.ResourceData) (*models.NewEburyAssociation, error) {
	association := models.NewEburyAssociation{
		Type: string(models.ResourceTypeEburyAssociations),
	}

	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	return &association, nil
}
