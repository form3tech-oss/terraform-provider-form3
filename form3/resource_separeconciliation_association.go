package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3SepaReconciliationAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSepaReconciliationAssociationCreate,
		Read:   resourceSepaReconciliationAssociationRead,
		Delete: resourceSepaReconciliationAssociationDelete,

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
			"iban": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bic": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"country": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sponsor_id": {
				Type:     schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceSepaReconciliationAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createSepaReconciliationNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create sepa sct association: %s", err)
	}

	createdAssociation, err := client.AssociationClient.Associations.PostSepareconciliation(associations.NewPostSepareconciliationParams().
		WithCreationRequest(&models.SepaReconciliationAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create sepa sct association: %s", err)
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] sepa sct association key: %s", d.Id())

	return resourceSepaReconciliationAssociationRead(d, meta)
}

func resourceSepaReconciliationAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	sepaReconciliationAssociation, err := client.AssociationClient.Associations.GetSepareconciliationID(associations.NewGetSepareconciliationIDParams().
		WithID(associationId))

	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find sepa sct association: %s", err)
	}

	_ = d.Set("organisation_id", sepaReconciliationAssociation.Payload.Data.OrganisationID.String())
	_ = d.Set("association_id", sepaReconciliationAssociation.Payload.Data.ID.String())
	_ = d.Set("bic", sepaReconciliationAssociation.Payload.Data.Attributes.Bic)
	_ = d.Set("iban", sepaReconciliationAssociation.Payload.Data.Attributes.Iban)
	_ = d.Set("country", sepaReconciliationAssociation.Payload.Data.Attributes.Country)
	_ = d.Set("name", sepaReconciliationAssociation.Payload.Data.Attributes.Name)

	return nil
}

func resourceSepaReconciliationAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	sepaReconciliationAssociation, err := client.AssociationClient.Associations.GetSepareconciliationID(associations.NewGetSepareconciliationIDParams().
		WithID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting sepa sct association: %s", err)
	}

	_, err = client.AssociationClient.Associations.DeleteSepareconciliationID(associations.NewDeleteSepareconciliationIDParams().
		WithID(sepaReconciliationAssociation.Payload.Data.ID).
		WithVersion(*sepaReconciliationAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting sepa sct association: %s", err)
	}

	return nil
}

func createSepaReconciliationNewAssociationFromResourceData(d *schema.ResourceData) (*models.NewSepaReconciliationAssociation, error) {

	association := models.NewSepaReconciliationAssociation{Attributes: &models.SepaReconciliationAssociationAttributes{}}
	association.Type = "sepaReconciliation_associations"
	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("bic"); ok {
		association.Attributes.Bic = attr.(string)
	}

	if attr, ok := d.GetOk("iban"); ok {
		association.Attributes.Iban = attr.(string)
	}

	if attr, ok := d.GetOk("country"); ok {
		association.Attributes.Country = attr.(string)
	}

	if attr, ok := d.GetOk("name"); ok {
		association.Attributes.Name = attr.(string)
	}

	return &association, nil
}
