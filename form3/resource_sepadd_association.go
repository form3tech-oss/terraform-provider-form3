package form3

import (
	"fmt"
	"log"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceForm3SepaDDAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSepaDDAssociationCreate,
		Read:   resourceSepaDDAssociationRead,
		Delete: resourceSepaDDAssociationDelete,

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
			"business_user": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"receiver_business_user": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
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
	_ = d.Set("bic", sepaDDAssociation.Payload.Data.Attributes.Bic)
	_ = d.Set("business_user", sepaDDAssociation.Payload.Data.Attributes.BusinessUser)
	_ = d.Set("receiver_business_user", sepaDDAssociation.Payload.Data.Attributes.ReceiverBusinessUser)

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

	association := models.NewSepaDDAssociation{Attributes: &models.SepaDDAssociationAttributes{}}
	association.Type = "sepadd_associations"
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

	return &association, nil
}
