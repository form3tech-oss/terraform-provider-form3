package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/platformsecurityapi"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

func TestAccSigningKeys_basic(t *testing.T) {
	var signingKeysResponse platformsecurityapi.GetPlatformSecuritySigningKeysSigningkeyIDOK
	organisationId := os.Getenv("FORM3_ORGANISATION_ID")
	signingKeyId := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSigningKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3SigningKeyConfigA, organisationId, signingKeyId),
				Check: resource.ComposeTestCheckFunc(testAccCheckSigningKeyExists("form3_signing_keys.signing_key", &signingKeysResponse),
					resource.TestCheckResourceAttr("form3_signing_keys.signing_key", "organisation_id", organisationId),
				),
			},
		},
	})
}

func testAccCheckSigningKeyExists(resourceKey string, p *platformsecurityapi.GetPlatformSecuritySigningKeysSigningkeyIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.PlatformClient.Platformsecurityapi.GetPlatformSecuritySigningKeysSigningkeyID(
			platformsecurityapi.NewGetPlatformSecuritySigningKeysSigningkeyIDParams().WithSigningkeyID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		p = foundRecord

		return nil
	}
}

func testAccCheckSigningKeyDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_signing_key" {
			continue
		}

		response, err := client.PlatformClient.Platformsecurityapi.GetPlatformSecuritySigningKeysSigningkeyID(
			platformsecurityapi.NewGetPlatformSecuritySigningKeysSigningkeyIDParams().
				WithSigningkeyID(strfmt.UUID(rs.Primary.ID)),
		)

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

const testForm3SigningKeyConfigA = `
resource "form3_signing_keys" "signing_key" {
	organisation_id  = "%s"
	signing_key_id               = "%s"
}
`
