package api

import (
	"context"
	"github.com/form3tech-oss/terraform-provider-form3/client/system"
	"github.com/form3tech-oss/terraform-provider-form3/models"
	"github.com/go-openapi/swag"
	"github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostKey(t *testing.T) {
	createResponse := createKey(t)
	defer deleteKey(createResponse, t)

	actualOrganisationId := createResponse.Payload.Data.OrganisationID.String()
	if actualOrganisationId != organisationId.String() {
		t.Fatalf("Expected %s, got %s", organisationId.String(), actualOrganisationId)
	}

	assert.Equal(t, "C=GB, O=Test Limited, OU=Test Bank, CN=12344321", createResponse.Payload.Data.Attributes.Subject)
	assert.Equal(t, "go-form3 testing", createResponse.Payload.Data.Attributes.Description)
	assert.Equal(t, "RSA", *createResponse.Payload.Data.Attributes.Type)
	assert.Contains(t, createResponse.Payload.Data.Attributes.PrivateKey, "BEGIN RSA PRIVATE KEY")
	assert.Contains(t, createResponse.Payload.Data.Attributes.PublicKey, "BEGIN PUBLIC KEY")
	assert.Contains(t, createResponse.Payload.Data.Attributes.CertificateSigningRequest, "CERTIFICATE")

}

func TestPostEllipticCurveKey(t *testing.T) {
	createResponse := createEcKey(t)
	defer deleteKey(createResponse, t)

	actualOrganisationId := createResponse.Payload.Data.OrganisationID.String()
	if actualOrganisationId != organisationId.String() {
		t.Fatalf("Expected %s, got %s", organisationId.String(), actualOrganisationId)
	}

	assert.Equal(t, "C=GB, O=Test Limited, OU=Test Bank, CN=12344321", createResponse.Payload.Data.Attributes.Subject)
	assert.Equal(t, "go-form3 testing", createResponse.Payload.Data.Attributes.Description)
	assert.Equal(t, "EC", *createResponse.Payload.Data.Attributes.Type)
	assert.Contains(t, createResponse.Payload.Data.Attributes.PrivateKey, "BEGIN EC PRIVATE KEY")
	assert.Contains(t, createResponse.Payload.Data.Attributes.PublicKey, "BEGIN PUBLIC KEY")
	assert.Contains(t, createResponse.Payload.Data.Attributes.CertificateSigningRequest, "CERTIFICATE")

}

func deleteKey(createResponse *system.PostKeysCreated, t *testing.T) {

	key, _ := auth.SystemClient.System.GetKeysKeyID(system.NewGetKeysKeyIDParams().WithKeyID(createResponse.Payload.Data.ID))

	_, err := auth.SystemClient.System.DeleteKeysKeyID(system.NewDeleteKeysKeyIDParams().
		WithKeyID(createResponse.Payload.Data.ID).WithVersion(*key.Payload.Data.Version),
	)
	assertNoErrorOccurred(err, t)
}

func createKey(t *testing.T) *system.PostKeysCreated {
	id, _ := uuid.NewV4()
	createResponse, err := auth.SystemClient.System.PostKeys(system.NewPostKeysParams().
		WithKeyCreationRequest(&models.KeyCreation{
			Data: &models.Key{
				ID:             *UUIDtoStrFmtUUID(id),
				OrganisationID: organisationId,
				Attributes: &models.KeyAttributes{
					Subject:     "C=GB, O=Test Limited, OU=Test Bank, CN=12344321",
					Description: "go-form3 testing",
				},
			},
		}))
	assertNoErrorOccurred(err, t)
	return createResponse
}

func createEcKey(t *testing.T) *system.PostKeysCreated {
	id, _ := uuid.NewV4()
	createResponse, err := auth.SystemClient.System.PostKeys(system.NewPostKeysParams().
		WithKeyCreationRequest(&models.KeyCreation{
			Data: &models.Key{
				ID:             *UUIDtoStrFmtUUID(id),
				OrganisationID: organisationId,
				Attributes: &models.KeyAttributes{
					Subject:     "C=GB, O=Test Limited, OU=Test Bank, CN=12344321",
					Description: "go-form3 testing",
					Type:        swag.String("EC"),
					Curve:       "PRIME256V1",
				},
			},
		}))
	assertNoErrorOccurred(err, t)
	return createResponse
}

func TestDeleteKey(t *testing.T) {
	createResponse := createKey(t)

	deleteKey(createResponse, t)

	list, err := auth.SystemClient.System.GetKeys(system.NewGetKeysParams())
	require.Nil(t, err)

	for _, certificateRequest := range list.Payload.Data {
		assert.NotEqual(t, createResponse.Payload.Data.ID, certificateRequest.ID)
	}
}

func TestGetKey(t *testing.T) {
	createResponse := createKey(t)
	defer deleteKey(createResponse, t)

	list, err := auth.SystemClient.System.GetKeys(system.NewGetKeysParams())
	require.Nil(t, err)

	foundRequest := false
	for _, certificateRequest := range list.Payload.Data {
		if createResponse.Payload.Data.ID == certificateRequest.ID {
			foundRequest = true
		}
	}
	assert.True(t, foundRequest)

	key, err := auth.SystemClient.System.GetKeysKeyID(system.NewGetKeysKeyIDParams().WithKeyID(createResponse.Payload.Data.ID))
	require.Nil(t, err)
	assert.Equal(t, "C=GB, O=Test Limited, OU=Test Bank, CN=12344321", key.Payload.Data.Attributes.Subject)

}

func TestPostKeyCertificate(t *testing.T) {
	createResponse := createKey(t)
	defer deleteKey(createResponse, t)

	createCertificateResponse := createCertificate(createResponse, t)
	defer deleteCertificate(createResponse, createCertificateResponse, t)

	assert.Equal(t, "Issuing Cert", createCertificateResponse.Payload.Data.Attributes.IssuingCertificates[0])
	assert.Equal(t, "Test Cert", *createCertificateResponse.Payload.Data.Attributes.Certificate)
}

func TestDeleteKeyCertificate(t *testing.T) {
	createResponse := createKey(t)
	defer deleteKey(createResponse, t)

	createCertificateResponse := createCertificate(createResponse, t)
	deleteCertificate(createResponse, createCertificateResponse, t)

	// check that delete worked by creating a new certificate - can only have one certificate per request
	createCertificateResponse2 := createCertificate(createResponse, t)
	defer deleteCertificate(createResponse, createCertificateResponse2, t)
}

func deleteCertificate(keysCreated *system.PostKeysCreated, certCreation *system.PostKeysKeyIDCertificatesCreated, t *testing.T) {

	cert, err := auth.SystemClient.System.GetKeysKeyIDCertificatesCertificateID(system.NewGetKeysKeyIDCertificatesCertificateIDParams().WithKeyID(keysCreated.Payload.Data.ID).
		WithCertificateID(certCreation.Payload.Data.ID))
	assertNoErrorOccurred(err, t)

	_, err = auth.SystemClient.System.DeleteKeysKeyIDCertificatesCertificateID(&system.DeleteKeysKeyIDCertificatesCertificateIDParams{
		KeyID:         keysCreated.Payload.Data.ID,
		CertificateID: certCreation.Payload.Data.ID,
		Version:       *cert.Payload.Data.Version,
		Context:       context.Background(),
	})
	assertNoErrorOccurred(err, t)
}

func createCertificate(createResponse *system.PostKeysCreated, t *testing.T) *system.PostKeysKeyIDCertificatesCreated {
	id, _ := uuid.NewV4()
	certName := "Test Cert"
	createCertResponse, err := auth.SystemClient.System.PostKeysKeyIDCertificates(system.NewPostKeysKeyIDCertificatesParams().
		WithKeyID(createResponse.Payload.Data.ID).
		WithCertificateCreationRequest(&models.CertificateCreation{
			Data: &models.Certificate{
				ID:             *UUIDtoStrFmtUUID(id),
				OrganisationID: organisationId,
				Attributes: &models.CertificateAttributes{
					Certificate:         &certName,
					IssuingCertificates: []string{"Issuing Cert"},
				},
			},
		}))
	assertNoErrorOccurred(err, t)
	return createCertResponse
}
