package form3

import (
	"fmt"
	"github.com/ewilde/go-form3"
	"github.com/ewilde/go-form3/client/subscriptions"
	"github.com/ewilde/go-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3Subscription() *schema.Resource {
	return &schema.Resource{
		Create: resourceSubscriptionCreate,
		Read:   resourceSubscriptionRead,
		Update: resourceSubscriptionUpdate,
		Delete: resourceSubscriptionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"organisation_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subscription_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"callback_transport": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"callback_uri": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"event_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"record_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSubscriptionCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	subscriptionId := d.Get("subscription_id").(string)
	log.Printf("[INFO] Creating subscription with id: %s", subscriptionId)

	subscription, err := createSubscriptionFromResourceData(d)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] subscription create: %#v", subscription)

	createdSubscription, err := client.NotificationClient.Subscriptions.PostSubscriptions(subscriptions.NewPostSubscriptionsParams().
		WithSubscriptionCreationRequest(&models.SubscriptionCreation{Data: subscription}))

	if err != nil {
		return fmt.Errorf("failed to create subscription: %s", err)
	}

	d.SetId(createdSubscription.Payload.Data.ID.String())
	log.Printf("[INFO] subscription key: %s", d.Id())

	return resourceSubscriptionRead(d, meta)
}

func resourceSubscriptionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	key := d.Id()
	subscriptionId, _ := GetUUIDOK(d, "subscription_id")

	if subscriptionId == "" {
		subscriptionId = strfmt.UUID(key)
		log.Printf("[INFO] Importing subscription for id: %s ", key)
	} else {
		log.Printf("[INFO] Reading subscription for id: %s ", subscriptionId)
	}

	subscription, err := client.NotificationClient.Subscriptions.GetSubscriptionsID(
		subscriptions.NewGetSubscriptionsIDParams().WithID(subscriptionId))

	if err != nil {
		apiError := err.(*runtime.APIError)
		if apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find subscription: %s", err)
	}

	d.Set("subscription_id", subscription.Payload.Data.ID.String())
	d.Set("organisation_id", subscription.Payload.Data.OrganisationID.String())
	d.Set("callback_transport", subscription.Payload.Data.Attributes.CallbackTransport)
	d.Set("callback_uri", subscription.Payload.Data.Attributes.CallbackURI)
	d.Set("event_type", subscription.Payload.Data.Attributes.EventType)
	d.Set("record_type", subscription.Payload.Data.Attributes.RecordType)

	return nil
}

func resourceSubscriptionUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(false)

	client := meta.(*form3.AuthenticatedClient)
	subscriptionFromResource, err := createSubscriptionFromResourceDataWithVersion(d, client)
	if err != nil {
		return fmt.Errorf("error updating subscription: %s", err)
	}

	_, err = client.NotificationClient.Subscriptions.PatchSubscriptionsID(subscriptions.NewPatchSubscriptionsIDParams().
		WithID(subscriptionFromResource.ID).
		WithSubscriptionUpdateRequest(&models.SubscriptionCreation{Data: subscriptionFromResource}))

	if err != nil {
		return fmt.Errorf("error updating subscription: %s", err)
	}

	return nil
}

func resourceSubscriptionDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	subscriptionFromResource, err := createSubscriptionFromResourceDataWithVersion(d, client)
	if err != nil {
		return fmt.Errorf("error deleting subscription: %s", err)
	}

	log.Printf("[INFO] Deleting subscription for id: %s", subscriptionFromResource.ID)

	_, err = client.NotificationClient.Subscriptions.DeleteSubscriptionsID(subscriptions.NewDeleteSubscriptionsIDParams().
		WithID(subscriptionFromResource.ID).
		WithVersion(*subscriptionFromResource.Version))

	if err != nil {
		return fmt.Errorf("error deleting subscription: %s", err)
	}

	return nil
}

func createSubscriptionFromResourceDataWithVersion(d *schema.ResourceData, client *form3.AuthenticatedClient) (*models.Subscription, error) {
	subscription, err := createSubscriptionFromResourceData(d)
	version, err := getSubscriptionVersion(client, subscription.ID)
	if err != nil {
		return nil, err
	}

	subscription.Version = &version

	return subscription, nil
}

func createSubscriptionFromResourceData(d *schema.ResourceData) (*models.Subscription, error) {

	subscription := models.Subscription{Attributes: &models.SubscriptionAttributes{}}
	subscription.Type = "subscriptions"
	if attr, ok := GetUUIDOK(d, "subscription_id"); ok {
		subscription.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		subscription.OrganisationID = attr
	}

	if attr, ok := d.GetOk("callback_transport"); ok {
		subscription.Attributes.CallbackTransport = attr.(string)
	}

	if attr, ok := d.GetOk("callback_uri"); ok {
		subscription.Attributes.CallbackURI = attr.(string)
	}

	if attr, ok := d.GetOk("event_type"); ok {
		subscription.Attributes.EventType = attr.(string)
	}

	if attr, ok := d.GetOk("record_type"); ok {
		subscription.Attributes.RecordType = attr.(string)
	}

	return &subscription, nil
}

func getSubscriptionVersion(client *form3.AuthenticatedClient, subscriptionId strfmt.UUID) (int64, error) {
	subscription, err := client.NotificationClient.Subscriptions.GetSubscriptionsID(subscriptions.NewGetSubscriptionsIDParams().
		WithID(subscriptionId))
	if err != nil {
		if err != nil {
			return -1, fmt.Errorf("error reading subscription: %s", err)
		}
	}

	return *subscription.Payload.Data.Version, nil
}
