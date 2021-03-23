package form3

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"golang.org/x/crypto/ssh"
	"fmt"
	"log"
	"strings"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/users"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceForm3CredentialPublicKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceCredentialPublicKeyCreate,
		Read:   resourceCredentialPublicKeyRead,
		Delete: resourceCredentialPublicKeyDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCredentialPublicKeyImport,
		},

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
				Required: true,
				ForceNew: true,
			},
			"public_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"public_key_fingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				ForceNew: false,
			},
		},
	}
}

func resourceCredentialPublicKeyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.Split(d.Id(), "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("public key import id must be in form '<userId>/<publicKeyId>'")
	}
	d.SetId(parts[1])
	d.Set("user_id", parts[0])
	return []*schema.ResourceData{d}, nil
}

func resourceCredentialPublicKeyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	userID, _ := GetUUIDOK(d, "user_id")
	log.Printf("[INFO] Creating credential public key for user id: %s", userID.String())

	publicKey, err := createCredentialPublicKeyFromResourceData(d)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] credential public key create: %#v", publicKey)

	createKeyParams := users.NewPostUsersUserIDCredentialsPublicKeyParams().WithUserID(userID).WithPublicKey(publicKey)
	createdPublicKey, err := client.SecurityClient.Users.PostUsersUserIDCredentialsPublicKey(createKeyParams)

	if err != nil {
		return fmt.Errorf("failed to create credential public key with id: %s error: %s", publicKey.ID, err)
	}

	d.SetId(createdPublicKey.Payload.Data.ID.String())
	log.Printf("[INFO] Credential public key with id: %s created", d.Id())
	return nil
}

func resourceCredentialPublicKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	keyID := d.Id()
	userID, _ := GetUUIDOK(d, "user_id")
	log.Printf("[INFO] Reading credential public key for id: %s", keyID)

	getKeyParams := users.NewGetUsersUserIDCredentialsPublicKeyPublicKeyIDParams().WithUserID(userID).WithPublicKeyID(strfmt.UUID(keyID))
	publicKey, err := client.SecurityClient.Users.GetUsersUserIDCredentialsPublicKeyPublicKeyID(getKeyParams)

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find credential public key for user %s with id:%s error %s", userID, keyID, err)
		}
		d.SetId("")
		return nil
	}

	d.SetId(publicKey.Payload.ID.String())
	d.Set("user_id", userID)
	d.Set("public_key_id", publicKey.Payload.ID)
	d.Set("organisation_id", publicKey.Payload.OrganisationID)
	d.Set("public_key", publicKey.Payload.Attributes.PublicKey)
	return nil
}

func resourceCredentialPublicKeyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	publicKey, err := createCredentialPublicKeyFromResourceData(d)
	if err != nil {
		return fmt.Errorf("error deleting credential public key: %s", form3.JsonErrorPrettyPrint(err))
	}

	userID, _ := GetUUIDOK(d, "user_id")
	log.Printf("[INFO] Deleting credential public key id: %s for user: %s", publicKey.ID.String(), userID.String())

	_, err = client.SecurityClient.Users.DeleteUsersUserIDCredentialsPublicKeyPublicKeyID(users.NewDeleteUsersUserIDCredentialsPublicKeyPublicKeyIDParams().
		WithUserID(userID).WithPublicKeyID(publicKey.ID))

	if err != nil {
		return fmt.Errorf("error deleting credential public key: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createCredentialPublicKeyFromResourceData(d *schema.ResourceData) (*models.PublicKey, error) {

	publicKey := models.PublicKey{Attributes: &models.PublicKeyAttributes{}}

	if attr, ok := GetUUIDOK(d, "public_key_id"); ok {
		publicKey.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		publicKey.OrganisationID = attr
	}

	if attr, ok := d.GetOk("public_key"); ok {
		keyString := attr.(string)

		// we only want to verify fingerprint when its specified with key
		// it is not required
		// the fingerprint is not added in to the external state
		if fingerprint, ok := d.GetOk("public_key_fingerprint"); ok {
			fingerprintCalculated, err := calculatePublicKeyFingerprint(keyString)
			if err != nil {
				return nil, err
			}

			if fingerprint != fingerprintCalculated {
				return nil, fmt.Errorf("the provided key doesn't match the fingerprint")
			}
		}

		publicKey.Attributes.PublicKey = keyString
	}

	return &publicKey, nil
}

func calculatePublicKeyFingerprint(key string) (string, error) {
	byteKey, _ := pem.Decode([]byte(key))

	var rawKey interface{}
	var err error
	rawKey, err = x509.ParsePKIXPublicKey(byteKey.Bytes)
	if err != nil {
		return "", fmt.Errorf("error when parsing the public key - %s", err)
	}

	var sshPublicKey ssh.PublicKey
	sshPublicKey, err = ssh.NewPublicKey(rawKey)
	if err != nil {
		return "", fmt.Errorf("error when calculating fingerprint of the key - %s", err)
	}

	return ssh.FingerprintLegacyMD5(sshPublicKey), nil
}
