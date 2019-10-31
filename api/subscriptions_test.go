package api

import (
	"github.com/form3tech-oss/terraform-provider-form3/client/subscriptions"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"testing"
)

func TestAccDeleteSubscription(t *testing.T) {

	createResponse, err := auth.NotificationClient.Subscriptions.PostSubscriptions(subscriptions.NewPostSubscriptionsParams().
		WithSubscriptionCreationRequest(&models.SubscriptionCreation{
			Data: &models.Subscription{
				OrganisationID: testOrganisationId,
				Type:           "subscriptions",
				ID:             strfmt.UUID("5e950680-1ea2-4898-ba0f-632214f51946"),
				Attributes: &models.SubscriptionAttributes{
					CallbackTransport: "queue",
					CallbackURI:       "https://sqs.eu-west-1.amazonaws.com/288840537196/notification-test",
					EventType:         "updated",
					RecordType:        "payments",
				},
			},
		}))

	assertNoErrorOccurred(err, t)

	_, err = auth.NotificationClient.Subscriptions.DeleteSubscriptionsID(subscriptions.NewDeleteSubscriptionsIDParams().
		WithID(createResponse.Payload.Data.ID),
	)

	if err != nil {
		t.Error(err)
	}

	_, err = auth.NotificationClient.Subscriptions.GetSubscriptionsID(subscriptions.NewGetSubscriptionsIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(err, t, 404)
}
