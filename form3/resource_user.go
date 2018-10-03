package form3

import (
	"fmt"
	"github.com/form3tech-oss/go-form3"
	"github.com/form3tech-oss/go-form3/client/users"
	"github.com/form3tech-oss/go-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3User() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"roles": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceUserCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	username := d.Get("user_name").(string)
	log.Printf("[INFO] Creating user with username: %s", username)

	user, err := createUserFromResourceData(d)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] user create: %#v", user)

	createdUser, err := client.SecurityClient.Users.PostUsers(users.NewPostUsersParams().
		WithUserCreationRequest(&models.UserCreation{Data: user}))

	if err != nil {
		return fmt.Errorf("failed to create user: %s", err)
	}

	d.SetId(createdUser.Payload.Data.ID.String())
	log.Printf("[INFO] user key: %s", d.Id())

	return resourceUserRead(d, meta)
}

func resourceUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	key := d.Id()
	userId, _ := GetUUIDOK(d, "user_id")
	userName := d.Get("user_name").(string)

	if userId == "" {
		userId = strfmt.UUID(key)
		log.Printf("[INFO] Importing user id: %s", key)
	} else {
		log.Printf("[INFO] Reading user for id: %s username: %s", key, userName)
	}

	user, err := client.SecurityClient.Users.GetUsersUserID(users.NewGetUsersUserIDParams().WithUserID(userId))
	if err != nil {
		apiError, ok := err.(*runtime.APIError)
		if ok && apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find user: %s", err)
	}

	d.Set("user_id", user.Payload.Data.ID.String())
	d.Set("user_name", user.Payload.Data.Attributes.Username)
	d.Set("email", user.Payload.Data.Attributes.Email)
	d.Set("organisation_id", user.Payload.Data.OrganisationID.String())
	d.Set("roles", readUserRoles(user.Payload.Data.Attributes.RoleIds))
	return nil
}

func resourceUserUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(false)

	if d.HasChange("email") || d.HasChange("roles") {
		client := meta.(*form3.AuthenticatedClient)
		userFromResource, err := createUserFromResourceDataWithVersion(d, client)
		if err != nil {
			return fmt.Errorf("error updating user: %s", err)
		}

		_, err = client.SecurityClient.Users.PatchUsersUserID(users.NewPatchUsersUserIDParams().
			WithUserID(userFromResource.ID).
			WithUserUpdateRequest(&models.UserCreation{Data: userFromResource}))

		if err != nil {
			return fmt.Errorf("error updating user: %s", err)
		}
	}

	return nil
}

func resourceUserDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*form3.AuthenticatedClient)

	userFromResource, err := createUserFromResourceDataWithVersion(d, client)
	if err != nil {
		return fmt.Errorf("error deleting user: %s", err)
	}

	log.Printf("[INFO] Deleting user for id: %s username: %s", userFromResource.ID, userFromResource.Attributes.Username)

	_, err = client.SecurityClient.Users.DeleteUsersUserID(users.NewDeleteUsersUserIDParams().
		WithUserID(userFromResource.ID).
		WithVersion(*userFromResource.Version))

	if err != nil {
		return fmt.Errorf("error deleting user: %s", err)
	}

	return nil
}

func createUserFromResourceDataWithVersion(d *schema.ResourceData, client *form3.AuthenticatedClient) (*models.User, error) {
	user, err := createUserFromResourceData(d)
	version, err := getUserVersion(client, user.ID)
	if err != nil {
		return nil, err
	}

	user.Version = &version

	return user, nil
}

func createUserFromResourceData(d *schema.ResourceData) (*models.User, error) {

	user := models.User{Attributes: &models.UserAttributes{}}
	user.Type = "users"
	if attr, ok := GetUUIDOK(d, "user_id"); ok {
		user.ID = attr
	}

	if attr, ok := d.GetOk("user_name"); ok {
		user.Attributes.Username = attr.(string)
	}

	if attr, ok := d.GetOk("email"); ok {
		user.Attributes.Email = attr.(string)
	}

	if attr, ok := GetUUIDOK(d, "organisation_id"); ok {
		user.OrganisationID = attr
	}

	if attr, ok := d.GetOk("roles"); ok {
		roles := []strfmt.UUID{}
		items := attr.([]interface{})
		for _, x := range items {
			item := x.(string)
			roles = append(roles, strfmt.UUID(item))
		}

		user.Attributes.RoleIds = roles
	}

	return &user, nil
}

func getUserVersion(client *form3.AuthenticatedClient, userId strfmt.UUID) (int64, error) {
	user, err := client.SecurityClient.Users.GetUsersUserID(users.NewGetUsersUserIDParams().WithUserID(userId))
	if err != nil {
		if err != nil {
			return -1, fmt.Errorf("error reading user: %s", err)
		}
	}

	return *user.Payload.Data.Version, nil
}

func readUserRoles(roles []strfmt.UUID) []string {
	result := make([]string, 0, len(roles))
	for _, role := range roles {
		result = append(result, role.String())
	}

	return result
}
