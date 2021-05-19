package form3

import (
	"fmt"
	"log"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/system"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceForm3Key() *schema.Resource {
	return &schema.Resource{
		Create: resourceKeyCreate,
		Read:   resourceKeyRead,
		Delete: resourceKeyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"key_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subject": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"certificate_signing_request": {
				Type:     schema.TypeString,
				Computed: true,
				Required: false,
				ForceNew: true,
			},
			"private_key": {
				Type:     schema.TypeString,
				Computed: true,
				Required: false,
				ForceNew: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Computed: true,
				Required: false,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "RSA",
				ForceNew: true,
			},
			"curve": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceKeyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	log.Print("[INFO] Creating Key ")

	certificateRequest, err := createKeyFromResourceData(d)
	if err != nil {
		return fmt.Errorf("failed to create Key: %s", form3.JsonErrorPrettyPrint(err))
	}

	createdKey, err := client.SystemClient.System.PostKeys(
		system.NewPostKeysParams().
			WithKeyCreationRequest(&models.KeyCreation{
				Data: certificateRequest,
			}))

	if err != nil {
		return fmt.Errorf("failed to create Key: %s", form3.JsonErrorPrettyPrint(err))
	}

	d.SetId(createdKey.Payload.Data.ID.String())
	log.Printf("[INFO] Key key: %s", d.Id())

	return resourceKeyRead(d, meta)
}

func resourceKeyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	keyId, _ := GetUUIDOK(d, "key_id")

	if keyId == "" {
		keyId = strfmt.UUID(d.Id())
		log.Printf("[INFO] Importing key with resource id: %s.", keyId)
	} else {
		log.Printf("[INFO] Reading key with resource id: %s.", keyId)
	}

	response, err := client.SystemClient.System.GetKeysKeyID(
		system.NewGetKeysKeyIDParams().WithKeyID(keyId))

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find key: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	d.Set("key_id", response.Payload.Data.ID.String())
	d.Set("subject", response.Payload.Data.Attributes.Subject)
	d.Set("organisation_id", response.Payload.Data.OrganisationID.String())
	d.Set("certificate_signing_request", response.Payload.Data.Attributes.CertificateSigningRequest)
	d.Set("public_key", response.Payload.Data.Attributes.PublicKey)
	d.Set("private_key", response.Payload.Data.Attributes.PrivateKey)
	d.Set("description", response.Payload.Data.Attributes.Description)
	d.Set("type", response.Payload.Data.Attributes.Type)
	d.Set("curve", response.Payload.Data.Attributes.Curve)

	return nil
}

func resourceKeyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	response, err := client.SystemClient.System.GetKeysKeyID(
		system.NewGetKeysKeyIDParams().WithKeyID(strfmt.UUID(d.Id())))

	if err != nil {
		return fmt.Errorf("error deleting Key: %s", form3.JsonErrorPrettyPrint(err))
	}

	log.Printf("[INFO] Deleting Key for id: %s ", response.Payload.Data.ID)

	_, err = client.SystemClient.System.DeleteKeysKeyID(
		system.NewDeleteKeysKeyIDParams().
			WithKeyID(response.Payload.Data.ID).
			WithVersion(*response.Payload.Data.Version))

	if err != nil {
		return fmt.Errorf("error deleting Key: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createKeyFromResourceData(d *schema.ResourceData) (*models.Key, error) {

	certificateRequest := &models.Key{
		Type:       "keys",
		Attributes: &models.KeyAttributes{},
	}

	if attr, ok := GetUUIDOK(d, "key_id"); ok {
		uuid := strfmt.UUID(attr.String())
		certificateRequest.ID = uuid
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		uuid := strfmt.UUID(attr.String())
		certificateRequest.OrganisationID = uuid
	}

	if attr, ok := d.GetOk("subject"); ok {
		certificateRequest.Attributes.Subject = attr.(string)
	}

	if attr, ok := d.GetOk("description"); ok {
		certificateRequest.Attributes.Description = attr.(string)
	}

	if attr, ok := d.GetOk("type"); ok {
		keyType := attr.(string)
		certificateRequest.Attributes.Type = &keyType
	}

	if attr, ok := d.GetOk("curve"); ok {
		certificateRequest.Attributes.Curve = attr.(string)
	}

	return certificateRequest, nil
}
