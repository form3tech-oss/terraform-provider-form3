package form3

import (
	"github.com/ewilde/go-form3/client/accounts"
	"github.com/ewilde/go-form3/models"
	"github.com/go-openapi/strfmt"
	"github.com/nu7hatch/gouuid"
	"testing"
)

func TestAccountConfigurationPost(t *testing.T) {
	createResponse, err := createAccountConfiguration()

	assertNoErrorOccurred(err, t)

	err = getAccountConfiguration(err, createResponse.Payload.Data.ID)

	assertNoErrorOccurred(err, t)
}

func TestAccountConfigurationGetList(t *testing.T) {
	_, err := createAccountConfiguration()

	assertNoErrorOccurred(err, t)

	getAllResponse, err := auth.AccountClient.Accounts.GetAccountconfigurations(accounts.NewGetAccountconfigurationsParams())

	assertNoErrorOccurred(err, t)

	if len(getAllResponse.Payload.Data) == 0 {
		t.Error("expected at least one account configuration")
	}
}

func TestAccountConfigurationGetID(t *testing.T) {
	createResponse, err := createAccountConfiguration()

	assertNoErrorOccurred(err, t)

	err = getAccountConfiguration(err, createResponse.Payload.Data.ID)

	assertNoErrorOccurred(err, t)
}

func TestAccountConfigurationUpdate(t *testing.T) {
	createResponse, err := createAccountConfiguration()

	assertNoErrorOccurred(err, t)

	var existingAccountGenerationConfiguration = createResponse.Payload.Data.Attributes.AccountGenerationConfiguration[0]
	var newAccountGenerationConfiguration = &models.AccountGenerationConfiguration{
		Country: "NL",
		ValidAccountRanges: models.AccountGenerationConfigurationValidAccountRanges{
			&models.AccountGenerationConfigurationValidAccountRangesItems{
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
				Version:        &version,
				Attributes: &models.AccountConfigurationAttributes{
					AccountGenerationConfiguration: models.AccountConfigurationAttributesAccountGenerationConfiguration{
						existingAccountGenerationConfiguration,
						newAccountGenerationConfiguration,
					},
				},
			},
		}),
	)

	assertNoErrorOccurred(err, t)

	getConfigurationResponse, err := auth.AccountClient.Accounts.GetAccountconfigurationsID(accounts.NewGetAccountconfigurationsIDParams().
		WithID(strfmt.UUID(createResponse.Payload.Data.ID)))

	if len(getConfigurationResponse.Payload.Data.Attributes.AccountGenerationConfiguration) != 2 {
		t.Error("expected to have two account generation configurations")
	}
}

func TestAccountConfigurationDelete(t *testing.T) {
	createResponse, err := createAccountConfiguration()

	assertNoErrorOccurred(err, t)

	err = deleteAccountConfiguration(err, createResponse.Payload.Data.ID)

	assertNoErrorOccurred(err, t)

	err = getAccountConfiguration(err, createResponse.Payload.Data.ID)

	assertStatusCode(err, t, 404)
}

func getAccountConfiguration(err error, accountConfigurationId strfmt.UUID) error {
	_, err = auth.AccountClient.Accounts.GetAccountconfigurationsID(accounts.NewGetAccountconfigurationsIDParams().
		WithID(strfmt.UUID(accountConfigurationId)))
	return err
}

func deleteAccountConfiguration(err error, accountConfigurationId strfmt.UUID) error {
	_, err = auth.AccountClient.Accounts.DeleteAccountconfigurationsID(accounts.NewDeleteAccountconfigurationsIDParams().
		WithID(accountConfigurationId),
	)
	return err
}

func createAccountConfiguration() (*accounts.PostAccountconfigurationsCreated, error) {
	newId, _ := uuid.NewV4()

	createResponse, err := auth.AccountClient.Accounts.PostAccountconfigurations(accounts.NewPostAccountconfigurationsParams().
		WithAccountConfigurationCreationRequest(&models.AccountConfigurationCreation{
			Data: &models.AccountConfiguration{
				OrganisationID: organisationId,
				Type:           "account_configurations",
				Version:        &version,
				ID:             strfmt.UUID(newId.String()),
				Attributes: &models.AccountConfigurationAttributes{
					AccountGenerationEnabled: true,
					AccountGenerationConfiguration: models.AccountConfigurationAttributesAccountGenerationConfiguration{
						&models.AccountGenerationConfiguration{
							Country: "US",
							ValidAccountRanges: models.AccountGenerationConfigurationValidAccountRanges{
								&models.AccountGenerationConfigurationValidAccountRangesItems{
									Minimum: 8400000000,
									Maximum: 8409999999,
								},
							},
						},
					},
				},
			},
		}))
	return createResponse, err
}
