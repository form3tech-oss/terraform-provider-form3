package form3

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"text/template"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccLhvAgencySynchronisation_basic(t *testing.T) {
	data := lhvAgencySynchronisationConfigData{
		OrganisationID:          uuid.New().String(),
		ParentOrganisationID:    os.Getenv("FORM3_ORGANISATION_ID"),
		AssociationId:           uuid.New().String(),
		ClientCode:              uuid.New().String(),
		AgencySynchronisationId: uuid.New().String(),
		Bic:                     generateTestBic(),
		BankId:                  []string{"040301", "040302", "040303"},
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLhvAgencySynchronisationDestroy,
		Steps: []resource.TestStep{
			{
				Config: lhvAgencySynchronisationConfig(data),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLhvAgencySynchronisationExists("form3_lhv_agency_synchronisation.agency_synchronisation"),
					resource.TestCheckResourceAttr("form3_lhv_agency_synchronisation.agency_synchronisation", "association_id", data.AssociationId),
					resource.TestCheckResourceAttr("form3_lhv_agency_synchronisation.agency_synchronisation", "organisation_id", data.OrganisationID),
					resource.TestCheckResourceAttr("form3_lhv_agency_synchronisation.agency_synchronisation", "agency_synchronisation_id", data.AgencySynchronisationId),
					resource.TestCheckResourceAttr("form3_lhv_agency_synchronisation.agency_synchronisation", "bic", data.Bic),
					resource.TestCheckResourceAttr("form3_lhv_agency_synchronisation.agency_synchronisation", "country", "GB"),
					resource.TestCheckResourceAttr("form3_lhv_agency_synchronisation.agency_synchronisation", "bank_id.#", "3"),
					resource.TestCheckResourceAttr("form3_lhv_agency_synchronisation.agency_synchronisation", "bank_id.0", "040301"),
					resource.TestCheckResourceAttr("form3_lhv_agency_synchronisation.agency_synchronisation", "bank_id.1", "040302"),
					resource.TestCheckResourceAttr("form3_lhv_agency_synchronisation.agency_synchronisation", "bank_id.2", "040303"),
				),
			},
		},
	})
}

func testAccCheckLhvAgencySynchronisationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_lhv_agency_synchronisation" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetLhvAssociationIDAgencySynchronisationAgencySynchronisationID(
			associations.NewGetLhvAssociationIDAgencySynchronisationAgencySynchronisationIDParams().
				WithAssociationID(strfmt.UUID(rs.Primary.Attributes["association_id"])).
				WithAgencySynchronisationID(strfmt.UUID(rs.Primary.Attributes["agency_synchronisation_id"])))
		if err == nil {
			return fmt.Errorf("lhv agency synchronisation record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckLhvAgencySynchronisationExists(resourceKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no lhv agency synchronisation ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetLhvAssociationIDAgencySynchronisationAgencySynchronisationID(
			associations.NewGetLhvAssociationIDAgencySynchronisationAgencySynchronisationIDParams().
				WithAssociationID(strfmt.UUID(rs.Primary.Attributes["association_id"])).
				WithAgencySynchronisationID(strfmt.UUID(rs.Primary.Attributes["agency_synchronisation_id"])))
		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("lhv lhv agency synchronisation not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		return nil
	}
}

type lhvAgencySynchronisationConfigData struct {
	OrganisationID          string
	ParentOrganisationID    string
	AssociationId           string
	ClientCode              string
	AgencySynchronisationId string
	Bic                     string
	BankId                  []string
}

func lhvAgencySynchronisationConfig(data lhvAgencySynchronisationConfigData) string {
	var sb strings.Builder

	err := template.Must(template.New("tpl").Parse(`
resource "form3_organisation" "organisation" {
	organisation_id        = "{{ .OrganisationID }}"
	parent_organisation_id = "{{ .ParentOrganisationID }}"
	name 		           = "terraform-organisation"
}

resource "form3_lhv_association" "association" {
	organisation_id = "${form3_organisation.organisation.organisation_id}"
	association_id  = "{{ .AssociationId }}"
	name            = "terraform-association"
	client_code     = "{{ .ClientCode }}"
	client_country  = "GB"
}

resource "form3_lhv_agency_synchronisation" "agency_synchronisation" {
	organisation_id           = "${form3_organisation.organisation.organisation_id}"
	association_id            = "${form3_lhv_association.association.association_id}"
	agency_synchronisation_id = "{{ .AgencySynchronisationId }}"
	bic                       = "{{ .Bic }}"
	bank_id                   = [{{ range .BankId }}"{{ . }}",{{ end }}]
	country                   = "GB"
}
`)).Execute(&sb, data)
	if err != nil {
		panic(err)
	}

	return sb.String()
}
