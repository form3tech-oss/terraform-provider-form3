package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/organisations"
	"github.com/form3tech-oss/terraform-provider-form3/client/system"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/satori/go.uuid"
	"os"
	"regexp"
	"testing"
)

func TestAccKey_basic(t *testing.T) {
	var response models.Key
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	keyId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3KeyConfigA, organisationId, parentOrganisationId, keyId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKeyExists("form3_key.test_key", &response),
					resource.TestCheckResourceAttr("form3_key.test_key", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_key.test_key", "key_id", keyId),
					resource.TestCheckResourceAttr("form3_key.test_key", "type", "RSA"),
					resource.TestCheckResourceAttr("form3_key.test_key", "subject", "CN=Terraform-test"),
					resource.TestCheckResourceAttr("form3_key.test_key", "description", "test key"),
					resource.TestMatchOutput("csr", regexp.MustCompile(".*-----BEGIN CERTIFICATE REQUEST-----.*")),
				),
			},
		},
	})
}

func TestAccKey_ellipticCurve(t *testing.T) {
	var response models.Key
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	keyId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3KeyConfigElliptic, organisationId, parentOrganisationId, keyId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKeyExists("form3_key.test_key", &response),
					resource.TestCheckResourceAttr("form3_key.test_key", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_key.test_key", "key_id", keyId),
					resource.TestCheckResourceAttr("form3_key.test_key", "subject", "CN=Terraform-test"),
					resource.TestCheckResourceAttr("form3_key.test_key", "type", "EC"),
					resource.TestCheckResourceAttr("form3_key.test_key", "curve", "PRIME256V1"),
					resource.TestCheckResourceAttr("form3_key.test_key", "description", "test key"),
					resource.TestMatchOutput("csr", regexp.MustCompile(".*-----BEGIN CERTIFICATE REQUEST-----.*")),
				),
			},
		},
	})
}

func TestAccKey_withCert(t *testing.T) {
	var response models.Key
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	keyId := uuid.NewV4().String()
	certificateId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3KeyConfigWithCert, organisationId, parentOrganisationId, keyId, certificateId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKeyExists("form3_key.test_key", &response),
					resource.TestCheckResourceAttr("form3_key.test_key", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_key.test_key", "key_id", keyId),
					resource.TestCheckResourceAttr("form3_key.test_key", "subject", "CN=Terraform-test-with-cert"),
					resource.TestMatchResourceAttr("form3_key.test_key", "certificate_signing_request", regexp.MustCompile(".*BEGIN CERTIFICATE REQUEST.*")),
					resource.TestCheckResourceAttr("form3_certificate.cert", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_certificate.cert", "key_id", keyId),
					resource.TestCheckResourceAttr("form3_certificate.cert", "certificate_id", certificateId),
					resource.TestMatchResourceAttr("form3_certificate.cert", "certificate", regexp.MustCompile(".*MIIGZzCCBU\\+gAwIBAgIQQAFoVhQdgReBTMSz0Ui/AjANBgkqhkiG9w0BAQsFADCB.*")),
					resource.TestCheckResourceAttr("form3_certificate.cert", "issuing_certificates.#", "3"),
					resource.TestMatchResourceAttr("form3_certificate.cert", "issuing_certificates.0", regexp.MustCompile(".*My Bank.*")),
					resource.TestMatchResourceAttr("form3_certificate.cert", "issuing_certificates.2", regexp.MustCompile(".*Root.*")),
				),
			},
		},
	})
}

func TestAccKey_withSelfSignedCert(t *testing.T) {
	var response models.Key
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	keyId := uuid.NewV4().String()
	certificateId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3KeyConfigWithSelfSignedCert, organisationId, parentOrganisationId, keyId, certificateId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKeyExists("form3_key.test_key", &response),
					resource.TestMatchResourceAttr("form3_certificate.cert", "actual_certificate", regexp.MustCompile(".*BEGIN CERTIFICATE.*")),
				),
			},
		},
	})
}

func TestAccKey_importExistingCert(t *testing.T) {
	var response models.Key
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	keyId := uuid.NewV4().String()
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

		_, err = client.SystemClient.System.PostKeys(system.NewPostKeysParams().
			WithKeyCreationRequest(&models.KeyCreation{
				Data: &models.Key{
					ID:             strfmt.UUID(keyId),
					OrganisationID: strfmt.UUID(organisationId),
					Attributes: &models.KeyAttributes{
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

		_, err = client.SystemClient.System.PostKeysKeyIDCertificates(system.NewPostKeysKeyIDCertificatesParams().
			WithKeyID(strfmt.UUID(keyId)).
			WithCertificateCreationRequest(&models.CertificateCreation{
				Data: &models.Certificate{
					ID:             strfmt.UUID(certificateId),
					OrganisationID: strfmt.UUID(organisationId),
					Attributes: &models.CertificateAttributes{
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
		CheckDestroy: testAccCheckKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3KeyConfigExistingKey, organisationId, parentOrganisationId, keyId),

				ResourceName:       "form3_organisation.organisation",
				ImportState:        true,
				ImportStateId:      organisationId,
				ImportStateVerify:  false,
				ExpectNonEmptyPlan: false,
			},
			{
				Config:            fmt.Sprintf(testForm3KeyConfigExistingKey, organisationId, parentOrganisationId, keyId),
				ResourceName:      "form3_key.test_key",
				ImportState:       true,
				ImportStateId:     keyId,
				ImportStateVerify: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKeyExists("form3_key.test_key", &response),
					resource.TestCheckResourceAttr("form3_key.test_key", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_key.test_key", "key_id", keyId),
					resource.TestCheckResourceAttr("form3_key.test_key", "subject", "CN=Terraform-test-existing-cert"),
					resource.TestCheckResourceAttr("form3_key.test_key", "private_key", "existing-key-101"),
					resource.TestCheckResourceAttr("form3_key.test_key", "public_key", "existing-key-102"),
					resource.TestMatchResourceAttr("form3_key.test_key", "certificate_signing_request", regexp.MustCompile(".*EXISTING CSR.*"))),
				ExpectNonEmptyPlan: false,
			},
			{
				Config:            fmt.Sprintf(testForm3KeyConfigExistingCert, organisationId, keyId, certificateId),
				ResourceName:      "form3_certificate.cert",
				ImportState:       true,
				ImportStateId:     keyId + "/" + certificateId,
				ImportStateVerify: false,

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("form3_certificate.cert", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_certificate.cert", "key_id", keyId),
					resource.TestCheckResourceAttr("form3_certificate.cert", "certificate_id", certificateId),
					resource.TestCheckResourceAttr("form3_certificate.cert", "certificate", "Existing Certificate"),
					resource.TestCheckResourceAttr("form3_certificate.cert", "issuing_certificates.#", "1"),
					resource.TestCheckResourceAttr("form3_certificate.cert", "issuing_certificates.0", "Existing Issuing Certificate"),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func ToStringPointer(s string) *string {
	return &s
}

func testAccCheckKeyDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_key" {
			continue
		}

		list, err := client.SystemClient.System.GetKeys(system.NewGetKeysParams())
		if err != nil {
			return err
		}

		for _, key := range list.Payload.Data {
			if rs.Primary.ID == string(key.ID) {
				return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, key)
			}
		}
	}

	return nil
}

func testAccCheckKeyExists(resourceKey string, cr *models.Key) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		list, err := client.SystemClient.System.GetKeys(system.NewGetKeysParams())
		if err != nil {
			return err
		}

		for _, key := range list.Payload.Data {
			if rs.Primary.ID == string(key.ID) {
				cr = key
			}
		}
		if cr == nil {
			return fmt.Errorf("record not found expected %s", rs.Primary.ID)
		}

		return nil
	}
}

const testForm3KeyConfigA = `

resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_key" "test_key" {
	organisation_id = "${form3_organisation.organisation.organisation_id}"
  subject         = "CN=Terraform-test"
  key_id          = "%s"
  description     = "test key"
}

output "csr" {
  value = "${form3_key.test_key.certificate_signing_request}"
}
`

const testForm3KeyConfigElliptic = `

resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_key" "test_key" {
	organisation_id  = "${form3_organisation.organisation.organisation_id}"
  subject          = "CN=Terraform-test"
  key_id           = "%s"
  description      = "test key"
  type             = "EC"
  curve            = "PRIME256V1"
}

output "csr" {
  value = "${form3_key.test_key.certificate_signing_request}"
}
`

const testForm3KeyConfigWithCert = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_key" "test_key" {
	organisation_id         = "${form3_organisation.organisation.organisation_id}"
  subject                 = "CN=Terraform-test-with-cert"
  key_id  = "%s"
}

resource "form3_certificate" "cert" {
	organisation_id         = "${form3_organisation.organisation.organisation_id}"
  key_id  = "${form3_key.test_key.key_id}"
  certificate_id          = "%s"
  certificate             = "-----BEGIN CERTIFICATE-----\nMIIGZzCCBU+gAwIBAgIQQAFoVhQdgReBTMSz0Ui/AjANBgkqhkiG9w0BAQsFADCB\nqzEnMCUGA1UECgw=\n-----END CERTIFICATE-----"
  issuing_certificates    = ["-----BEGIN CERTIFICATE-----\nMy Bank\n-----END CERTIFICATE-----",
                              "-----BEGIN CERTIFICATE-----\nMy Bank's Bank'\n-----END CERTIFICATE-----",
                              "-----BEGIN CERTIFICATE-----\nRoot'\n-----END CERTIFICATE-----"
                            ]
}
`

const testForm3KeyConfigWithSelfSignedCert = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_key" "test_key" {
	organisation_id         = "${form3_organisation.organisation.organisation_id}"
  subject                 = "CN=Terraform-test-selfsigned"
  key_id  = "%s"
}

resource "form3_certificate" "cert" {
	organisation_id         = "${form3_organisation.organisation.organisation_id}"
  key_id  = "${form3_key.test_key.key_id}"
  certificate_id          = "%s"
}
`

const testForm3KeyConfigExistingKey = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_key" "test_key" {
	organisation_id         = "${form3_organisation.organisation.organisation_id}"
  subject                 = "CN=Terraform-test-existing"
  key_id  = "%s"
}
`

const testForm3KeyConfigExistingCert = `
resource "form3_certificate" "cert" {
	organisation_id         = "%s"
  key_id                  = "%s"
  certificate_id          = "%s"
  certificate             = "Existing Certificate"
  issuing_certificates    = ["Existing Issuing Certificate"]
}`
