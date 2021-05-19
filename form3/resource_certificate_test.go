package form3

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
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

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3KeyConfigWithCert(organisationID, parentOrganisationID, testOrgName, keyID, certificateID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKeyExists("form3_key.test_key", &response),
					resource.TestCheckResourceAttr("form3_key.test_key", "organisation_id", organisationID),
					resource.TestCheckResourceAttr("form3_key.test_key", "key_id", keyID),
					resource.TestCheckResourceAttr("form3_key.test_key", "subject", "CN=Terraform-test-with-cert"),
					resource.TestMatchResourceAttr("form3_key.test_key", "certificate_signing_request", regexp.MustCompile(".*BEGIN CERTIFICATE REQUEST.*")),
					resource.TestCheckResourceAttr("form3_certificate.cert", "organisation_id", organisationID),
					resource.TestCheckResourceAttr("form3_certificate.cert", "key_id", keyID),
					resource.TestCheckResourceAttr("form3_certificate.cert", "certificate_id", certificateID),
					resource.TestMatchResourceAttr("form3_certificate.cert", "certificate", regexp.MustCompile(`.*hx60i6ONiYT9H4NxnSVRvUL8As\+9iaqYqUQMjfEc.*`)),
					//resource.TestCheckResourceAttr("form3_certificate.cert", "issuing_certificates.#", "3"),
					//resource.TestMatchResourceAttr("form3_certificate.cert", "issuing_certificates.0", regexp.MustCompile(".*My Bank.*")),
					//resource.TestMatchResourceAttr("form3_certificate.cert", "issuing_certificates.2", regexp.MustCompile(".*Root.*")),
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

		_, err = client.SystemClient.System.PostKeysKeyIDCertificates(system.NewPostKeysKeyIDCertificatesParams().
			WithKeyID(strfmt.UUID(keyID)).
			WithCertificateCreationRequest(&models.CertificateCreation{
				Data: &models.Certificate{
					ID:             strfmt.UUID(certificateID),
					OrganisationID: strfmt.UUID(organisationID),
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

func getTestForm3KeyConfigWithCert(orgID, parOrgID, orgName, keyID, certID string) string {
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
	  certificate             = "-----BEGIN CERTIFICATE-----\nMIIDVTCCAj2gAwIBAgIUHopfB4GiImqSfQ1hN5jTr5AKm0YwDQYJKoZIhvcNAQEL\nBQAwOjELMAkGA1UEBhMCQVUxCzAJBgNVBAgMAkdCMR4wHAYDVQQKDBVGb3JtMyBG\naW5hbmNpYWwgQ2xvdWQwHhcNMjEwNTE5MTUzOTM2WhcNMjEwNjE4MTUzOTM2WjA6\nMQswCQYDVQQGEwJBVTELMAkGA1UECAwCR0IxHjAcBgNVBAoMFUZvcm0zIEZpbmFu\nY2lhbCBDbG91ZDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAOk8KTdf\n82QvBp6dvXj9jywkRWPOqccPc9STD0zRghE0ONmKv7St1WvPxJAbnYxL1pSTvLHD\nXrtAIpW0tTINyTzj/nS/TAb9qIk6xRu3vxzfuETXloxo/rRKkp/K7WHpwep4u+Oz\nS96bdSUl0LD8K14gphG+kl/L2leuY3916hZHNjwfgJqQKtXg75RahQH1bj1/6WcA\n8I4/JR0Tfd58PYY71SS6fEnutGjaeKaT49jT7viHKBnl8NTyjeXmh4e6aSYP5Nur\nftgWI0c0E979StcGZQSznz7L81DBWXQY12/Zj+MloaIMgw5wh7WFoXhpqGwijuVW\nXSn46lR2syLXsEkCAwEAAaNTMFEwHQYDVR0OBBYEFHqnwpAjaLablj1r2te6BX2y\nmEVfMB8GA1UdIwQYMBaAFHqnwpAjaLablj1r2te6BX2ymEVfMA8GA1UdEwEB/wQF\nMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAIIPujZX8/5ofbb/9BpVHE5iOYe9Bdmh\n409MHdNiLJJ9H3Z+mKy07kf7/NfQg75nTi7EClpVmjHEA6lqhxXf5tXiEXknW7Za\nJLlLvGUjIUDUMUvMF0cKyeO/hx60i6ONiYT9H4NxnSVRvUL8As+9iaqYqUQMjfEc\nsKuD4251r4/0kuc3h+V9oUSEF+F32xnmNmR/n0UyuiTQ0zZKsvWuf6fuqlL0B16x\nQCoRZmrLVUGsUY9ZMeUiqLyc0hVLnXeRUj0Osdpl1ye93zU3oe6RG+kBF6DX7T7p\nL5EzEqeqsdvHLyuME4Qd85Nn98SnItnFUxonzXB07CqKkD8w3TJ/P2U=\n-----END CERTIFICATE-----"
	}`, orgID, parOrgID, orgName, keyID, certID)
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
