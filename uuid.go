package form3

import (
	"github.com/go-openapi/strfmt"
	"github.com/nu7hatch/gouuid"
)

func ConvertUUIDtoPointer(s strfmt.UUID) *strfmt.UUID { return &s }
func ConvertStringtoUUID(s string) *strfmt.UUID       { return ConvertUUIDtoPointer(strfmt.UUID(s)) }
func ConvertUUIDtoStrFmtUUID(s *uuid.UUID) *strfmt.UUID {
	return ConvertUUIDtoPointer(strfmt.UUID(s.String()))
}
func NewStrFmtUUID() *strfmt.UUID {
	id, _ := uuid.NewV4()
	return ConvertUUIDtoStrFmtUUID(id)
}
