package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/accounts"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3Account() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountCreate,
		Read:   resourceAccountRead,
		Delete: resourceAccountDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"bank_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bank_id_code": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"iban": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
		},
	}
}

func resourceAccountCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*form3.AuthenticatedClient)

	accountResourceID := d.Get("account_id").(string)
	log.Printf("[INFO] Creating account with id: %s", accountResourceID)

	account, err := createAccountFromResourceData(d)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] account create: %#v", account)

	createdAccount, err := client.AccountClient.Accounts.PostAccounts(accounts.NewPostAccountsParams().
		WithAccountCreationRequest(&models.AccountCreation{
			Data: account,
		}))

	if err != nil {
		return fmt.Errorf("failed to create account: %s", err)
	}

	d.SetId(createdAccount.Payload.Data.ID.String())
	log.Printf("[INFO] account key: %s", d.Id())

	return resourceAccountRead(d, meta)
}

func resourceAccountRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	accountResourceID, _ := GetUUIDOK(d, "account_id")

	if accountResourceID == "" {
		accountResourceID = strfmt.UUID(d.Id())
		log.Printf("[INFO] Importing account with resource id: %s.", accountResourceID)
	} else {
		log.Printf("[INFO] Reading account with resource id: %s.", accountResourceID)
	}

	account, err := client.AccountClient.Accounts.GetAccountsID(accounts.NewGetAccountsIDParams().
		WithID(accountResourceID))

	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find account: %s", err)
	}

	d.Set("account_id", account.Payload.Data.ID.String())
	d.Set("organisation_id", account.Payload.Data.OrganisationID.String())
	d.Set("bank_id", account.Payload.Data.Attributes.BankID)
	d.Set("bank_id_code", account.Payload.Data.Attributes.BankIDCode)
	d.Set("bic", account.Payload.Data.Attributes.Bic)
	d.Set("country", account.Payload.Data.Attributes.Country)
	if _, ok := d.GetOk("iban"); ok {
		d.Set("iban", account.Payload.Data.Attributes.Iban)
	} else {
		d.Set("iban", nil)
	}
	if _, ok := d.GetOk("account_number"); ok {
		d.Set("account_number", account.Payload.Data.Attributes.AccountNumber)
	} else {
		d.Set("account_number", nil)
	}
	return nil
}

func resourceAccountDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	account, err := createAccountFromResourceDataWithVersion(d, client)
	if err != nil {
		return fmt.Errorf("error deleting account: %s", err)
	}

	log.Printf("[INFO] Deleting account with resource id: %s account: %s", account.ID, account.Attributes.AccountNumber)

	_, err = client.AccountClient.Accounts.DeleteAccountsID(accounts.NewDeleteAccountsIDParams().
		WithID(form3.UUIDValue(account.ID)).
		WithVersion(*account.Version))

	if err != nil {
		return fmt.Errorf("error deleting account: %s", err)
	}

	return nil
}

func createAccountFromResourceData(d *schema.ResourceData) (*models.Account, error) {

	account := models.Account{Attributes: &models.AccountAttributes{}}
	account.Type = "accounts"
	if attr, ok := GetUUIDOK(d, "account_id"); ok {
		account.ID = form3.UUID(attr)
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		account.OrganisationID = form3.UUID(attr)
	}

	if attr, ok := d.GetOk("account_number"); ok {
		account.Attributes.AccountNumber = attr.(string)
	}

	if attr, ok := d.GetOk("bank_id"); ok {
		account.Attributes.BankID = attr.(string)
	}

	if attr, ok := d.GetOk("bank_id_code"); ok {
		account.Attributes.BankIDCode = attr.(string)
	}

	if attr, ok := d.GetOk("bic"); ok {
		account.Attributes.Bic = attr.(string)
	}

	if attr, ok := d.GetOk("country"); ok {
		account.Attributes.Country = form3.String(attr.(string))
	}

	if attr, ok := d.GetOk("iban"); ok {
		account.Attributes.Iban = attr.(string)
	}

	return &account, nil
}

func createAccountFromResourceDataWithVersion(d *schema.ResourceData, client *form3.AuthenticatedClient) (*models.Account, error) {
	account, err := createAccountFromResourceData(d)
	version, err := getAccountVersion(client, form3.UUIDValue(account.ID))
	if err != nil {
		return nil, err
	}

	account.Version = &version

	return account, nil
}

func getAccountVersion(client *form3.AuthenticatedClient, id strfmt.UUID) (int64, error) {
	account, err := client.AccountClient.Accounts.GetAccountsID(accounts.NewGetAccountsIDParams().WithID(id))
	if err != nil {
		if err != nil {
			return -1, fmt.Errorf("error reading account: %s", err)
		}
	}

	return *account.Payload.Data.Version, nil
}
