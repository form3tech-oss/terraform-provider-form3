package form3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"text/template"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccBacsAssociation_basic(t *testing.T) {
	var bacsResponse associations.GetBacsIDOK
	var inputKeyResponse models.Key
	var outputKeyResponse models.Key

	configData := associationConfigWithCerts{
		OrgID:           uuid.New().String(),
		OrgParentID:     os.Getenv("FORM3_ORGANISATION_ID"),
		AssociationID:   uuid.New().String(),
		InputKeyID:      uuid.New().String(),
		InputCertID:     uuid.New().String(),
		MessagingKeyID:  uuid.New().String(),
		MessagingCertID: uuid.New().String(),
		OutputKeyID:     uuid.New().String(),
		OutputCertID:    uuid.New().String(),
	}
	defer verifyOrgDoesNotExist(t, configData.OrgID)

	config, err := makeTestForm3BacsAssociationConfigWithCerts(configData)
	if err != nil {
		t.Fatalf("make configuration failed: %v", err)
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: func(state *terraform.State) error {
			if err := testAccCheckBacsAssociationDestroy(state); err != nil {
				return err
			}

			return testAccCheckKeyDestroy(state)
		},
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBacsAssociationExists("form3_bacs_association.association", &bacsResponse),
					testAccCheckKeyExists("form3_key.inputKey", &inputKeyResponse),
					testAccCheckKeyExists("form3_key.outputKey", &outputKeyResponse),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "service_user_number", "112238"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "account_number", "12345678"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "sorting_code", "123456"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "account_type", "1"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "organisation_id", configData.OrgID),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "association_id", configData.AssociationID),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "input_key_id", configData.InputKeyID),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "input_certificate_id", configData.InputCertID),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "input_tsu_number", "B12345"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "messaging_key_id", configData.MessagingKeyID),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "messaging_certificate_id", configData.MessagingCertID),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "messaging_tsu_number", "B12346"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "output_key_id", configData.OutputKeyID),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "output_certificate_id", configData.OutputCertID),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "output_tsu_number", "B12347"),
				),
			},
		},
	})
}

func TestAccBacsAssociation_zeroAccountType(t *testing.T) {
	var bacsResponse associations.GetBacsIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationId)

	associationId := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBacsAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3BacsAssociationConfigZeroAccountType, organisationId, parentOrganisationId, associationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBacsAssociationExists("form3_bacs_association.association", &bacsResponse),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "service_user_number", "112233"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "account_number", "87654321"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "sorting_code", "654321"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "account_type", "0"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "association_id", associationId),
				),
			},
		},
	})
}

func TestAccBacsAssociation_withBankIdAndCentre(t *testing.T) {
	var bacsResponse associations.GetBacsIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationId)

	associationId := uuid.New().String()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBacsAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3BacsAssociationConfigWithBankIdAndCentre, organisationId, parentOrganisationId, associationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBacsAssociationExists("form3_bacs_association.association", &bacsResponse),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "bank_code", "1234"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "centre_number", "42"),
				),
			},
		},
	})
}

func TestAccBacsAssociation_withTestFileSubmissionFlag(t *testing.T) {
	var bacsResponse associations.GetBacsIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationId)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBacsAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3BacsAssociationConfigWithTestFileFlag, organisationId, parentOrganisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBacsAssociationExists("form3_bacs_association.association", &bacsResponse),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "test_file_submission", "true"),
				),
			},
		},
	})
}

func TestAccBacsAssociation_withMultiSunConfig(t *testing.T) {

	var bacsResponse associations.GetBacsIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationId)
	associationId := uuid.New().String()

	sun := "223344"
	sunConfig := models.BacsServiceUserNumber{
		ServiceUserNumber:   &sun,
		AutoReturnSortCode:  "123456",
		ContraAccountNumber: "12345678",
		ContraSortCode:      "123456",
	}
	sunConfigJson, jsonErr := json.Marshal(sunConfig)
	if jsonErr != nil {
		panic(jsonErr)
	}
	secondSun := strings.ReplaceAll(string(sunConfigJson), "\"", "\\\"")
	multiSunConfig := "[\"" + secondSun + "\"]"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBacsAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3BacsAssociationWithSunConfig, organisationId, parentOrganisationId, associationId, multiSunConfig),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBacsAssociationExists("form3_bacs_association.association", &bacsResponse),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "service_user_number", "112233"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "account_number", "87654321"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "sorting_code", "654321"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "account_type", "0"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "service_user_numbers_config.#", "1"),
					resource.TestCheckResourceAttr("form3_bacs_association.association", "service_user_numbers_config.0", string(sunConfigJson)),
				),
			},
		},
	})
}

func testAccCheckBacsAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_bacs_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("bacs record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckBacsAssociationExists(resourceKey string, association *associations.GetBacsIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("bacs not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no bacs Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("bacs record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		association = foundRecord

		return nil
	}
}

type associationConfigWithCerts struct {
	OrgID           string
	OrgParentID     string
	AssociationID   string
	InputKeyID      string
	InputCertID     string
	MessagingKeyID  string
	MessagingCertID string
	OutputKeyID     string
	OutputCertID    string
}

func makeTestForm3BacsAssociationConfigWithCerts(data associationConfigWithCerts) (string, error) {
	tfeTemplate := `
resource "form3_organisation" "organisation" {
	organisation_id        = "{{ .OrgID }}"
	parent_organisation_id = "{{ .OrgParentID }}"
	name 		           = "terraform-provider-form3-test-organisation"
}

resource "form3_key" "inputKey" {
	organisation_id = "${form3_organisation.organisation.organisation_id}"
	subject         = "CN=Terraform-test-with-selfsigned"
	key_id          = "{{ .InputKeyID }}"
}

resource "form3_certificate" "inputCert" {
	organisation_id         = "${form3_organisation.organisation.organisation_id}"
	key_id                  = "${form3_key.inputKey.key_id}"
	certificate_id          = "{{ .InputCertID }}"
}

resource "form3_key" "outputKey" {
	organisation_id = "${form3_organisation.organisation.organisation_id}"
	subject         = "CN=Terraform-test-with-selfsigned"
	key_id          = "{{ .OutputKeyID }}"
}

resource "form3_certificate" "outputCert" {
	organisation_id         = "${form3_organisation.organisation.organisation_id}"
	key_id                  = "${form3_key.outputKey.key_id}"
	certificate_id          = "{{ .OutputCertID }}"
}

resource "form3_bacs_association" "association" {
	organisation_id                  = "${form3_organisation.organisation.organisation_id}"
	association_id                   = "{{ .AssociationID }}"
	service_user_number              = "112238"
    account_number                   = "12345678"
    sorting_code                     = "123456"
    account_type                     = 1

    input_key_id                     = "{{ .InputKeyID }}"
    input_certificate_id             = "{{ .InputCertID }}"
    input_tsu_number                 = "B12345"

    messaging_key_id                 = "{{ .MessagingKeyID }}"
    messaging_certificate_id         = "{{ .MessagingCertID }}"
    messaging_tsu_number             = "B12346"

    output_key_id                    = "{{ .OutputKeyID }}"
    output_certificate_id            = "{{ .OutputCertID }}"
    output_tsu_number                = "B12347"
}`

	tpl := template.Must(template.New("tpl").Parse(tfeTemplate))

	var buf bytes.Buffer
	err := tpl.Execute(&buf, data)

	return buf.String(), err
}

const testForm3BacsAssociationConfigZeroAccountType = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		           = "terraform-provider-form3-test-organisation"
}

resource "form3_bacs_association" "association" {
	organisation_id                  = "${form3_organisation.organisation.organisation_id}"
	association_id                   = "%s"
	service_user_number              = "112233"
    account_number                   = "87654321"
    sorting_code                     = "654321"
    account_type                     = 0
}`

const testForm3BacsAssociationConfigWithBankIdAndCentre = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		           = "terraform-provider-form3-test-organisation"
}

resource "form3_bacs_association" "association" {
	organisation_id                  = "${form3_organisation.organisation.organisation_id}"
	association_id                   = "%s"
	service_user_number              = "112233"
    account_number                   = "87654321"
    sorting_code                     = "654321"
    account_type                     = 0
    bank_code                        = "1234"
    centre_number                    = "42"
}`

const testForm3BacsAssociationConfigWithTestFileFlag = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		           = "terraform-provider-form3-test-organisation"
}

resource "form3_bacs_association" "association" {
	organisation_id                  = "${form3_organisation.organisation.organisation_id}"
	association_id                   = "e7373962-b030-492f-b73b-68ca1e5c800e"
	service_user_number              = "112233"
    account_number                   = "87654321"
    sorting_code                     = "654321"
    account_type                     = 0
    bank_code                        = "1234"
    centre_number                    = "42"
    test_file_submission             = true
}`

const testForm3BacsAssociationWithSunConfig = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		           = "terraform-provider-form3-test-organisation"
}

resource "form3_bacs_association" "association" {
	organisation_id                  = "${form3_organisation.organisation.organisation_id}"
	association_id                   = "%s"
	service_user_number              = "112233"
    account_number                   = "87654321"
    sorting_code                     = "654321"
    account_type                     = 0
	service_user_numbers_config     = %s
}`
