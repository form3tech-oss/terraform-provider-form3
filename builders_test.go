package form3

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewSchemeTransactionID(t *testing.T) {
	builder := PaymentBuilder{}
	assert.Equal(t, 31, len(builder.NewSchemeTransactionID()))
}
