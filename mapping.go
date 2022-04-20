package ent2ogen

import (
	"fmt"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/jsonschema"
)

// Mapping is used to render templates.
type Mapping struct {
	From          *gen.Type
	To            *ir.Type
	FieldMappings []FieldMapping
	EdgeMappings  []EdgeMapping

	ext *Extension
}

type FieldMapping struct {
	From *gen.Field
	To   *ir.Field
}

type EdgeMapping struct {
	From *gen.Edge
	To   *ir.Field
}

func (e *Extension) createMapping(from *gen.Type, to *ir.Type) error {
	if _, ok := e.cfg.Mappings[from]; ok {
		return fmt.Errorf("mapping already created")
	}

	m := &Mapping{
		From: from,
		To:   to,
		ext:  e,
	}

	if err := m.checkCompatibility(); err != nil {
		return err
	}

	e.cfg.Mappings[from] = m
	return nil
}

func (m *Mapping) checkCompatibility() error {
	if m.To.Kind != ir.KindStruct {
		return fmt.Errorf("schema must be an object")
	}

	for _, field := range m.To.Fields {
		if field.Spec == nil {
			return fmt.Errorf("field %q has no spec", field.Name)
		}

		f, ok := m.lookupField(field.Spec.Name)
		if !ok {
			e, ok := m.lookupEdge(field.Spec.Name)
			if !ok {
				return fmt.Errorf("property %q not found in ent schema", field.Spec.Name)
			}

			if err := m.createEdgeMapping(e, field); err != nil {
				return fmt.Errorf("edge %q: %w", f.Name, err)
			}

			continue
		}

		if err := m.createFieldMapping(f, field); err != nil {
			return fmt.Errorf("field %q: %w", f.Name, err)
		}
	}

	return nil
}

func (m *Mapping) lookupField(name string) (*gen.Field, bool) {
	for _, f := range m.EntFields() {
		if f.Name == name {
			return f, true
		}
	}

	return nil, false
}

func (m *Mapping) createFieldMapping(entField *gen.Field, ogenField *ir.Field) error {
	if ogenField.Spec == nil {
		return nil
	}

	ogenSchema := ogenField.Spec.Schema

	if entField.Optional != !ogenField.Spec.Required {
		return fmt.Errorf("optionality mismatch")
	}

	if entField.Nillable != ogenSchema.Nullable {
		return fmt.Errorf("nullability mismatch")
	}

	if entField.Optional && !entField.Nillable {
		return fmt.Errorf("optional fields are not supported")
	}

	type tf struct {
		Type   jsonschema.SchemaType
		Format string
	}

	mapping := map[field.Type]tf{
		field.TypeBool:   {jsonschema.Boolean, ""},
		field.TypeString: {jsonschema.String, ""},
		field.TypeInt:    {jsonschema.Integer, "int32"},
		field.TypeInt64:  {jsonschema.Integer, "int64"},
		field.TypeTime:   {jsonschema.String, "date-time"},
		field.TypeUUID:   {jsonschema.String, "uuid"},
	}

	v, ok := mapping[entField.Type.Type]
	if !ok {
		return fmt.Errorf("unsupported ent type: %q", entField.Type.Type)
	}

	if ogenSchema.Type != v.Type {
		return fmt.Errorf("type mismatch: expected %q but have %q", v.Type, ogenSchema.Type)
	}

	if ogenSchema.Format != v.Format {
		return fmt.Errorf("type format mismatch: expected %q but have %q", v.Format, ogenSchema.Format)
	}

	if entField.Enums != nil {
		return fmt.Errorf("enum is not supported")
	}

	m.FieldMappings = append(m.FieldMappings, FieldMapping{
		From: entField,
		To:   ogenField,
	})

	return nil
}

func (m *Mapping) lookupEdge(name string) (*gen.Edge, bool) {
	for _, e := range m.From.Edges {
		if e.Name == name {
			return e, true
		}
	}

	return nil, false
}

func (m *Mapping) createEdgeMapping(from *gen.Edge, to *ir.Field) error {
	if to.Spec == nil || to.Spec.Schema == nil {
		return fmt.Errorf("spec cannot be nil")
	}

	if from.Optional != false {
		return fmt.Errorf("optional edges are not supported")
	}

	if !from.Unique {
		return fmt.Errorf("only unique edges are supported")
	}

	if err := m.ext.createMapping(from.Type, to.Type); err != nil {
		return fmt.Errorf("edge %q: %w", from.Name, err)
	}

	m.EdgeMappings = append(m.EdgeMappings, EdgeMapping{
		From: from,
		To:   to,
	})

	return nil
}

// EntFields returns ent schema fields.
func (m *Mapping) EntFields() []*gen.Field {
	fields := make([]*gen.Field, 0, len(m.From.Fields)+1)
	fields = append(fields, m.From.ID)
	fields = append(fields, m.From.Fields...)
	return fields
}
