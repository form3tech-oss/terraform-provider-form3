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
				Config: tf_file{
					[]string{
						organisation_resource("organisation", organisationId, parentOrganisationId, testOrgName),
						sepainstant_association_resource{
							resource_name:               "association",
							dependant_organisation_name: "organisation",
							association_id:              &associationId,
							business_user_dn:            &businessUserDN,
							transport_profile_id:        &testProfile1,
							bic:                         &bic,
						}.String(),
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
				Config: tf_file{
					[]string{
						organisation_resource("organisation", organisationId, parentOrganisationId, testOrgName),
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
						}.String(),
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
				Config: tf_file{
					[]string{
						organisation_resource("organisation", organisationId, parentOrganisationId, testOrgName),
						sepainstant_association_resource{
							resource_name:               "association",
							dependant_organisation_name: "organisation",
							association_id:              &associationId,
							business_user_dn:            &businessUserDN,
							transport_profile_id:        &testProfile1,
							bic:                         &bic,
						}.String(),
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
				Config: tf_file{
					[]string{
						organisation_resource("organisation", organisationId, parentOrganisationId, testOrgName),
						organisation_resource("organisation_sponsored", sponsoredOrganisationId, parentOrganisationId, testOrgName),
						sepainstant_association_resource{
							resource_name:               "association",
							dependant_organisation_name: "organisation",
							association_id:              &associationId,
							business_user_dn:            &businessUserDN,
							transport_profile_id:        &testProfile1,
							bic:                         &bic,
						}.String(),
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
						}.String(),
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
				Config: tf_file{
					[]string{
						organisation_resource("organisation", organisationId, parentOrganisationId, testOrgName),
						organisation_resource("organisation_sponsored", sponsoredOrganisationId, parentOrganisationId, testOrgName),
						sepainstant_association_resource{
							resource_name:               "association",
							dependant_organisation_name: "organisation",
							association_id:              &associationId,
							business_user_dn:            &businessUserDN,
							transport_profile_id:        &testProfile1,
							bic:                         &bic,
						}.String(),
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
						}.String(),
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

type tf_file struct {
	resources []string
}

func (r tf_file) String() string {
	var sb strings.Builder
	sb.WriteString("\n")
	for _, r := range r.resources {
		sb.WriteString(fmt.Sprintf("%s\n\n", r))
	}
	return sb.String()
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
	addStringArgument("organisation_id", &dependent_org_id, &sb)

	addStringArgument("association_id", s.association_id, &sb)
	addStringArgument("parent_organisation_id", s.parent_organisation_id, &sb)
	addStringArgument("organisation_sponsor_id", s.organisation_sponsor_id, &sb)
	addStringArgument("sponsored_association_id", s.sponsored_association_id, &sb)
	addStringArgument("bic", s.bic, &sb)
	addStringArgument("business_user_dn", s.business_user_dn, &sb)
	addStringArgument("transport_profile_id", s.transport_profile_id, &sb)
	addStringArgument("sponsor_id", s.sponsor_id, &sb)

	addStringArgument("clearing_system", s.clearing_system, &sb)
	addBoolArgument("enable_customer_check", s.enable_customer_check, &sb)
	addBoolArgument("enable_customer_admission_decision", s.enable_customer_admission_decision, &sb)
	addBoolArgument("simulator_only", s.simulator_only, &sb)
	addBoolArgument("disable_outbound_payments", s.disable_outbound_payments, &sb)

	addStringSliceArgument("reachable_bics", s.reachable_bics, &sb)
	addStringSliceArgument("depends_on", s.depends_on, &sb)

	sb.WriteString("}")
	return sb.String()
}

func addStringArgument(argumentName string, value *string, sb *strings.Builder) {
	if value != nil {
		sb.WriteString(fmt.Sprintf("    %s = \"%s\"\n", argumentName, *value))
	}
}

func addStringSliceArgument(argumentName string, value []string, sb *strings.Builder) {
	if value != nil {
		items := make([]string, len(value))
		for i, v := range value {
			items[i] = fmt.Sprintf("\"%s\"", v)
		}
		sb.WriteString(fmt.Sprintf("    %s = [%s]\n", argumentName, strings.Join(items, ",")))
	}
}

func addBoolArgument(argumentName string, value *bool, sb *strings.Builder) {
	if value != nil {
		sb.WriteString(fmt.Sprintf("    %s = %s\n", argumentName, strconv.FormatBool(*value)))
	}
}

func organisation_resource(resource_name, organisation_id, parent_organisation_id, name string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("resource \"form3_organisation\" \"%s\" {\n", resource_name))
	addStringArgument("organisation_id", &organisation_id, &sb)
	addStringArgument("parent_organisation_id", &parent_organisation_id, &sb)
	addStringArgument("name", &name, &sb)
	sb.WriteString("}")
	return sb.String()
}
