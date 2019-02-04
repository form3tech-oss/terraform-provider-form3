package form3

import (
	"fmt"
	"github.com/form3tech-oss/go-form3"
	"github.com/form3tech-oss/go-form3/client/system"
	"github.com/form3tech-oss/go-form3/models"
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
