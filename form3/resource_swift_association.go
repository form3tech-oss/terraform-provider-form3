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

func resourceForm3SwiftAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwiftAssociationCreate,
		Read:   resourceSwiftAssociationRead,
		Delete: resourceSwiftAssociationDelete,
		Update: resourceSwiftAssociationUpdate,
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
				ForceNew: false,
			},
		},
	}
}

func resourceSwiftAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createSwiftNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create swift association: %s", form3.JsonErrorPrettyPrint(err))
	}

	createdAssociation, err := client.AssociationClient.Associations.PostSwift(associations.NewPostSwiftParams().
		WithCreationRequest(&models.SwiftAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create swift association: %s", form3.JsonErrorPrettyPrint(err))
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] swift association key: %s", d.Id())

	return resourceSwiftAssociationRead(d, meta)
}

func createSwiftNewAssociationFromResourceData(d *schema.ResourceData) (*models.SwiftNewAssociation, error) {
	association := models.SwiftNewAssociation{
		Type:       string(models.ResourceTypeSwiftAssociations),
		Attributes: &models.SwiftAssociationAttributes{},
	}

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

func resourceSwiftAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	swiftAssociation, err := client.AssociationClient.Associations.GetSwiftID(associations.NewGetSwiftIDParams().
		WithID(associationId))

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find swift association: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	_ = d.Set("association_id", swiftAssociation.Payload.Data.ID.String())
	_ = d.Set("organisation_id", swiftAssociation.Payload.Data.OrganisationID.String())
	_ = d.Set("bic", swiftAssociation.Payload.Data.Attributes.Bic)

	return nil
}

func resourceSwiftAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	swiftAssociation, err := client.AssociationClient.Associations.GetSwiftID(associations.NewGetSwiftIDParams().
		WithID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting swift association: %s", form3.JsonErrorPrettyPrint(err))
	}

	_, err = client.AssociationClient.Associations.DeleteSwiftID(associations.NewDeleteSwiftIDParams().
		WithID(swiftAssociation.Payload.Data.ID).
		WithVersion(*swiftAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting swift association: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func resourceSwiftAssociationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createSwiftUpdateAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to update swift association: %s", form3.JsonErrorPrettyPrint(err))
	}

	existingAssociation, err := client.AssociationClient.Associations.GetSwiftID(associations.NewGetSwiftIDParams().
		WithID(association.ID))
	if err != nil {
		return fmt.Errorf("error fetching swift association: %s", form3.JsonErrorPrettyPrint(err))
	}

	if existingAssociation == nil || existingAssociation.Payload == nil {
		return fmt.Errorf("swift association with id %s is nil or has a nil payload", association.ID)
	}

	association.Version = *existingAssociation.Payload.Data.Version

	log.Printf("[INFO] Updating swift association with id: %s", association.ID)
	_, err = client.AssociationClient.Associations.PatchSwiftID(associations.NewPatchSwiftIDParams().
		WithID(association.ID).
		WithPayload(&models.SwiftAssociationPatch{
			Data: association,
		}))
	if err != nil {
		return fmt.Errorf("failed to update swift association: %s", form3.JsonErrorPrettyPrint(err))
	}
	return nil
}

func createSwiftUpdateAssociationFromResourceData(d *schema.ResourceData) (*models.SwiftUpdateAssociation, error) {
	association := &models.SwiftUpdateAssociation{
		Type:       string(models.ResourceTypeSwiftAssociations),
		Attributes: &models.SwiftUpdateAssociationAttributes{},
	}
	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("bic"); ok {
		association.Attributes.Bic = attr.(string)
	}
	return association, nil
}
