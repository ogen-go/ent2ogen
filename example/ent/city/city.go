// Code generated by entc, DO NOT EDIT.

package city

import (
	"fmt"
)

const (
	// Label holds the string label denoting the city type in the database.
	Label = "city"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldRequiredEnum holds the string denoting the required_enum field in the database.
	FieldRequiredEnum = "required_enum"
	// FieldNullableEnum holds the string denoting the nullable_enum field in the database.
	FieldNullableEnum = "nullable_enum"
	// Table holds the table name of the city in the database.
	Table = "cities"
)

// Columns holds all SQL columns for city fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldRequiredEnum,
	FieldNullableEnum,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// RequiredEnum defines the type for the "required_enum" enum field.
type RequiredEnum string

// RequiredEnum values.
const (
	RequiredEnumA RequiredEnum = "a"
	RequiredEnumB RequiredEnum = "b"
)

func (re RequiredEnum) String() string {
	return string(re)
}

// RequiredEnumValidator is a validator for the "required_enum" field enum values. It is called by the builders before save.
func RequiredEnumValidator(re RequiredEnum) error {
	switch re {
	case RequiredEnumA, RequiredEnumB:
		return nil
	default:
		return fmt.Errorf("city: invalid enum value for required_enum field: %q", re)
	}
}

// NullableEnum defines the type for the "nullable_enum" enum field.
type NullableEnum string

// NullableEnum values.
const (
	NullableEnumC NullableEnum = "c"
	NullableEnumD NullableEnum = "d"
)

func (ne NullableEnum) String() string {
	return string(ne)
}

// NullableEnumValidator is a validator for the "nullable_enum" field enum values. It is called by the builders before save.
func NullableEnumValidator(ne NullableEnum) error {
	switch ne {
	case NullableEnumC, NullableEnumD:
		return nil
	default:
		return fmt.Errorf("city: invalid enum value for nullable_enum field: %q", ne)
	}
}
