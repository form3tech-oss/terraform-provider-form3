package form3

import (
	"fmt"
	"log"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/ace"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceForm3Ace() *schema.Resource {
	return &schema.Resource{
		Create: resourceAceCreate,
		Read:   resourceAceRead,
		Delete: resourceAceDelete,

		Schema: map[string]*schema.Schema{
			"ace_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"role_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organisation_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"record_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	aceId := d.Get("ace_id").(string)
	log.Printf("[INFO] Creating ace with id: %s", aceId)

	aceResource, err := createRoleAceFromResourceData(d)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] ace create: %#v", aceResource)

	createdRole, err := client.SecurityClient.Ace.PostRolesRoleIDAces(ace.NewPostRolesRoleIDAcesParams().
		WithRoleID(aceResource.Attributes.RoleID).
		WithAceCreationRequest(&models.AceCreation{
			Data: aceResource,
		}))

	if err != nil {
		return fmt.Errorf("failed to create ace: %s", form3.JsonErrorPrettyPrint(err))
	}

	d.SetId(createdRole.Payload.Data.ID.String())
	log.Printf("[INFO] ace key: %s", d.Id())

	return resourceAceRead(d, meta)
}

func resourceAceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	aceId, _ := GetUUIDOK(d, "ace_id")
	roleId, _ := GetUUIDOK(d, "role_id")

	log.Printf("[INFO] Reading ace for id: %s role id: %s", aceId, roleId)

	aceResponse, err := client.SecurityClient.Ace.GetRolesRoleIDAcesAceID(ace.NewGetRolesRoleIDAcesAceIDParams().
		WithAceID(aceId).
		WithRoleID(roleId))

	if err != nil {
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find ace: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
	}

	_ = d.Set("ace_id", aceResponse.Payload.Data.ID.String())
	_ = d.Set("role_id", aceResponse.Payload.Data.Attributes.RoleID.String())
	_ = d.Set("record_type", aceResponse.Payload.Data.Attributes.RecordType)
	_ = d.Set("action", aceResponse.Payload.Data.Attributes.Action)
	return nil
}

func resourceAceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	aceFromResource, err := createRoleAceFromResourceData(d)
	if err != nil {
		return fmt.Errorf("error deleting ace: %s", form3.JsonErrorPrettyPrint(err))
	}

	log.Printf("[INFO] Deleting ace for id: %s role id: %s", aceFromResource.ID, aceFromResource.Attributes.RoleID)

	_, err = client.SecurityClient.Ace.DeleteRolesRoleIDAcesAceID(ace.NewDeleteRolesRoleIDAcesAceIDParams().
		WithAceID(aceFromResource.ID).
		WithRoleID(aceFromResource.Attributes.RoleID))

	apiError, ok := err.(*runtime.APIError)
	if ok && apiError.Code == 404 {
		return nil
	}

	if err != nil {
		return fmt.Errorf("error deleting ace: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func createRoleAceFromResourceData(d *schema.ResourceData) (*models.Ace, error) {

	ace := models.Ace{Attributes: &models.AceAttributes{}}
	ace.Type = "ace"
	if attr, ok := GetUUIDOK(d, "ace_id"); ok {
		ace.ID = attr
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		ace.OrganisationID = attr
	}

	if attr, ok := GetUUIDOK(d, "role_id"); ok {
		ace.Attributes.RoleID = attr
	}

	if attr, ok := d.GetOk("action"); ok {
		ace.Attributes.Action = attr.(string)
	}

	if attr, ok := d.GetOk("record_type"); ok {
		ace.Attributes.RecordType = attr.(string)
	}

	return &ace, nil
}
