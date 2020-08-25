package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/users"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceForm3CredentialSso() *schema.Resource {
	return &schema.Resource{
		Create: resourceCredentialSsoCreate,
		Read:   resourceCredentialSsoRead,
		Delete: resourceCredentialSsoDelete,

		Schema: map[string]*schema.Schema{
			"sso_user_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCredentialSsoCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	ssoUser, err := createSsoCredentialFromResourceData(d)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Creating sso user credential for user: %s, sso_user_id: %s", ssoUser.Attributes.UserID, ssoUser.Attributes.SsoID)

	params := users.NewPostUsersUserIDCredentialsSsoParams().WithUserID(ssoUser.Attributes.UserID).WithSsoUserCreation(&models.SsoUserCreation{Data: ssoUser})
	created, err := client.SecurityClient.Users.PostUsersUserIDCredentialsSso(params)

	if err != nil {
		return fmt.Errorf("failed to create sso user credential for user: %s error: %s", ssoUser.Attributes.UserID, err)
	}

	d.SetId(created.Payload.Data.Attributes.SsoID)
	log.Printf("[INFO] Sso credential with id: %s created", d.Id())
	return nil
}

func resourceCredentialSsoRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	ssoUser, err := createSsoCredentialFromResourceData(d)
	if err != nil {
		return err
	}
	log.Printf("[INFO] Reading sso credential for user: %s, sso_user_id: %s", ssoUser.Attributes.UserID, ssoUser.Attributes.SsoID)

	params := users.NewGetUsersUserIDCredentialsSsoSsoUserIDParams().WithUserID(ssoUser.Attributes.UserID).WithSsoUserID(ssoUser.Attributes.SsoID)
	response, err := client.SecurityClient.Users.GetUsersUserIDCredentialsSsoSsoUserID(params)

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find sso credential for user: %s, sso_user_id: %s error: %s", ssoUser.Attributes.UserID, ssoUser.Attributes.SsoID, err)
		}
		d.SetId("")
		return nil
	}

	d.SetId(response.Payload.Data.Attributes.SsoID)
	d.Set("user_id", response.Payload.Data.Attributes.UserID)
	d.Set("organisation_id", response.Payload.Data.OrganisationID)
	d.Set("sso_user_id", response.Payload.Data.Attributes.SsoID)
	return nil
}

func resourceCredentialSsoDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	ssoUser, err := createSsoCredentialFromResourceData(d)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Deleting sso user credential for user_id: %s, sso_user_id: %s", ssoUser.Attributes.UserID, ssoUser.Attributes.SsoID)

	_, err = client.SecurityClient.Users.DeleteUsersUserIDCredentialsSsoSsoUserID(users.NewDeleteUsersUserIDCredentialsSsoSsoUserIDParams().
		WithUserID(ssoUser.Attributes.UserID).WithSsoUserID(ssoUser.Attributes.SsoID))

	if err != nil {
		return fmt.Errorf("error deleting sso user credential: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createSsoCredentialFromResourceData(d *schema.ResourceData) (*models.SsoUser, error) {
	ssoUser := models.SsoUser{ID: form3.NewUUIDValue(), Attributes: &models.SsoUserAttributes{}}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		ssoUser.OrganisationID = attr
	}

	if attr, ok := GetUUIDOK(d, "user_id"); ok {
		ssoUser.Attributes.UserID = attr
	}

	if attr, ok := d.GetOk("sso_user_id"); ok {
		ssoUser.Attributes.SsoID = attr.(string)
	}

	return &ssoUser, nil
}
