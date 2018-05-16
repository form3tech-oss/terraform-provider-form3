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

func resourceForm3BacsAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceBacsAssociationCreate,
		Read:   resourceBacsAssociationRead,
		Delete: resourceBacsAssociationDelete,

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
			"service_user_number": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_number": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sorting_code": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_type": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBacsAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	serviceUserNumber := d.Get("service_user_number").(string)
	log.Printf("[INFO] Creating Bacs association with service user number: %s", serviceUserNumber)

	association, err := createBacsNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create Bacs association: %s", err)
	}

	createdAssociation, err := client.AssociationClient.Associations.PostBacs(associations.NewPostBacsParams().
		WithCreationRequest(&models.BacsAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create Bacs association: %s", err)
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] Bacs association key: %s", d.Id())

	return resourceBacsAssociationRead(d, meta)
}

func resourceBacsAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	key := d.Id()
	associationId, _ := GetUUIDOK(d, "association_id")
	serviceUserNumber := d.Get("service_user_number").(string)
	log.Printf("[INFO] Reading Bacs association for id: %s service user number: %s", key, serviceUserNumber)

	bacsAssociation, err := client.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().
		WithID(associationId))

	if err != nil {
		apiError := err.(*runtime.APIError)
		if apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find Bacs association: %s", err)
	}

	d.Set("association_id", bacsAssociation.Payload.Data.ID.String())
	d.Set("service_user_number", bacsAssociation.Payload.Data.Attributes.ServiceUserNumber)
	d.Set("account_number", bacsAssociation.Payload.Data.Attributes.AccountNumber)
	d.Set("sorting_code", bacsAssociation.Payload.Data.Attributes.SortingCode)
	d.Set("account_type", bacsAssociation.Payload.Data.Attributes.AccountType)
	return nil
}

func resourceBacsAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationFromResource, err := createBacsAssociationFromResourceDataWithVersion(d, client)
	if err != nil {
		return fmt.Errorf("error deleting Bacs association: %s", err)
	}

	log.Printf("[INFO] Deleting Bacs association for id: %s service user number: %s", associationFromResource.ID, associationFromResource.Attributes.ServiceUserNumber)

	_, err = client.AssociationClient.Associations.DeleteBacsID(associations.NewDeleteBacsIDParams().
		WithID(associationFromResource.ID).
		WithVersion(*associationFromResource.Version))

	if err != nil {
		return fmt.Errorf("error deleting Bacs association: %s", err)
	}

	return nil
}

func createBacsAssociationFromResourceDataWithVersion(d *schema.ResourceData, client *form3.AuthenticatedClient) (*models.BacsAssociation, error) {
	association, err := createBacsAssociationFromResourceData(d)
	version, err := getBacsAssociation(client, association.ID)
	if err != nil {
		return nil, err
	}

	association.Version = &version

	return association, nil
}

func createBacsAssociationFromResourceData(d *schema.ResourceData) (*models.BacsAssociation, error) {
	association := models.BacsAssociation{Attributes: &models.BacsAssociationAttributes{}}
	association.Type = "associations"
	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("service_user_number"); ok {
		association.Attributes.ServiceUserNumber = attr.(string)
	}

	if attr, ok := d.GetOk("account_number"); ok {
		association.Attributes.AccountNumber = attr.(string)
	}

	if attr, ok := d.GetOk("sorting_code"); ok {
		association.Attributes.SortingCode = attr.(string)
	}

	if attr, ok := d.GetOk("account_type"); ok {
		association.Attributes.AccountType = int64(attr.(int))
	}

	return &association, nil
}

func createBacsNewAssociationFromResourceData(d *schema.ResourceData) (*models.BacsNewAssociation, error) {

	association := models.BacsNewAssociation{Attributes: &models.BacsAssociationAttributes{}}
	association.Type = "associations"
	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("service_user_number"); ok {
		association.Attributes.ServiceUserNumber = attr.(string)
	}

	if attr, ok := d.GetOk("account_number"); ok {
		association.Attributes.AccountNumber = attr.(string)
	}

	if attr, ok := d.GetOk("sorting_code"); ok {
		association.Attributes.SortingCode = attr.(string)
	}

	if attr, ok := d.GetOk("account_type"); ok {
		association.Attributes.AccountType = int64(attr.(int))
	}

	return &association, nil
}

func getBacsAssociation(client *form3.AuthenticatedClient, associationId strfmt.UUID) (int64, error) {
	association, err := client.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().WithID(associationId))
	if err != nil {
		if err != nil {
			return -1, fmt.Errorf("error reading Bacs association: %s", err)
		}
	}

	return *association.Payload.Data.Version, nil
}
