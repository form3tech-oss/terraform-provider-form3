package form3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccSepaInstantAssociation_CreateAndPatch(t *testing.T) {

	/*
		This test creates an organisation with a basic sepainstant assocation resource
		where most of the parameters are not included in the tf file and so assume the default values.
		It then patches the resource with those fields set to their non default values.
		Finally it	patches again with those fields removed to show that they are then returned to their default values
	*/

	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.New().String()
	associationId := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationId)
	bic := generateTestBic()
	businessUserDN := fmt.Sprintf("cn=%s,ou=pilot,ou=eba_ips,o=88331,dc=sianet,dc=sia,dc=eu", strings.ToLower(bic))
	testProfile1 := "TEST_PROFILE_1"
	tips := "tips"
	bTrue := true

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSepaInstantAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testTFFile{
					[]fmt.Stringer{
						organisation_resource{resource_name: "organisation", organisation_id: organisationId, parent_organisation_id: parentOrganisationId, name: testOrgName},
						sepainstant_association_resource{
							resource_name:               "association",
							dependant_organisation_name: "organisation",
							association_id:              &associationId,
							business_user_dn:            &businessUserDN,
							transport_profile_id:        &testProfile1,
							bic:                         &bic,
						},
					},
				}.String(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaInstantAssociationExists("form3_sepainstant_association.association"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "business_user_dn", businessUserDN),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "transport_profile_id", testProfile1),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "bic", bic),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "sponsor_id", ""),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "reachable_bics.#", "0"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "clearing_system", "auto"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "disable_outbound_payments", "false"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "simulator_only", "false"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "enable_customer_check", "false"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "enable_customer_admission_decision", "false"),
				),
			},
			{
				Config: testTFFile{
					[]fmt.Stringer{
						organisation_resource{resource_name: "organisation", organisation_id: organisationId, parent_organisation_id: parentOrganisationId, name: testOrgName},
						sepainstant_association_resource{
							resource_name:               "association",
							dependant_organisation_name: "organisation",
							association_id:              &associationId,
							business_user_dn:            &businessUserDN,
							transport_profile_id:        &testProfile1,
							bic:                         &bic,
							clearing_system:             &tips,
							disable_outbound_payments:   &bTrue,
							//simulator_only:                     &bTrue,
							enable_customer_admission_decision: &bTrue,
							enable_customer_check:              &bTrue,
						},
					},
				}.String(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaInstantAssociationExists("form3_sepainstant_association.association"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "business_user_dn", businessUserDN),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "transport_profile_id", testProfile1),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "bic", bic),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "sponsor_id", ""),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "reachable_bics.#", "0"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "clearing_system", "tips"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "disable_outbound_payments", "true"),
					//resource.TestCheckResourceAttr("form3_sepainstant_association.association", "simulator_only", "true"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "enable_customer_check", "true"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "enable_customer_admission_decision", "true"),
				),
			},
			{
				Config: testTFFile{
					[]fmt.Stringer{
						organisation_resource{resource_name: "organisation", organisation_id: organisationId, parent_organisation_id: parentOrganisationId, name: testOrgName},
						sepainstant_association_resource{
							resource_name:               "association",
							dependant_organisation_name: "organisation",
							association_id:              &associationId,
							business_user_dn:            &businessUserDN,
							transport_profile_id:        &testProfile1,
							bic:                         &bic,
						},
					},
				}.String(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaInstantAssociationExists("form3_sepainstant_association.association"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "business_user_dn", businessUserDN),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "transport_profile_id", testProfile1),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "bic", bic),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "sponsor_id", ""),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "reachable_bics.#", "0"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "clearing_system", "auto"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "disable_outbound_payments", "false"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "simulator_only", "false"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "enable_customer_check", "false"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association", "enable_customer_admission_decision", "false"),
				),
			},
		},
	})
}

func TestAccSepaInstantAssociation_SponsoredAssociation_CreateAndPatchReachableBics(t *testing.T) {

	/*
		This test creates an organisation with an association resource and a
		sponsored organisation with an association.
		It then patches the sponsored association resource with changes to reachable bics.
	*/

	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.New().String()
	sponsoredOrganisationId := uuid.New().String()
	associationId := uuid.New().String()
	sponsoredAssociationId := uuid.New().String()
	defer verifyOrgDoesNotExist(t, organisationId)
	defer verifyOrgDoesNotExist(t, sponsoredOrganisationId)
	bic := generateTestBic()
	sponsoredBic := generateTestBic()
	businessUserDN := fmt.Sprintf("cn=%s,ou=pilot,ou=eba_ips,o=88331,dc=sianet,dc=sia,dc=eu", strings.ToLower(bic))
	testProfile1 := "TEST_PROFILE_1"
	emptyString := ""
	sponsorId := "${form3_sepainstant_association.association.association_id}"
	reachableBics := []string{generateTestBicWithLength(11)}
	patchedReachableBics := []string{generateTestBicWithLength(11), generateTestBicWithLength(11)}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSepaInstantAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testTFFile{
					[]fmt.Stringer{
						organisation_resource{resource_name: "organisation", organisation_id: organisationId, parent_organisation_id: parentOrganisationId, name: testOrgName},
						organisation_resource{resource_name: "organisation_sponsored", organisation_id: sponsoredOrganisationId, parent_organisation_id: parentOrganisationId, name: testOrgName},
						sepainstant_association_resource{
							resource_name:               "association",
							dependant_organisation_name: "organisation",
							association_id:              &associationId,
							business_user_dn:            &businessUserDN,
							transport_profile_id:        &testProfile1,
							bic:                         &bic,
						},
						sepainstant_association_resource{
							resource_name:               "association_sponsored",
							dependant_organisation_name: "organisation_sponsored",
							association_id:              &sponsoredAssociationId,
							transport_profile_id:        &emptyString,
							business_user_dn:            &emptyString,
							bic:                         &sponsoredBic,
							reachable_bics:              reachableBics,
							sponsor_id:                  &sponsorId,
							depends_on:                  []string{"form3_sepainstant_association.association"},
						},
					},
				}.String(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaInstantAssociationExists("form3_sepainstant_association.association"),
					testAccCheckSepaInstantAssociationExists("form3_sepainstant_association.association_sponsored"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "association_id", sponsoredAssociationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "organisation_id", sponsoredOrganisationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "business_user_dn", ""),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "transport_profile_id", ""),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "bic", sponsoredBic),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "sponsor_id", associationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "reachable_bics.#", strconv.Itoa(len(reachableBics))),
				),
			},
			{
				Config: testTFFile{
					[]fmt.Stringer{
						organisation_resource{resource_name: "organisation", organisation_id: organisationId, parent_organisation_id: parentOrganisationId, name: testOrgName},
						organisation_resource{resource_name: "organisation_sponsored", organisation_id: sponsoredOrganisationId, parent_organisation_id: parentOrganisationId, name: testOrgName},
						sepainstant_association_resource{
							resource_name:               "association",
							dependant_organisation_name: "organisation",
							association_id:              &associationId,
							business_user_dn:            &businessUserDN,
							transport_profile_id:        &testProfile1,
							bic:                         &bic,
						},
						sepainstant_association_resource{
							resource_name:               "association_sponsored",
							dependant_organisation_name: "organisation_sponsored",
							association_id:              &sponsoredAssociationId,
							transport_profile_id:        &emptyString,
							business_user_dn:            &emptyString,
							bic:                         &sponsoredBic,
							reachable_bics:              patchedReachableBics,
							sponsor_id:                  &sponsorId,
							depends_on:                  []string{"form3_sepainstant_association.association"},
						},
					},
				}.String(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaInstantAssociationExists("form3_sepainstant_association.association"),
					testAccCheckSepaInstantAssociationExists("form3_sepainstant_association.association_sponsored"),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "association_id", sponsoredAssociationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "organisation_id", sponsoredOrganisationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "business_user_dn", ""),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "transport_profile_id", ""),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "bic", sponsoredBic),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "sponsor_id", associationId),
					resource.TestCheckResourceAttr("form3_sepainstant_association.association_sponsored", "reachable_bics.#", strconv.Itoa(len(patchedReachableBics))),
				),
			},
		},
	})
}

func testAccCheckSepaInstantAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_sepainstant_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetSepainstantID(associations.NewGetSepainstantIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("sepa instant record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckSepaInstantAssociationExists(resourceKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no bacs Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetSepainstantID(associations.NewGetSepainstantIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("sepa instant record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		return nil
	}
}

type sepainstant_association_resource struct {
	resource_name               string
	dependant_organisation_name string

	association_id           *string
	parent_organisation_id   *string
	organisation_sponsor_id  *string
	sponsored_association_id *string
	bic                      *string
	business_user_dn         *string
	transport_profile_id     *string
	sponsor_id               *string
	reachable_bics           []string

	clearing_system                    *string
	enable_customer_check              *bool
	enable_customer_admission_decision *bool
	simulator_only                     *bool
	disable_outbound_payments          *bool

	depends_on []string
}

func (s sepainstant_association_resource) String() string {
	dependent_org_id := fmt.Sprintf("${form3_organisation.%s.organisation_id}", s.dependant_organisation_name)

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("resource \"%s\" \"%s\" {\n", "form3_sepainstant_association", s.resource_name))
	addTestResourceStringArgument("organisation_id", &dependent_org_id, &sb)

	addTestResourceStringArgument("association_id", s.association_id, &sb)
	addTestResourceStringArgument("parent_organisation_id", s.parent_organisation_id, &sb)
	addTestResourceStringArgument("organisation_sponsor_id", s.organisation_sponsor_id, &sb)
	addTestResourceStringArgument("sponsored_association_id", s.sponsored_association_id, &sb)
	addTestResourceStringArgument("bic", s.bic, &sb)
	addTestResourceStringArgument("business_user_dn", s.business_user_dn, &sb)
	addTestResourceStringArgument("transport_profile_id", s.transport_profile_id, &sb)
	addTestResourceStringArgument("sponsor_id", s.sponsor_id, &sb)

	addTestResourceStringArgument("clearing_system", s.clearing_system, &sb)
	addTestResourceBoolArgument("enable_customer_check", s.enable_customer_check, &sb)
	addTestResourceBoolArgument("enable_customer_admission_decision", s.enable_customer_admission_decision, &sb)
	addTestResourceBoolArgument("simulator_only", s.simulator_only, &sb)
	addTestResourceBoolArgument("disable_outbound_payments", s.disable_outbound_payments, &sb)

	addTestResourceStringSliceArgument("reachable_bics", s.reachable_bics, &sb)
	addTestResourceStringSliceArgument("depends_on", s.depends_on, &sb)

	sb.WriteString("}")
	return sb.String()
}


