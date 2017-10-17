package form3

import (
	"fmt"
	"github.com/ewilde/go-form3"
	"github.com/ewilde/go-form3/client/users"
	"github.com/ewilde/go-form3/models"
	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func resourceForm3User() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,

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

	createdUser, err := client.ApiClients.Users.PostUsers(users.NewPostUsersParams().
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
	log.Printf("[INFO] Reading user for id: %s username: %s", key, userName)

	user, err := client.ApiClients.Users.GetUsersUserID(users.NewGetUsersUserIDParams().WithUserID(userId))
	if err != nil {
		apiError := err.(*runtime.APIError)
		if apiError.Code == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("couldn't find user: %s", err)
	}

	d.Set("user_id", user.Payload.Data.ID.String())
	d.Set("user_name", user.Payload.Data.Attributes.Username)
	d.Set("email", user.Payload.Data.Attributes.Email)
	d.Set("organisation_id", user.Payload.Data.OrganisationID.String())
	return nil
}

func resourceUserUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Partial(false)
	userFromResource, err := createUserFromResourceData(d)
	if err != nil {
		return fmt.Errorf("error updating user: %s", err)
	}

	if d.HasChange("description") {
		client := meta.(*form3.AuthenticatedClient)
		_, err = client.ApiClients.Users.PatchUsersUserID(users.NewPatchUsersUserIDParams().
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

	key := d.Id()
	userId, _ := GetUUIDOK(d, "user_id")
	userName := d.Get("user_name").(string)
	log.Printf("[INFO] Deleting user for id: %s username: %s", key, userName)

	_, err := client.ApiClients.Users.DeleteUsersUserID(users.NewDeleteUsersUserIDParams().WithUserID(userId))
	if err != nil {
		return fmt.Errorf("error deleting user: %s", err)
	}

	return nil
}

func createUserFromResourceData(d *schema.ResourceData) (*models.User, error) {

	user := models.User{Attributes: &models.UserAttributes{}}
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

	return &user, nil
}
