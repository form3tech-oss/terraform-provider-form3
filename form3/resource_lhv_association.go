package form3

import (
	"fmt"
	"log"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceForm3LhvAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceLhvAssociationCreate,
		Read:   resourceLhvAssociationRead,
		Delete: resourceLhvAssociationDelete,
		Update: resourceLhvAssociationUpdate,

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
				ForceNew: false,
			},
			"client_code": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"client_country": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
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
		return fmt.Errorf("failed to create lhv association: %s", form3.JsonErrorPrettyPrint(err))
	}

	createdAssociation, err := client.AssociationClient.Associations.PostLhv(associations.NewPostLhvParams().
		WithCreationRequest(&models.LhvAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create lhv association: %s", form3.JsonErrorPrettyPrint(err))
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
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find lhv sct association: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
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
		return fmt.Errorf("error deleting lhv sct association: %s", form3.JsonErrorPrettyPrint(err))
	}

	_, err = client.AssociationClient.Associations.DeleteLhvAssociationID(associations.NewDeleteLhvAssociationIDParams().
		WithAssociationID(lhvAssociation.Payload.Data.ID).
		WithVersion(*lhvAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting lhv association: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func resourceLhvAssociationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createLhvUpdateAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create update lhv association: %s", form3.JsonErrorPrettyPrint(err))
	}

	existingAssociation, err := client.AssociationClient.Associations.GetLhvAssociationID(associations.NewGetLhvAssociationIDParams().
		WithAssociationID(association.ID))
	if err != nil {
		return fmt.Errorf("error fetching lhv association: %s", form3.JsonErrorPrettyPrint(err))
	}

	if existingAssociation == nil || existingAssociation.Payload == nil {
		return fmt.Errorf("lhv association with id %s is nil or has a nil payload", association.ID)
	}

	association.Version = existingAssociation.Payload.Data.Version

	log.Printf("[INFO] Updating lhv association with id: %s", association.ID)
	_, err = client.AccountClient.Associations.PatchLhvAssociationID(associations.NewPatchLhvAssociationIDParams().
		WithAssociationID(association.ID).
		WithPayload(&models.LhvAssociationPatch{
			Data: association,
		}))
	if err != nil {
		return fmt.Errorf("failed to update lhv association: %s", form3.JsonErrorPrettyPrint(err))
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

	if attr, ok := d.GetOk("name"); ok {
		association.Attributes.Name = attr.(string)
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
	return &association, nil
}

func createLhvUpdateAssociationFromResourceData(d *schema.ResourceData) (*models.LhvUpdateAssociation, error) {
	association := models.LhvUpdateAssociation{
		Type:       models.LhvAssociationTypeLhvgatewayAssociations,
		Attributes: &models.LhvUpdateAssociationAttributes{},
	}

	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("name"); ok {
		association.Attributes.Name = attr.(string)
	}

	if attr, ok := d.GetOk("client_code"); ok {
		association.Attributes.ClientCode = attr.(string)
	}

	if attr, ok := d.GetOk("client_country"); ok {
		association.Attributes.ClientCountry = attr.(string)
	}
	if attr, ok := d.GetOk("use_simulator"); ok {
		res := attr.(bool)
		association.Attributes.UseSimulator = &res
	}
	return &association, nil
}
