package form3

import (
	"fmt"
	"github.com/form3tech-oss/go-form3"
	"github.com/form3tech-oss/go-form3/client/organisations"
	"github.com/form3tech-oss/go-form3/client/system"
	"github.com/form3tech-oss/go-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/satori/go.uuid"
	"os"
	"regexp"
	"testing"
)

func TestAccVocalinkReportCertificateRequest_basic(t *testing.T) {
	var response models.VocalinkReportCertificateRequest
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	certificateRequestId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVocalinkReportCertificateRequestDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3VocalinkReportCertificateRequestConfigA, organisationId, parentOrganisationId, certificateRequestId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocalinkReportCertificateRequestExists("form3_vocalink_report_certificate_request.cert_req", &response),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate_request.cert_req", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate_request.cert_req", "certificate_request_id", certificateRequestId),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate_request.cert_req", "subject", "CN=Terraform-test"),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate_request.cert_req", "description", "vocalink contact name is hsm1234"),
					resource.TestMatchOutput("csr", regexp.MustCompile(".*-----BEGIN CERTIFICATE REQUEST-----.*")),
				),
			},
		},
	})
}

func TestAccVocalinkReportCertificateRequest_withCert(t *testing.T) {
	var response models.VocalinkReportCertificateRequest
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	certificateRequestId := uuid.NewV4().String()
	certificateId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVocalinkReportCertificateRequestDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3VocalinkReportCertificateRequestConfigWithCert, organisationId, parentOrganisationId, certificateRequestId, certificateId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocalinkReportCertificateRequestExists("form3_vocalink_report_certificate_request.cert_req", &response),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate_request.cert_req", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate_request.cert_req", "certificate_request_id", certificateRequestId),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate_request.cert_req", "subject", "CN=Terraform-test-with-cert"),
					resource.TestMatchResourceAttr("form3_vocalink_report_certificate_request.cert_req", "certificate_signing_request", regexp.MustCompile(".*BEGIN CERTIFICATE REQUEST.*")),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate.cert", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate.cert", "certificate_request_id", certificateRequestId),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate.cert", "certificate_id", certificateId),
					resource.TestMatchResourceAttr("form3_vocalink_report_certificate.cert", "certificate", regexp.MustCompile(".*MIIGZzCCBU\\+gAwIBAgIQQAFoVhQdgReBTMSz0Ui/AjANBgkqhkiG9w0BAQsFADCB.*")),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate.cert", "issuing_certificates.#", "3"),
					resource.TestMatchResourceAttr("form3_vocalink_report_certificate.cert", "issuing_certificates.0", regexp.MustCompile(".*My Bank.*")),
					resource.TestMatchResourceAttr("form3_vocalink_report_certificate.cert", "issuing_certificates.2", regexp.MustCompile(".*Root.*")),
				),
			},
		},
	})
}

func TestAccVocalinkReportCertificateRequest_withSelfSignedCert(t *testing.T) {
	var response models.VocalinkReportCertificateRequest
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	certificateRequestId := uuid.NewV4().String()
	certificateId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVocalinkReportCertificateRequestDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3VocalinkReportCertificateRequestConfigWithSelfSignedCert, organisationId, parentOrganisationId, certificateRequestId, certificateId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocalinkReportCertificateRequestExists("form3_vocalink_report_certificate_request.cert_req", &response),
					resource.TestMatchResourceAttr("form3_vocalink_report_certificate.cert", "actual_certificate", regexp.MustCompile(".*BEGIN CERTIFICATE.*")),
				),
			},
		},
	})
}

func TestAccVocalinkReportCertificateRequest_importExistingCert(t *testing.T) {
	var response models.VocalinkReportCertificateRequest
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	certificateRequestId := uuid.NewV4().String()
	certificateId := uuid.NewV4().String()

	if acc, ok := os.LookupEnv("TF_ACC"); ok && acc == "1" {
		// Setup existing resources to be imported.
		config := Config{
			ApiHost:      os.Getenv("FORM3_HOST"),
			ClientId:     os.Getenv("FORM3_CLIENT_ID"),
			ClientSecret: os.Getenv("FORM3_CLIENT_SECRET"),
		}
		client, err := config.Client()
		if err != nil {
			t.Fail()
		}

		_, err = client.OrganisationClient.Organisations.PostUnits(organisations.NewPostUnitsParams().
			WithOrganisationCreationRequest(&models.OrganisationCreation{Data: &models.Organisation{
				OrganisationID: strfmt.UUID(parentOrganisationId),
				ID:             strfmt.UUID(organisationId),
				Type:           "organisations",
				Attributes: &models.OrganisationAttributes{
					Name: "terraform-organisation",
				},
			}}))
		if err != nil {
			t.Fail()
		}

		_, err = client.SystemClient.System.PostVocalinkreportCertificateRequests(system.NewPostVocalinkreportCertificateRequestsParams().
			WithCertificateRequestCreationRequest(&models.VocalinkReportCertificateRequestCreation{
				Data: &models.VocalinkReportCertificateRequest{
					ID:             strfmt.UUID(certificateRequestId),
					OrganisationID: strfmt.UUID(organisationId),
					Attributes: &models.VocalinkReportCertificateRequestAttributes{
						CertificateSigningRequest: "EXISTING CSR",
						Description:               "",
						PrivateKey:                "existing-key-101",
						PublicKey:                 "existing-key-102",
						Subject:                   "CN=Terraform-test-existing-cert",
					},
				},
			}))
		if err != nil {
			t.Fail()
		}

		_, err = client.SystemClient.System.PostVocalinkreportCertificateRequestsCertificateRequestIDCertificate(system.NewPostVocalinkreportCertificateRequestsCertificateRequestIDCertificateParams().
			WithCertificateRequestID(strfmt.UUID(certificateRequestId)).
			WithCertificateCreationRequest(&models.VocalinkReportCertificateCreation{
				Data: &models.VocalinkReportCertificate{
					ID:             strfmt.UUID(certificateId),
					OrganisationID: strfmt.UUID(organisationId),
					Attributes: &models.VocalinkReportCertificateAttributes{
						Certificate:         ToStringPointer("Existing Certificate"),
						IssuingCertificates: []string{"Existing Issuing Certificate"},
					},
				},
			}))
		if err != nil {
			t.Fail()
		}
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVocalinkReportCertificateRequestDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3VocalinkReportCertificateRequestConfigExistingCertReq, organisationId, parentOrganisationId, certificateRequestId),

				ResourceName:      "form3_organisation.organisation",
				ImportState:       true,
				ImportStateId:     organisationId,
				ImportStateVerify: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVocalinkReportCertificateRequestExists("form3_vocalink_report_certificate_request.cert_req", &response),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate_request.cert_req", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate_request.cert_req", "certificate_request_id", certificateRequestId),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate_request.cert_req", "subject", "CN=Terraform-test-existing-cert"),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate_request.cert_req", "private_key", "existing-key-101"),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate_request.cert_req", "public_key", "existing-key-102"),
					resource.TestMatchResourceAttr("form3_vocalink_report_certificate_request.cert_req", "certificate_signing_request", regexp.MustCompile(".*EXISTING CSR.*"))),
			},
			{
				Config:        fmt.Sprintf(testForm3VocalinkReportCertificateRequestConfigExistingCert, organisationId, certificateRequestId, certificateId),
				ResourceName:  "form3_vocalink_report_certificate_request.cert",
				ImportState:   true,
				ImportStateId: certificateId,

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate.cert", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate.cert", "certificate_request_id", certificateRequestId),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate.cert", "certificate_id", certificateId),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate.cert", "certificate", "Existing Certificate"),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate.cert", "issuing_certificates.#", "1"),
					resource.TestCheckResourceAttr("form3_vocalink_report_certificate.cert", "issuing_certificates.0", "Existing Issuing Certificate"),
				),
			},
		},
	})
}

func ToStringPointer(s string) *string {
	return &s
}

func testAccCheckVocalinkReportCertificateRequestDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_vocalink_report_certificate_request" {
			continue
		}

		list, err := client.SystemClient.System.GetVocalinkreportCertificateRequests(system.NewGetVocalinkreportCertificateRequestsParams())
		if err != nil {
			return err
		}

		for _, certificateRequest := range list.Payload.Data {
			if rs.Primary.ID == string(certificateRequest.ID) {
				return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, certificateRequest)
			}
		}
	}

	return nil
}

func testAccCheckVocalinkReportCertificateRequestExists(resourceKey string, cr *models.VocalinkReportCertificateRequest) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		list, err := client.SystemClient.System.GetVocalinkreportCertificateRequests(system.NewGetVocalinkreportCertificateRequestsParams())
		if err != nil {
			return err
		}

		for _, certificateRequest := range list.Payload.Data {
			if rs.Primary.ID == string(certificateRequest.ID) {
				cr = certificateRequest
			}
		}
		if cr == nil {
			return fmt.Errorf("record not found expected %s", rs.Primary.ID)
		}

		return nil
	}
}

const testForm3VocalinkReportCertificateRequestConfigA = `

resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_vocalink_report_certificate_request" "cert_req" {
	organisation_id         = "${form3_organisation.organisation.organisation_id}"
  subject                 = "CN=Terraform-test"
  certificate_request_id  = "%s"
  description             = "vocalink contact name is hsm1234"
}

output "csr" {
  value = "${form3_vocalink_report_certificate_request.cert_req.certificate_signing_request}"
}
`

const testForm3VocalinkReportCertificateRequestConfigWithCert = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_vocalink_report_certificate_request" "cert_req" {
	organisation_id         = "${form3_organisation.organisation.organisation_id}"
  subject                 = "CN=Terraform-test-with-cert"
  certificate_request_id  = "%s"
}

resource "form3_vocalink_report_certificate" "cert" {
	organisation_id         = "${form3_organisation.organisation.organisation_id}"
  certificate_request_id  = "${form3_vocalink_report_certificate_request.cert_req.certificate_request_id}"
  certificate_id          = "%s"
  certificate             = "-----BEGIN CERTIFICATE-----\nMIIGZzCCBU+gAwIBAgIQQAFoVhQdgReBTMSz0Ui/AjANBgkqhkiG9w0BAQsFADCB\nqzEnMCUGA1UECgw=\n-----END CERTIFICATE-----"
  issuing_certificates    = ["-----BEGIN CERTIFICATE-----\nMy Bank\n-----END CERTIFICATE-----",
                              "-----BEGIN CERTIFICATE-----\nMy Bank's Bank'\n-----END CERTIFICATE-----",
                              "-----BEGIN CERTIFICATE-----\nRoot'\n-----END CERTIFICATE-----"
                            ]
}
`

const testForm3VocalinkReportCertificateRequestConfigWithSelfSignedCert = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_vocalink_report_certificate_request" "cert_req" {
	organisation_id         = "${form3_organisation.organisation.organisation_id}"
  subject                 = "CN=Terraform-test-selfsigned"
  certificate_request_id  = "%s"
}

resource "form3_vocalink_report_certificate" "cert" {
	organisation_id         = "${form3_organisation.organisation.organisation_id}"
  certificate_request_id  = "${form3_vocalink_report_certificate_request.cert_req.certificate_request_id}"
  certificate_id          = "%s"
}
`

const testForm3VocalinkReportCertificateRequestConfigExistingCertReq = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_vocalink_report_certificate_request" "cert_req" {
	organisation_id         = "${form3_organisation.organisation.organisation_id}"
  subject                 = "CN=Terraform-test-existing"
  certificate_request_id  = "%s"
}
`

const testForm3VocalinkReportCertificateRequestConfigExistingCert = `
resource "form3_vocalink_report_certificate" "cert" {
	organisation_id         = "%s"
  certificate_request_id  = "%s"
  certificate_id          = "%s"
  certificate             = "Existing Certificate"
  issuing_certificates    = ["Existing Issuing Certificate"]
}`
