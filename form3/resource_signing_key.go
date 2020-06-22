package form3

import (
	"fmt"
	"github.com/form3tech-oss/terraform-provider-form3/client/platformsecurityapi"
	"log"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceForm3SigningKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceSigningKeyCreate,
		Read:   resourceSigningKeyRead,
		Delete: resourceSigningKeyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"public_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"organisation_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
		},
	}
}

func resourceSigningKeyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	id := d.Get("id").(string)
	log.Printf("[INFO] Creating signing_key with id: %s", id)

	signingKey, err := createSigningKeyFromResourceData(d)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] signing_key create: %#v", signingKey)

	createdSigningKey, err := client.SystemClient.Platformsecurityapi.PostPlatformSecuritySigningKeys(platformsecurityapi.NewPostPlatformSecuritySigningKeysParams().WithData(
		&models.SigningKeysCreation{
			Data: signingKey,
		},
	))

	if err != nil {
		return fmt.Errorf("failed to create signing_key: %s", form3.JsonErrorPrettyPrint(err))
	}

	d.SetId(createdSigningKey.Payload.Data.ID.String())
	log.Printf("[INFO] signing_key id: %s", d.Id())

	return resourceSigningKeyRead(d, meta)
}

func resourceSigningKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	key := d.Id()
	signingKeyId, _ := GetUUIDOK(d, "signing_key_id")

	if signingKeyId == "" {
		signingKeyId = strfmt.UUID(key)
		log.Printf("[INFO] Importing signing_key id: %s ", key)
	} else {
		log.Printf("[INFO] Reading signing_key id: %s", key)
	}

	signingKey, err := client.SystemClient.Platformsecurityapi.GetPlatformSecuritySigningKeysSigningkeyID(
		platformsecurityapi.NewGetPlatformSecuritySigningKeysSigningkeyIDParams().WithSigningkeyID(signingKeyId))

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find signing_key: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	d.Set("id", signingKey.Payload.Data.ID.String())
	return nil
}

func resourceSigningKeyUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(false)

	if d.HasChange("id") {
		return fmt.Errorf("error updating signing_key as they can not be changed")
	}

	return nil
}

func resourceSigningKeyDelete(d *schema.ResourceData, meta interface{}) error {
	// you can't delete a signing key, just remove it from state?
	log.Printf("[INFO] Deleting signing_key id: %s",
		d.Get("id"))

	return nil
}

func createSigningKeyFromResourceData(d *schema.ResourceData) (*models.SigningKeysRequestData, error) {
	signingKey := models.SigningKeysRequestData{}
	if attr, ok := GetUUIDOK(d, "id"); ok {
		id := attr.String()
		signingKey.ID = &id
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		orgId := attr.String()
		signingKey.OrganisationID = &orgId
	}
	return &signingKey, nil
}
