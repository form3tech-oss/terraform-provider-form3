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

func resourceForm3SepaInstantAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSepaInstantAssociationCreate,
		Read:   resourceSepaInstantAssociationRead,
		Delete: resourceSepaInstantAssociationDelete,
		Update: resourceSepaInstantAssociationUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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
				ForceNew: false,
			},
			"transport_profile_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"bic": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"simulator_only": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  false,
			},
			"sponsor_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
				Default:  "",
			},
			"disable_outbound_payments": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  false,
			},
			"reachable_bics": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				ForceNew: false,
			},
		},
	}
}

func resourceSepaInstantAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createSepaInstantNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create sepa instant association from resource data: %s", form3.JsonErrorPrettyPrint(err))
	}

	createdAssociation, err := client.AssociationClient.Associations.PostSepainstant(associations.NewPostSepainstantParams().
		WithCreationRequest(&models.SepaInstantAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("error when posting new sepa instant association resource: %s", form3.JsonErrorPrettyPrint(err))
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
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find sepa instant association: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
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
	if reachableBics := sepaInstantAssociation.Payload.Data.Attributes.ReachableBics; reachableBics != nil {
		d.Set("reachable_bics", reachableBics)
	}
	return nil
}

func resourceSepaInstantAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	sepaInstantAssociation, err := client.AssociationClient.Associations.GetSepainstantID(associations.NewGetSepainstantIDParams().
		WithID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting sepa instant association: %s", form3.JsonErrorPrettyPrint(err))
	}

	_, err = client.AssociationClient.Associations.DeleteSepainstantID(associations.NewDeleteSepainstantIDParams().
		WithID(sepaInstantAssociation.Payload.Data.ID).
		WithVersion(*sepaInstantAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting sepa instant association: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func resourceSepaInstantAssociationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createSepaInstantUpdateAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create an updated sepa instant association resource: %s", form3.JsonErrorPrettyPrint(err))
	}

	existingAssociation, err := client.AssociationClient.Associations.GetSepainstantID(
		associations.NewGetSepainstantIDParams().WithID(association.ID))

	if err != nil {
		return fmt.Errorf("could not get sepa instant association with id: %s", association.ID.String())
	}

	if existingAssociation == nil || existingAssociation.Payload == nil {
		return fmt.Errorf("sepa instant association with id %s is nil or has a nil payload", association.ID)
	}

	_, err = client.AssociationClient.Associations.PatchSepainstantID(
		associations.NewPatchSepainstantIDParams().
			WithVersion(*existingAssociation.Payload.Data.Version).
			WithID(association.ID).WithPayload(&models.SepaInstantAssociationPatch{
			Data: &models.UpdateSepaInstantAssociation{
				ID:             association.ID,
				OrganisationID: association.OrganisationID,
				Type:           models.SepaInstantAssociationReferenceTypeSepainstantAssociations,
				Attributes: &models.UpdateSepaInstantAssociationAttributes{
					Bic:                     association.Attributes.Bic,
					TransportProfileID:      association.Attributes.TransportProfileID,
					BusinessUserDn:          association.Attributes.BusinessUserDn,
					DisableOutboundPayments: association.Attributes.DisableOutboundPayments,
					SimulatorOnly:           association.Attributes.SimulatorOnly,
					ReachableBics:           association.Attributes.ReachableBics,
				},
			},
		}))
	if err != nil {
		return fmt.Errorf("failed to patch sepa instant association: %s", form3.JsonErrorPrettyPrint(err))
	}

	log.Printf("[INFO] sepa instant association key: #{d.ID}")

	return nil
}

func createSepaInstantUpdateAssociationFromResourceData(d *schema.ResourceData) (*models.UpdateSepaInstantAssociation, error) {
	association := models.UpdateSepaInstantAssociation{Attributes: &models.UpdateSepaInstantAssociationAttributes{}}

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

	if attr, ok := d.GetOk("disable_outbound_payments"); ok {
		b := attr.(bool)
		association.Attributes.DisableOutboundPayments = &b
	}

	if attr, ok := d.GetOk("simulator_only"); ok {
		b := attr.(bool)
		association.Attributes.SimulatorOnly = &b
	}

	if attr, ok := d.GetOk("reachable_bics"); ok {
		rawList := attr.([]interface{})
		bicList := make([]string, 0, len(rawList))
		for _, e := range rawList {
			bicList = append(bicList, e.(string))
		}

		association.Attributes.ReachableBics = bicList
	}

	return &association, nil
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

	if attr, ok := d.GetOk("reachable_bics"); ok {
		rawList := attr.([]interface{})
		bicList := make([]string, 0, len(rawList))
		for _, e := range rawList {
			bicList = append(bicList, e.(string))
		}

		association.Attributes.ReachableBics = bicList
	}

	return &association, nil
}
