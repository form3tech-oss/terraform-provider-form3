package form3

import (
	"fmt"
	"github.com/form3tech-oss/go-form3"
	"github.com/form3tech-oss/go-form3/client/system"
	"github.com/form3tech-oss/go-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3Certificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceCertificateCreate,
		Read:   resourceCertificateRead,
		Delete: resourceCertificateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"key_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"certificate_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"certificate": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"issuing_certificates": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				ForceNew: true,
			},
			"actual_certificate": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCertificateCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	log.Print("[INFO] Creating Certificate")

	certificate, err := createVocalinkReportNewCertificateFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create Certificate : %s", err)
	}
	certRequestId, _ := GetUUIDOK(d, "key_id")

	createdCertificate, err := client.SystemClient.System.PostCertificateRequestsCertificateRequestIDCertificate(
		system.NewPostCertificateRequestsCertificateRequestIDCertificateParams().
			WithCertificateRequestID(certRequestId).
			WithCertificateCreationRequest(&models.CertificateCreation{
				Data: certificate,
			}))

	if err != nil {
		return fmt.Errorf("failed to create certificate : %s", err)
	}

	d.SetId(createdCertificate.Payload.Data.ID.String())
	log.Printf("[INFO] Certificate  key: %s", d.Id())

	return resourceCertificateRead(d, meta)
}

func resourceCertificateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	certRequestId, _ := GetUUIDOK(d, "key_id")
	certId, _ := GetUUIDOK(d, "certificate__id")

	if certId == "" {
		certId = strfmt.UUID(d.Id())
		log.Printf("[INFO] Importing certificate with resource id: %s.", certId)
	} else {
		log.Printf("[INFO] Reading certificate with resource id: %s.", certId)
	}

	response, err := client.SystemClient.System.GetCertificateRequestsCertificateRequestIDCertificate(
		system.NewGetCertificateRequestsCertificateRequestIDCertificateParams().
			WithCertificateRequestID(certRequestId))

	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		} else {
			return err
		}
	}

	var cert *models.Certificate
	for _, certificate := range response.Payload.Data {
		if certId == certificate.ID {
			cert = certificate
		}
	}
	if cert == nil {
		d.SetId("")
		return nil
	}

	d.Set("actual_certificate", cert.Attributes.Certificate)
	d.Set("issuing_certificates", cert.Attributes.IssuingCertificates)
	d.Set("organisation_id", cert.OrganisationID.String())
	d.Set("certificate_id", cert.ID.String())
	d.Set("key_id", certRequestId.String())

	return nil
}

func resourceCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	certRequestId, _ := GetUUIDOK(d, "key_id")

	response, err := client.SystemClient.System.GetCertificateRequestsCertificateRequestIDCertificate(
		system.NewGetCertificateRequestsCertificateRequestIDCertificateParams().
			WithCertificateRequestID(strfmt.UUID(certRequestId)))
	var cert *models.Certificate
	for _, certificate := range response.Payload.Data {
		if strfmt.UUID(d.Id()) == certificate.ID {
			cert = certificate
		}
	}

	if cert == nil {
		return fmt.Errorf("unable to delete certificate as it does not exist")
	}

	log.Printf("[INFO] Deleting Certificate request for id: %s ", strfmt.UUID(d.Id()))

	_, err = client.SystemClient.System.DeleteCertificateRequestsCertificateRequestIDCertificateCertificateID(

		system.NewDeleteCertificateRequestsCertificateRequestIDCertificateCertificateIDParams().
			WithCertificateRequestID(certRequestId).
			WithVersion(*cert.Version).
			WithCertificateID(cert.ID))

	if err != nil {
		return fmt.Errorf("error deleting Certificate: %s", err)
	}

	return nil
}

func createVocalinkReportNewCertificateFromResourceData(d *schema.ResourceData) (*models.Certificate, error) {

	certificate := &models.Certificate{
		Type:       "certificate",
		Attributes: &models.CertificateAttributes{},
	}

	if attr, ok := GetUUIDOK(d, "certificate_id"); ok {
		uuid := strfmt.UUID(attr.String())
		certificate.ID = uuid
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		uuid := strfmt.UUID(attr.String())
		certificate.OrganisationID = uuid
	}

	if attr, ok := d.GetOk("certificate"); ok {
		s := attr.(string)
		certificate.Attributes.Certificate = &s
	}

	if attr, ok := d.GetOk("issuing_certificates"); ok {
		arr := attr.([]interface{})
		var issuingCerts []string
		for _, v := range arr {
			issuingCerts = append(issuingCerts, v.(string))
		}
		certificate.Attributes.IssuingCertificates = issuingCerts
	}

	return certificate, nil
}
