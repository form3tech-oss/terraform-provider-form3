package form3

import (
	"fmt"
	"log"

	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/terraform-provider-form3/models"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceForm3EburyAssociationAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceEburyAssociationAccountCreate,
		Read:   resourceEburyAssociationAccountRead,
		Delete: resourceEburyAssociationAccountDelete,
		Schema: map[string]*schema.Schema{
			"association_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"association_account_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_number": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_number_code": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_with_bic": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_with_bank_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_with_bank_id_code": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"account_labels": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
			"currency": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"country": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"iban": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
		SchemaVersion: 0,
	}
}

func resourceEburyAssociationAccountCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationAccount, err := createEburyAssociationAccountFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create ebury association account: %s", form3.JsonErrorPrettyPrint(err))
	}

	createdAssociationAccount, err := client.AssociationClient.Associations.PostEburyAssociationIDAccounts(
		associations.NewPostEburyAssociationIDAccountsParams().
			WithAssociationID(strfmt.UUID(d.Get("association_id").(string))).
			WithCreationRequest(&models.EburyAssociationAccountCreation{Data: associationAccount}))
	if err != nil {
		return fmt.Errorf("failed to create ebury association account: %s", form3.JsonErrorPrettyPrint(err))
	}

	d.SetId(createdAssociationAccount.Payload.Data.ID.String())
	log.Printf("[INFO] ebury association account key: %s", d.Id())

	return resourceEburyAssociationAccountRead(d, meta)
}

func resourceEburyAssociationAccountRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	associationID, _ := GetUUIDOK(d, "association_id")
	associationAccountId, _ := GetUUIDOK(d, "association_account_id")

	eburyAssociationAccount, err := client.AssociationClient.Associations.GetEburyAssociationIDAccountsID(
		associations.NewGetEburyAssociationIDAccountsIDParams().WithAssociationID(associationID).WithID(associationAccountId))

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find ebury gateway associationAccount: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	_ = d.Set("account_number", eburyAssociationAccount.Payload.Data.Attributes.AccountNumber)
	_ = d.Set("account_number_code", eburyAssociationAccount.Payload.Data.Attributes.AccountNumberCode)
	_ = d.Set("account_labels", eburyAssociationAccount.Payload.Data.Attributes.AccountLabels)
	_ = d.Set("currency", eburyAssociationAccount.Payload.Data.Attributes.Currency)
	_ = d.Set("country", eburyAssociationAccount.Payload.Data.Attributes.Country)
	_ = d.Set("iban", eburyAssociationAccount.Payload.Data.Attributes.Iban)

	if eburyAssociationAccount.Payload.Data.Attributes.AccountWith != nil {
		_ = d.Set("account_with_bic", eburyAssociationAccount.Payload.Data.Attributes.AccountWith.Bic)
		_ = d.Set("account_with_bank_id", eburyAssociationAccount.Payload.Data.Attributes.AccountWith.BankID)
		_ = d.Set("account_with_bank_id_code", eburyAssociationAccount.Payload.Data.Attributes.AccountWith.BankIDCode)
	}

	return nil
}

func resourceEburyAssociationAccountDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	eburyAssociationAccount, err := client.AssociationClient.Associations.GetEburyAssociationIDAccountsID(associations.NewGetEburyAssociationIDAccountsIDParams().
		WithID(strfmt.UUID(d.Id())).
		WithAssociationID(strfmt.UUID(d.Get("association_id").(string))))
	if err != nil {
		return fmt.Errorf("error deleting ebury gateway associationAccount: %s", form3.JsonErrorPrettyPrint(err))
	}

	_, err = client.AssociationClient.Associations.DeleteEburyAssociationIDAccountsID(associations.NewDeleteEburyAssociationIDAccountsIDParams().
		WithID(eburyAssociationAccount.Payload.Data.ID).
		WithAssociationID(strfmt.UUID(d.Get("association_id").(string))).
		WithVersion(*eburyAssociationAccount.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting ebury gateway associationAccount: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createEburyAssociationAccountFromResourceData(d *schema.ResourceData) (*models.NewEburyAssociationAccount, error) {
	associationAccount := models.NewEburyAssociationAccount{
		Type: string(models.ResourceTypeEburyAssociationAccounts),
		Attributes: &models.EburyAssociationAccountAttributes{
			AccountWith: &models.EburyAssociationAccountAttributesAccountWith{},
		},
	}

	if attr, ok := GetUUIDOK(d, "association_account_id"); ok {
		associationAccount.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		associationAccount.OrganisationID = attr
	}

	if attr, ok := d.GetOk("account_number"); ok {
		associationAccount.Attributes.AccountNumber = attr.(string)
	}

	if attr, ok := d.GetOk("account_number_code"); ok {
		associationAccount.Attributes.AccountNumberCode = attr.(string)
	}

	if attr, ok := d.GetOk("account_with_bic"); ok {
		associationAccount.Attributes.AccountWith.Bic = attr.(string)
	}

	if attr, ok := d.GetOk("account_with_bank_id"); ok {
		associationAccount.Attributes.AccountWith.BankID = attr.(string)
	}

	if attr, ok := d.GetOk("account_with_bank_id_code"); ok {
		associationAccount.Attributes.AccountWith.BankIDCode = attr.(string)
	}

	if attr, ok := d.GetOk("account_labels"); ok {
		coll := attr.([]interface{})
		for _, v := range coll {
			associationAccount.Attributes.AccountLabels = append(associationAccount.Attributes.AccountLabels, v.(string))
		}
	}

	if attr, ok := d.GetOk("currency"); ok {
		associationAccount.Attributes.Currency = attr.(string)
	}

	if attr, ok := d.GetOk("country"); ok {
		associationAccount.Attributes.Country = attr.(string)
	}

	if attr, ok := d.GetOk("iban"); ok {
		associationAccount.Attributes.Iban = attr.(string)
	}

	return &associationAccount, nil
}
