package ent2ogen

import (
	"fmt"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"entgo.io/ent/schema/field"
	"github.com/iancoleman/strcase"
	"github.com/ogen-go/ogen"
)

type Mapping struct {
	From     *load.Schema // Ent schema
	FromType *gen.Type    // Ent go type
	To       *ogen.Schema // OpenAPI schema
}

func (m *Mapping) checkConvertability() error {
	if m.To.Type != "object" {
		return fmt.Errorf("schema must be an object")
	}

	searchField := func(name string) (*ogen.Schema, bool) {
		for _, prop := range m.To.Properties {
			if prop.Name == name {
				return prop.Schema, true
			}
		}
		return nil, false
	}

	required := func(name string) bool {
		for _, f := range m.To.Required {
			if f == name {
				return true
			}
		}
		return false
	}

	for _, field := range m.From.Fields {
		s, ok := searchField(field.Name)
		if !ok {
			fmt.Printf("type %q: field %q not found in schema object\n", m.From.Name, field.Name)
			continue
		}

		if err := m.checkField(field, required(field.Name), s); err != nil {
			return fmt.Errorf("field %q: %w", field.Name, err)
		}
	}

	return nil
}

func (m *Mapping) checkField(f *load.Field, required bool, s *ogen.Schema) error {
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

	v, ok := mapping[f.Info.Type]
	if !ok {
		return fmt.Errorf("unsupported ent type: %q", f.Info.Type)
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
	name := strcase.ToSnake(f.Name)

	for _, prop := range m.To.Properties {
		if prop.Name == name {
			return true
		}
	}

	return false
}
