package form3

import (
	"fmt"
	"log"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
)

func resourceForm3SepaSctAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceSepaSctAssociationCreate,
		Read:   resourceSepaSctAssociationRead,
		Delete: resourceSepaSctAssociationDelete,

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
				Optional: true,
				ForceNew: true,
			},
			"business_user": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"receiver_business_user": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"can_sponsor": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"sponsor_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"reachable_bics": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSepaSctAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createSepaSctNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create sepa sct association: %s", form3.JsonErrorPrettyPrint(err))
	}

	createdAssociation, err := client.AssociationClient.Associations.PostSepasct(associations.NewPostSepasctParams().
		WithCreationRequest(&models.SepaSctAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create sepa sct association: %s", form3.JsonErrorPrettyPrint(err))
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] sepa sct association key: %s", d.Id())

	return resourceSepaSctAssociationRead(d, meta)
}

func resourceSepaSctAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	sepaSctAssociation, err := client.AssociationClient.Associations.GetSepasctID(associations.NewGetSepasctIDParams().
		WithID(associationId))

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find sepa sct association: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	_ = d.Set("association_id", sepaSctAssociation.Payload.Data.ID.String())
	_ = d.Set("bic", sepaSctAssociation.Payload.Data.Attributes.Bic)
	_ = d.Set("business_user", sepaSctAssociation.Payload.Data.Attributes.BusinessUser)
	_ = d.Set("receiver_business_user", sepaSctAssociation.Payload.Data.Attributes.ReceiverBusinessUser)

	return nil
}

func resourceSepaSctAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	sepaSctAssociation, err := client.AssociationClient.Associations.GetSepasctID(associations.NewGetSepasctIDParams().
		WithID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting sepa sct association: %s", form3.JsonErrorPrettyPrint(err))
	}

	_, err = client.AssociationClient.Associations.DeleteSepasctID(associations.NewDeleteSepasctIDParams().
		WithID(sepaSctAssociation.Payload.Data.ID).
		WithVersion(*sepaSctAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting sepa sct association: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createSepaSctNewAssociationFromResourceData(d *schema.ResourceData) (*models.NewSepaSctAssociation, error) {

	association := models.NewSepaSctAssociation{Attributes: &models.SepaSctAssociationAttributes{}}
	association.Type = "sepasct_associations"
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

	if attr, ok := d.GetOk("reachable_bics"); ok {
		rawList := attr.([]interface{})
		bicList := make([]string, 0, len(rawList))
		for _, e := range rawList {
			bicList = append(bicList, e.(string))
		}

		association.Attributes.ReachableBics = bicList
	}

	if attr, ok := d.GetOk("can_sponsor"); ok {
		association.Attributes.CanSponsor = attr.(bool)
	}

	if attr, ok := d.GetOk("sponsor_id"); ok {
		association.Relationships = &models.SepaSctAssociationRelationships{
			Sponsor: &models.SepaSctAssociationRelationshipsSponsor{
				Data: &models.SepaSctSponsorAssociationReference{
					ID:   attr.(string),
					Type: association.Type,
				},
			},
		}
	}

	return &association, nil
}
