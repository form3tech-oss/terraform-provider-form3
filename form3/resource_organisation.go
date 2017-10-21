package form3

import (
	"fmt"
	"github.com/ewilde/go-form3"
	"github.com/ewilde/go-form3/client/organisations"
	"github.com/ewilde/go-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3Organisation() *schema.Resource {
	return &schema.Resource{
		Create: resourceOrganisationCreate,
		Read:   resourceOrganisationRead,
		Update: resourceOrganisationUpdate,
		Delete: resourceOrganisationDelete,

		Schema: map[string]*schema.Schema{
			"organisation_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parent_organisation_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceOrganisationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	name := d.Get("name").(string)
	log.Printf("[INFO] Creating organisation with name: %s", name)

	organisation, err := createOrganisationFromResourceData(d)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] organisation create: %#v", organisation)

	createdOrganisation, err := client.OrganisationClient.Organisations.PostUnits(organisations.NewPostUnitsParams().
		WithOrganisationCreationRequest(&models.OrganisationCreation{Data: organisation}))

	if err != nil {
		return fmt.Errorf("failed to create organisation: %s", err)
	}

	d.SetId(createdOrganisation.Payload.Data.ID.String())
	log.Printf("[INFO] organisation key: %s", d.Id())

	return resourceOrganisationRead(d, meta)
}

func resourceOrganisationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	key := d.Id()
	organisationId, _ := GetUUIDOK(d, "organisation_id")
	organisationName := d.Get("name").(string)
	log.Printf("[INFO] Reading organisation for id: %s name: %s", key, organisationName)

	organisation, err := client.OrganisationClient.Organisations.GetUnitsID(
		organisations.NewGetUnitsIDParams().WithID(organisationId))

	if err != nil {
		apiError := err.(*runtime.APIError)
		if apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find organisation: %s", err)
	}

	d.Set("organisation_id", organisation.Payload.Data.ID.String())
	d.Set("parent_organisation_id", organisation.Payload.Data.OrganisationID.String())
	d.Set("name", organisation.Payload.Data.Attributes.Name)
	return nil
}

func resourceOrganisationUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(false)

	if d.HasChange("name") {
		client := meta.(*form3.AuthenticatedClient)
		organisationFromResource, err := createOrganisationFromResourceDataWithVersion(d, client)
		if err != nil {
			return fmt.Errorf("error updating organisation: %s", err)
		}

		_, err = client.OrganisationClient.Organisations.PatchUnitsID(organisations.NewPatchUnitsIDParams().
			WithID(organisationFromResource.ID).
			WithOrganisationUpdateRequest(&models.OrganisationUpdate{Data: organisationFromResource}))

		if err != nil {
			return fmt.Errorf("error updating organisation: %s", err)
		}
	}

	return nil
}

func resourceOrganisationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	organisationFromResource, err := createOrganisationFromResourceDataWithVersion(d, client)
	if err != nil {
		return fmt.Errorf("error deleting organisation: %s", err)
	}

	log.Printf("[INFO] Deleting organisation for id: %s organisationname: %s",
		organisationFromResource.ID, organisationFromResource.Attributes.Name)

	_, err = client.OrganisationClient.Organisations.DeleteUnitsID(organisations.NewDeleteUnitsIDParams().
		WithID(organisationFromResource.ID).
		WithVersion(*organisationFromResource.Version))

	if err != nil {
		return fmt.Errorf("error deleting organisation: %s", err)
	}

	return nil
}

func createOrganisationFromResourceDataWithVersion(d *schema.ResourceData, client *form3.AuthenticatedClient) (*models.Organisation, error) {
	organisation, err := createOrganisationFromResourceData(d)
	version, err := getOrganisationVersion(client, organisation.ID)
	if err != nil {
		return nil, err
	}

	organisation.Version = &version

	return organisation, nil
}

func createOrganisationFromResourceData(d *schema.ResourceData) (*models.Organisation, error) {

	organisation := models.Organisation{Attributes: &models.OrganisationAttributes{}}
	organisation.Type = "organisations"
	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		organisation.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "parent_organisation_id"); ok {
		organisation.OrganisationID = attr
	}

	if attr, ok := d.GetOk("name"); ok {
		organisation.Attributes.Name = attr.(string)
	}

	return &organisation, nil
}

func getOrganisationVersion(client *form3.AuthenticatedClient, organisationId strfmt.UUID) (int64, error) {
	organisation, err := client.OrganisationClient.Organisations.GetUnitsID(organisations.NewGetUnitsIDParams().
		WithID(organisationId))
	if err != nil {
		if err != nil {
			return -1, fmt.Errorf("error reading organisation: %s", err)
		}
	}

	return *organisation.Payload.Data.Version, nil
}
