package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/users"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/satori/go.uuid"
	"log"
	"os"
	"testing"
)

func TestAccCredentialPublicKey_multipleKeys_sequential(t *testing.T) {
	log.SetOutput(os.Stdout)
	organisationID := os.Getenv("FORM3_ORGANISATION_ID")
	publicKeyIDOne := uuid.NewV4().String()
	publicKeyIDTwo := uuid.NewV4().String()
	publicKeyIDThree := uuid.NewV4().String()
	userID := uuid.NewV4().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCredentialPublicKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3CredentialPublicKeyConfigSingle, "test_public_key_single_one", publicKeyIDOne, organisationID, userID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialPublicKeyExists("form3_credential_public_key.test_public_key_single_one", publicKeyIDOne),
				),
			},
			{
				Config: fmt.Sprintf(testForm3CredentialPublicKeyConfigSingle, "test_public_key_single_two", publicKeyIDTwo, organisationID, userID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialPublicKeyExists("form3_credential_public_key.test_public_key_single_two", publicKeyIDTwo),
				),
			},
			{
				Config: fmt.Sprintf(testForm3CredentialPublicKeyConfigSingle, "test_public_key_single_three", publicKeyIDThree, organisationID, userID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialPublicKeyExists("form3_credential_public_key.test_public_key_single_three", publicKeyIDThree),
				),
			},
		},
	})
}

func TestAccCredentialPublicKey_singleKey(t *testing.T) {
	log.SetOutput(os.Stdout)
	organisationID := os.Getenv("FORM3_ORGANISATION_ID")
	publicKeyID := uuid.NewV4().String()
	userID := uuid.NewV4().String()
	keyName := "test_public_key"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCredentialPublicKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3CredentialPublicKeyConfigSingle, keyName, publicKeyID, organisationID, userID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialPublicKeyExists(fmt.Sprintf("form3_credential_public_key.%s", keyName), publicKeyID)),
			},
		},
	})
}

func testAccCheckCredentialPublicKeyDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_credential_public_key" {
			continue
		}

		response, err := client.SecurityClient.Users.GetUsersUserIDCredentialsPublicKey(users.NewGetUsersUserIDCredentialsPublicKeyParams().
			WithUserID(strfmt.UUID(strfmt.UUID(rs.Primary.Attributes["user_id"]))))

		if err == nil {
			return fmt.Errorf("public key %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckCredentialPublicKeyExists(resourceKey string, publicKeyID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		log.Printf("[INFO] Checking that public key with public key id: %s exists", publicKeyID)
		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		publicKey, err := client.SecurityClient.Users.GetUsersUserIDCredentialsPublicKeyPublicKeyID(users.NewGetUsersUserIDCredentialsPublicKeyPublicKeyIDParams().
			WithUserID(strfmt.UUID(rs.Primary.Attributes["user_id"])).WithPublicKeyID(strfmt.UUID(publicKeyID)))

		if err != nil {
			return err
		}

		found := false

		if publicKey.Payload.ID.String() == publicKeyID {
			found = true
		}

		if !found {
			return fmt.Errorf("public key with id %s not found", rs.Primary.ID)
		}

		return nil
	}
}

const testForm3CredentialPublicKeyConfigSingle = `
resource "form3_credential_public_key" "%s" {
	user_id 		= "${form3_user.public_key_test_user.user_id}"
	organisation_id = "${form3_user.public_key_test_user.organisation_id}"
	public_key_id   = "%s"
	public_key      = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4JNqRbybmYHd9jlnbQwu\nw8Rg1O21IC9bns9oeeah5ZU605taCfSJUk/sEd1IKS/n4mqIi8Pm8JLiumvh1sK3\nxnqqPhxGiLLiUt9dnK3xT2WU9YEzlxRY4BbMJV12cAKI4Fu26OKrPfumud0yQLX8\nHEQSBldq0tE9tFxZi7ruzMVP7J0cNRdPtM2F97dFMeLIyh2MzXz5vIzsKprh7jaQ\nUCC2YTrpU+ZKbpvGN5Ql3KTJroiirtqQT/ZxUzLB4ChMfOLkbKTofieeNnsU2hSV\nb1Okcv5i26rzrKW2jjrIhi/QU0R/YLEc5+A06fc9Ua9U9uqyWadHkMso6xszY2Za\nEwIDAQAB\n-----END PUBLIC KEY-----\n"
}

resource "form3_user" "public_key_test_user" {
	organisation_id = "%s"
	user_id 		= "%s"
	user_name 	    = "terraform-user"
	email 			= "terraform-user@form3.tech"
	roles 			= ["32881d6b-a000-4258-b779-56c59970590f"]
}`
