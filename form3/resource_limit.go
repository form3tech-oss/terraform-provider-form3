package form3

import (
	"fmt"
	"log"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/limits"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceForm3Limit() *schema.Resource {
	return &schema.Resource{
		Create: resourceLimitCreate,
		Read:   resourceLimitRead,
		Delete: resourceLimitDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"limit_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"amount": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"gateway": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"scheme": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"settlement_cycle_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceLimitCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	amount := d.Get("amount").(string)
	log.Printf("[INFO] Creating limit with amount: %s", amount)

	limit, err := createLimitFromResourceData(d)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] limit create: %#v", limit)

	createdLimit, err := client.LimitsClient.Limits.PostLimits(limits.NewPostLimitsParams().
		WithLimitCreationRequest(&models.LimitCreation{
			Data: limit,
		}))

	if err != nil {
		return fmt.Errorf("failed to create limit: %s", form3.JsonErrorPrettyPrint(err))
	}

	d.SetId(createdLimit.Payload.Data.ID.String())
	log.Printf("[INFO] limit key: %s", d.Id())

	return resourceLimitRead(d, meta)
}

func resourceLimitRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	key := d.Id()
	limitId, _ := GetUUIDOK(d, "limit_id")
	amount := d.Get("amount").(string)

	if limitId == "" {
		limitId = strfmt.UUID(key)
		log.Printf("[INFO] Importing limit id: %s", limitId)
	} else {
		log.Printf("[INFO] Reading limit for id: %s amount: %s", key, amount)
	}

	limit, err := client.LimitsClient.Limits.GetLimitsID(limits.NewGetLimitsIDParams().WithID(limitId))
	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find limit: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	d.Set("limit_id", limit.Payload.Data.ID.String())
	d.Set("amount", limit.Payload.Data.Attributes.Amount)
	d.Set("organisation_id", limit.Payload.Data.OrganisationID.String())
	d.Set("gateway", limit.Payload.Data.Attributes.Gateway)
	d.Set("scheme", limit.Payload.Data.Attributes.Scheme)
	d.Set("settlement_cycle_type", limit.Payload.Data.Attributes.SettlementCycleType)

	return nil
}

func resourceLimitDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	limitFromResource, err := createLimitFromResourceDataWithVersion(d, client)
	if err != nil {
		return fmt.Errorf("error deleting limit: %s", form3.JsonErrorPrettyPrint(err))
	}

	log.Printf("[INFO] Deleting limit for id: %s amount: %s", limitFromResource.ID, limitFromResource.Attributes.Amount)

	_, err = client.LimitsClient.Limits.DeleteLimitsID(limits.NewDeleteLimitsIDParams().
		WithID(limitFromResource.ID).
		WithVersion(*limitFromResource.Version))

	if err != nil {
		return fmt.Errorf("error deleting limit: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createLimitFromResourceDataWithVersion(d *schema.ResourceData, client *form3.AuthenticatedClient) (*models.Limit, error) {
	limit, err := createLimitFromResourceData(d)
	if err != nil {
		return nil, err
	}
	version, err := getLimitVersion(client, limit.ID)
	if err != nil {
		return nil, err
	}

	limit.Version = &version

	return limit, nil
}

func createLimitFromResourceData(d *schema.ResourceData) (*models.Limit, error) {

	limit := models.Limit{Attributes: &models.LimitAttributes{}}
	limit.Type = "limits"
	if attr, ok := GetUUIDOK(d, "limit_id"); ok {
		limit.ID = attr
	}

	if attr, ok := d.GetOk("amount"); ok {
		limit.Attributes.Amount = attr.(string)
	}

	if attr, ok := d.GetOk("gateway"); ok {
		limit.Attributes.Gateway = attr.(string)
	}

	if attr, ok := d.GetOk("scheme"); ok {
		limit.Attributes.Scheme = attr.(string)
	}

	if attr, ok := d.GetOk("settlement_cycle_type"); ok {
		val := attr.(string)
		limit.Attributes.SettlementCycleType = models.SettlementCycleType(val)
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		limit.OrganisationID = attr
	}

	return &limit, nil
}

func getLimitVersion(client *form3.AuthenticatedClient, limitId strfmt.UUID) (int64, error) {
	limit, err := client.LimitsClient.Limits.GetLimitsID(limits.NewGetLimitsIDParams().WithID(limitId))
	if err != nil {
		return -1, fmt.Errorf("error reading limit: %s", form3.JsonErrorPrettyPrint(err))
	}

	return *limit.Payload.Data.Version, nil
}
