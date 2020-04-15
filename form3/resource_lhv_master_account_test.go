package form3

import (
	"fmt"
	"os"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccLhvMasterAccount_basic(t *testing.T) {
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.New().String()
	associationId := uuid.New().String()
	clientCode := uuid.New().String()
	masterAccountId := uuid.New().String()
	iban := uuid.New().String()
	bic := generateTestBic()
	bicID := uuid.New().String()
	bankidID := uuid.New().String()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLhvMasterAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3LhvMasterAccountConfigA, organisationId, parentOrganisationId, associationId, masterAccountId, clientCode, iban, bic, bicID, bankidID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLhvMasterAccountExists("form3_lhv_master_account.master_account"),
					resource.TestCheckResourceAttr("form3_lhv_master_account.master_account", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_lhv_master_account.master_account", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_lhv_master_account.master_account", "master_account_id", associationId),
					resource.TestCheckResourceAttr("form3_lhv_master_account.master_account", "bank_id", "999999"),
					resource.TestCheckResourceAttr("form3_lhv_master_account.master_account", "iban", iban),
					resource.TestCheckResourceAttr("form3_lhv_master_account.master_account", "bic", bic),
					resource.TestCheckResourceAttr("form3_lhv_master_account.master_account", "country", "UK"),
					resource.TestCheckResourceAttr("form3_lhv_master_account.master_account", "requires_direct_account", "false"),
				),
			},
		},
	})
}

func testAccCheckLhvMasterAccountDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_lhv_master_account" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetLhvAssociationIDMasterAccountsMasterAccountID(
			associations.NewGetLhvAssociationIDMasterAccountsMasterAccountIDParams().
				WithAssociationID(strfmt.UUID(rs.Primary.Attributes["association_id"])).
				WithMasterAccountID(strfmt.UUID(rs.Primary.Attributes["master_account_id"])))
		if err == nil {
			return fmt.Errorf("lhv master account record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckLhvMasterAccountExists(resourceKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no lhv master account ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetLhvAssociationIDMasterAccountsMasterAccountID(
			associations.NewGetLhvAssociationIDMasterAccountsMasterAccountIDParams().
				WithAssociationID(strfmt.UUID(rs.Primary.Attributes["association_id"])).
				WithMasterAccountID(strfmt.UUID(rs.Primary.Attributes["master_account_id"])))
		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("lhv master account not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		return nil
	}
}

const testForm3LhvMasterAccountConfigA = `
locals {
	parentId        = "%s"
	organisationId  = "%s"
	associationId   = "%s"
	masterAccountId = "%s"
	clientCode      = "%s"
	iban            = "%s"
	bic             = "%s"
	bic_id          = "%s"
	bankid_id       = "%s"
}

resource "form3_organisation" "organisation" {
	organisation_id        = "${local.organisationId}"
	parent_organisation_id = "${local.parentId}"
	name 		           = "terraform-organisation"
}

resource "form3_lhv_association" "association" {
	organisation_id = "${form3_organisation.organisation.organisation_id}"
	association_id  = "${local.associationId}"
	name            = "terraform-association"
	client_code     = "${local.clientCode}"
	client_country  = "UK"
}

resource "form3_bic" "bic" {
	organisation_id = "${form3_organisation.organisation.organisation_id}"
	bic_id          = "${local.bic_id}"
	bic       	    = "${local.bic}"
}

resource "form3_bank_id" "bank_id" {
	organisation_id  = "${form3_organisation.organisation.organisation_id}"
	bank_resource_id = "${local.bankid_id}"
	bank_id       	 = "999999"
	bank_id_code     = "GBDSC"
	country          = "GB"
}

resource "form3_lhv_master_account" "master_account" {
	organisation_id         = "${form3_organisation.organisation.organisation_id}"
	association_id          = "${form3_lhv_association.association.association_id}"
	master_account_id       = "${local.masterAccountId}"
	iban                    = "${local.iban}"
	bic                     = "${local.bic}"
	country                 = "GB"
	requires_direct_account = false
}
`
