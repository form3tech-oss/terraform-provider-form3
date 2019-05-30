package api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSchemeTransactionID(t *testing.T) {
	builder := PaymentBuilder{}
	assert.Equal(t, 17, len(NewSchemeTransactionID()))
	id1 := NewSchemeTransactionID()
	id2 := NewSchemeTransactionID()

	fmt.Println(id1)
	fmt.Println(id2)

	assert.NotEqual(t, id1, id2)
}
