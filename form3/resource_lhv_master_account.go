package form3

import (
	"errors"
	"fmt"
	"log"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceForm3LhvMasterAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceLhvMasterAccountCreate,
		Read:   resourceLhvMasterAccountRead,
		Delete: resourceLhvMasterAccountDelete,

		Schema: map[string]*schema.Schema{
			"association_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"master_account_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"iban": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bic": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bank_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"country": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"requires_direct_account": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  false,
			},
		},
	}
}

func resourceLhvMasterAccountCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, found := GetUUIDOK(d, "association_id")
	if !found {
		return errors.New("failed to read lhv master account association_id")
	}

	masterAccountResource, err := createLhvNewMasterAccountFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create lhv master account: %s", form3.JsonErrorPrettyPrint(err))
	}

	createdMasterAccount, err := client.AssociationClient.Associations.PostLhvAssociationIDMasterAccounts(
		associations.NewPostLhvAssociationIDMasterAccountsParams().
			WithAssociationID(associationId).
			WithCreationRequest(&models.LhvMasterAccountCreation{
				Data: masterAccountResource,
			}))
	if err != nil {
		return fmt.Errorf("failed to create lhv master account: %s", form3.JsonErrorPrettyPrint(err))
	}

	d.SetId(createdMasterAccount.Payload.Data.ID.String())
	log.Printf("[INFO] lhv master account id: %s", d.Id())

	return resourceLhvMasterAccountRead(d, meta)
}

func resourceLhvMasterAccountRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, found := GetUUIDOK(d, "association_id")
	if !found {
		return errors.New("failed to read lhv master account association_id")
	}

	masterAccount, err := client.AssociationClient.Associations.GetLhvAssociationIDMasterAccountsMasterAccountID(
		associations.NewGetLhvAssociationIDMasterAccountsMasterAccountIDParams().
			WithAssociationID(associationId).
			WithMasterAccountID(strfmt.UUID(d.Id())))

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find lhv sct association: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	_ = d.Set("country", masterAccount.Payload.Data.Attributes.Country)
	_ = d.Set("bic", masterAccount.Payload.Data.Attributes.Bic)
	_ = d.Set("bank_id", masterAccount.Payload.Data.Attributes.BankID)
	_ = d.Set("iban", masterAccount.Payload.Data.Attributes.Iban)
	_ = d.Set("requires_direct_account", masterAccount.Payload.Data.Attributes.RequiresDirectAccount)
	return nil
}

func resourceLhvMasterAccountDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationId, found := GetUUIDOK(d, "association_id")
	if !found {
		return errors.New("failed to read lhv master account association_id")
	}

	masterAccount, err := client.AssociationClient.Associations.GetLhvAssociationIDMasterAccountsMasterAccountID(
		associations.NewGetLhvAssociationIDMasterAccountsMasterAccountIDParams().
			WithAssociationID(associationId).
			WithMasterAccountID(strfmt.UUID(d.Id())))
	if err != nil {
		return fmt.Errorf("error deleting lhv master account: %s", form3.JsonErrorPrettyPrint(err))
	}

	_, err = client.AssociationClient.Associations.DeleteLhvAssociationIDMasterAccountsMasterAccountID(
		associations.NewDeleteLhvAssociationIDMasterAccountsMasterAccountIDParams().
			WithAssociationID(associationId).
			WithMasterAccountID(masterAccount.Payload.Data.ID).
			WithVersion(masterAccount.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting lhv sct association: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createLhvNewMasterAccountFromResourceData(d *schema.ResourceData) (*models.LhvMasterAccount, error) {
	masterAccount := models.LhvMasterAccount{
		Type: models.LhvMasterAccountTypeLhvgatewayMasterAccounts,
	}

	if attr, ok := GetUUIDOK(d, "master_account_id"); ok {
		masterAccount.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		masterAccount.OrganisationID = attr
	}

	if attr, ok := d.GetOk("country"); ok {
		masterAccount.Attributes.Country = attr.(string)
	}

	if attr, ok := d.GetOk("bic"); ok {
		masterAccount.Attributes.Bic = attr.(string)
	}

	if attr, ok := d.GetOk("bank_id"); ok {
		masterAccount.Attributes.BankID = attr.(string)
	}

	if attr, ok := d.GetOk("iban"); ok {
		masterAccount.Attributes.Iban = attr.(string)
	}

	if attr, ok := d.GetOk("requires_direct_account"); ok {
		masterAccount.Attributes.RequiresDirectAccount = attr.(*bool)
	}

	return &masterAccount, nil
}
