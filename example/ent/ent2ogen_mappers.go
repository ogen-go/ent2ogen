// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"

	openapi "github.com/ogen-go/ent2ogen/example/openapi"
)

type CitySlice []*City

func (s CitySlice) ToOpenAPI() ([]openapi.City, error) {
	return s.toOpenAPI()
}

func (s CitySlice) toOpenAPI() (_ []openapi.City, err error) {
	result := make([]openapi.City, len(s))
	for i, v := range s {
		result[i], err = v.toOpenAPI()
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (e *City) ToOpenAPI() (*openapi.City, error) {
	t, err := e.toOpenAPI()
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (e *City) toOpenAPI() (t openapi.City, err error) {
	t.Name = e.Name
	t.RequiredEnum = openapi.CityRequiredEnum(e.RequiredEnum)
	if e.NullableEnum != nil {
		t.NullableEnum.SetTo(openapi.CityNullableEnum(*e.NullableEnum))
	} else {
		t.NullableEnum.Null = true
	}
	return t, nil
}

type UserSlice []*User

// Following edges must be loaded:
//
//	required_city
//	optional_city
//	friend_list:
//	  required_city
//	  optional_city
//	  friend_list...
func (s UserSlice) ToOpenAPI() ([]openapi.User, error) {
	return s.toOpenAPI()
}

func (s UserSlice) toOpenAPI() (_ []openapi.User, err error) {
	result := make([]openapi.User, len(s))
	for i, v := range s {
		result[i], err = v.toOpenAPI()
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Following edges must be loaded:
//
//	required_city
//	optional_city
//	friend_list:
//	  required_city
//	  optional_city
//	  friend_list...
func (e *User) ToOpenAPI() (*openapi.User, error) {
	t, err := e.toOpenAPI()
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (e *User) toOpenAPI() (t openapi.User, err error) {
	t.ID = e.ID
	t.FirstName = e.FirstName
	t.LastName = e.LastName
	t.Username = e.UserName
	if e.OptionalNullableBool != nil {
		t.OptionalNullableBool.SetTo(*e.OptionalNullableBool)
	} else {
		t.OptionalNullableBool.Set, t.OptionalNullableBool.Null = false, true
	}
	// Edge 'required_city'.
	if err := func() error {
		v, err := e.Edges.RequiredCityOrErr()
		if err != nil {
			return fmt.Errorf("load: %w", err)
		}
		openapiType, err := v.toOpenAPI()
		if err != nil {
			return fmt.Errorf("convert to openapi: %w", err)
		}
		t.RequiredCity = openapiType
		return nil
	}(); err != nil {
		return t, fmt.Errorf("edge 'required_city': %w", err)
	}
	// Edge 'optional_city'.
	if err := func() error {
		v, err := e.Edges.OptionalCityOrErr()
		if err != nil {
			if IsNotFound(err) {
				return nil
			}
			return fmt.Errorf("load: %w", err)
		}
		openapiType, err := v.toOpenAPI()
		if err != nil {
			return fmt.Errorf("convert to openapi: %w", err)
		}
		t.OptionalCity.SetTo(openapiType)
		return nil
	}(); err != nil {
		return t, fmt.Errorf("edge 'optional_city': %w", err)
	}
	// Edge 'friend_list'.
	if err := func() error {
		v, err := e.Edges.FriendListOrErr()
		if err != nil {
			if IsNotFound(err) {
				return nil
			}
			return fmt.Errorf("load: %w", err)
		}
		openapiType, err := UserSlice(v).toOpenAPI()
		if err != nil {
			return fmt.Errorf("convert to openapi: %w", err)
		}
		t.Friends = openapiType
		return nil
	}(); err != nil {
		return t, fmt.Errorf("edge 'friend_list': %w", err)
	}
	return t, nil
}