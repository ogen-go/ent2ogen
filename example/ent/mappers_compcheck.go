// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/google/uuid"
	"github.com/ogen-go/ent2ogen/example/ent/city"
	openapi "github.com/ogen-go/ent2ogen/example/openapi"
)

func _() {
	_ = struct {
		Name         string
		RequiredEnum openapi.CityRequiredEnum
		NullableEnum openapi.NilCityNullableEnum
	}(openapi.City{})
	_ = map[bool]struct{}{
		string(openapi.CityRequiredEnumA) == string(city.RequiredEnumA): {},
		false: {},
	}
	_ = map[bool]struct{}{
		string(openapi.CityRequiredEnumB) == string(city.RequiredEnumB): {},
		false: {},
	}
	_ = map[bool]struct{}{
		string(openapi.CityNullableEnumC) == string(city.NullableEnumC): {},
		false: {},
	}
	_ = map[bool]struct{}{
		string(openapi.CityNullableEnumD) == string(city.NullableEnumD): {},
		false: {},
	}
}

func _() {
	_ = struct {
		ID                   uuid.UUID
		FirstName            string
		LastName             string
		Username             string
		OptionalNullableBool openapi.OptNilBool
		City                 openapi.City
		Friends              []openapi.User
	}(openapi.User{})
}
