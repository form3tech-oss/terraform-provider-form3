package form3

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/accounts"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccBic_basic(t *testing.T) {
	var bicResponse accounts.GetBicsIDOK
	organisationId := os.Getenv("FORM3_ORGANISATION_ID")
	bicId := uuid.New().String()
	bic := generateTestBic()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBicDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3BicConfigA, organisationId, bicId, bic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBicExists("form3_bic.bic", &bicResponse),
					resource.TestCheckResourceAttr(
						"form3_bic.bic", "bic", bic),
				),
			},
		},
	})
}

func TestAccBic_importBasic(t *testing.T) {

	organisationId := os.Getenv("FORM3_ORGANISATION_ID")
	bicId := uuid.New().String()
	bic := generateTestBic()

	resourceName := "form3_bic.bic"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBicDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testForm3BicConfigA, organisationId, bicId, bic),
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func generateTestBic() string {
	return generateTestBicWithLength(8)
}

func generateTestBicWithLength(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	var characters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 6)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}

	characters = []rune("23456789")
	c := make([]rune, length-6)
	for i := range c {
		c[i] = characters[rand.Intn(len(characters))]
	}

	return string(b) + string(c)
}

func testAccCheckBicDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_bic" {
			continue
		}

		response, err := client.AccountClient.Accounts.GetBicsID(accounts.NewGetBicsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckBicExists(resourceKey string, bicResponse *accounts.GetBicsIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AccountClient.Accounts.GetBicsID(accounts.NewGetBicsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		bicResponse = foundRecord

		return nil
	}
}

const testForm3BicConfigA = `
resource "form3_bic" "bic" {
	organisation_id = "%s"
  bic_id          = "%s"
	bic       	    = "%s"
}
`
