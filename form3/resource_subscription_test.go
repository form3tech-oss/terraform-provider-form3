package form3

import (
	"fmt"
	"os"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/subscriptions"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccSubscription_basic(t *testing.T) {
	var subscriptionResponse subscriptions.GetSubscriptionsIDOK
	organisationId := os.Getenv("FORM3_ORGANISATION_ID")
	subscriptionId := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3SubscriptionConfigA, organisationId, subscriptionId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriptionExists("form3_subscription.subscription", &subscriptionResponse),
					resource.TestCheckResourceAttr(
						"form3_subscription.subscription", "callback_transport", "queue"),
					resource.TestCheckResourceAttr(
						"form3_subscription.subscription", "callback_uri", "https://sqs.eu-west-1.amazonaws.com/288840537196/notification-test"),
					resource.TestCheckResourceAttr(
						"form3_subscription.subscription", "event_type", "Updated"),
					resource.TestCheckResourceAttr(
						"form3_subscription.subscription", "record_type", "PaymentAdmission"),
				),
			},
			{
				Config: fmt.Sprintf(testForm3SubscriptionConfigAUpdate, organisationId, subscriptionId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriptionExists("form3_subscription.subscription", &subscriptionResponse),
					resource.TestCheckResourceAttr(
						"form3_subscription.subscription", "callback_uri", "https://sqs.eu-west-1.amazonaws.com/288840537196/notification-test-2"),
				),
			},
		},
	})
}

func TestAccSubscription_importBasic(t *testing.T) {

	organisationId := os.Getenv("FORM3_ORGANISATION_ID")
	subscriptionId := uuid.New().String()

	resourceName := "form3_subscription.subscription"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSubscriptionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testForm3SubscriptionConfigA, organisationId, subscriptionId),
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckSubscriptionDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_subscription" {
			continue
		}

		response, err := client.NotificationClient.Subscriptions.GetSubscriptionsID(subscriptions.NewGetSubscriptionsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckSubscriptionExists(resourceKey string, subscription *subscriptions.GetSubscriptionsIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.NotificationClient.Subscriptions.GetSubscriptionsID(subscriptions.NewGetSubscriptionsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		subscription = foundRecord

		return nil
	}
}

const testForm3SubscriptionConfigA = `
resource "form3_subscription" "subscription" {
	organisation_id    = "%s"
	subscription_id    = "%s"
	callback_transport = "queue"
  	callback_uri       = "https://sqs.eu-west-1.amazonaws.com/288840537196/notification-test"
  	event_type         = "Updated"
  	record_type        = "PaymentAdmission"
}`

const testForm3SubscriptionConfigAUpdate = `
resource "form3_subscription" "subscription" {
	organisation_id    = "%s"
	subscription_id    = "%s"
	callback_transport = "queue"
  	callback_uri       = "https://sqs.eu-west-1.amazonaws.com/288840537196/notification-test-2"
  	event_type         = "Updated"
  	record_type        = "PaymentAdmission"
}`
