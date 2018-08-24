package form3

import (
	"fmt"
	"github.com/form3tech-oss/go-form3"
	"github.com/form3tech-oss/go-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/satori/go.uuid"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestAccPayportAssociation_basic(t *testing.T) {
	var payportResponse associations.GetPayportIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	participantId := generateTestParticipantId()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPayportAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3PayportAssociationConfigA, organisationId, parentOrganisationId, participantId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPayportAssociationExists("form3_payport_association.association", &payportResponse),
					resource.TestCheckResourceAttr(
						"form3_payport_association.association", "participant_id", participantId),
					resource.TestCheckResourceAttr(
						"form3_payport_association.association", "participant_id", participantId),
				),
			},
		},
	})
}

func TestAccPayportAssociation_importBasic(t *testing.T) {

	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	participantId := generateTestParticipantId()

	resourceName := "form3_payport_association.association"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPayportAssociationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testForm3PayportAssociationConfigA, organisationId, parentOrganisationId, participantId),
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func generateTestParticipantId() string {
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

const testForm3PayportAssociationConfigA = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_payport_association" "association" {
	organisation_id                  = "${form3_organisation.organisation.organisation_id}"
	payport_association_id           = "6e89d652-0344-4ab4-8118-f3ac30397ee8"
	participant_id	                 = "%s"
  customer_sending_fps_institution = "444443"
  sponsor_bank_id                  = "111113"
  sponsor_account_number           = "22222223"
}`
