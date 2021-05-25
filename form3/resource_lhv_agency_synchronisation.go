package form3

import (
	"errors"
	"fmt"
	"log"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceForm3LhvAgencySynchronisation() *schema.Resource {
	return &schema.Resource{
		Create: resourceLhvAgencySynchronisationCreate,
		Read:   resourceLhvAgencySynchronisationRead,
		Delete: resourceLhvAgencySynchronisationDelete,
		Schema: map[string]*schema.Schema{
			"agency_synchronisation_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
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
			"country": {
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
				Optional: true,
				ForceNew: true,
			},
			"bank_id": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceLhvAgencySynchronisationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, found := GetUUIDOK(d, "association_id")
	if !found {
		return errors.New("failed to read lhv agency synchronisation association_id")
	}

	agencyResource, err := createLhvNewAgencySynchronisationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create lhv agency synchronisation: %s", form3.JsonErrorPrettyPrint(err))
	}

	createdAgency, err := client.AssociationClient.Associations.PostLhvAssociationIDAgencySynchronisations(
		associations.NewPostLhvAssociationIDAgencySynchronisationsParams().
			WithAssociationID(associationId).
			WithCreationRequest(&models.LhvAgencySynchronisationCreation{
				Data: agencyResource,
			}))
	if err != nil {
		return fmt.Errorf("failed to create lhv agency synchronisation: %s", form3.JsonErrorPrettyPrint(err))
	}

	d.SetId(createdAgency.Payload.Data.ID.String())
	log.Printf("[INFO] lhv agency synchronisation id: %s", d.Id())

	return resourceLhvAgencySynchronisationRead(d, meta)
}

func resourceLhvAgencySynchronisationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, found := GetUUIDOK(d, "association_id")
	if !found {
		return errors.New("failed to read lhv agency synchronisation association_id")
	}

	agencySynchronisation, err := client.AssociationClient.Associations.GetLhvAssociationIDAgencySynchronisationsAgencySynchronisationID(
		associations.NewGetLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams().
			WithAssociationID(associationId).
			WithAgencySynchronisationID(strfmt.UUID(d.Id())))

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find lhv agency sychronisation association: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	_ = d.Set("country", agencySynchronisation.Payload.Data.Attributes.Country)
	_ = d.Set("bic", agencySynchronisation.Payload.Data.Attributes.Bic)
	_ = d.Set("bank_id", agencySynchronisation.Payload.Data.Attributes.BankID)
	if agencySynchronisation.Payload.Data.Attributes.Iban != "" {
		_ = d.Set("iban", agencySynchronisation.Payload.Data.Attributes.Iban)
	}
	return nil
}

func resourceLhvAgencySynchronisationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, found := GetUUIDOK(d, "association_id")
	if !found {
		return errors.New("failed to read lhv agency synchronisation association_id")
	}

	agencySynchronisation, err := client.AssociationClient.Associations.GetLhvAssociationIDAgencySynchronisationsAgencySynchronisationID(
		associations.NewGetLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams().
			WithAssociationID(associationId).
			WithAgencySynchronisationID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting lhv agency synchronisation: %s", form3.JsonErrorPrettyPrint(err))
	}

	_, err = client.AssociationClient.Associations.DeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationID(
		associations.NewDeleteLhvAssociationIDAgencySynchronisationsAgencySynchronisationIDParams().
			WithAssociationID(associationId).
			WithAgencySynchronisationID(agencySynchronisation.Payload.Data.ID).
			WithVersion(*agencySynchronisation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting lhv agency sychronisation association: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createLhvNewAgencySynchronisationFromResourceData(d *schema.ResourceData) (*models.LhvAgencySynchronisation, error) {
	var zero int64
	agencySynchronisation := models.LhvAgencySynchronisation{
		Type:    models.LhvAgencySynchronisationTypeLhvgatewayAgencySynchronisations,
		Version: &zero,
	}

	if attr, ok := GetUUIDOK(d, "agency_synchronisation_id"); ok {
		agencySynchronisation.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		agencySynchronisation.OrganisationID = attr
	}

	if attr, ok := d.GetOk("country"); ok {
		agencySynchronisation.Attributes.Country = attr.(string)
	}

	if attr, ok := d.GetOk("bic"); ok {
		agencySynchronisation.Attributes.Bic = attr.(string)
	}

	if attr, ok := d.GetOk("iban"); ok {
		agencySynchronisation.Attributes.Iban = attr.(string)
	}

	if attr, ok := d.GetOk("bank_id"); ok {
		arr := attr.([]interface{})
		for _, v := range arr {
			agencySynchronisation.Attributes.BankID = append(agencySynchronisation.Attributes.BankID, v.(string))
		}
	}

	return &agencySynchronisation, nil
}
