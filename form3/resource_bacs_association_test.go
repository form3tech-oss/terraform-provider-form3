package form3

import (
  "fmt"
  "github.com/ewilde/go-form3"
  "github.com/ewilde/go-form3/client/associations"
  "github.com/go-openapi/strfmt"
  "github.com/hashicorp/terraform/helper/resource"
  "github.com/hashicorp/terraform/terraform"
  "github.com/satori/go.uuid"
  "math/rand"
  "os"
  "testing"
  "time"
)

func TestAccBacsAssociation_basic(t *testing.T) {
  var bacsResponse associations.GetBacsIDOK
  parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
  organisationId := uuid.NewV4().String()
  participantId := generateBacsTestParticipantId()

  resource.Test(t, resource.TestCase{
    PreCheck:     func() { testAccPreCheck(t) },
    Providers:    testAccProviders,
    CheckDestroy: testAccCheckBacsAssociationDestroy,
    Steps: []resource.TestStep{
      {
        Config: fmt.Sprintf(testForm3BacsAssociationConfigA, organisationId, parentOrganisationId, participantId),
        Check: resource.ComposeTestCheckFunc(
          testAccCheckBacsAssociationExists("form3_bacs_association.association", &bacsResponse),
          resource.TestCheckResourceAttr(
            "form3_bacs_association.association", "participant_id", participantId),
          resource.TestCheckResourceAttr(
            "form3_bacs_association.association", "participant_id", participantId),
        ),
      },
    },
  })
}

func TestAccBacsAssociation_importBasic(t *testing.T) {

  parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
  organisationId := uuid.NewV4().String()
  participantId := generateBacsTestParticipantId()

  resourceName := "form3_bacs_association.association"

  resource.Test(t, resource.TestCase{
    PreCheck:     func() { testAccPreCheck(t) },
    Providers:    testAccProviders,
    CheckDestroy: testAccCheckBacsAssociationDestroy,
    Steps: []resource.TestStep{
      resource.TestStep{
        Config: fmt.Sprintf(testForm3BacsAssociationConfigA, organisationId, parentOrganisationId, participantId),
      },

      resource.TestStep{
        ResourceName:      resourceName,
        ImportState:       true,
        ImportStateVerify: true,
      },
    },
  })
}

func generateBacsTestParticipantId() string {
  rand.Seed(time.Now().UTC().UnixNano())
  var characters = []rune("0123456789")
  b := make([]rune, 8)
  for i := range b {
    b[i] = characters[rand.Intn(len(characters))]
  }

  return string(b)
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
      return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
    }
  }

  return nil
}

func testAccCheckBacsAssociationExists(resourceKey string, association *associations.GetBacsIDOK) resource.TestCheckFunc {
  return func(s *terraform.State) error {
    rs, ok := s.RootModule().Resources[resourceKey]

    if !ok {
      return fmt.Errorf("not found: %s", resourceKey)
    }

    if rs.Primary.ID == "" {
      return fmt.Errorf("no Record ID is set")
    }

    client := testAccProvider.Meta().(*form3.AuthenticatedClient)

    foundRecord, err := client.AssociationClient.Associations.GetBacsID(associations.NewGetBacsIDParams().
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

const testForm3BacsAssociationConfigA = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_bacs_association" "association" {
	organisation_id                  = "${form3_organisation.organisation.organisation_id}"
	bacs_association_id           = "6e89d652-0344-4ab4-8118-f3ac30397ee8"
	participant_id	                 = "%s"
  customer_sending_fps_institution = "444443"
  sponsor_bank_id                  = "111113"
  sponsor_account_number           = "22222223"
}`
