package form3

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewSchemeTransactionID(t *testing.T) {
	builder := PaymentBuilder{}
	assert.Equal(t, 17, len(builder.NewSchemeTransactionID()))
}
