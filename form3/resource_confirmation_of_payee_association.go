package form3

import (
	"fmt"
	"log"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceForm3ConfirmationOfPayeeAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfirmationOfPayeeAssociationCreate,
		Read:   resourceConfirmationOfPayeeAssociationRead,
		Delete: resourceConfirmationOfPayeeAssociationDelete,
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

			"open_banking_organisation_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"open_banking_public_key_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"signing_key_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"signing_dn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"signing_certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"exact_match_threshold": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"close_match_threshold": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceConfirmationOfPayeeAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	log.Print("[INFO] Creating ConfirmationOfPayee association ")

	association, err := createConfirmationOfPayeeNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create ConfirmationOfPayee association: %s", form3.JsonErrorPrettyPrint(err))
	}

	createdAssociation, err := client.AssociationClient.Associations.PostConfirmationOfPayee(associations.NewPostConfirmationOfPayeeParams().
		WithCreationRequest(&models.CoPAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create ConfirmationOfPayee association: %s", form3.JsonErrorPrettyPrint(err))
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] ConfirmationOfPayee association key: %s", d.Id())

	return resourceConfirmationOfPayeeAssociationRead(d, meta)
}

func resourceConfirmationOfPayeeAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	if associationId == "" {
		associationId = strfmt.UUID(d.Id())
		log.Printf("[INFO] Importing association with resource id: %s.", associationId)
	} else {
		log.Printf("[INFO] Reading association with resource id: %s.", associationId)
	}

	association, err := client.AssociationClient.Associations.GetConfirmationOfPayeeID(associations.NewGetConfirmationOfPayeeIDParams().
		WithID(associationId))

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find confirmation of payee association: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	_ = d.Set("association_id", association.Payload.Data.ID.String())
	_ = d.Set("organisation_id", association.Payload.Data.OrganisationID.String())
	_ = d.Set("open_banking_organisation_id", association.Payload.Data.Attributes.OpenBankingOrganisationID)
	_ = d.Set("open_banking_public_key_id", association.Payload.Data.Attributes.PublicKeyID)
	_ = d.Set("signing_key_id", association.Payload.Data.Relationships.SigningCertificate.Data.KeyID)
	_ = d.Set("signing_dn", association.Payload.Data.Relationships.SigningCertificate.Data.Dn)
	_ = d.Set("signing_certificate_id", association.Payload.Data.Relationships.SigningCertificate.Data.ID)
	_ = d.Set("exact_match_threshold", association.Payload.Data.Attributes.MatchingCriteria.ExactMatchThreshold)
	_ = d.Set("close_match_threshold", association.Payload.Data.Attributes.MatchingCriteria.CloseMatchThreshold)

	return nil
}

func resourceConfirmationOfPayeeAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	association, err := client.AssociationClient.Associations.GetConfirmationOfPayeeID(associations.NewGetConfirmationOfPayeeIDParams().
		WithID(strfmt.UUID(d.Id())))

	if err != nil {
		return fmt.Errorf("error deleting ConfirmationOfPayee association: %s", form3.JsonErrorPrettyPrint(err))
	}

	log.Printf("[INFO] Deleting ConfirmationOfPayee association for id: %s ", association.Payload.Data.ID)

	params := associations.NewDeleteConfirmationOfPayeeIDParams().
		WithID(*association.Payload.Data.ID)

	if association.Payload.Data.Version != nil {
		params.WithVersion(*association.Payload.Data.Version)
	}

	_, err = client.AssociationClient.Associations.DeleteConfirmationOfPayeeID(params)

	if err != nil {
		return fmt.Errorf("error deleting ConfirmationOfPayee association: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createConfirmationOfPayeeNewAssociationFromResourceData(d *schema.ResourceData) (*models.CoPAssociation, error) {

	association := &models.CoPAssociation{
		Type: "confirmation_of_payee_associations",
		Relationships: &models.CoPAssociationRelationships{
			SigningCertificate: &models.SigningCertificate{
				Data: &models.SigningCertificateData{
					Type: "certificates",
				},
			},
		},
		Attributes: &models.CoPAssociationAttributes{},
	}

	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		uuid := strfmt.UUID(attr.String())
		association.ID = &uuid
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		uuid := strfmt.UUID(attr.String())
		association.OrganisationID = &uuid
	}

	if attr, ok := d.GetOk("open_banking_organisation_id"); ok {
		association.Attributes.OpenBankingOrganisationID = swag.String(attr.(string))
	}

	if attr, ok := d.GetOk("open_banking_public_key_id"); ok {
		association.Attributes.PublicKeyID = swag.String(attr.(string))
	}

	if attr, ok := GetUUIDOK(d, "signing_key_id"); ok {
		uuid := strfmt.UUID(attr.String())
		association.Relationships.SigningCertificate.Data.KeyID = &uuid
	}

	if attr, ok := GetUUIDOK(d, "signing_certificate_id"); ok {
		uuid := strfmt.UUID(attr.String())
		association.Relationships.SigningCertificate.Data.ID = &uuid
	}

	if attr, ok := d.GetOk("signing_dn"); ok {
		association.Relationships.SigningCertificate.Data.Dn = swag.String(attr.(string))
	}

	association.Attributes.MatchingCriteria = &models.MatchingCriteria{}
	if attr, ok := d.GetOk("exact_match_threshold"); ok {
		association.Attributes.MatchingCriteria.ExactMatchThreshold = swag.String(attr.(string))
	}

	if attr, ok := d.GetOk("close_match_threshold"); ok {
		association.Attributes.MatchingCriteria.CloseMatchThreshold = swag.String(attr.(string))
	}

	return association, nil
}
