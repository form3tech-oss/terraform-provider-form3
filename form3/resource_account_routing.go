package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/account_routings"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3AccountRouting() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountRoutingCreate,
		Read:   resourceAccountRoutingRead,
		Delete: resourceAccountRoutingDelete,

		Schema: map[string]*schema.Schema{
			"account_routing_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_generator": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_provisioner": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAccountRoutingCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	accountRoutingResourceID := d.Get("account_routing_id").(string)
	log.Printf("[INFO] Creating account routing with id: %s", accountRoutingResourceID)

	accountRouting, err := createAccountRoutingFromResourceData(d)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] account routing create: %#v", accountRouting)

	createdAccountRouting, err := client.AccountClient.AccountRoutings.PostAccountRoutings(account_routings.NewPostAccountRoutingsParams().
		WithAccountRoutingCreationRequest(&models.AccountRoutingCreation{
			Data: accountRouting,
		}))

	if err != nil {
		return fmt.Errorf("failed to create account routing: %s", err)
	}
	d.SetId(createdAccountRouting.Payload.Data.ID.String())
	log.Printf("[INFO] account routing id: %s", d.Id())

	return resourceAccountRoutingRead(d, meta)
}

func resourceAccountRoutingRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	accountRoutingResourceID, _ := GetUUIDOK(d, "account_routing_id")

	if accountRoutingResourceID == "" {
		accountRoutingResourceID = strfmt.UUID(d.Id())
		log.Printf("[INFO] Importing account routing with resource id: %s.", accountRoutingResourceID)
	} else {
		log.Printf("[INFO] Reading account routing with resource id: %s.", accountRoutingResourceID)
	}

	account, err := client.AccountClient.AccountRoutings.GetAccountRoutingsID(account_routings.NewGetAccountRoutingsIDParams().
		WithID(accountRoutingResourceID))

	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find account routing : %s", err)
	}

	d.Set("account_routing_id", account.Payload.Data.ID.String())
	d.Set("organisation_id", account.Payload.Data.OrganisationID.String())
	d.Set("account_generator", account.Payload.Data.Attributes.AccountGenerator)
	d.Set("account_provisioner", account.Payload.Data.Attributes.AccountProvisioner)
	d.Set("match", account.Payload.Data.Attributes.Match)
	d.Set("priority", account.Payload.Data.Attributes.Priority)
	return nil
}

func resourceAccountRoutingDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	accountRouting, err := createAccountRoutingFromResourceData(d)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Deleting account routing with resource id: %s", accountRouting.ID)

	_, err = client.AccountClient.AccountRoutings.DeleteAccountRoutingsID(account_routings.NewDeleteAccountRoutingsIDParams().
		WithID(form3.UUIDValue(accountRouting.ID)))

	if err != nil {
		return fmt.Errorf("error deleting account routing: %s", err)
	}

	return nil
}

func createAccountRoutingFromResourceData(d *schema.ResourceData) (*models.AccountRouting, error) {

	account := models.AccountRouting{Attributes: &models.AccountRoutingAttributes{}}
	account.Type = "account_routings"
	if attr, ok := GetUUIDOK(d, "account_routing_id"); ok {
		account.ID = form3.UUID(attr)
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		account.OrganisationID = form3.UUID(attr)
	}

	if attr, ok := d.GetOk("account_generator"); ok {
		gen := attr.(string)
		account.Attributes.AccountGenerator = &gen
	}

	if attr, ok := d.GetOk("account_provisioner"); ok {
		gen := attr.(string)
		account.Attributes.AccountProvisioner = &gen
	}

	if attr, ok := d.GetOk("match"); ok {
		gen := attr.(string)
		account.Attributes.Match = &gen
	}

	if attr, ok := d.GetOk("priority"); ok {
		gen := int64(attr.(int))
		account.Attributes.Priority = &gen
	}
	return &account, nil
}
