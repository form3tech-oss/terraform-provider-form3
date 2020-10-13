package form3

import (
	"fmt"
	"log"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/form3tech-oss/terraform-provider-form3/client/associations"

	"github.com/form3tech-oss/terraform-provider-form3/models"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceForm3EburyAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceEburyAssociationCreate,
		Read:   resourceEburyAssociationRead,
		Update: resourceEburyAssociationUpdate,
		Delete: resourceEburyAssociationDelete,
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
			"organisation_location": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"funding_currency": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"ebury_contact_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"ebury_client_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"party_payment_fee": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"organisation_payment_fee": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"organisation_kyc_model": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"party_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"party_address": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: false,
			},
			"party_city": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"party_post_code": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
		},
		SchemaVersion: 0,
	}
}

func resourceEburyAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := createEburyAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create ebury association: %s", form3.JsonErrorPrettyPrint(err))
	}

	createdAssociation, err := client.AssociationClient.Associations.PostEbury(
		associations.NewPostEburyParams().WithCreationRequest(&models.EburyAssociationCreation{Data: association}))
	if err != nil {
		return fmt.Errorf("failed to create ebury association: %s", form3.JsonErrorPrettyPrint(err))
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] ebury association key: %s", d.Id())

	return resourceEburyAssociationRead(d, meta)
}

func resourceEburyAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	eburyAssociation, err := client.AssociationClient.Associations.GetEburyID(
		associations.NewGetEburyIDParams().WithID(associationId))

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find ebury gateway association: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	_ = d.Set("organisation_location", eburyAssociation.Payload.Data.Attributes.OrganisationLocation)
	_ = d.Set("ebury_client_id", eburyAssociation.Payload.Data.Attributes.EburyClientID)
	_ = d.Set("ebury_contact_id", eburyAssociation.Payload.Data.Attributes.EburyContactID)
	_ = d.Set("association_id", eburyAssociation.Payload.Data.ID.String())
	_ = d.Set("party_payment_fee", eburyAssociation.Payload.Data.Attributes.PartyPaymentFee)
	_ = d.Set("organisation_payment_fee", eburyAssociation.Payload.Data.Attributes.OrganisationPaymentFee)
	_ = d.Set("organisation_kyc_model", eburyAssociation.Payload.Data.Attributes.OrganisationKycModel)
	_ = d.Set("party_name", eburyAssociation.Payload.Data.Attributes.PartyName)
	_ = d.Set("party_address", eburyAssociation.Payload.Data.Attributes.PartyAddress)
	_ = d.Set("party_city", eburyAssociation.Payload.Data.Attributes.PartyCity)
	_ = d.Set("party_post_code", eburyAssociation.Payload.Data.Attributes.PartyPostCode)

	return nil
}

func resourceEburyAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	eburyAssociation, err := client.AssociationClient.Associations.GetEburyID(associations.NewGetEburyIDParams().
		WithID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting ebury gateway association: %s", form3.JsonErrorPrettyPrint(err))
	}

	_, err = client.AssociationClient.Associations.DeleteEburyID(associations.NewDeleteEburyIDParams().
		WithID(eburyAssociation.Payload.Data.ID).
		WithVersion(*eburyAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting ebury gateway association: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createEburyAssociationFromResourceData(d *schema.ResourceData) (*models.NewEburyAssociation, error) {
	association := models.NewEburyAssociation{
		Type:       string(models.ResourceTypeEburyAssociations),
		Attributes: &models.EburyAssociationAttributes{},
	}

	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("organisation_location"); ok {
		oloc := swag.String(attr.(string))
		association.Attributes.OrganisationLocation = oloc
	}

	if attr, ok := d.GetOk("funding_currency"); ok {
		bc := swag.String(attr.(string))
		association.Attributes.FundingCurrency = bc
	}

	if attr, ok := d.GetOk("ebury_contact_id"); ok {
		ct_id := swag.String(attr.(string))
		association.Attributes.EburyContactID = ct_id
	}

	if attr, ok := d.GetOk("ebury_client_id"); ok {
		cl_id := swag.String(attr.(string))
		association.Attributes.EburyClientID = cl_id
	}

	if attr, ok := d.GetOk("party_payment_fee"); ok {
		association.Attributes.PartyPaymentFee = attr.(string)
	}

	if attr, ok := d.GetOk("organisation_payment_fee"); ok {
		association.Attributes.OrganisationPaymentFee = attr.(string)
	}

	if attr, ok := d.GetOk("organisation_kyc_model"); ok {
		association.Attributes.OrganisationKycModel = attr.(string)
	}

	if attr, ok := d.GetOk("party_name"); ok {
		association.Attributes.PartyName = attr.(string)
	}

	if attr, ok := d.GetOk("party_address"); ok {
		coll := attr.([]interface{})
		for _, v := range coll {
			association.Attributes.PartyAddress = append(association.Attributes.PartyAddress, v.(string))
		}
	}

	if attr, ok := d.GetOk("party_city"); ok {
		association.Attributes.PartyCity = attr.(string)
	}

	if attr, ok := d.GetOk("party_post_code"); ok {
		association.Attributes.PartyPostCode = attr.(string)
	}

	return &association, nil
}

func resourceEburyAssociationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	updated, err := createEburyUpdateAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create ebury association: %s", form3.JsonErrorPrettyPrint(err))
	}

	current, err := client.AssociationClient.Associations.GetEburyID(
		associations.NewGetEburyIDParams().WithID(updated.ID))
	if err != nil {
		return fmt.Errorf("failed to fetch ebury association for ID: %s, %s", updated.ID, form3.JsonErrorPrettyPrint(err))
	} else if current == nil || current.Payload == nil {
		return fmt.Errorf("ebury association nil for ID: %s", updated.ID)
	}

	updated.Version = current.Payload.Data.Version

	_, err = client.AssociationClient.Associations.PatchEburyID(
		associations.NewPatchEburyIDParams().
			WithID(updated.ID).
			WithPatchBody(&models.EburyAssociationAmendment{
				Data: updated,
			}))
	if err != nil {
		return fmt.Errorf("failed to update ebury association: %s", form3.JsonErrorPrettyPrint(err))
	}

	log.Printf("[INFO] updated ebury association key: %s", d.Id())
	return resourceEburyAssociationRead(d, meta)
}

func createEburyUpdateAssociationFromResourceData(d *schema.ResourceData) (*models.EburyAssociationUpdate, error) {
	association := models.EburyAssociationUpdate{
		Type:       string(models.ResourceTypeEburyAssociations),
		Attributes: &models.EburyAssociationPatchAttributes{},
	}

	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("organisation_location"); ok {
		oloc := swag.String(attr.(string))
		association.Attributes.OrganisationLocation = oloc
	}

	if attr, ok := d.GetOk("funding_currency"); ok {
		bc := swag.String(attr.(string))
		association.Attributes.FundingCurrency = bc
	}

	if attr, ok := d.GetOk("ebury_contact_id"); ok {
		ct_id := swag.String(attr.(string))
		association.Attributes.EburyContactID = ct_id
	}

	if attr, ok := d.GetOk("ebury_client_id"); ok {
		cl_id := swag.String(attr.(string))
		association.Attributes.EburyClientID = cl_id
	}

	if attr, ok := d.GetOk("party_payment_fee"); ok {
		ppf := swag.String(attr.(string))
		association.Attributes.PartyPaymentFee = ppf
	}

	if attr, ok := d.GetOk("organisation_payment_fee"); ok {
		opf := swag.String(attr.(string))
		association.Attributes.OrganisationPaymentFee = opf
	}

	if attr, ok := d.GetOk("organisation_kyc_model"); ok {
		okm := swag.String(attr.(string))
		association.Attributes.OrganisationKycModel = okm
	}

	if attr, ok := d.GetOk("party_name"); ok {
		pn := swag.String(attr.(string))
		association.Attributes.PartyName = pn
	}

	if attr, ok := d.GetOk("party_address"); ok {
		coll := attr.([]interface{})
		for _, v := range coll {
			association.Attributes.PartyAddress = append(association.Attributes.PartyAddress, v.(string))
		}
	}

	if attr, ok := d.GetOk("party_city"); ok {
		pc := swag.String(attr.(string))
		association.Attributes.PartyCity = pc
	}

	if attr, ok := d.GetOk("party_post_code"); ok {
		ppc := swag.String(attr.(string))
		association.Attributes.PartyPostCode = ppc
	}

	return &association, nil
}
