package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSchemeTransactionID(t *testing.T) {
	initOrgs, _ := auth.OrganisationClient.Organisations.GetUnits(nil)
	defer assertNoOrgLeak(t, auth, initOrgs.Payload.Data)
	builder := PaymentBuilder{}
	assert.Equal(t, 17, len(builder.NewSchemeTransactionID()))
	id1 := builder.NewSchemeTransactionID()
	id2 := builder.NewSchemeTransactionID()

	assert.NotEqual(t, id1, id2)
}
