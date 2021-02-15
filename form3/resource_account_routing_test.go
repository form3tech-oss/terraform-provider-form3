package form3

import (
	"fmt"
	"os"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/account_routings"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccAccountRouting_basic(t *testing.T) {
	var accountRoutingResponse account_routings.GetAccountRoutingsIDOK
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)
	accountRoutingID := uuid.New().String()
	accountGenerator := "accountapi"
	accountProvisioner := "accountapi"
	match := "*"
	priority := 1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAccountRoutingDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3AccountRoutingTFConfig(organisationID, parentOrganisationID, testOrgName, accountRoutingID, accountGenerator, accountProvisioner, match, priority),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAccountRoutingExists("form3_account_routing.account_routing", &accountRoutingResponse),
					resource.TestCheckResourceAttr("form3_account_routing.account_routing", "account_routing_id", accountRoutingID),
					resource.TestCheckResourceAttr("form3_account_routing.account_routing", "organisation_id", organisationID),
					resource.TestCheckResourceAttr("form3_account_routing.account_routing", "account_generator", accountGenerator),
					resource.TestCheckResourceAttr("form3_account_routing.account_routing", "account_provisioner", accountProvisioner),
					resource.TestCheckResourceAttr("form3_account_routing.account_routing", "match", match),
					resource.TestCheckResourceAttr("form3_account_routing.account_routing", "priority", fmt.Sprintf("%d", priority)),
				),
			},
		},
	})
}

func TestAccAccountRouting_importBasic(t *testing.T) {

	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	organisationID := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationID)
	accountRoutingID := uuid.New().String()
	accountGenerator := "accountapi"
	accountProvisioner := "accountapi"
	match := "*"
	priority := 1

	resourceName := "form3_account_routing.account_routing"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAccountRoutingDestroy,
		Steps: []resource.TestStep{
			{
				Config: getTestForm3AccountRoutingTFConfig(organisationID, parentOrganisationID, testOrgName, accountRoutingID, accountGenerator, accountProvisioner, match, priority),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckAccountRoutingExists(resourceKey string, accountRoutingResponse *account_routings.GetAccountRoutingsIDOK) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AccountClient.AccountRoutings.GetAccountRoutingsID(account_routings.NewGetAccountRoutingsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		accountRoutingResponse = foundRecord
		return nil
	}
}

func testAccCheckAccountRoutingDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_account_routing" {
			continue
		}

		response, err := client.AccountClient.AccountRoutings.GetAccountRoutingsID(account_routings.NewGetAccountRoutingsIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func getTestForm3AccountRoutingTFConfig(organisationID, parentOrganisationID, orgName, accountRoutingID, accountGenerator, accountProvisioner, match string, priority int) string {
	return fmt.Sprintf(`resource "form3_organisation" "organisation" {
		organisation_id        = "%s"
		parent_organisation_id = "%s"
		name 		           = "%s"
	}
	
	resource "form3_account_routing" "account_routing" {
		account_routing_id   = "%s"
		organisation_id      = "${form3_organisation.organisation.organisation_id}"
		account_generator 	 = "%s"
		account_provisioner  = "%s"
		match                = "%s"
		priority             = %d
	}
	`, organisationID, parentOrganisationID, orgName, accountRoutingID, accountGenerator, accountProvisioner, match, priority)
}
