package form3

import (
	"github.com/go-openapi/strfmt"
	"github.com/nu7hatch/gouuid"
)

func UUID(s strfmt.UUID) *strfmt.UUID { return &s }

func UUIDValue(s *strfmt.UUID) strfmt.UUID { return *s }

func UUIDtoStrFmtUUID(s *uuid.UUID) *strfmt.UUID {
	return UUID(strfmt.UUID(s.String()))
}

func NewUUID() *strfmt.UUID {
	id, _ := uuid.NewV4()
	return UUIDtoStrFmtUUID(id)
}

func NewUUIDValue() strfmt.UUID {
	id, _ := uuid.NewV4()
	return *UUIDtoStrFmtUUID(id)
}

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// StringValue returns the value of the string pointer passed in or
// "" if the pointer is nil.
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}
