package ent2ogen

import (
	"fmt"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ogen"
)

type Mapping struct {
	From *gen.Type    // Ent go type
	To   *ogen.Schema // OpenAPI schema

	fromSchema *load.Schema
}

func (m *Mapping) checkConvertability() error {
	if m.To.Type != "object" {
		return fmt.Errorf("schema must be an object")
	}

	fields := m.EntFields()
	lookupEntField := func(name string) (*gen.Field, bool) {
		for _, f := range fields {
			if f.Name == name {
				return f, true
			}
		}
		return nil, false
	}

	propRequired := func(name string) bool {
		for _, f := range m.To.Required {
			if f == name {
				return true
			}
		}
		return false
	}

	for _, prop := range m.To.Properties {
		f, ok := lookupEntField(prop.Name)
		if !ok {
			return fmt.Errorf("property %q not found in ent schema", prop.Name)
		}

		if err := m.checkField(f, propRequired(prop.Name), prop.Schema); err != nil {
			return fmt.Errorf("property %q: %w", f.Name, err)
		}
	}

	return nil
}

func (m *Mapping) checkField(f *gen.Field, required bool, s *ogen.Schema) error {
	if f.Optional != !required {
		return fmt.Errorf("optionality mismatch")
	}

	if f.Nillable != s.Nullable {
		return fmt.Errorf("nullability mismatch")
	}

	type tf struct {
		Type   string
		Format string
	}

	mapping := map[field.Type]tf{
		field.TypeBool:   {"bool", ""},
		field.TypeString: {"string", ""},
		field.TypeInt:    {"integer", "int32"},
		field.TypeInt64:  {"integer", "int64"},
		field.TypeTime:   {"string", "date-time"},
		field.TypeUUID:   {"string", "uuid"},
	}

	v, ok := mapping[f.Type.Type]
	if !ok {
		return fmt.Errorf("unsupported ent type: %q", f.Type.Type)
	}

	if s.Type != v.Type {
		return fmt.Errorf("type mismatch: expected %q but have %q", v.Type, s.Type)
	}

	if s.Format != v.Format {
		return fmt.Errorf("type format mismatch: expected %q but have %q", v.Format, s.Format)
	}

	if f.Enums != nil {
		return fmt.Errorf("enum is not supported")
	}

	return nil
}

func (m *Mapping) HasOpenAPIField(f *gen.Field) bool {
	for _, prop := range m.To.Properties {
		if prop.Name == f.Name {
			return true
		}
	}

	return false
}

// EntFields returns ent schema fields.
func (m *Mapping) EntFields() []*gen.Field {
	fields := make([]*gen.Field, 0, len(m.From.Fields)+1)
	fields = append(fields, m.From.ID)
	fields = append(fields, m.From.Fields...)
	return fields
}
