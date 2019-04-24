package form3

import (
	"fmt"
	"github.com/form3tech-oss/go-form3"
	"github.com/form3tech-oss/go-form3/client/roles"
	"github.com/form3tech-oss/go-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3Role() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoleCreate,
		Read:   resourceRoleRead,
		Delete: resourceRoleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parent_role_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRoleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	rolename := d.Get("name").(string)
	log.Printf("[INFO] Creating role with name: %s", rolename)

	role, err := createRoleFromResourceData(d)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] role create: %#v", role)

	createdRole, err := client.SecurityClient.Roles.PostRoles(roles.NewPostRolesParams().
		WithRoleCreationRequest(&models.RoleCreation{
			Data: role,
		}))

	if err != nil {
		return fmt.Errorf("failed to create role: %s", err)
	}

	d.SetId(createdRole.Payload.Data.ID.String())
	log.Printf("[INFO] role key: %s", d.Id())

	return resourceRoleRead(d, meta)
}

func resourceRoleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	key := d.Id()
	roleId, _ := GetUUIDOK(d, "role_id")
	roleName := d.Get("name").(string)

	if roleId == "" {
		roleId = strfmt.UUID(key)
		log.Printf("[INFO] Importing role id: %s", roleId)
	} else {
		log.Printf("[INFO] Reading role for id: %s rolename: %s", key, roleName)
	}

	role, err := client.SecurityClient.Roles.GetRolesRoleID(roles.NewGetRolesRoleIDParams().WithRoleID(roleId))
	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find role: %s", err)
	}

	d.Set("role_id", role.Payload.Data.ID.String())
	d.Set("name", role.Payload.Data.Attributes.Name)
	d.Set("organisation_id", role.Payload.Data.OrganisationID.String())

	if role.Payload.Data.Attributes.ParentRoleID != nil {
		d.Set("parent_role_id", role.Payload.Data.Attributes.ParentRoleID.String())
	}

	return nil
}

func resourceRoleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	roleFromResource, err := createRoleFromResourceDataWithVersion(d, client)
	if err != nil {
		return fmt.Errorf("error deleting role: %s", err)
	}

	log.Printf("[INFO] Deleting role for id: %s rolename: %s", roleFromResource.ID, roleFromResource.Attributes.Name)

	_, err = client.SecurityClient.Roles.DeleteRolesRoleID(roles.NewDeleteRolesRoleIDParams().
		WithRoleID(roleFromResource.ID).
		WithVersion(*roleFromResource.Version))

	if err != nil {
		return fmt.Errorf("error deleting role: %s", err)
	}

	return nil
}

func createRoleFromResourceDataWithVersion(d *schema.ResourceData, client *form3.AuthenticatedClient) (*models.Role, error) {
	role, err := createRoleFromResourceData(d)
	version, err := getRoleVersion(client, role.ID)
	if err != nil {
		return nil, err
	}

	role.Version = &version

	return role, nil
}

func createRoleFromResourceData(d *schema.ResourceData) (*models.Role, error) {

	role := models.Role{Attributes: &models.RoleAttributes{}}
	role.Type = "roles"
	if attr, ok := GetUUIDOK(d, "role_id"); ok {
		role.ID = attr
	}

	if attr, ok := d.GetOk("name"); ok {
		role.Attributes.Name = attr.(string)
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		role.OrganisationID = attr
	}

	if attr, ok := GetUUIDOK(d, "parent_role_id"); ok {
		role.Attributes.ParentRoleID = &attr
	}

	return &role, nil
}

func getRoleVersion(client *form3.AuthenticatedClient, roleId strfmt.UUID) (int64, error) {
	role, err := client.SecurityClient.Roles.GetRolesRoleID(roles.NewGetRolesRoleIDParams().WithRoleID(roleId))
	if err != nil {
		if err != nil {
			return -1, fmt.Errorf("error reading role: %s", err)
		}
	}

	return *role.Payload.Data.Version, nil
}
