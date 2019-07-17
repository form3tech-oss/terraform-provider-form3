package form3

import (
	"fmt"
	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/users"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"log"
	"os"
	"testing"
)

func TestAccCredentialPublicKey_basic(t *testing.T) {
	var publicKeyresponse models.PublicKey
	organisationId := os.Getenv("FORM3_ORGANISATION_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCredentialPublicKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3CredentialPublicKeyConfigA, organisationId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCredentialPublicKeyExists("form3_credential_public_key.public_key", &publicKeyresponse)),
			},
		},
	})
}

func testAccCheckCredentialPublicKeyDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_credential" {
			continue
		}

		_, err := client.SecurityClient.Users.GetUsersUserIDCredentialsPublicKey(users.NewGetUsersUserIDCredentialsPublicKeyParams().
			WithUserID(strfmt.UUID(strfmt.UUID(rs.Primary.Attributes["user_id"]))))

		if err == nil {
			return fmt.Errorf("error listing credentials: %s", err)
		}
	}

	return nil
}

func testAccCheckCredentialPublicKeyExists(resourceKey string, publicKey *models.PublicKey) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		log.Printf("[INFO] Checking that public key with public key id: %s exists", rs.Primary.ID)
		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		publicKeyList, err := client.SecurityClient.Users.GetUsersUserIDCredentialsPublicKey(users.NewGetUsersUserIDCredentialsPublicKeyParams().
			WithUserID(strfmt.UUID(rs.Primary.Attributes["user_id"])))

		if err != nil {
			return err
		}

		found := false
		for _, element := range publicKeyList.Payload.Data {
			log.Printf("[DEBUG] Checking that public key with id: %s exists comparing with %s", rs.Primary.ID, element.PublicKeyID.String())
			if element.PublicKeyID.String() == rs.Primary.ID {
				found = true
			}
		}

		if !found {
			return fmt.Errorf("credential record not found")
		}

		return nil
	}
}

const testForm3CredentialPublicKeyConfigA = `
resource "form3_credential_public_key" "public_key" {
	user_id 		= "${form3_user.user.user_id}"
	organisation_Id = "${form3_user.user.organisation_id}"
	public_key_id   = "${uuid()}"
	public_key   = ""-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4JNqRbybmYHd9jlnbQwu\nw8Rg1O21IC9bns9oeeah5ZU605taCfSJUk/sEd1IKS/n4mqIi8Pm8JLiumvh1sK3\nxnqqPhxGiLLiUt9dnK3xT2WU9YEzlxRY4BbMJV12cAKI4Fu26OKrPfumud0yQLX8\nHEQSBldq0tE9tFxZi7ruzMVP7J0cNRdPtM2F97dFMeLIyh2MzXz5vIzsKprh7jaQ\nUCC2YTrpU+ZKbpvGN5Ql3KTJroiirtqQT/ZxUzLB4ChMfOLkbKTofieeNnsU2hSV\nb1Okcv5i26rzrKW2jjrIhi/QU0R/YLEc5+A06fc9Ua9U9uqyWadHkMso6xszY2Za\nEwIDAQAB\n-----END PUBLIC KEY-----\n"
}

resource "form3_user" "user" {
	organisation_id = "%s"
	user_id 		= "${uuid()}"
	user_name 	    = "terraform-user"
	email 			= "terraform-user@form3.tech"
	roles 			= ["32881d6b-a000-4258-b779-56c59970590f"]

  lifecycle {
    ignore_changes = ["user_id"]
  }

}`
