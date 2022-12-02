// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"

	openapi "github.com/ogen-go/ent2ogen/internal/test/api"
)

type SchemaASlice []*SchemaA

// Following edges must be loaded:
//
//	edge_schemab_unique_required
//	edge_schemab_unique_required_bindto_bs
//	edge_schemab_unique_optional
//	edge_schemab
//	edge_schemaa_recursive:
//	  edge_schemab_unique_required
//	  edge_schemab_unique_required_bindto_bs
//	  edge_schemab_unique_optional
//	  edge_schemab
//	  edge_schemaa_recursive...
func (s SchemaASlice) ToOpenAPI() ([]openapi.SchemaA, error) {
	return s.toOpenAPI()
}

func (s SchemaASlice) toOpenAPI() (_ []openapi.SchemaA, err error) {
	result := make([]openapi.SchemaA, len(s))
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
//	edge_schemab_unique_required
//	edge_schemab_unique_required_bindto_bs
//	edge_schemab_unique_optional
//	edge_schemab
//	edge_schemaa_recursive:
//	  edge_schemab_unique_required
//	  edge_schemab_unique_required_bindto_bs
//	  edge_schemab_unique_optional
//	  edge_schemab
//	  edge_schemaa_recursive...
func (e *SchemaA) ToOpenAPI() (*openapi.SchemaA, error) {
	t, err := e.toOpenAPI()
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (e *SchemaA) toOpenAPI() (t openapi.SchemaA, err error) {
	t.Int64 = e.Int64
	t.StringFoobarBind = e.StringBindtoFoobar
	if e.StringOptionalNullable != nil {
		t.StringOptionalNullable.SetTo(*e.StringOptionalNullable)
	} else {
		t.StringOptionalNullable.Null = true
	}
	if e.OptionalNullableBool != nil {
		t.OptionalNullableBool.SetTo(*e.OptionalNullableBool)
	} else {
		t.OptionalNullableBool.Set, t.OptionalNullableBool.Null = false, true
	}
	t.JsontypeStrings = e.JsontypeStrings
	t.JsontypeStringsOptional = e.JsontypeStringsOptional
	t.RequiredEnum = openapi.SchemaARequiredEnum(e.RequiredEnum)
	if e.OptionalNullableEnum != nil {
		t.OptionalNullableEnum.SetTo(openapi.SchemaAOptionalNullableEnum(*e.OptionalNullableEnum))
	} else {
		t.OptionalNullableEnum.Set, t.OptionalNullableEnum.Null = false, true
	}
	// Edge 'edge_schemab_unique_required'.
	if err := func() error {
		v, err := e.Edges.EdgeSchemabUniqueRequiredOrErr()
		if err != nil {
			return fmt.Errorf("load: %w", err)
		}
		openapiType, err := v.toOpenAPI()
		if err != nil {
			return fmt.Errorf("convert to openapi: %w", err)
		}
		t.EdgeSchemabUniqueRequired = openapiType
		return nil
	}(); err != nil {
		return t, fmt.Errorf("edge 'edge_schemab_unique_required': %w", err)
	}
	// Edge 'edge_schemab_unique_required_bindto_bs'.
	if err := func() error {
		v, err := e.Edges.EdgeSchemabUniqueRequiredBindtoBsOrErr()
		if err != nil {
			return fmt.Errorf("load: %w", err)
		}
		openapiType, err := v.toOpenAPI()
		if err != nil {
			return fmt.Errorf("convert to openapi: %w", err)
		}
		t.EdgeSchemabUniqueRequiredBsBind = openapiType
		return nil
	}(); err != nil {
		return t, fmt.Errorf("edge 'edge_schemab_unique_required_bindto_bs': %w", err)
	}
	// Edge 'edge_schemab_unique_optional'.
	if err := func() error {
		v, err := e.Edges.EdgeSchemabUniqueOptionalOrErr()
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
		t.EdgeSchemabUniqueOptional.SetTo(openapiType)
		return nil
	}(); err != nil {
		return t, fmt.Errorf("edge 'edge_schemab_unique_optional': %w", err)
	}
	// Edge 'edge_schemab'.
	if err := func() error {
		v, err := e.Edges.EdgeSchemabOrErr()
		if err != nil {
			if IsNotFound(err) {
				return nil
			}
			return fmt.Errorf("load: %w", err)
		}
		openapiType, err := SchemaBSlice(v).toOpenAPI()
		if err != nil {
			return fmt.Errorf("convert to openapi: %w", err)
		}
		t.EdgeSchemab = openapiType
		return nil
	}(); err != nil {
		return t, fmt.Errorf("edge 'edge_schemab': %w", err)
	}
	// Edge 'edge_schemaa_recursive'.
	if err := func() error {
		v, err := e.Edges.EdgeSchemaaRecursiveOrErr()
		if err != nil {
			if IsNotFound(err) {
				return nil
			}
			return fmt.Errorf("load: %w", err)
		}
		openapiType, err := SchemaASlice(v).toOpenAPI()
		if err != nil {
			return fmt.Errorf("convert to openapi: %w", err)
		}
		t.EdgeSchemaaRecursive = openapiType
		return nil
	}(); err != nil {
		return t, fmt.Errorf("edge 'edge_schemaa_recursive': %w", err)
	}
	return t, nil
}

type SchemaBSlice []*SchemaB

func (s SchemaBSlice) ToOpenAPI() ([]openapi.SchemaB, error) {
	return s.toOpenAPI()
}

func (s SchemaBSlice) toOpenAPI() (_ []openapi.SchemaB, err error) {
	result := make([]openapi.SchemaB, len(s))
	for i, v := range s {
		result[i], err = v.toOpenAPI()
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (e *SchemaB) ToOpenAPI() (*openapi.SchemaB, error) {
	t, err := e.toOpenAPI()
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (e *SchemaB) toOpenAPI() (t openapi.SchemaB, err error) {
	t.ID = e.ID
	return t, nil
}
