package form3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
)

func TestAccSepaSctAssociation_basic(t *testing.T) {
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationId)

	associationId := uuid.New().String()
	bic := generateTestBic()

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },

		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSepaSctAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3SepaSctAssociationConfig, organisationId, parentOrganisationId, associationId, bic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaSctAssociationExists("form3_sepasct_association.association", false),
					resource.TestCheckResourceAttr("form3_sepasct_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_sepasct_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_sepasct_association.association", "bic", bic),
					resource.TestCheckResourceAttr("form3_sepasct_association.association", "business_user", "PR344567"),
					resource.TestCheckResourceAttr("form3_sepasct_association.association", "receiver_business_user", "PR344568"),
				),
			},
		},
	})
}

func TestAccSepaSctAssociation_reachable(t *testing.T) {
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	sponsorOrganisationId := uuid.New().String()
	sponsorAssociationId := uuid.New().String()
	sponsorBic := generateTestBic()
	sponsorResourceKey := "form3_sepasct_association.sponsor_association"

	sponsoredOrganisationId := uuid.New().String()
	sponsoredAssociationId := uuid.New().String()
	sponsoredBicList := []string{generateTestBicWithLength(11), generateTestBicWithLength(11), generateTestBicWithLength(11)}
	sponsoredResourceKey := "form3_sepasct_association.sponsored_association"

	config := fmt.Sprintf(
		testForm3SepaSctSponsoredAssociationConfig,
		sponsorOrganisationId, parentOrganisationId, sponsorAssociationId, sponsorBic,
		sponsoredOrganisationId, parentOrganisationId, sponsoredAssociationId, strings.Join(sponsoredBicList, "\",\""),
	)
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },

		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSepaSctAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaSctAssociationExists(sponsorResourceKey, true),
					resource.TestCheckResourceAttr(sponsorResourceKey, "association_id", sponsorAssociationId),
					resource.TestCheckResourceAttr(sponsorResourceKey, "organisation_id", sponsorOrganisationId),
					resource.TestCheckResourceAttr(sponsorResourceKey, "bic", sponsorBic),
					resource.TestCheckResourceAttr(sponsorResourceKey, "can_sponsor", "true"),
					resource.TestCheckResourceAttr(sponsorResourceKey, "business_user", "PR344569"),
					resource.TestCheckResourceAttr(sponsorResourceKey, "receiver_business_user", "PR344570"),

					testAccCheckSepaSctAssociationExists(sponsoredResourceKey, false),
					resource.TestCheckResourceAttr(sponsoredResourceKey, "association_id", sponsoredAssociationId),
					resource.TestCheckResourceAttr(sponsoredResourceKey, "organisation_id", sponsoredOrganisationId),
					resource.TestCheckResourceAttr(sponsoredResourceKey, "sponsor_id", sponsorAssociationId),
					resource.TestCheckResourceAttr(sponsoredResourceKey, "reachable_bics.#", strconv.Itoa(len(sponsoredBicList))),
				),
			},
		},
	})
}

func testAccCheckSepaSctAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_sepasct_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetSepasctID(associations.NewGetSepasctIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("sepa sct record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckSepaSctAssociationExists(resourceKey string, canSponsor bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no sepasct Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetSepasctID(associations.NewGetSepasctIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("sepa sct record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		if foundRecord.Payload.Data.Attributes.CanSponsor != canSponsor {
			return fmt.Errorf("sepa sct association can_sponsor mismatch, want %v found %v", canSponsor, foundRecord.Payload.Data.Attributes.CanSponsor)
		}

		return nil
	}
}

const (
	testForm3SepaSctAssociationConfig = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		               = "terraform-provider-form3-test-organisation"
}

resource "form3_sepasct_association" "association" {
	organisation_id        = "${form3_organisation.organisation.organisation_id}"
	association_id         = "%s"
	bic                    = "%s"
  business_user          = "PR344567"
  receiver_business_user = "PR344568"
}`
	testForm3SepaSctSponsoredAssociationConfig = `
resource "form3_organisation" "sponsor_organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		           = "terraform-provider-form3-test-organisation"
}

resource "form3_sepasct_association" "sponsor_association" {
	organisation_id        = "${form3_organisation.sponsor_organisation.organisation_id}"
	association_id         = "%s"
	bic                    = "%s"
	can_sponsor            = true
    business_user          = "PR344569"
    receiver_business_user = "PR344570"
}

resource "form3_organisation" "sponsored_organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		           = "terraform-provider-form3-test-organisation"
}

resource "form3_sepasct_association" "sponsored_association" {
	organisation_id        = "${form3_organisation.sponsored_organisation.organisation_id}"
	association_id         = "%s"
	reachable_bics         = ["%s"]
	sponsor_id             = "${form3_sepasct_association.sponsor_association.association_id}"
}`
)
