package api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSchemeTransactionID(t *testing.T) {
	builder := PaymentBuilder{}
	assert.Equal(t, 17, len(builder.NewSchemeTransactionID()))
	id1 := builder.NewSchemeTransactionID()
	id2 := builder.NewSchemeTransactionID()

	fmt.Println(id1)
	fmt.Println(id2)

	assert.NotEqual(t, id1, id2)
}
