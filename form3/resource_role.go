package form3

import (
	"fmt"
	"github.com/form3tech-oss/terraform-provider-form3/client/users"
	"github.com/giantswarm/retry-go"
	"github.com/go-openapi/swag"
	"log"
	"time"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/roles"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
		return fmt.Errorf("failed to create role: %s", form3.JsonErrorPrettyPrint(err))
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
		if !form3.IsJsonErrorStatusCode(err, 404) {
			return fmt.Errorf("couldn't find role: %s", form3.JsonErrorPrettyPrint(err))
		}
		d.SetId("")
		return nil
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
		return fmt.Errorf("error deleting role: %s", form3.JsonErrorPrettyPrint(err))
	}

	log.Printf("[INFO] Deleting role for id: %s rolename: %s", roleFromResource.ID, roleFromResource.Attributes.Name)

	associatedUsers, err := getAssociatedUsers(client, roleFromResource)
	if err != nil {
		return err
	}

	err = removeRoleAssociation(associatedUsers, client, roleFromResource)
	if err != nil {
		return err
	}

	_, err = client.SecurityClient.Roles.DeleteRolesRoleID(roles.NewDeleteRolesRoleIDParams().
		WithRoleID(roleFromResource.ID).
		WithVersion(*roleFromResource.Version))

	if err != nil {
		return fmt.Errorf("error deleting role: %s", form3.JsonErrorPrettyPrint(err))
	}

	return nil
}

func removeRoleAssociation(associatedUsers []*models.User, client *form3.AuthenticatedClient, roleFromResource *models.Role) error {
	for _, originalUser := range associatedUsers {
		err := retry.Do(func() error {
			// We need to retry getting the user and patching it, in-case it has been updated between when we paged
			// through all users and when we get round to patching it. If the user has been updated in the interim,
			// we need to patch its current version.
			currentUser, err := client.SecurityClient.Users.GetUsersUserID(users.NewGetUsersUserIDParams().WithUserID(originalUser.ID))
			if err != nil {
				return fmt.Errorf("error deleting role (patching users roles): %s", form3.JsonErrorPrettyPrint(err))
			}
			_, err = client.SecurityClient.Users.PatchUsersUserID(users.NewPatchUsersUserIDParams().
				WithUserID(originalUser.ID).
				WithUserUpdateRequest(&models.UserCreation{
					Data: &models.User{
						Attributes: &models.UserAttributes{
							RoleIds: without(currentUser.Payload.Data.Attributes.RoleIds, roleFromResource.ID),
						},
						ID:             originalUser.ID,
						OrganisationID: originalUser.OrganisationID,
						Type:           originalUser.Type,
						Version:        currentUser.Payload.Data.Version,
					},
				}))
			if err != nil {
				return fmt.Errorf("error deleting role (patching users roles): %s", form3.JsonErrorPrettyPrint(err))
			}
			return nil
		}, retry.Sleep(1*time.Second), retry.MaxTries(5))
		if err != nil {
			return err
		}
	}
	return nil
}

func getAssociatedUsers(client *form3.AuthenticatedClient, roleFromResource *models.Role) ([]*models.User, error) {
	associatedUsers := []*models.User{}
	// TODO: we would need to page through all the pages here to GET all users, so that we can manually filter them for associated roles here.
	allUsers, err := client.SecurityClient.Users.GetUsers(users.NewGetUsersParams().WithPageSize(swag.Int64(1000)))
	if err != nil {
		return nil, fmt.Errorf("error getting all users: %s", form3.JsonErrorPrettyPrint(err))
	}
	for _, user := range allUsers.Payload.Data {
		for _, roleId := range user.Attributes.RoleIds {
			if roleId == roleFromResource.ID {
				associatedUsers = append(associatedUsers, user)
			}
		}
	}
	return associatedUsers, nil
}

func without(ids []strfmt.UUID, id strfmt.UUID) []strfmt.UUID {
	result := []strfmt.UUID{}
	for _, i := range ids {
		if i == id {
			continue
		}
		result = append(result, i)
	}
	return result
}

func createRoleFromResourceDataWithVersion(d *schema.ResourceData, client *form3.AuthenticatedClient) (*models.Role, error) {
	role, err := createRoleFromResourceData(d)
	if err != nil {
		return nil, err
	}
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
			return -1, fmt.Errorf("error reading role: %s", form3.JsonErrorPrettyPrint(err))
		}
	}

	return *role.Payload.Data.Version, nil
}
