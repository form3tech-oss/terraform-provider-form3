package form3

import (
	"fmt"
	"github.com/ewilde/go-form3"
	"github.com/ewilde/go-form3/client/associations"
	"github.com/ewilde/go-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3PayportAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourcePayportAssociationCreate,
		Read:   resourcePayportAssociationRead,
		Delete: resourcePayportAssociationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"payport_association_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"participant_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"customer_sending_fps_institution": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sponsor_bank_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sponsor_account_number": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePayportAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	participantId := d.Get("participant_id").(string)
	log.Printf("[INFO] Creating payport association with participant id: %s", participantId)

	association, err := createPayportAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create payport association: %s", err)
	}

	createdAssociation, err := client.AssociationClient.Associations.PostPayport(associations.NewPostPayportParams().
		WithCreationRequest(&models.PayportAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create payport association: %s", err)
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] association key: %s", d.Id())

	return resourcePayportAssociationRead(d, meta)
}

func resourcePayportAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	key := d.Id()
	associationId, _ := GetUUIDOK(d, "payport_association_id")
	participantId := d.Get("participant_id").(string)

	if associationId == "" {
		associationId = strfmt.UUID(key)
		log.Printf("[INFO] Importing payport association for id: %s participant id: %s", associationId, participantId)
	} else {
		log.Printf("[INFO] Reading payport association for id: %s participant id: %s", associationId, participantId)
	}

	association, err := client.AssociationClient.Associations.GetPayportID(associations.NewGetPayportIDParams().
		WithID(associationId))

	if err != nil {
		apiError := err.(*runtime.APIError)
		if apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find payport association: %s", err)
	}

	d.Set("payport_association_id", association.Payload.Data.ID.String())
	d.Set("participant_id", association.Payload.Data.Attributes.ParticipantID)
	d.Set("customer_sending_fps_institution", association.Payload.Data.Attributes.CustomerSendingFpsInstitution)
	d.Set("sponsor_bank_id", association.Payload.Data.Attributes.SponsorBankID)
	d.Set("sponsor_account_number", association.Payload.Data.Attributes.SponsorAccountNumber)
	d.Set("organisation_id", association.Payload.Data.OrganisationID.String())
	return nil
}

func resourcePayportAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationFromResource, err := createPayportAssociationFromResourceDataWithVersion(d, client)
	if err != nil {
		return fmt.Errorf("error deleting payport association: %s", err)
	}

	log.Printf("[INFO] Deleting payport association for id: %s participant id: %s", associationFromResource.ID, associationFromResource.Attributes.ParticipantID)

	_, err = client.AssociationClient.Associations.DeletePayportID(associations.NewDeletePayportIDParams().
		WithID(associationFromResource.ID).
		WithVersion(*associationFromResource.Version))

	if err != nil {
		return fmt.Errorf("error deleting payport association: %s", err)
	}

	return nil
}

func createPayportAssociationFromResourceDataWithVersion(d *schema.ResourceData, client *form3.AuthenticatedClient) (*models.PayportAssociation, error) {
	association, err := createPayportAssociationFromResourceData(d)
	version, err := getPayportAssociationVersion(client, association.ID)
	if err != nil {
		return nil, err
	}

	association.Version = &version

	return association, nil
}

func createPayportAssociationFromResourceData(d *schema.ResourceData) (*models.PayportAssociation, error) {

	association := models.PayportAssociation{Attributes: &models.PayportAssociationAttributes{}}
	association.Type = "associations"
	if attr, ok := GetUUIDOK(d, "payport_association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("participant_id"); ok {
		association.Attributes.ParticipantID = attr.(string)
	}

	if attr, ok := d.GetOk("customer_sending_fps_institution"); ok {
		association.Attributes.CustomerSendingFpsInstitution = attr.(string)
	}

	if attr, ok := d.GetOk("sponsor_bank_id"); ok {
		association.Attributes.SponsorBankID = attr.(string)
	}

	if attr, ok := d.GetOk("sponsor_account_number"); ok {
		association.Attributes.SponsorAccountNumber = attr.(string)
	}

	return &association, nil
}

func getPayportAssociationVersion(client *form3.AuthenticatedClient, associationId strfmt.UUID) (int64, error) {
	association, err := client.AssociationClient.Associations.GetPayportID(associations.NewGetPayportIDParams().WithID(associationId))
	if err != nil {
		if err != nil {
			return -1, fmt.Errorf("error reading payport association: %s", err)
		}
	}

	return *association.Payload.Data.Version, nil
}
