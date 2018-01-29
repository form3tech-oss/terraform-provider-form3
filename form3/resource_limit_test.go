package form3

import (
	"fmt"
	"github.com/ewilde/go-form3"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/satori/go.uuid"
	"os"
	"testing"
  "github.com/ewilde/go-form3/client/payments"
)

func TestAccLimit_basic(t *testing.T) {
	var limitResponse payments.GetLimitsIDOK
	organisationId := os.Getenv("FORM3_ORGANISATION_ID")
	limitId := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLimitDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3LimitConfigA, organisationId, limitId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLimitExists("form3_limit.limit", &limitResponse),
					resource.TestCheckResourceAttr("form3_limit.limit", "amount", "1000"),
          resource.TestCheckResourceAttr("form3_limit.limit", "gateway", "payport_interface"),
          resource.TestCheckResourceAttr("form3_limit.limit", "scheme", "FPS"),
          resource.TestCheckResourceAttr("form3_limit.limit", "settlement_cycle_type", "daily"),
				),
			},
		},
	})
}

func TestAccLimit_importBasic(t *testing.T) {

	organisationId := os.Getenv("FORM3_ORGANISATION_ID")
	limitId := uuid.NewV4().String()

	resourceName := "form3_limit.limit"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLimitDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testForm3LimitConfigA, organisationId, limitId),
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckLimitDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_limit" {
			continue
		}

		response, err := client.PaymentsClient.Payments.GetLimitsID(payments.NewGetLimitsIDParams().WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckLimitExists(resourceKey string, limit *payments.GetLimitsIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.PaymentsClient.Payments.GetLimitsID(payments.NewGetLimitsIDParams().WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		limit = foundRecord

		return nil
	}
}

const testForm3LimitConfigA = `
resource "form3_limit" "limit" {
	organisation_id       = "%s"
	limit_id     	      	= "%s"
	amount     	        	= "1000"
  gateway               = "payport_interface"
  scheme                = "FPS"
  settlement_cycle_type = "daily"
}`
