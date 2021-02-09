package api

import (
	"testing"

	"github.com/form3tech-oss/terraform-provider-form3/client/accounts"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/strfmt"
)

var accountConfigurationVersion = int64(0)

func TestAccountConfigurationPost(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	createResponse, err := createAccountConfiguration()

	assertNoErrorOccurred(t, err)

	err = getAccountConfiguration(createResponse.Payload.Data.ID)

	assertNoErrorOccurred(t, err)

	err = deleteAccountConfiguration(createResponse.Payload.Data.ID)

	assertNoErrorOccurred(t, err)
}

func TestAccountConfigurationGetList(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	createResponse, err := createAccountConfiguration()

	assertNoErrorOccurred(t, err)

	getAllResponse, err := auth.AccountClient.Accounts.GetAccountconfigurations(accounts.NewGetAccountconfigurationsParams())

	assertNoErrorOccurred(t, err)

	if len(getAllResponse.Payload.Data) == 0 {
		t.Error("expected at least one account configuration")
	}

	err = deleteAccountConfiguration(createResponse.Payload.Data.ID)

	assertNoErrorOccurred(t, err)
}

func TestAccountConfigurationGetID(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	createResponse, err := createAccountConfiguration()

	assertNoErrorOccurred(t, err)

	err = getAccountConfiguration(createResponse.Payload.Data.ID)

	assertNoErrorOccurred(t, err)

	err = deleteAccountConfiguration(createResponse.Payload.Data.ID)

	assertNoErrorOccurred(t, err)
}

func TestAccountConfigurationUpdate(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	createResponse, err := createAccountConfiguration()

	assertNoErrorOccurred(t, err)

	var existingAccountGenerationConfiguration = createResponse.Payload.Data.Attributes.AccountGenerationConfiguration[0]
	var newAccountGenerationConfiguration = &models.AccountGenerationConfiguration{
		Country: "NL",
		ValidAccountRanges: []*models.AccountGenerationConfigurationValidAccountRangesItems{
			{
				Minimum: 1234567890,
				Maximum: 9999999999,
			},
		},
	}

	_, err = auth.AccountClient.Accounts.PatchAccountconfigurationsID(accounts.NewPatchAccountconfigurationsIDParams().
		WithID(createResponse.Payload.Data.ID).
		WithConfigAmendRequest(&models.ConfigurationAmendment{
			Data: &models.AccountConfiguration{
				ID:             createResponse.Payload.Data.ID,
				OrganisationID: organisationId,
				Version:        &accountConfigurationVersion,
				Attributes: &models.AccountConfigurationAttributes{
					AccountGenerationConfiguration: []*models.AccountGenerationConfiguration{
						existingAccountGenerationConfiguration,
						newAccountGenerationConfiguration,
					},
				},
			},
		}),
	)

	assertNoErrorOccurred(t, err)

	getConfigurationResponse, err := auth.AccountClient.Accounts.GetAccountconfigurationsID(accounts.NewGetAccountconfigurationsIDParams().
		WithID(createResponse.Payload.Data.ID))

	assertNoErrorOccurred(t, err)

	if len(getConfigurationResponse.Payload.Data.Attributes.AccountGenerationConfiguration) != 2 {
		t.Error("expected to have two account generation configurations")
	}

	accountConfigurationVersion = *getConfigurationResponse.Payload.Data.Version

	err = deleteAccountConfiguration(createResponse.Payload.Data.ID)

	assertNoErrorOccurred(t, err)
}

func TestAccountConfigurationDelete(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	createResponse, err := createAccountConfiguration()

	assertNoErrorOccurred(t, err)

	err = deleteAccountConfiguration(createResponse.Payload.Data.ID)

	assertNoErrorOccurred(t, err)

	err = getAccountConfiguration(createResponse.Payload.Data.ID)

	assertStatusCode(t, err, 404)
}

func getAccountConfiguration(accountConfigurationId strfmt.UUID) error {
	_, err := auth.AccountClient.Accounts.GetAccountconfigurationsID(accounts.NewGetAccountconfigurationsIDParams().
		WithID(accountConfigurationId))
	return err
}

func deleteAccountConfiguration(accountConfigurationId strfmt.UUID) error {
	_, err := auth.AccountClient.Accounts.DeleteAccountconfigurationsID(accounts.NewDeleteAccountconfigurationsIDParams().
		WithID(accountConfigurationId).
		WithVersion(accountConfigurationVersion),
	)
	return err
}

func createAccountConfiguration() (*accounts.PostAccountconfigurationsCreated, error) {
	newId := NewUUID()

	createResponse, err := auth.AccountClient.Accounts.PostAccountconfigurations(accounts.NewPostAccountconfigurationsParams().
		WithAccountConfigurationCreationRequest(&models.AccountConfigurationCreation{
			Data: &models.AccountConfiguration{
				OrganisationID: testOrganisationId,
				Type:           "account_configurations",
				Version:        &accountConfigurationVersion,
				ID:             *newId,
				Attributes: &models.AccountConfigurationAttributes{
					AccountGenerationEnabled: true,
					AccountGenerationConfiguration: []*models.AccountGenerationConfiguration{
						{
							Country: "US",
							ValidAccountRanges: []*models.AccountGenerationConfigurationValidAccountRangesItems{
								{
									Minimum: 84000000,
									Maximum: 84099999,
								},
							},
						},
					},
				},
			},
		}))
	return createResponse, err
}
