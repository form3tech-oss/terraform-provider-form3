package form3

import (
	"fmt"
	"log"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/users"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceForm3Credential() *schema.Resource {
	return &schema.Resource{
		Create: resourceCredentialCreate,
		Read:   resourceCredentialRead,
		Delete: resourceCredentialDelete,

		Schema: map[string]*schema.Schema{
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_secret": &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceCredentialCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	userId, _ := GetUUIDOK(d, "user_id")
	log.Printf("[INFO] Creating credential for user id: %s", userId.String())

	credential, err := createCredentialFromResourceData(d)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] credential create: %#v", credential)

	createdCredential, err := client.SecurityClient.Users.PostUsersUserIDCredentials(users.NewPostUsersUserIDCredentialsParams().
		WithUserID(userId))

	if err != nil {
		return fmt.Errorf("failed to create credential: %s", form3.JsonErrorPrettyPrint(err))
	}

	d.SetId(createdCredential.Payload.Data.ClientID.String())
	log.Printf("[INFO] credential id: %s", d.Id())

	d.Set("client_id", createdCredential.Payload.Data.ClientID.String())
	d.Set("client_secret", createdCredential.Payload.Data.ClientSecret)
	return nil
}

func resourceCredentialRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	key := d.Id()
	userId, _ := GetUUIDOK(d, "user_id")
	log.Printf("[INFO] Reading credential for client id: %s", key)

	credentials, err := client.SecurityClient.Users.GetUsersUserIDCredentials(users.NewGetUsersUserIDCredentialsParams().
		WithUserID(userId))

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find credential: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	for _, element := range credentials.Payload.Data {
		if element.ClientID.String() == key {
			d.Set("user_id", userId.String())
			d.Set("client_id", key)
			return nil
		}
	}

	d.SetId("")
	return nil
}

func resourceCredentialDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	credential, err := createCredentialFromResourceData(d)
	if err != nil {
		return fmt.Errorf("error deleting credential: %s", form3.JsonErrorPrettyPrint(err))
	}

	userId, _ := GetUUIDOK(d, "user_id")
	log.Printf("[INFO] Deleting credtial id: %s for user: %s", credential.ClientID.String(), userId.String())

	_, err = client.SecurityClient.Users.DeleteUsersUserIDCredentialsClientID(users.NewDeleteUsersUserIDCredentialsClientIDParams().
		WithUserID(userId).
		WithClientID(credential.ClientID))

	if err != nil {
		return fmt.Errorf("error deleting credential: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createCredentialFromResourceData(d *schema.ResourceData) (*models.CredentialSecret, error) {

	credential := &models.CredentialSecret{}

	if attr, ok := GetUUIDOK(d, "client_id"); ok {
		credential.ClientID = attr
	}

	if attr, ok := d.GetOk("client_secret"); ok {
		credential.ClientSecret = attr.(string)
	}

	return credential, nil
}
