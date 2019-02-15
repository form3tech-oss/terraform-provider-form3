package form3

import (
	"fmt"
	"github.com/form3tech-oss/go-form3"
	"github.com/form3tech-oss/go-form3/client/payment_defaults"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/satori/go.uuid"
	"os"
	"testing"
)

func TestAccPaymentDefaults_basic(t *testing.T) {
	var paymentDefaultsResponse payment_defaults.GetPaymentdefaultsIDOK
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.NewV4().String()
	id := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPaymentDefaultsDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3PaymentDefaultsConfig, organisationId, parentOrganisationId, id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPaymentDefaultsExists("form3_payment_defaults.payment_defaults", &paymentDefaultsResponse),
					resource.TestCheckResourceAttr("form3_payment_defaults.payment_defaults", "default_payment_scheme", "FPS"),
					resource.TestCheckResourceAttr("form3_payment_defaults.payment_defaults", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_payment_defaults.payment_defaults", "payment_defaults_id", id),
				),
			},
		},
	})
}

func testAccCheckPaymentDefaultsDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_payment_defaults" {
			continue
		}

		response, err := client.PaymentdefaultsClient.PaymentDefaults.GetPaymentdefaultsID(payment_defaults.NewGetPaymentdefaultsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("payment defaults record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckPaymentDefaultsExists(resourceKey string, paymentdefaultsIDOK *payment_defaults.GetPaymentdefaultsIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("payment defaults not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no payment defaults Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.PaymentdefaultsClient.PaymentDefaults.GetPaymentdefaultsID(payment_defaults.NewGetPaymentdefaultsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("payment defaults record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		paymentdefaultsIDOK = foundRecord

		return nil
	}
}

const testForm3PaymentDefaultsConfig = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-organisation"
}

resource "form3_payment_defaults" "payment_defaults" {
	organisation_id                  = "${form3_organisation.organisation.organisation_id}"
	payment_defaults_id              = "%s"
  default_payment_scheme           = "FPS"
}`
