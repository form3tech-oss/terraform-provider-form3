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

func resourceForm3VocalinkReportCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceVocalinkReportCertificateCreate,
		Read:   resourceVocalinkReportCertificateRead,
		Delete: resourceVocalinkReportCertificateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"certificate_request_id": {
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

func resourceVocalinkReportCertificateCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	log.Print("[INFO] Creating VocalinkReport Certificate  ")

	certificate, err := createVocalinkReportNewCertificateFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create VocalinkReport certificate : %s", err)
	}
	certRequestId, _ := GetUUIDOK(d, "certificate_request_id")

	createdCertificate, err := client.SystemClient.System.PostVocalinkreportCertificateRequestsCertificateRequestIDCertificate(
		system.NewPostVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams().
			WithCertificateRequestID(certRequestId).
			WithCertificateCreationRequest(&models.VocalinkReportCertificateCreation{
				Data: certificate,
			}))

	if err != nil {
		return fmt.Errorf("failed to create VocalinkReport certificate : %s", err)
	}

	d.SetId(createdCertificate.Payload.Data.ID.String())
	log.Printf("[INFO] VocalinkReport certificate  key: %s", d.Id())

	return resourceVocalinkReportCertificateRead(d, meta)
}

func resourceVocalinkReportCertificateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	certRequestId, _ := GetUUIDOK(d, "certificate_request_id")
	certId, _ := GetUUIDOK(d, "certificate__id")

	if certId == "" {
		certId = strfmt.UUID(d.Id())
		log.Printf("[INFO] Importing certificate with resource id: %s.", certId)
	} else {
		log.Printf("[INFO] Reading certificate with resource id: %s.", certId)
	}

	response, err := client.SystemClient.System.GetVocalinkreportCertificateRequestsCertificateRequestIDCertificate(
		system.NewGetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams().
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

	var cert *models.VocalinkReportCertificate
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
	d.Set("certificate_request_id", certRequestId.String())

	return nil
}

func resourceVocalinkReportCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	certRequestId, _ := GetUUIDOK(d, "certificate_request_id")

	response, err := client.SystemClient.System.GetVocalinkreportCertificateRequestsCertificateRequestIDCertificate(
		system.NewGetVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams().
			WithCertificateRequestID(strfmt.UUID(certRequestId)))
	var cert *models.VocalinkReportCertificate
	for _, certificate := range response.Payload.Data {
		if strfmt.UUID(d.Id()) == certificate.ID {
			cert = certificate
		}
	}

	if cert == nil {
		return fmt.Errorf("unable to delete certificate as it does not exist")
	}

	log.Printf("[INFO] Deleting VocalinkReport certificate request for id: %s ", strfmt.UUID(d.Id()))

	_, err = client.SystemClient.System.DeleteVocalinkreportCertificateRequestsCertificateRequestIDCertificateCertificateID(

		system.NewDeleteVocalinkreportCertificateRequestsCertificateRequestIDCertificateCertificateIDParams().
			WithCertificateRequestID(certRequestId).
			WithVersion(*cert.Version).
			WithCertificateID(cert.ID))

	if err != nil {
		return fmt.Errorf("error deleting VocalinkReport certificate: %s", err)
	}

	return nil
}

func createVocalinkReportNewCertificateFromResourceData(d *schema.ResourceData) (*models.VocalinkReportCertificate, error) {

	certificate := &models.VocalinkReportCertificate{
		Type:       "vocalink_report_certificate",
		Attributes: &models.VocalinkReportCertificateAttributes{},
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
