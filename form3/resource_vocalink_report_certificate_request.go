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

func resourceForm3VocalinkReportCertificateRequest() *schema.Resource {
	return &schema.Resource{
		Create: resourceVocalinkReportCertificateRequestCreate,
		Read:   resourceVocalinkReportCertificateRequestRead,
		Delete: resourceVocalinkReportCertificateRequestDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"certificate_request_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subject": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"certificate_signing_request": {
				Type:     schema.TypeString,
				Computed: true,
				Required: false,
				ForceNew: true,
			},
			"private_key": {
				Type:     schema.TypeString,
				Computed: true,
				Required: false,
				ForceNew: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Computed: true,
				Required: false,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceVocalinkReportCertificateRequestCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	log.Print("[INFO] Creating VocalinkReport Certificate Request ")

	certificateRequest, err := createVocalinkReportNewCertificateRequestFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create VocalinkReport certificate Request: %s", err)
	}

	createdCertificateRequest, err := client.SystemClient.System.PostVocalinkreportCertificateRequests(
		system.NewPostVocalinkreportCertificateRequestsParams().
			WithCertificateRequestCreationRequest(&models.VocalinkReportCertificateRequestCreation{
				Data: certificateRequest,
			}))

	if err != nil {
		return fmt.Errorf("failed to create VocalinkReport ertificate Request: %s", err)
	}

	d.SetId(createdCertificateRequest.Payload.Data.ID.String())
	log.Printf("[INFO] VocalinkReport certificate Request key: %s", d.Id())

	return resourceVocalinkReportCertificateRequestRead(d, meta)
}

func resourceVocalinkReportCertificateRequestRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	certRequestId, _ := GetUUIDOK(d, "certificate_request_id")

	if certRequestId == "" {
		certRequestId = strfmt.UUID(d.Id())
		log.Printf("[INFO] Importing certificate request with resource id: %s.", certRequestId)
	} else {
		log.Printf("[INFO] Reading certificate request with resource id: %s.", certRequestId)
	}

	response, err := client.SystemClient.System.GetVocalinkreportCertificateRequestsCertificateRequestID(
		system.NewGetVocalinkreportCertificateRequestsCertificateRequestIDParams().WithCertificateRequestID(certRequestId))

	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		} else {
			return err
		}
	}

	d.Set("certificate_request_id", response.Payload.Data.ID.String())
	d.Set("subject", response.Payload.Data.Attributes.Subject)
	d.Set("organisation_id", response.Payload.Data.OrganisationID.String())
	d.Set("certificate_signing_request", response.Payload.Data.Attributes.CertificateSigningRequest)
	d.Set("public_key", response.Payload.Data.Attributes.PublicKey)
	d.Set("private_key", response.Payload.Data.Attributes.PrivateKey)
	d.Set("description", response.Payload.Data.Attributes.Description)

	return nil
}

func resourceVocalinkReportCertificateRequestDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	response, err := client.SystemClient.System.GetVocalinkreportCertificateRequestsCertificateRequestID(
		system.NewGetVocalinkreportCertificateRequestsCertificateRequestIDParams().WithCertificateRequestID(strfmt.UUID(d.Id())))

	if err != nil {
		return fmt.Errorf("error deleting VocalinkReport certificate request: %s", err)
	}

	log.Printf("[INFO] Deleting VocalinkReport certificate request for id: %s ", response.Payload.Data.ID)

	_, err = client.SystemClient.System.DeleteVocalinkreportCertificateRequestsCertificateRequestID(
		system.NewDeleteVocalinkreportCertificateRequestsCertificateRequestIDParams().
			WithCertificateRequestID(response.Payload.Data.ID).
			WithVersion(*response.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting VocalinkReport certificate request: %s", err)
	}

	return nil
}

func createVocalinkReportNewCertificateRequestFromResourceData(d *schema.ResourceData) (*models.VocalinkReportCertificateRequest, error) {

	certificateRequest := &models.VocalinkReportCertificateRequest{
		Type:       "vocalink_report_certificate_requests",
		Attributes: &models.VocalinkReportCertificateRequestAttributes{},
	}

	if attr, ok := GetUUIDOK(d, "certificate_request_id"); ok {
		uuid := strfmt.UUID(attr.String())
		certificateRequest.ID = uuid
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		uuid := strfmt.UUID(attr.String())
		certificateRequest.OrganisationID = uuid
	}

	if attr, ok := d.GetOk("subject"); ok {
		certificateRequest.Attributes.Subject = attr.(string)
	}

	if attr, ok := d.GetOk("description"); ok {
		certificateRequest.Attributes.Description = attr.(string)
	}

	return certificateRequest, nil
}
