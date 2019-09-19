package form3

import (
	"fmt"
	"log"

	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/runtime"

	"github.com/form3tech-oss/terraform-provider-form3/client/associations"

	"github.com/form3tech-oss/terraform-provider-form3/models"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceForm3GocardlessAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceGocardlessAssociationCreate,
		Read:   resourceGocardlessAssociationRead,
		Delete: resourceGocardlessAssociationDelete,
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
			"schemes": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
		},
		SchemaVersion: 0,
	}
}

func resourceGocardlessAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createGocardlessAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create gocardless association: %s", err)
	}

	createdAssociation, err := client.AssociationClient.Associations.PostGocardless(
		associations.NewPostGocardlessParams().WithCreationRequest(&models.GocardlessAssociationCreation{Data: association}))
	if err != nil {
		return fmt.Errorf("failed to create gocardless association: %s", err)
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] gocardless association key: %s", d.Id())

	return resourceGocardlessAssociationRead(d, meta)
}

func resourceGocardlessAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	gocardlessAssociation, err := client.AssociationClient.Associations.GetGocardlessID(
		associations.NewGetGocardlessIDParams().WithID(associationId))
	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("couldn't find gocardless gateway association: %s", err)
	}

	_ = d.Set("association_id", gocardlessAssociation.Payload.ID.String())
	_ = d.Set("schemes", gocardlessAssociation.Payload.Attributes.Schemes)
	return nil
}

func resourceGocardlessAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	gocardlessAssociation, err := client.AssociationClient.Associations.GetGocardlessID(associations.NewGetGocardlessIDParams().
		WithID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting gocardless gateway association: %s", err)
	}

	_, err = client.AssociationClient.Associations.DeleteGocardlessID(associations.NewDeleteGocardlessIDParams().
		WithID(gocardlessAssociation.Payload.ID).
		WithVersion(*gocardlessAssociation.Payload.Version))

	if err != nil {
		return fmt.Errorf("error deleting gocardless gateway association: %s", err)
	}

	return nil
}

func createGocardlessAssociationFromResourceData(d *schema.ResourceData) (*models.GocardlessAssociation, error) {
	association := models.GocardlessAssociation{Attributes: &models.GocardlessAssociationAttributes{}}

	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("schemes"); ok {
		arr := attr.([]interface{})
		var schemes []string
		for _, v := range arr {
			schemes = append(schemes, v.(string))
		}
		association.Attributes.Schemes = schemes
	}

	return &association, nil
}