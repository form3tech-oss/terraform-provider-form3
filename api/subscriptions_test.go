package api

import (
	"log"
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/client/subscriptions"
	"github.com/form3tech-oss/terraform-provider-form3/models"
)

func TestAccDeleteSubscription(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)

	id := NewUUID()

	defer func() {
		if t.Failed() {
			if _, err := auth.NotificationClient.Subscriptions.DeleteSubscriptionsID(subscriptions.NewDeleteSubscriptionsIDParams().
				WithID(*id).WithVersion(0),
			); err != nil {
				log.Printf("[CLEANUP] Did not delete subscription id, error %s\n", JsonErrorPrettyPrint(err))
			} else {
				log.Printf("[CLEANUP] Successfully subscription id\n")
			}
		}
	}()
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)

	createResponse, err := auth.NotificationClient.Subscriptions.PostSubscriptions(subscriptions.NewPostSubscriptionsParams().
		WithSubscriptionCreationRequest(&models.SubscriptionCreation{
			Data: &models.Subscription{
				OrganisationID: testOrganisationId,
				Type:           "subscriptions",
				ID:             *id,
				Attributes: &models.SubscriptionAttributes{
					CallbackTransport: "queue",
					CallbackURI:       "https://sqs.eu-west-1.amazonaws.com/288840537196/notification-test",
					EventType:         "updated",
					RecordType:        "payments",
				},
			},
		}))

	assertNoErrorOccurred(t, err)

	_, err = auth.NotificationClient.Subscriptions.DeleteSubscriptionsID(subscriptions.NewDeleteSubscriptionsIDParams().
		WithID(createResponse.Payload.Data.ID).WithVersion(0),
	)

	assertNoErrorOccurred(t, err)

	_, err = auth.NotificationClient.Subscriptions.GetSubscriptionsID(subscriptions.NewGetSubscriptionsIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertStatusCode(t, err, 404)
}
