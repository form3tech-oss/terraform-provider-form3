package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/users"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3CredentialPublicKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceCredentialPublicKeyCreate,
		Read:   resourceCredentialPublicKeyRead,
		Delete: resourceCredentialPublicKeyDelete,

		Schema: map[string]*schema.Schema{
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
			"public_key_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCredentialPublicKeyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	userId, _ := GetUUIDOK(d, "user_id")
	log.Printf("[INFO] Creating credential public key for user id: %s", userId.String())

	publicKey, err := createCredentialPublicKeyFromResourceData(d)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] credential public key create: %#v", publicKey)

	createdPublicKey, err := client.SecurityClient.Users.PostUsersUserIDCredentialsPublicKey(users.NewPostUsersUserIDCredentialsPublicKeyParams().
		WithUserID(userId))

	if err != nil {
		return fmt.Errorf("failed to create credential public key: %s", err)
	}

	d.SetId(createdPublicKey.Payload.Data.ID.String())
	log.Printf("[INFO] credential public key id: %s", d.Id())

	d.Set("public_key_id", createdPublicKey.Payload.Data.ID.String())
	return nil
}

func resourceCredentialPublicKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	key := d.Id()
	userId, _ := GetUUIDOK(d, "user_id")
	log.Printf("[INFO] Reading credential public key for client id: %s", key)

	credentials, err := client.SecurityClient.Users.GetUsersUserIDCredentialsPublicKey(users.NewGetUsersUserIDCredentialsPublicKeyParams().
		WithUserID(userId))

	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find credential public key: %s", err)
	}

	for _, element := range credentials.Payload.Data {
		if element.PublicKeyID.String() == key {
			d.Set("user_id", userId.String())
			d.Set("public_key_id", key)
			return nil
		}
	}

	d.SetId("")
	return nil
}

func resourceCredentialPublicKeyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	publicKey, err := createCredentialPublicKeyFromResourceData(d)
	if err != nil {
		return fmt.Errorf("error deleting credential public key: %s", err)
	}

	userId, _ := GetUUIDOK(d, "user_id")
	log.Printf("[INFO] Deleting credtial public key id: %s for user: %s", publicKey.ID.String(), userId.String())

	_, err = client.SecurityClient.Users.DeleteUsersUserIDCredentialsPublicKeyPublicKeyID(users.NewDeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams().
		WithUserID(userId).WithPublicKeyID(publicKey.ID))

	if err != nil {
		return fmt.Errorf("error deleting credential public key: %s", err)
	}

	return nil
}

func createCredentialPublicKeyFromResourceData(d *schema.ResourceData) (*models.PublicKey, error) {

	publicKey := &models.PublicKey{Attributes: &models.PublicKeyAttributes{}}

	if attr, ok := GetUUIDOK(d, "public_key_id"); ok {
		publicKey.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		publicKey.OrganisationID = attr
	}

	if attr, ok := d.GetOk("public_key"); ok {
		publicKey.Attributes.PublicKey = attr.(string)
	}

	return publicKey, nil
}
