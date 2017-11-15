package form3

import (
	"github.com/ewilde/go-form3/client/subscriptions"
	"github.com/ewilde/go-form3/client/users"
	"github.com/ewilde/go-form3/models"
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
					UserID:            getUserId(),
					CallbackTransport: "queue",
					CallbackURI:       "https://sqs.eu-west-1.amazonaws.com/134201431238/notification-test",
					EventType:         "updated",
					RecordType:        "foo_bar",
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
func getUserId() strfmt.UUID {
	response, err := auth.SecurityClient.Users.GetUsers(users.NewGetUsersParams())
	if err == nil {
		return ""
	}

	return response.Payload.Data[0].ID
}
