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

func resourceForm3BacsAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceBacsAssociationCreate,
		Read:   resourceBacsAssociationRead,
		Delete: resourceBacsAssociationDelete,
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
			"service_user_number": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_number": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sorting_code": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_type": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"bank_code": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"centre_number": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"input_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"input_certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"output_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"output_certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"messaging_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"messaging_certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"use_test_file_submission": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceBacsAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	serviceUserNumber := d.Get("service_user_number").(string)
	log.Printf("[INFO] Creating Bacs association with service user number: %s", serviceUserNumber)

	association, err := createBacsNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create Bacs association: %s", err)
	}

	createdAssociation, err := client.AssociationClient.Associations.PostBacs(associations.NewPostBacsParams().
		WithCreationRequest(&models.BacsAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create Bacs association: %s", err)
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] Bacs association key: %s", d.Id())

	return resourceBacsAssociationRead(d, meta)
}

func resourceBacsAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	if associationId == "" {
		associationId = strfmt.UUID(d.Id())
		log.Printf("[INFO] Importing bacs association with resource id: %s.", associationId)
	} else {
		log.Printf("[INFO] Reading bacs association with resource id: %s.", associationId)
	}

	bacsAssociation, err := client.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().
		WithID(associationId))

	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		} else {
			return err
		}
	}

	d.Set("association_id", bacsAssociation.Payload.Data.ID.String())
	d.Set("organisation_id", bacsAssociation.Payload.Data.OrganisationID.String())
	d.Set("service_user_number", bacsAssociation.Payload.Data.Attributes.ServiceUserNumber)
	d.Set("account_number", bacsAssociation.Payload.Data.Attributes.AccountNumber)
	d.Set("sorting_code", bacsAssociation.Payload.Data.Attributes.SortingCode)
	d.Set("account_type", bacsAssociation.Payload.Data.Attributes.AccountType)
	d.Set("bank_code", bacsAssociation.Payload.Data.Attributes.BankCode)
	d.Set("centre_number", bacsAssociation.Payload.Data.Attributes.CentreNumber)
	d.Set("use_test_file_submission", bacsAssociation.Payload.Data.Attributes.UseTestFileSubmission)

	if bacsAssociation.Payload.Data.Relationships != nil {
		if bacsAssociation.Payload.Data.Relationships.InputCertificate != nil && bacsAssociation.Payload.Data.Relationships.InputCertificate.Data != nil {
			d.Set("input_key_id", bacsAssociation.Payload.Data.Relationships.InputCertificate.Data.KeyID)
			d.Set("input_certificate_id", bacsAssociation.Payload.Data.Relationships.InputCertificate.Data.CertificateID)
		}
		if bacsAssociation.Payload.Data.Relationships.OutputCertificate != nil && bacsAssociation.Payload.Data.Relationships.OutputCertificate.Data != nil {
			d.Set("output_key_id", bacsAssociation.Payload.Data.Relationships.OutputCertificate.Data.KeyID)
			d.Set("output_certificate_id", bacsAssociation.Payload.Data.Relationships.OutputCertificate.Data.CertificateID)
		}
		if bacsAssociation.Payload.Data.Relationships.MessagingCertificate != nil && bacsAssociation.Payload.Data.Relationships.MessagingCertificate.Data != nil {
			d.Set("messaging_key_id", bacsAssociation.Payload.Data.Relationships.MessagingCertificate.Data.KeyID)
			d.Set("messaging_certificate_id", bacsAssociation.Payload.Data.Relationships.MessagingCertificate.Data.CertificateID)
		}
	}

	return nil
}

func resourceBacsAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	bacsAssociation, err := client.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().
		WithID(strfmt.UUID(d.Id())))

	if err != nil {
		return fmt.Errorf("error deleting Bacs association: %s", err)
	}

	log.Printf("[INFO] Deleting Bacs association for id: %s service user number: %s", bacsAssociation.Payload.Data.ID, bacsAssociation.Payload.Data.Attributes.ServiceUserNumber)

	_, err = client.AssociationClient.Associations.DeleteBacsID(associations.NewDeleteBacsIDParams().
		WithID(bacsAssociation.Payload.Data.ID).
		WithVersion(*bacsAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting Bacs association: %s", err)
	}

	return nil
}

func createBacsNewAssociationFromResourceData(d *schema.ResourceData) (*models.BacsNewAssociation, error) {

	association := models.BacsNewAssociation{
		Attributes:    &models.BacsAssociationAttributes{},
		Relationships: &models.BacsAssociationRelationships{},
	}

	association.Type = "associations"
	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		association.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		association.OrganisationID = attr
	}

	if attr, ok := d.GetOk("service_user_number"); ok {
		association.Attributes.ServiceUserNumber = attr.(string)
	}

	if attr, ok := d.GetOk("account_number"); ok {
		association.Attributes.AccountNumber = attr.(string)
	}

	if attr, ok := d.GetOk("sorting_code"); ok {
		association.Attributes.SortingCode = attr.(string)
	}

	if attr, ok := d.GetOkExists("account_type"); ok {
		accountType := int64(attr.(int))
		association.Attributes.AccountType = &accountType
	}

	if attr, ok := d.GetOk("bank_code"); ok {
		association.Attributes.BankCode = attr.(string)
	}

	if attr, ok := d.GetOk("use_test_file_submission"); ok {
		association.Attributes.UseTestFileSubmission = attr.(bool)
	}

	association.Relationships.InputCertificate = buildRelationship(d, "input")
	association.Relationships.OutputCertificate = buildRelationship(d, "output")
	association.Relationships.MessagingCertificate = buildRelationship(d, "messaging")

	return &association, nil
}

func buildRelationship(d *schema.ResourceData, relation string) *models.BacsAssociationCertificateRelationship {
	if keyId, ok := GetUUIDOK(d, relation+"_key_id"); ok {
		if certId, certOk := GetUUIDOK(d, relation+"_certificate_id"); certOk {
			return &models.BacsAssociationCertificateRelationship{
				Data: &models.BacsAssociationCertificateRelationshipData{KeyID: keyId, CertificateID: certId},
			}
		}
		return &models.BacsAssociationCertificateRelationship{
			Data: &models.BacsAssociationCertificateRelationshipData{KeyID: keyId},
		}
	}
	return nil
}
