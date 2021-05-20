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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
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
					resource.TestMatchResourceAttr("form3_certificate.cert", "certificate", regexp.MustCompile(`.*DTJYYuo47rgUr2szykNbxTdLl0uYbW43DNzy\+q/77\+xg4to/F6rqqrvI6Q\+ETJw6.*`)),
					resource.TestCheckResourceAttr("form3_certificate.cert", "issuing_certificates.#", "3"),
					resource.TestMatchResourceAttr("form3_certificate.cert", "issuing_certificates.0", regexp.MustCompile(`.*b7HnblLHCJu4Se/Tw9S0dvgU6L8oZ9D92\+LXjraph0mmmQ0/Yemt8bDvH9rRAgMB.*`)),
					resource.TestMatchResourceAttr("form3_certificate.cert", "issuing_certificates.2", regexp.MustCompile(`.*Wc2I/20pb6nXkvP83rXNexHNMYmDarovpAzzNK\+Zu2X3sD93u9iO2jFbjFllAgMB.*`)),
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
		name 		               = "%s"
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
	  certificate             =  "-----BEGIN CERTIFICATE-----\nMIICejCCAeOgAwIBAgIUHRbxjJt+4jVXTsla1mO7wUZrZzEwDQYJKoZIhvcNAQEL\nBQAwTzELMAkGA1UEBhMCVUsxDzANBgNVBAgMBkxvbmRvbjEPMA0GA1UEBwwGTG9u\nZG9uMQ4wDAYDVQQKDAVmb3JtMzEOMAwGA1UECwwFZm9ybTMwHhcNMjEwNTIwMDk1\nMzU3WhcNMjIwNTIwMDk1MzU3WjBPMQswCQYDVQQGEwJVSzEPMA0GA1UECAwGTG9u\nZG9uMQ8wDQYDVQQHDAZMb25kb24xDjAMBgNVBAoMBWZvcm0zMQ4wDAYDVQQLDAVm\nb3JtMzCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA9hZxqbnPjoWXAbMuvkll\niQLvARL202q7Wmh5EhWjN7P2A42/7arYbPWbpGg7YFglDgbx9lOp2uyGhZMInrE1\nDTJYYuo47rgUr2szykNbxTdLl0uYbW43DNzy+q/77+xg4to/F6rqqrvI6Q+ETJw6\nimfStsrwaG85U0uLOWBd9t8CAwEAAaNTMFEwHQYDVR0OBBYEFINj9JnsYm+ABL85\nUx0LJFxKGOHaMB8GA1UdIwQYMBaAFINj9JnsYm+ABL85Ux0LJFxKGOHaMA8GA1Ud\nEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADgYEAPJbkJvjxyVuODFAYuFvwkcCO\nLmPj1AoMtrxi6pCSwEkwzmE/SXsuGxzsJ5W/XM9tJk4xsLmP7zo3HW5LEHjN66kP\nP0eETHy6uAL7mR4+qR3u7V5dZ5ezAD1P03ebQPCwZgshEg0YuuCn9FUGaRKTDZqB\n2e9LzTyGh7HY0P+JOS0=\n-----END CERTIFICATE-----"
	  issuing_certificates    = ["-----BEGIN CERTIFICATE-----\nMIICZjCCAc+gAwIBAgIUX/NtNTyBvvxcXud8RrYc43AW628wDQYJKoZIhvcNAQEL\nBQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM\nGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yMTA1MjAwOTU4MzNaFw0yMjA1\nMjAwOTU4MzNaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw\nHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwgZ8wDQYJKoZIhvcNAQEB\nBQADgY0AMIGJAoGBAMdiz4nORaGAnjpvJXXEQqtlbXL4VBErY/dIu9WiEa+XohQg\nh8QvYFA88Q6PmyP/7hNUWNG3wVkeLhj/1MB08FuaX4yWKWCp+29y7FVZqbmSAZta\nb7HnblLHCJu4Se/Tw9S0dvgU6L8oZ9D92+LXjraph0mmmQ0/Yemt8bDvH9rRAgMB\nAAGjUzBRMB0GA1UdDgQWBBQdGqdpjO0msCUfTfHEJJpaNRGqOTAfBgNVHSMEGDAW\ngBQdGqdpjO0msCUfTfHEJJpaNRGqOTAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3\nDQEBCwUAA4GBAKoaGZ8I712W1P3o3LKouy3thiQ97hdy0JZmj3xuqcgbyOfeyomM\nHE9YUR54F/EBnf8Xv7aqQUYG0CmAbilmAKmSMzHB4AJri38RWFDlQoyiQKb9Vopb\nXF3j5zElK4KkLn08rK/VChZv9cc/OVRAYql363DQaNc6Wj2A1PQUHNuL\n-----END CERTIFICATE-----",
								 "-----BEGIN CERTIFICATE-----\nMIICZjCCAc+gAwIBAgIUevd6x+OpPwNPQOdcnBnpw7n3PVIwDQYJKoZIhvcNAQEL\nBQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM\nGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yMTA1MjAwOTU4NTJaFw0yMjA1\nMjAwOTU4NTJaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw\nHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwgZ8wDQYJKoZIhvcNAQEB\nBQADgY0AMIGJAoGBAKrNXcVAHvptcqAulMLrIugvfjhDfpR9FG63XiECkBNHwl/2\nMEYuGjssUQ9vuwM06lRhKnjfzv72M2XfDQyamiU7QJMDDIcBDlNz8LQ0O51peuHy\nz6ynm1RvJuEgXDKZ8DZX9wPGHUg+y+URdPIjdO9Ue1l0nzka3zlz9TZ3jQYxAgMB\nAAGjUzBRMB0GA1UdDgQWBBQJYVHNE2ve/7BI+Sfnriuq2nCjCTAfBgNVHSMEGDAW\ngBQJYVHNE2ve/7BI+Sfnriuq2nCjCTAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3\nDQEBCwUAA4GBAI6h5j8br041yMWAMmrWYSwrb1flcVkZzGeUSqxQrRZXjrTCZ2wm\niXPu/8xz8yd27xQ8Tbi/osUhjX0bLy9KXJTwxdOUVvkCWo7qigyg0TTTZniWCFMl\nqWisoSLErkcw6CgpGRodyZNmitRvpXPKxCv8qoUhpzkg5M6OshR5PrBi\n-----END CERTIFICATE-----",
								 "-----BEGIN CERTIFICATE-----\nMIICZjCCAc+gAwIBAgIUKmAhNwgZwzXFPBDFzNnrQhYscb4wDQYJKoZIhvcNAQEL\nBQAwRTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM\nGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yMTA1MjAwOTU4NTlaFw0yMjA1\nMjAwOTU4NTlaMEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw\nHwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwgZ8wDQYJKoZIhvcNAQEB\nBQADgY0AMIGJAoGBANbse3Sn19sON88LoRuMdLV3+N+Ddn+IreInGXZrm7TRjmkM\nputWs5w4vDzilJunQojfZGxLzo2R7fV2aNO+dC41cINnN3TzMJeW9WN4F24yWk7V\nWc2I/20pb6nXkvP83rXNexHNMYmDarovpAzzNK+Zu2X3sD93u9iO2jFbjFllAgMB\nAAGjUzBRMB0GA1UdDgQWBBQc2Ngkis1BDi7hFf5S1JShjjogljAfBgNVHSMEGDAW\ngBQc2Ngkis1BDi7hFf5S1JShjjogljAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3\nDQEBCwUAA4GBAK/TpNREAIWV/wB3VtMNrCNkWp5w3IVrC3EmitZXWzmT4bxjaI5M\nP0XQtVDJfhig3+8iCftzkFZyP9hVkGyVL2O3UoEl1IYKUTisfzz22cQqu0cFyjlk\n/yxo3rLDoaSVjduX7pfl81BftpY+3Y0cpeTBa4x09JJHYYw2CY7XcHnN\n-----END CERTIFICATE-----"
								]
	}`, orgID, parOrgID, orgName, keyID, certID)
}

func getTestForm3KeyConfigWithSelfSignedCert(orgID, parOrgID, orgName, keyID, certID string) string {
	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		               = "%s"
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
	`, orgID, parOrgID, orgName, keyID, certID)
}

func getTestForm3KeyConfigExistingKey(orgID, parOrgID, orgName, keyID string) string {
	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		               = "%s"
	}

	resource "form3_key" "test_key" {
		organisation_id         = "${form3_organisation.organisation.organisation_id}"
	  subject                 = "CN=Terraform-test-existing"
	  key_id  = "%s"
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
