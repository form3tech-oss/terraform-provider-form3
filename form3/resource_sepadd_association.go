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

func resourceForm3SepaDDAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSepaDDAssociationCreate,
		Read:   resourceSepaDDAssociationRead,
		Delete: resourceSepaDDAssociationDelete,
		Update: resourceSepaDDAssociationUpdate,
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
			"business_user": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"receiver_business_user": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"local_instrument": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"allow_submissions": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  false,
			},
		},
	}
}

func createSepaDDUpdateAssociationFromResourceData(d *schema.ResourceData) (*models.SepaDDAssociationUpdate, error) {
	association := &models.SepaDDAssociationUpdate{
		Type:       string(models.ResourceTypeSepaddAssociations),
		Attributes: &models.SepaDDAssociationAttributes{},
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

	if attr, ok := d.GetOk("business_user"); ok {
		association.Attributes.BusinessUser = attr.(string)
	}

	if attr, ok := d.GetOk("receiver_business_user"); ok {
		association.Attributes.ReceiverBusinessUser = attr.(string)
	}
	if attr, ok := d.GetOk("local_instrument"); ok {
		association.Attributes.LocalInstrument = attr.(string)
	}
	if attr, ok := d.GetOk("allow_submissions"); ok {
		association.Attributes.AllowSubmissions = attr.(bool)
	}
	return association, nil
}

func resourceSepaDDAssociationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createSepaDDUpdateAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to update sepadd association: %s", form3.JsonErrorPrettyPrint(err))
	}

	existingAssociation, err := client.AssociationClient.Associations.GetSepaddID(associations.NewGetSepaddIDParams().
		WithID(association.ID))
	if err != nil {
		return fmt.Errorf("error fetching sepadd association: %s", form3.JsonErrorPrettyPrint(err))
	}

	if existingAssociation == nil || existingAssociation.Payload == nil {
		return fmt.Errorf("sepadd association with id %s is nil or has a nil payload", association.ID)
	}

	association.Version = existingAssociation.Payload.Data.Version

	log.Printf("[INFO] Updating sepadd association with id: %s", association.ID)
	_, err = client.AssociationClient.Associations.PatchSepaddID(associations.NewPatchSepaddIDParams().
		WithID(association.ID).
		WithUpdateRequest(&models.SepaDDAssociationPatch{
			Data: association,
		}))
	if err != nil {
		return fmt.Errorf("failed to update sepadd association: %s", form3.JsonErrorPrettyPrint(err))
	}
	return nil
}
func resourceSepaDDAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createSepaDDNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create sepa dd association: %s", form3.JsonErrorPrettyPrint(err))
	}

	createdAssociation, err := client.AssociationClient.Associations.PostSepadd(associations.NewPostSepaddParams().
		WithCreationRequest(&models.SepaDDAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create sepa dd association: %s", form3.JsonErrorPrettyPrint(err))
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] sepa dd association key: %s", d.Id())

	return resourceSepaDDAssociationRead(d, meta)
}

func resourceSepaDDAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	sepaDDAssociation, err := client.AssociationClient.Associations.GetSepaddID(associations.NewGetSepaddIDParams().
		WithID(associationId))

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find sepa dd association: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	_ = d.Set("association_id", sepaDDAssociation.Payload.Data.ID.String())
	_ = d.Set("organisation_id", sepaDDAssociation.Payload.Data.OrganisationID.String())
	_ = d.Set("bic", sepaDDAssociation.Payload.Data.Attributes.Bic)
	_ = d.Set("business_user", sepaDDAssociation.Payload.Data.Attributes.BusinessUser)
	_ = d.Set("receiver_business_user", sepaDDAssociation.Payload.Data.Attributes.ReceiverBusinessUser)
	_ = d.Set("local_instrument", sepaDDAssociation.Payload.Data.Attributes.LocalInstrument)
	_ = d.Set("allow_submissions", sepaDDAssociation.Payload.Data.Attributes.AllowSubmissions)

	return nil
}

func resourceSepaDDAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	sepaDDAssociation, err := client.AssociationClient.Associations.GetSepaddID(associations.NewGetSepaddIDParams().
		WithID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting sepa dd association: %s", form3.JsonErrorPrettyPrint(err))
	}

	_, err = client.AssociationClient.Associations.DeleteSepaddID(associations.NewDeleteSepaddIDParams().
		WithID(sepaDDAssociation.Payload.Data.ID).
		WithVersion(*sepaDDAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting sepa dd association: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createSepaDDNewAssociationFromResourceData(d *schema.ResourceData) (*models.NewSepaDDAssociation, error) {
	association := models.NewSepaDDAssociation{
		Type:       string(models.ResourceTypeSepaddAssociations),
		Attributes: &models.SepaDDAssociationAttributes{},
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

	if attr, ok := d.GetOk("business_user"); ok {
		association.Attributes.BusinessUser = attr.(string)
	}

	if attr, ok := d.GetOk("receiver_business_user"); ok {
		association.Attributes.ReceiverBusinessUser = attr.(string)
	}

	if attr, ok := d.GetOk("local_instrument"); ok {
		association.Attributes.LocalInstrument = attr.(string)
	}

	if attr, ok := d.GetOk("allow_submissions"); ok {
		association.Attributes.AllowSubmissions = attr.(bool)
	}

	return &association, nil
}
