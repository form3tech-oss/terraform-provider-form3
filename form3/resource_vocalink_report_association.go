package form3

import (
	"fmt"
	"github.com/form3tech-oss/go-form3"
	"github.com/form3tech-oss/go-form3/client/associations"
	"github.com/form3tech-oss/go-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3VocalinkReportAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceVocalinkReportAssociationCreate,
		Read:   resourceVocalinkReportAssociationRead,
		Delete: resourceVocalinkReportAssociationDelete,
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
			"fps_member_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"fps_member_certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"bacs_member_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"bacs_member_certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"bacs_service_user_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"bacs_service_user_certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceVocalinkReportAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	log.Print("[INFO] Creating VocalinkReport association ")

	association, err := createVocalinkReportNewAssociationFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create VocalinkReport association: %s", err)
	}

	createdAssociation, err := client.AssociationClient.Associations.PostVocalinkreport(associations.NewPostVocalinkreportParams().
		WithCreationRequest(&models.VocalinkReportAssociationCreation{
			Data: association,
		}))

	if err != nil {
		return fmt.Errorf("failed to create VocalinkReport association: %s", err)
	}

	d.SetId(createdAssociation.Payload.Data.ID.String())
	log.Printf("[INFO] VocalinkReport association key: %s", d.Id())

	return resourceVocalinkReportAssociationRead(d, meta)
}

func resourceVocalinkReportAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, _ := GetUUIDOK(d, "association_id")

	if associationId == "" {
		associationId = strfmt.UUID(d.Id())
		log.Printf("[INFO] Importing bacs association with resource id: %s.", associationId)
	} else {
		log.Printf("[INFO] Reading bacs association with resource id: %s.", associationId)
	}

	bacsAssociation, err := client.AssociationClient.Associations.GetVocalinkreportID(associations.NewGetVocalinkreportIDParams().
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

	if bacsAssociation.Payload.Data.Relationships != nil {
		if bacsAssociation.Payload.Data.Relationships.FpsMemberCertificate != nil && bacsAssociation.Payload.Data.Relationships.FpsMemberCertificate.Data != nil {
			d.Set("fps_member_key_id", bacsAssociation.Payload.Data.Relationships.FpsMemberCertificate.Data.KeyID)
			d.Set("fps_member_certificate_id", bacsAssociation.Payload.Data.Relationships.FpsMemberCertificate.Data.CertificateID)
		}
		if bacsAssociation.Payload.Data.Relationships.BacsMemberCertificate != nil && bacsAssociation.Payload.Data.Relationships.BacsMemberCertificate.Data != nil {
			d.Set("bacs_member_key_id", bacsAssociation.Payload.Data.Relationships.BacsMemberCertificate.Data.KeyID)
			d.Set("bacs_member_certificate_id", bacsAssociation.Payload.Data.Relationships.BacsMemberCertificate.Data.CertificateID)
		}
		if bacsAssociation.Payload.Data.Relationships.BacsServiceUserCertificate != nil && bacsAssociation.Payload.Data.Relationships.BacsServiceUserCertificate.Data != nil {
			d.Set("bacs_service_user_key_id", bacsAssociation.Payload.Data.Relationships.BacsServiceUserCertificate.Data.KeyID)
			d.Set("bacs_service_user_certificate_id", bacsAssociation.Payload.Data.Relationships.BacsServiceUserCertificate.Data.CertificateID)
		}
	}

	return nil
}

func resourceVocalinkReportAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	bacsAssociation, err := client.AssociationClient.Associations.GetVocalinkreportID(associations.NewGetVocalinkreportIDParams().
		WithID(strfmt.UUID(d.Id())))

	if err != nil {
		return fmt.Errorf("error deleting VocalinkReport association: %s", err)
	}

	log.Printf("[INFO] Deleting VocalinkReport association for id: %s ", bacsAssociation.Payload.Data.ID)

	_, err = client.AssociationClient.Associations.DeleteVocalinkreportID(associations.NewDeleteVocalinkreportIDParams().
		WithID(bacsAssociation.Payload.Data.ID).
		WithVersion(*bacsAssociation.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting VocalinkReport association: %s", err)
	}

	return nil
}

func createVocalinkReportNewAssociationFromResourceData(d *schema.ResourceData) (*models.NewVocalinkReportAssociation, error) {

	association := &models.NewVocalinkReportAssociation{
		Relationships: &models.VocalinkReportAssociationRelationships{},
	}

	association.Type = "associations"
	if attr, ok := GetUUIDOK(d, "association_id"); ok {
		uuid := strfmt.UUID(attr.String())
		association.ID = &uuid
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		uuid := strfmt.UUID(attr.String())
		association.OrganisationID = &uuid
	}

	association.Relationships.FpsMemberCertificate = buildVocalinkRelationship(d, "fps_member")
	association.Relationships.BacsMemberCertificate = buildVocalinkRelationship(d, "bacs_member")
	association.Relationships.BacsServiceUserCertificate = buildVocalinkRelationship(d, "bacs_service_user")

	return association, nil
}

func buildVocalinkRelationship(d *schema.ResourceData, relation string) *models.VocalinkReportAssociationCertificateRelationship {
	if keyId, ok := GetUUIDOK(d, relation+"_key_id"); ok {
		if certId, certOk := GetUUIDOK(d, relation+"_certificate_id"); certOk {
			return &models.VocalinkReportAssociationCertificateRelationship{
				Data: &models.VocalinkReportAssociationCertificateRelationshipData{KeyID: keyId, CertificateID: certId},
			}
		}
		return &models.VocalinkReportAssociationCertificateRelationship{
			Data: &models.VocalinkReportAssociationCertificateRelationshipData{KeyID: keyId},
		}
	}
	return nil
}
