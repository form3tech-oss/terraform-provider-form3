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

func resourceForm3StarlingAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceStarlingAssociationCreate,
		Read:   resourceStarlingAssociationRead,
		Delete: resourceStarlingAssociationDelete,

		Schema: map[string]*schema.Schema{
			"association_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"starling_account_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"starling_account_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceStarlingAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	name := d.Get("starling_account_name").(string)
	log.Printf("[INFO] Creating association with account name: %s", name)

	association, err := createNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create association: %s", err)
	}

	createdAssociation, err := client.AssociationClient.Associations.PostStarling(associations.NewPostStarlingParams().
		WithCreationRequest(&models.AssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create association: %s", err)
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] association key: %s", d.Id())

	return resourceStarlingAssociationRead(d, meta)
}

func resourceStarlingAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	key := d.Id()
	associationId, _ := GetUUIDOK(d, "association_id")
	associationName := d.Get("starling_account_name").(string)
	log.Printf("[INFO] Reading association for id: %s account name: %s", key, associationName)

	association, err := client.AssociationClient.Associations.GetStarlingID(associations.NewGetStarlingIDParams().
		WithID(associationId))

	if err != nil {
		apiError := err.(*runtime.APIError)
		if apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find association: %s", err)
	}

	d.Set("association_id", association.Payload.Data.ID.String())
	d.Set("starling_account_name", association.Payload.Data.Attributes.StarlingAccountName)
	d.Set("starling_account_id", association.Payload.Data.Attributes.StarlingAccountUID)
	return nil
}

func resourceStarlingAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationFromResource, err := createAssociationFromResourceDataWithVersion(d, client)
	if err != nil {
		return fmt.Errorf("error deleting association: %s", err)
	}

	log.Printf("[INFO] Deleting association for id: %s account name: %s", associationFromResource.ID, associationFromResource.Attributes.StarlingAccountName)

	_, err = client.AssociationClient.Associations.DeleteStarlingID(associations.NewDeleteStarlingIDParams().
		WithID(associationFromResource.ID).
		WithVersion(*associationFromResource.Version))

	if err != nil {
		return fmt.Errorf("error deleting association: %s", err)
	}

	return nil
}

func createAssociationFromResourceDataWithVersion(d *schema.ResourceData, client *form3.AuthenticatedClient) (*models.Association, error) {
	association, err := createAssociationFromResourceData(d)
	version, err := getAssociationVersion(client, association.ID)
	if err != nil {
		return nil, err
	}

	association.Version = &version

	return association, nil
}

func createAssociationFromResourceData(d *schema.ResourceData) (*models.Association, error) {

	association := models.Association{Attributes: &models.AssociationAttributes{}}
	association.Type = "associations"
	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("starling_account_name"); ok {
		association.Attributes.StarlingAccountName = attr.(string)
	}

	if attr, ok := GetUUIDOK(d, "starling_account_id"); ok {
		association.Attributes.StarlingAccountUID = attr
	}

	return &association, nil
}

func createNewAssociationFromResourceData(d *schema.ResourceData) (*models.NewAssociation, error) {

	association := models.NewAssociation{Attributes: &models.NewAssociationAttributes{}}
	association.Type = "associations"
	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("starling_account_name"); ok {
		association.Attributes.StarlingAccountName = attr.(string)
	}

	return &association, nil
}

func getAssociationVersion(client *form3.AuthenticatedClient, associationId strfmt.UUID) (int64, error) {
	association, err := client.AssociationClient.Associations.GetStarlingID(associations.NewGetStarlingIDParams().WithID(associationId))
	if err != nil {
		if err != nil {
			return -1, fmt.Errorf("error reading association: %s", err)
		}
	}

	return *association.Payload.Data.Version, nil
}
