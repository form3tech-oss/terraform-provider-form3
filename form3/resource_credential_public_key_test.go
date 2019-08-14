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
			// Given I have created an user
			// When i add first public key to this user
			// When i add second public key to this user
			// When i add third public key to this user
			// Then I can see public keys added
			{
				Config: fmt.Sprintf(testForm3CredentialPublicKeyConfigMulti, organisationID, userID, publicKeyIDOne, publicKeyIDTwo, publicKeyIDThree),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialPublicKeysExists(
						[]string{
							"form3_credential_public_key.test_public_key_multi_one",
							"form3_credential_public_key.test_public_key_multi_two",
							"form3_credential_public_key.test_public_key_multi_three",
						}, []string{publicKeyIDOne, publicKeyIDTwo, publicKeyIDThree}),
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

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCredentialPublicKeyDestroy,
		Steps: []resource.TestStep{
			// Given I have created an user
			// When i add public key to this user
			// Then I can see a public key was added
			{
				Config: fmt.Sprintf(testForm3CredentialPublicKeyConfigSingle, organisationID, userID, publicKeyID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialPublicKeysExists([]string{"form3_credential_public_key.test_public_key_single"}, []string{publicKeyID})),
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

func testAccCheckCredentialPublicKeysExists(resourceKeys []string, publicKeyIDs []string) resource.TestCheckFunc {

	if len(resourceKeys) != len(publicKeyIDs) {
		return func(s *terraform.State) error {
			return fmt.Errorf("incorrect parameters - resource keys list len doesnt match public key ids len")
		}
	}

	return func(s *terraform.State) error {
		var i int

		for i = 0; i < len(resourceKeys); i++ {
			resourceKey := resourceKeys[i]
			publicKeyID := publicKeyIDs[i]

			rs, ok := s.RootModule().Resources[resourceKey]

			if !ok {
				return fmt.Errorf("resource not found: %s", resourceKey)
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

		}

		return nil
	}
}

const testForm3CredentialPublicKeyConfigMulti = `
resource "form3_user" "public_key_test_user" {
	organisation_id = "%s"
	user_id 		= "%s"
	user_name 	    = "terraform-user"
	email 			= "terraform-user@form3.tech"
	roles 			= ["32881d6b-a000-4258-b779-56c59970590f"]
}

resource "form3_credential_public_key" "test_public_key_multi_one" {
	user_id 		= "${form3_user.public_key_test_user.user_id}"
	organisation_id = "${form3_user.public_key_test_user.organisation_id}"
	public_key_id   = "%s"
	public_key      = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4JNqRbybmYHd9jlnbQwu\nw8Rg1O21IC9bns9oeeah5ZU605taCfSJUk/sEd1IKS/n4mqIi8Pm8JLiumvh1sK3\nxnqqPhxGiLLiUt9dnK3xT2WU9YEzlxRY4BbMJV12cAKI4Fu26OKrPfumud0yQLX8\nHEQSBldq0tE9tFxZi7ruzMVP7J0cNRdPtM2F97dFMeLIyh2MzXz5vIzsKprh7jaQ\nUCC2YTrpU+ZKbpvGN5Ql3KTJroiirtqQT/ZxUzLB4ChMfOLkbKTofieeNnsU2hSV\nb1Okcv5i26rzrKW2jjrIhi/QU0R/YLEc5+A06fc9Ua9U9uqyWadHkMso6xszY2Za\nEwIDAQAB\n-----END PUBLIC KEY-----\n"
}

resource "form3_credential_public_key" "test_public_key_multi_two" {
	user_id 		= "${form3_user.public_key_test_user.user_id}"
	organisation_id = "${form3_user.public_key_test_user.organisation_id}"
	public_key_id   = "%s"
	public_key      = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4JNqRbybmYHd9jlnbQwu\nw8Rg1O21IC9bns9oeeah5ZU605taCfSJUk/sEd1IKS/n4mqIi8Pm8JLiumvh1sK3\nxnqqPhxGiLLiUt9dnK3xT2WU9YEzlxRY4BbMJV12cAKI4Fu26OKrPfumud0yQLX8\nHEQSBldq0tE9tFxZi7ruzMVP7J0cNRdPtM2F97dFMeLIyh2MzXz5vIzsKprh7jaQ\nUCC2YTrpU+ZKbpvGN5Ql3KTJroiirtqQT/ZxUzLB4ChMfOLkbKTofieeNnsU2hSV\nb1Okcv5i26rzrKW2jjrIhi/QU0R/YLEc5+A06fc9Ua9U9uqyWadHkMso6xszY2Za\nEwIDAQAB\n-----END PUBLIC KEY-----\n"

depends_on = [ "form3_credential_public_key.test_public_key_multi_one" ]
}

resource "form3_credential_public_key" "test_public_key_multi_three" {
	user_id 		= "${form3_user.public_key_test_user.user_id}"
	organisation_id = "${form3_user.public_key_test_user.organisation_id}"
	public_key_id   = "%s"
	public_key      = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4JNqRbybmYHd9jlnbQwu\nw8Rg1O21IC9bns9oeeah5ZU605taCfSJUk/sEd1IKS/n4mqIi8Pm8JLiumvh1sK3\nxnqqPhxGiLLiUt9dnK3xT2WU9YEzlxRY4BbMJV12cAKI4Fu26OKrPfumud0yQLX8\nHEQSBldq0tE9tFxZi7ruzMVP7J0cNRdPtM2F97dFMeLIyh2MzXz5vIzsKprh7jaQ\nUCC2YTrpU+ZKbpvGN5Ql3KTJroiirtqQT/ZxUzLB4ChMfOLkbKTofieeNnsU2hSV\nb1Okcv5i26rzrKW2jjrIhi/QU0R/YLEc5+A06fc9Ua9U9uqyWadHkMso6xszY2Za\nEwIDAQAB\n-----END PUBLIC KEY-----\n"

depends_on = [ "form3_credential_public_key.test_public_key_multi_two" ]
}`

const testForm3CredentialPublicKeyConfigSingle = `
resource "form3_user" "public_key_test_user" {
	organisation_id = "%s"
	user_id 		= "%s"
	user_name 	    = "terraform-user"
	email 			= "terraform-user@form3.tech"
	roles 			= ["32881d6b-a000-4258-b779-56c59970590f"]
}

resource "form3_credential_public_key" "test_public_key_single" {
	user_id 		= "${form3_user.public_key_test_user.user_id}"
	organisation_id = "${form3_user.public_key_test_user.organisation_id}"
	public_key_id   = "%s"
	public_key      = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4JNqRbybmYHd9jlnbQwu\nw8Rg1O21IC9bns9oeeah5ZU605taCfSJUk/sEd1IKS/n4mqIi8Pm8JLiumvh1sK3\nxnqqPhxGiLLiUt9dnK3xT2WU9YEzlxRY4BbMJV12cAKI4Fu26OKrPfumud0yQLX8\nHEQSBldq0tE9tFxZi7ruzMVP7J0cNRdPtM2F97dFMeLIyh2MzXz5vIzsKprh7jaQ\nUCC2YTrpU+ZKbpvGN5Ql3KTJroiirtqQT/ZxUzLB4ChMfOLkbKTofieeNnsU2hSV\nb1Okcv5i26rzrKW2jjrIhi/QU0R/YLEc5+A06fc9Ua9U9uqyWadHkMso6xszY2Za\nEwIDAQAB\n-----END PUBLIC KEY-----\n"
}`
