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
	From  *gen.Field
	To    *ir.Field
	Enums []EnumMapping // only for enum fields
}

type EnumMapping struct {
	From gen.Enum
	To   *ir.EnumVariant
}

type EdgeMapping struct {
	From *gen.Edge
	To   *ir.Field
}

func (e *Extension) createMapping(from *gen.Type, to *ir.Type) error {
	if _, ok := e.recur[from]; ok {
		return nil
	}

	e.recur[from] = struct{}{}
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

		f, ok, err := m.lookupField(field.Spec.Name)
		if err != nil {
			return fmt.Errorf("lookup for field with name %q: %w", field.Spec.Name, err)
		}

		if !ok {
			e, ok, err := m.lookupEdge(field.Spec.Name)
			if err != nil {
				return fmt.Errorf("lookup for edge with name %q: %w", field.Spec.Name, err)
			}

			if !ok {
				return fmt.Errorf("property %q not found in ent schema", field.Spec.Name)
			}

			if err := m.createEdgeMapping(e, field); err != nil {
				return fmt.Errorf("edge %q: %w", e.Name, err)
			}

			continue
		}

		if err := m.createFieldMapping(f, field); err != nil {
			return fmt.Errorf("field %q: %w", f.Name, err)
		}
	}

	return nil
}

func (m *Mapping) lookupField(name string) (*gen.Field, bool, error) {
	var matches []*gen.Field
	for _, f := range m.EntFields() {
		ant, err := annotation(f.Annotations)
		if err != nil {
			return nil, false, fmt.Errorf("read field %q annotation: %w", f.Name, err)
		}

		if ant != nil && ant.BindTo == name {
			matches = append(matches, f)
			continue
		}

		if f.Name == name {
			matches = append(matches, f)
		}
	}

	switch {
	case len(matches) == 0:
		return nil, false, nil

	case len(matches) == 1:
		return matches[0], true, nil

	default:
		var names []string
		for _, m := range matches {
			names = append(names, m.Name)
		}
		return nil, false, fmt.Errorf("matched multiple fields: %v", names)
	}
}

func (m *Mapping) createFieldMapping(entField *gen.Field, ogenField *ir.Field) error {
	if ogenField.Spec == nil {
		return nil
	}

	ogenSchema := ogenField.Spec.Schema

	if entField.Optional && !entField.Nillable {
		return fmt.Errorf("optional fields are not supported")
	}

	if entField.Nillable != ogenSchema.Nullable {
		return fmt.Errorf("nullability mismatch")
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
		field.TypeEnum:   {jsonschema.String, ""},
	}

	v, ok := mapping[entField.Type.Type]
	if !ok {
		return fmt.Errorf("unsupported ent type: %s", entField.Type.ConstName())
	}

	if ogenSchema.Type != v.Type {
		return fmt.Errorf("type mismatch: expected %q but have %q", v.Type, ogenSchema.Type)
	}

	if ogenSchema.Format != v.Format {
		return fmt.Errorf("type format mismatch: expected %q but have %q", v.Format, ogenSchema.Format)
	}

	var enumMappings []EnumMapping
	if entField.Type.Type == field.TypeEnum {
		if len(ogenSchema.Enum) != len(entField.EnumValues()) {
			return fmt.Errorf("enum mismatch")
		}

		dbEnums := map[string]gen.Enum{}
		for _, enum := range entField.Enums {
			dbEnums[enum.Value] = enum
		}

		typ := ogenField.Type
		if typ.IsGeneric() { // Nullable enums.
			typ = typ.GenericOf
		}

		for _, ogenEnum := range typ.EnumVariants {
			val, ok := ogenEnum.Value.(string)
			if !ok {
				return fmt.Errorf("unexpected enum value type: %T", ogenEnum.Value)
			}

			dbEnum, ok := dbEnums[val]
			if !ok {
				return fmt.Errorf("enum value %q not found in ent schema", val)
			}

			enumMappings = append(enumMappings, EnumMapping{
				From: dbEnum,
				To:   ogenEnum,
			})
		}
	}

	m.FieldMappings = append(m.FieldMappings, FieldMapping{
		From:  entField,
		To:    ogenField,
		Enums: enumMappings,
	})

	return nil
}

func (m *Mapping) lookupEdge(name string) (*gen.Edge, bool, error) {
	var matches []*gen.Edge
	for _, e := range m.From.Edges {
		ant, err := annotation(e.Annotations)
		if err != nil {
			return nil, false, fmt.Errorf("read edge %q annotation: %w", e.Name, err)
		}

		if ant != nil && ant.BindTo == name {
			matches = append(matches, e)
			continue
		}

		if e.Name == name {
			matches = append(matches, e)
		}
	}

	switch {
	case len(matches) == 0:
		return nil, false, nil

	case len(matches) == 1:
		return matches[0], true, nil

	default:
		var names []string
		for _, m := range matches {
			names = append(names, m.Name)
		}
		return nil, false, fmt.Errorf("matched multiple edges: %v", names)
	}
}

func (m *Mapping) createEdgeMapping(edge *gen.Edge, field *ir.Field) error {
	if field.Spec == nil || field.Spec.Schema == nil {
		return fmt.Errorf("spec cannot be nil")
	}

	typ := field.Type
	switch {
	case edge.Optional && edge.Unique: // Single optional type.
		if !typ.IsGeneric() {
			return fmt.Errorf("edge is optional, generic type is expected, not %q", typ.Kind)
		}

		typ = typ.GenericOf

	case edge.Optional && !edge.Unique: // Multiple optional types.
		if !typ.IsArray() {
			return fmt.Errorf("edge is not unique, schema must be an array, not %q", typ.Kind)
		}

		typ = typ.Item

	case !edge.Optional && edge.Unique: // Required unique type.
		// Use field type.

	case !edge.Optional && !edge.Unique: // Required multiple types.
		if !typ.IsArray() {
			return fmt.Errorf("edge is not unique, schema must be an array, not %q", typ.Kind)
		}

		typ = typ.Item

	default:
		panic("unreachable")
	}

	if err := m.ext.createMapping(edge.Type, typ); err != nil {
		return err
	}

	m.EdgeMappings = append(m.EdgeMappings, EdgeMapping{
		From: edge,
		To:   field,
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

func assign(dst *ir.Field, src *gen.Field) (string, error) {
	var (
		assignT = "t." + dst.Name
		srcT    = "e." + src.StructField()
	)

	if src.Nillable {
		srcT = "*" + srcT
	}

	if dst.Type.IsGeneric() {
		if dst.Type.GenericOf.IsPrimitive() {
			return assignT + ".SetTo(" + srcT + ")", nil
		}

		gotyp := dst.Type.GenericOf.Go()
		return fmt.Sprintf("%s.SetTo(openapi.%s(%s))", assignT, gotyp, srcT), nil
	}

	if dst.Type.IsEnum() {
		gotyp := dst.Type.Go()
		return fmt.Sprintf("%s = openapi.%s(%s)", assignT, gotyp, srcT), nil
	}

	return fmt.Sprintf("%s = %s", assignT, srcT), nil
}
