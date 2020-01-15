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
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bic": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"iban": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address_street": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address_building_number": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address_city": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address_country": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sponsor_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "",
			},
		},
	}
}

func resourceSepaReconciliationAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createSepaReconciliationNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create sepa reconciliation association: %s", err)
	}

	createdAssociation, err := client.AssociationClient.Associations.PostSepaReconciliation(associations.NewPostSepaReconciliationParams().
		WithCreationRequest(&models.SepaReconciliationAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create sepa reconciliation association: %s", err)
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] sepa reconciliation assocation key: %s", d.Id())

	return resourceSepaReconciliationAssociationRead(d, meta)
}

func resourceSepaReconciliationAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	sepaReconciliationAssociation, err := client.AssociationClient.Associations.GetSepaReconciliationID(associations.NewGetSepaReconciliationIDParams().
		WithID(associationId))

	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find sepa reconciliation assocation: %s", err)
	}

	_ = d.Set("association_id", sepaReconciliationAssociation.Payload.Data.ID.String())
	_ = d.Set("name", sepaReconciliationAssociation.Payload.Data.Attributes.Name)
	_ = d.Set("bic", sepaReconciliationAssociation.Payload.Data.Attributes.Bic)
	_ = d.Set("iban", sepaReconciliationAssociation.Payload.Data.Attributes.Iban)
	_ = d.Set("address_street", sepaReconciliationAssociation.Payload.Data.Attributes.Address.Street)
	_ = d.Set("address_building_number", sepaReconciliationAssociation.Payload.Data.Attributes.Address.BuildingNumber)
	_ = d.Set("address_city", sepaReconciliationAssociation.Payload.Data.Attributes.Address.City)
	_ = d.Set("address_country", sepaReconciliationAssociation.Payload.Data.Attributes.Address.Country)
	if sepaReconciliationAssociation.Payload.Data.Relationships == nil {
		d.Set("sponsor_id", "")
	} else {
		d.Set("sponsor_id", sepaReconciliationAssociation.Payload.Data.Relationships.Sponsor.Data.ID.String())
	}
	return nil
}

func resourceSepaReconciliationAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	SepaReconciliationAssociation, err := client.AssociationClient.Associations.GetSepaReconciliationID(associations.NewGetSepaReconciliationIDParams().
		WithID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting sepa reconciliation assocation: %s", err)
	}

	_, err = client.AssociationClient.Associations.DeleteSepaReconciliationID(associations.NewDeleteSepaReconciliationIDParams().
		WithID(SepaReconciliationAssociation.Payload.Data.ID).
		WithVersion(*SepaReconciliationAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting sepa reconciliation assocation: %s", err)
	}

	return nil
}

func createSepaReconciliationNewAssociationFromResourceData(d *schema.ResourceData) (*models.NewSepaReconciliationAssociation, error) {

	association := models.NewSepaReconciliationAssociation{Attributes: &models.SepaReconciliationAssociationAttributes{}}
	association.Type = "separeconciliation_associations"
	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("name"); ok {
		association.Attributes.Name = attr.(string)
	}

	if attr, ok := d.GetOk("bic"); ok {
		association.Attributes.Bic = attr.(string)
	}

	if attr, ok := d.GetOk("iban"); ok {
		association.Attributes.Iban = attr.(string)
	}

	if attr, ok := d.GetOk("address_street"); ok {
		association.Attributes.Address.Street = attr.(string)
	}

	if attr, ok := d.GetOk("address_building_number"); ok {
		association.Attributes.Address.BuildingNumber = attr.(string)
	}

	if attr, ok := d.GetOk("address_city"); ok {
		association.Attributes.Address.City = attr.(string)
	}

	if attr, ok := d.GetOk("address_country"); ok {
		association.Attributes.Address.Country = attr.(string)
	}

	if attr, ok := GetUUIDOK(d, "sponsor_id"); ok {
		association.Relationships = &models.SepaReconciliationAssociationRelationships{
			Sponsor: models.SepaReconciliationAssociationRelationshipsSponsor{
				Data: models.SepaReconciliationRelationshipData{
					ID:   attr,
					Type: "separeconciliation_associations",
				},
			},
		}
	}
	return &association, nil
}
