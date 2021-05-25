package form3

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPayportAssociation_basic_non_settling(t *testing.T) {
	var payportResponse associations.GetPayportIDOK
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)

	participantID := generateTestParticipantID()
	associationID := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPayportAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3PayportAssociationConfigNonSettling(organisationID, parentOrganisationID, testOrgName, associationID, participantID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPayportAssociationExists("form3_payport_association.association", &payportResponse),
					resource.TestCheckResourceAttrSet("form3_payport_association.association", "payport_association_id"),
					resource.TestCheckResourceAttr("form3_payport_association.association", "organisation_id", organisationID),
					resource.TestCheckResourceAttr("form3_payport_association.association", "participant_id", participantID),
					resource.TestCheckResourceAttr("form3_payport_association.association", "participant_type", "non_settling"),
					resource.TestCheckResourceAttr("form3_payport_association.association", "customer_sending_fps_institution", "444443"),
					resource.TestCheckResourceAttr("form3_payport_association.association", "sponsor_bank_id", "111113"),
					resource.TestCheckResourceAttr("form3_payport_association.association", "sponsor_account_number", "22222223"),
				),
			},
		},
	})
}

func TestAccPayportAssociation_basic_settling(t *testing.T) {
	var payportResponse associations.GetPayportIDOK
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)

	participantID := generateTestParticipantID()
	associationID := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPayportAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3PayportAssociationConfigSettling(organisationID, parentOrganisationID, testOrgName, associationID, participantID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPayportAssociationExists("form3_payport_association.association", &payportResponse),
					resource.TestCheckResourceAttrSet("form3_payport_association.association", "payport_association_id"),
					resource.TestCheckResourceAttr("form3_payport_association.association", "organisation_id", organisationID),
					resource.TestCheckResourceAttr("form3_payport_association.association", "participant_id", participantID),
					resource.TestCheckResourceAttr("form3_payport_association.association", "participant_type", "settling"),
					resource.TestCheckResourceAttr("form3_payport_association.association", "customer_sending_fps_institution", "444443"),
					resource.TestCheckResourceAttr("form3_payport_association.association", "sponsor_bank_id", ""),
					resource.TestCheckResourceAttr("form3_payport_association.association", "sponsor_account_number", ""),
				),
			},
		},
	})
}

func TestAccPayportAssociation_importBasic(t *testing.T) {
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)

	participantID := generateTestParticipantID()
	associationID := uuid.New().String()

	resourceName := "form3_payport_association.association"

	cases := map[string]func(orgID, parOrgID, orgName, payassocID, partID string) string{
		"non_settling": getTestForm3PayportAssociationConfigNonSettling,
		"settling":     getTestForm3PayportAssociationConfigSettling,
	}
	for name, config := range cases {
		t.Run(name, func(t *testing.T) {

			resource.Test(t, resource.TestCase{
				PreCheck:     func() { testAccPreCheck(t) },
				Providers:    testAccProviders,
				CheckDestroy: testAccCheckPayportAssociationDestroy,
				Steps: []resource.TestStep{
					{
						Config: config(organisationID, parentOrganisationID, testOrgName, associationID, participantID),
					},
					{
						ResourceName:      resourceName,
						ImportState:       true,
						ImportStateVerify: true,
					},
				},
			})

		})
	}
}

func generateTestParticipantID() string {
	rand.Seed(time.Now().UTC().UnixNano())
	var characters = []rune("0123456789")
	b := make([]rune, 8)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}

	return string(b)
}

func testAccCheckPayportAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_payport_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetPayportID(associations.NewGetPayportIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckPayportAssociationExists(resourceKey string, association *associations.GetPayportIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]
		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetPayportID(associations.NewGetPayportIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))
		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		association = foundRecord

		return nil
	}
}

func getTestForm3PayportAssociationConfigNonSettling(orgID, parOrgID, orgName, payassocID, partID string) string {
	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		               = "%s"
	}

	resource "form3_payport_association" "association" {
		organisation_id                  = "${form3_organisation.organisation.organisation_id}"
		payport_association_id           = "%s"
		participant_id	                 = "%s"
		participant_type                 = "non_settling"
	  customer_sending_fps_institution = "444443"
	  sponsor_bank_id                  = "111113"
	  sponsor_account_number           = "22222223"
	}`, orgID, parOrgID, orgName, payassocID, partID)
}

func getTestForm3PayportAssociationConfigSettling(orgID, parOrgID, orgName, payassocID, partID string) string {
	return fmt.Sprintf(`
	resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		               = "%s"
	}

	resource "form3_payport_association" "association" {
		organisation_id                  = "${form3_organisation.organisation.organisation_id}"
		payport_association_id           = "%s"
		participant_id	                 = "%s"
		participant_type                 = "settling"
	  customer_sending_fps_institution = "444443"
	}`, orgID, parOrgID, orgName, payassocID, partID)
}
