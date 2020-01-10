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

func resourceForm3SepaInstantAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSepaInstantAssociationCreate,
		Read:   resourceSepaInstantAssociationRead,
		Delete: resourceSepaInstantAssociationDelete,
		Update: resourceSepaInstantAssociationUpdate,

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
			"business_user_dn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"transport_profile_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bic": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"simulator_only": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  false,
			},
			"sponsor_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "",
			},
			"disable_outbound_payments": {
				Type:     schema.TypeBool,
				Required: false,
				ForceNew: false,
				Default:  false,
			},
		},
	}
}

func resourceSepaInstantAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createSepaInstantNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create sepa instant association: %s", err)
	}

	createdAssociation, err := client.AssociationClient.Associations.PostSepainstant(associations.NewPostSepainstantParams().
		WithCreationRequest(&models.SepaInstantAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create sepa instant association: %s", err)
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] sepa instant association key: %s", d.Id())

	return resourceSepaInstantAssociationRead(d, meta)
}

func resourceSepaInstantAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	sepaInstantAssociation, err := client.AssociationClient.Associations.GetSepainstantID(associations.NewGetSepainstantIDParams().
		WithID(associationId))

	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find sepa instant association: %s", err)
	}

	d.Set("association_id", sepaInstantAssociation.Payload.Data.ID.String())
	d.Set("business_user_dn", sepaInstantAssociation.Payload.Data.Attributes.BusinessUserDn)
	d.Set("transport_profile_id", sepaInstantAssociation.Payload.Data.Attributes.TransportProfileID)
	d.Set("bic", sepaInstantAssociation.Payload.Data.Attributes.Bic)
	d.Set("simulator_only", sepaInstantAssociation.Payload.Data.Attributes.SimulatorOnly)
	if sepaInstantAssociation.Payload.Data.Relationships == nil {
		d.Set("sponsor_id", "")
	} else {
		d.Set("sponsor_id", sepaInstantAssociation.Payload.Data.Relationships.Sponsor.Data.ID.String())
	}
	return nil
}

func resourceSepaInstantAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	sepaInstantAssociation, err := client.AssociationClient.Associations.GetSepainstantID(associations.NewGetSepainstantIDParams().
		WithID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting sepa instant association: %s", err)
	}

	_, err = client.AssociationClient.Associations.DeleteSepainstantID(associations.NewDeleteSepainstantIDParams().
		WithID(sepaInstantAssociation.Payload.Data.ID).
		WithVersion(*sepaInstantAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting sepa instant association: %s", err)
	}

	return nil
}

func resourceSepaInstantAssociationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createSepaInstantNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create sepa instant association: %s", err)
	}

	return nil
}

func createSepaInstantNewAssociationFromResourceData(d *schema.ResourceData) (*models.NewSepaInstantAssociation, error) {

	association := models.NewSepaInstantAssociation{Attributes: &models.SepaInstantAssociationAttributes{}}

	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("bic"); ok {
		association.Attributes.Bic = attr.(string)
	}

	if attr, ok := d.GetOk("business_user_dn"); ok {
		association.Attributes.BusinessUserDn = attr.(string)
	}

	if attr, ok := d.GetOk("transport_profile_id"); ok {
		association.Attributes.TransportProfileID = attr.(string)
	}

	if attr, ok := d.GetOk("simulator_only"); ok {
		b := attr.(bool)
		association.Attributes.SimulatorOnly = &b
	}

	if attr, ok := d.GetOk("disable_outbound_payments"); ok {
		b := attr.(bool)
		association.Attributes.DisableOutboundPayments = &b
	}

	if attr, ok := GetUUIDOK(d, "sponsor_id"); ok {
		association.Relationships = &models.SepaInstantAssociationRelationships{
			Sponsor: models.SepaInstantAssociationRelationshipsSponsor{
				Data: models.SepaInstantAssociationReference{
					ID:   attr,
					Type: "sepainstant_associations",
				},
			},
		}
	}
	return &association, nil
}
