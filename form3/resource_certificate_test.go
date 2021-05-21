package form3

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/api/support"
	"github.com/form3tech-oss/terraform-provider-form3/client/organisations"
	"github.com/form3tech-oss/terraform-provider-form3/client/system"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccKey_basic(t *testing.T) {
	var response models.Key
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)
	keyID := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3KeyConfig(organisationID, parentOrganisationID, testOrgName, keyID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKeyExists("form3_key.test_key", &response),
					resource.TestCheckResourceAttr("form3_key.test_key", "organisation_id", organisationID),
					resource.TestCheckResourceAttr("form3_key.test_key", "key_id", keyID),
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
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)
	keyID := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3KeyConfigElliptic(organisationID, parentOrganisationID, testOrgName, keyID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKeyExists("form3_key.test_key", &response),
					resource.TestCheckResourceAttr("form3_key.test_key", "organisation_id", organisationID),
					resource.TestCheckResourceAttr("form3_key.test_key", "key_id", keyID),
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
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)
	keyID := uuid.New().String()
	certificateID := uuid.New().String()

	certificate, err := support.GenerateSelfSignedCert()
	if err != nil {
		t.Fail()
	}

	var issuers []string
	for i := 0; i < 3; i++ {
		issuer, err := support.GenerateSelfSignedCert()
		if err != nil {
			t.Fail()
		}
		issuers = append(issuers, issuer)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3KeyConfigWithCert(organisationID, parentOrganisationID, testOrgName, keyID, certificateID, certificate, issuers),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKeyExists("form3_key.test_key", &response),
					resource.TestCheckResourceAttr("form3_key.test_key", "organisation_id", organisationID),
					resource.TestCheckResourceAttr("form3_key.test_key", "key_id", keyID),
					resource.TestCheckResourceAttr("form3_key.test_key", "subject", "CN=Terraform-test-with-cert"),
					resource.TestMatchResourceAttr("form3_key.test_key", "certificate_signing_request", regexp.MustCompile(".*BEGIN CERTIFICATE REQUEST.*")),
					resource.TestCheckResourceAttr("form3_certificate.cert", "organisation_id", organisationID),
					resource.TestCheckResourceAttr("form3_certificate.cert", "key_id", keyID),
					resource.TestCheckResourceAttr("form3_certificate.cert", "certificate_id", certificateID),
					resource.TestCheckResourceAttr("form3_certificate.cert", "certificate", certificate),
					resource.TestCheckResourceAttr("form3_certificate.cert", "issuing_certificates.#", "3"),
					resource.TestCheckResourceAttr("form3_certificate.cert", "issuing_certificates.0", issuers[0]),
					resource.TestCheckResourceAttr("form3_certificate.cert", "issuing_certificates.1", issuers[1]),
					resource.TestCheckResourceAttr("form3_certificate.cert", "issuing_certificates.2", issuers[2]),
				),
			},
		},
	})
}

func TestAccKey_withSelfSignedCert(t *testing.T) {
	var response models.Key
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)
	keyID := uuid.New().String()
	certificateID := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3KeyConfigWithSelfSignedCert(organisationID, parentOrganisationID, testOrgName, keyID, certificateID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKeyExists("form3_key.test_key", &response),
					resource.TestMatchResourceAttr("form3_certificate.cert", "actual_certificate", regexp.MustCompile(".*BEGIN CERTIFICATE.*")),
				),
			},
		},
	})
}

func deleteOrganisation(client *form3.AuthenticatedClient, organisationId string) {
	fmt.Printf("[INFO] Deleting test organisation %v", organisationId)

	if _, err := client.OrganisationClient.Organisations.DeleteUnitsID(organisations.NewDeleteUnitsIDParams().
		WithID(strfmt.UUID(organisationId)).WithVersion(0)); err != nil {
		fmt.Printf("[WARN] Failed to delete test organisation %v", organisationId)
	}

	fmt.Printf("[INFO] Sucessfuly deleted test organisation %v", organisationId)
}

func TestAccKey_importExistingCert(t *testing.T) {
	testAccPreCheck(t)

	var response models.Key
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)
	keyID := uuid.New().String()
	certificateID := uuid.New().String()

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
				OrganisationID: strfmt.UUID(parentOrganisationID),
				ID:             strfmt.UUID(organisationID),
				Type:           "organisations",
				Attributes: &models.OrganisationAttributes{
					Name: testOrgName,
				},
			}}))
		if err != nil {
			t.Fail()
		}

		defer deleteOrganisation(client, organisationID)

		_, err = client.SystemClient.System.PostKeys(system.NewPostKeysParams().
			WithKeyCreationRequest(&models.KeyCreation{
				Data: &models.Key{
					ID:             strfmt.UUID(keyID),
					OrganisationID: strfmt.UUID(organisationID),
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

		newCertificate, err := support.GenerateSelfSignedCert()
		if err != nil {
			t.Fail()
		}
		issuer, err := support.GenerateSelfSignedCert()
		if err != nil {
			t.Fail()
		}

		_, err = client.SystemClient.System.PostKeysKeyIDCertificates(system.NewPostKeysKeyIDCertificatesParams().
			WithKeyID(strfmt.UUID(keyID)).
			WithCertificateCreationRequest(&models.CertificateCreation{
				Data: &models.Certificate{
					ID:             strfmt.UUID(certificateID),
					OrganisationID: strfmt.UUID(organisationID),
					Attributes: &models.CertificateAttributes{
						Certificate:         &newCertificate,
						IssuingCertificates: []string{issuer},
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
				Config: getTestForm3KeyConfigExistingKey(organisationID, parentOrganisationID, testOrgName, keyID),

				ResourceName:       "form3_organisation.organisation",
				ImportState:        true,
				ImportStateId:      organisationID,
				ImportStateVerify:  false,
				ExpectNonEmptyPlan: false,
			},
			{
				Config:            getTestForm3KeyConfigExistingKey(organisationID, parentOrganisationID, testOrgName, keyID),
				ResourceName:      "form3_key.test_key",
				ImportState:       true,
				ImportStateId:     keyID,
				ImportStateVerify: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKeyExists("form3_key.test_key", &response),
					resource.TestCheckResourceAttr("form3_key.test_key", "organisation_id", organisationID),
					resource.TestCheckResourceAttr("form3_key.test_key", "key_id", keyID),
					resource.TestCheckResourceAttr("form3_key.test_key", "subject", "CN=Terraform-test-existing-cert"),
					resource.TestCheckResourceAttr("form3_key.test_key", "private_key", "existing-key-101"),
					resource.TestCheckResourceAttr("form3_key.test_key", "public_key", "existing-key-102"),
					resource.TestMatchResourceAttr("form3_key.test_key", "certificate_signing_request", regexp.MustCompile(".*EXISTING CSR.*"))),
				ExpectNonEmptyPlan: false,
			},
			{
				Config:            getTestForm3KeyConfigExistingCert(organisationID, keyID, certificateID),
				ResourceName:      "form3_certificate.cert",
				ImportState:       true,
				ImportStateId:     keyID + "/" + certificateID,
				ImportStateVerify: false,

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("form3_certificate.cert", "organisation_id", organisationID),
					resource.TestCheckResourceAttr("form3_certificate.cert", "key_id", keyID),
					resource.TestCheckResourceAttr("form3_certificate.cert", "certificate_id", certificateID),
					resource.TestCheckResourceAttr("form3_certificate.cert", "certificate", "Existing Certificate"),
					resource.TestCheckResourceAttr("form3_certificate.cert", "issuing_certificates.#", "1"),
					resource.TestCheckResourceAttr("form3_certificate.cert", "issuing_certificates.0", "Existing Issuing Certificate"),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
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

func getTestForm3KeyConfig(orgID, parOrgID, orgName, keyID string) string {
	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		               = "%s"
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
	`, orgID, parOrgID, orgName, keyID)
}

func getTestForm3KeyConfigElliptic(orgID, parOrgID, orgName, keyID string) string {
	return fmt.Sprintf(`

	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		               = "%s"
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
	}`, orgID, parOrgID, orgName, keyID)
}

func getTestForm3KeyConfigWithCert(orgID, parOrgID, orgName, keyID, certID, certificate string, issuers []string) string {
	var certIssuers []string
	for _, issuer := range issuers {
		certIssuers = append(certIssuers, getTestForm3CertTerraformFormat(issuer))
	}
	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		           = "%s"
	}

	resource "form3_key" "test_key" {
	  organisation_id         = "${form3_organisation.organisation.organisation_id}"
	  subject                 = "CN=Terraform-test-with-cert"
	  key_id                  = "%s"
	}

	resource "form3_certificate" "cert" {
	  organisation_id         = "${form3_organisation.organisation.organisation_id}"
	  key_id                  = "${form3_key.test_key.key_id}"
	  certificate_id          = "%s"
	  certificate             =  "%s"
	  issuing_certificates    = [%s]
	}`, orgID, parOrgID, orgName, keyID, certID, getTestForm3CertTerraformFormat(certificate), `"`+strings.Join(certIssuers, "\",\n\"")+`"`)
}

func getTestForm3KeyConfigWithSelfSignedCert(orgID, parOrgID, orgName, keyID, certID string) string {
	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		           = "%s"
	}

	resource "form3_key" "test_key" {
	  organisation_id         = "${form3_organisation.organisation.organisation_id}"
	  subject                 = "CN=Terraform-test-selfsigned"
	  key_id                  = "%s"
	}

	resource "form3_certificate" "cert" {
	  organisation_id         = "${form3_organisation.organisation.organisation_id}"
	  key_id                  = "${form3_key.test_key.key_id}"
	  certificate_id          = "%s"
	}
	`, orgID, parOrgID, orgName, keyID, certID)
}

func getTestForm3KeyConfigExistingKey(orgID, parOrgID, orgName, keyID string) string {
	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		           = "%s"
	}

	resource "form3_key" "test_key" {
		organisation_id         = "${form3_organisation.organisation.organisation_id}"
		subject                 = "CN=Terraform-test-existing"
		key_id  				= "%s"
	}
	`, orgID, parOrgID, orgName, keyID)
}

func getTestForm3KeyConfigExistingCert(orgID, keyID, certID string) string {
	return fmt.Sprintf(`
	resource "form3_certificate" "cert" {
		organisation_id         = "%s"
	  key_id                  = "%s"
	  certificate_id          = "%s"
	  certificate             = "Existing Certificate"
	  issuing_certificates    = ["Existing Issuing Certificate"]
	}`, orgID, keyID, certID)
}

func getTestForm3CertTerraformFormat(cert string) string {
	return regexp.MustCompile("\n").ReplaceAllString(cert, `\n`)
}
