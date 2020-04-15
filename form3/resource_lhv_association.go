package form3

import (
	"fmt"
	"log"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"client_code": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"client_country": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"use_simulator": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceLhvAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createLhvNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create lhv association: %s", err)
	}

	createdAssociation, err := client.AssociationClient.Associations.PostLhv(associations.NewPostLhvParams().
		WithCreationRequest(&models.LhvAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create lhv association: %s", err)
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] lhv sct association key: %s", d.Id())

	return resourceLhvAssociationRead(d, meta)
}

func resourceLhvAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	lhvAssociation, err := client.AssociationClient.Associations.GetLhvAssociationID(associations.NewGetLhvAssociationIDParams().
		WithAssociationID(associationId))

	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find lhv sct association: %s", err)
	}

	_ = d.Set("association_id", lhvAssociation.Payload.Data.ID.String())
	_ = d.Set("name", lhvAssociation.Payload.Data.Attributes.Name)
	_ = d.Set("client_code", lhvAssociation.Payload.Data.Attributes.ClientCode)
	_ = d.Set("client_country", lhvAssociation.Payload.Data.Attributes.ClientCountry)
	return nil
}

func resourceLhvAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	lhvAssociation, err := client.AssociationClient.Associations.GetLhvAssociationID(associations.NewGetLhvAssociationIDParams().
		WithAssociationID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting lhv sct association: %s", err)
	}

	_, err = client.AssociationClient.Associations.DeleteLhvAssociationID(associations.NewDeleteLhvAssociationIDParams().
		WithAssociationID(lhvAssociation.Payload.Data.ID).
		WithVersion(*lhvAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting lhv sct association: %s", err)
	}

	return nil
}

func createLhvNewAssociationFromResourceData(d *schema.ResourceData) (*models.NewLhvAssociation, error) {

	association := models.NewLhvAssociation{
		Type:       models.LhvAssociationTypeLhvgatewayAssociations,
		Attributes: &models.LhvAssociationAttributes{},
	}

	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("client_code"); ok {
		association.Attributes.ClientCode = attr.(string)
	}
	if attr, ok := d.GetOk("client_country"); ok {
		association.Attributes.ClientCountry = attr.(string)
	}

	if attr, ok := d.GetOk("use_simulator"); ok {
		association.Attributes.UseSimulator = attr.(bool)
	}

	if attr, ok := d.GetOk("name"); ok {
		association.Attributes.Name = attr.(string)
	}

	return &association, nil
}
