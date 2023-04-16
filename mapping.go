package ent2ogen

import (
	"log"
	"strings"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"github.com/go-faster/errors"
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
		return errors.New("schema must be an object")
	}

	for _, field := range m.To.Fields {
		if field.Spec == nil {
			return errors.Errorf("field %q has no spec", field.Name)
		}

		f, ok, err := m.lookupField(field.Spec.Name)
		if err != nil {
			return errors.Wrapf(err, "lookup for field with name %q", field.Spec.Name)
		}

		if !ok {
			e, ok, err := m.lookupEdge(field.Spec.Name)
			if err != nil {
				return errors.Wrapf(err, "lookup for edge with name %q", field.Spec.Name)
			}

			if !ok {
				return errors.Errorf("property %q not found in ent schema", field.Spec.Name)
			}

			if err := m.createEdgeMapping(e, field); err != nil {
				return errors.Wrapf(err, "edge %q", e.Name)
			}

			continue
		}

		if err := m.createFieldMapping(f, field); err != nil {
			return errors.Wrapf(err, "field %q", f.Name)
		}
	}

	return nil
}

func (m *Mapping) lookupField(name string) (*gen.Field, bool, error) {
	var matches []*gen.Field
	for _, f := range m.EntFields() {
		ant, err := annotation(f.Annotations)
		if err != nil {
			return nil, false, errors.Wrapf(err, "read field %q annotation", f.Name)
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
		names := make([]string, 0, len(matches))
		for _, m := range matches {
			names = append(names, m.Name)
		}
		return nil, false, errors.Errorf("matched multiple fields: %v", names)
	}
}

func (m *Mapping) createFieldMapping(entField *gen.Field, ogenField *ir.Field) error {
	if ogenField.Spec == nil {
		return nil
	}

	var (
		et = entField.Type
		js = ogenField.Spec.Schema
	)

	assertJS := func(js *jsonschema.Schema, typ jsonschema.SchemaType, format string) error {
		if js == nil {
			return errors.New("unexpected null or empty schema")
		}
		if js.Type != typ {
			return errors.Errorf("type mismatch: expected %q but have %q", typ, js.Type)
		}
		if js.Format != format {
			if format == "" {
				return errors.Errorf("unexpected format %q", js.Format)
			}
			return errors.Errorf("format mismatch: expected %q but have %q", format, js.Format)
		}
		return nil
	}

	switch {
	case entField.Optional && !entField.Nillable:
		switch {
		case et.Type == field.TypeJSON && et.Ident == "[]string":
		case et.Type == field.TypeJSON && et.Ident == "[]int":
		default:
			return errors.New("optional ent fields are not supported - you need to make the field either optional and nullable or required")
		}

	case entField.Optional && entField.Nillable:
		if !ogenField.Type.IsGeneric() {
			return errors.Errorf("ent field is optional, ogen type must be generic, not %q", ogenField.Type.Kind)
		}

	case !entField.Optional:
		if ogenField.Type.IsGeneric() {
			return errors.New("openapi field must be required")
		}

	default:
		panic("unreachable")
	}

	checks := map[field.Type]func() error{
		field.TypeBool: func() error { return assertJS(js, jsonschema.Boolean, "") },

		field.TypeString: func() error { return assertJS(js, jsonschema.String, "") },
		field.TypeTime:   func() error { return assertJS(js, jsonschema.String, "date-time") },
		field.TypeUUID:   func() error { return assertJS(js, jsonschema.String, "uuid") },
		field.TypeEnum:   func() error { return assertJS(js, jsonschema.String, "") },
		field.TypeBytes:  func() error { return assertJS(js, jsonschema.String, "byte") },

		field.TypeInt:    func() error { return assertJS(js, jsonschema.Integer, "") },
		field.TypeInt16:  func() error { return assertJS(js, jsonschema.Integer, "int16") },
		field.TypeInt32:  func() error { return assertJS(js, jsonschema.Integer, "int32") },
		field.TypeInt64:  func() error { return assertJS(js, jsonschema.Integer, "int64") },
		field.TypeUint:   func() error { return assertJS(js, jsonschema.Integer, "uint") },
		field.TypeUint16: func() error { return assertJS(js, jsonschema.Integer, "uint16") },
		field.TypeUint32: func() error { return assertJS(js, jsonschema.Integer, "uint32") },
		field.TypeUint64: func() error { return assertJS(js, jsonschema.Integer, "uint64") },

		field.TypeJSON: func() error {
			switch et.Ident {
			case "[]string":
				if err := assertJS(js, jsonschema.Array, ""); err != nil {
					return err
				}
				if err := assertJS(js.Item, jsonschema.String, ""); err != nil {
					return errors.Wrap(err, "items")
				}
			case "[]int":
				if err := assertJS(js, jsonschema.Array, ""); err != nil {
					return err
				}
				if err := assertJS(js.Item, jsonschema.Integer, ""); err != nil {
					return errors.Wrap(err, "items")
				}
			default:
				return errors.Errorf("unsupported ent json type: %s", et.Ident)
			}

			return nil
		},
	}

	if f, ok := checks[et.Type]; ok {
		if err := f(); err != nil {
			return err
		}
	} else {
		return errors.Errorf("unsupported ent type: %s", entField.Type.ConstName())
	}

	fm := FieldMapping{
		From: entField,
		To:   ogenField,
	}

	if entField.Type.Type == field.TypeEnum {
		enumMappings, err := m.createEnumMappings(entField, ogenField)
		if err != nil {
			return errors.Wrap(err, "enum")
		}

		fm.Enums = enumMappings
	}

	m.FieldMappings = append(m.FieldMappings, fm)
	return nil
}

func (m *Mapping) createEnumMappings(ef *gen.Field, gf *ir.Field) ([]EnumMapping, error) {
	var (
		enumMappings []EnumMapping

		et = ef.Type
		js = gf.Spec.Schema
	)

	if et.Type != field.TypeEnum || len(js.Enum) == 0 {
		return nil, errors.New("bad schema")
	}

	if len(js.Enum) != len(ef.EnumValues()) {
		return nil, errors.New("enum mismatch")
	}

	dbEnums := make(map[string]gen.Enum, len(ef.Enums))
	for _, enum := range ef.Enums {
		dbEnums[enum.Value] = enum
	}

	typ := gf.Type
	if typ.IsGeneric() { // Nullable enums.
		typ = typ.GenericOf
	}

	for _, ogenEnum := range typ.EnumVariants {
		val, ok := ogenEnum.Value.(string)
		if !ok {
			return nil, errors.Errorf("unexpected enum value type: %T", ogenEnum.Value)
		}

		dbEnum, ok := dbEnums[val]
		if !ok {
			return nil, errors.Errorf("enum value %q not found in ent schema", val)
		}

		enumMappings = append(enumMappings, EnumMapping{
			From: dbEnum,
			To:   ogenEnum,
		})
	}

	return enumMappings, nil
}

func (m *Mapping) lookupEdge(name string) (*gen.Edge, bool, error) {
	var matches []*gen.Edge
	for _, e := range m.From.Edges {
		ant, err := annotation(e.Annotations)
		if err != nil {
			return nil, false, errors.Wrapf(err, "read edge %q annotation", e.Name)
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
		names := make([]string, 0, len(matches))
		for _, m := range matches {
			names = append(names, m.Name)
		}
		return nil, false, errors.Errorf("matched multiple edges: %v", names)
	}
}

func (m *Mapping) createEdgeMapping(edge *gen.Edge, field *ir.Field) error {
	if field.Spec == nil || field.Spec.Schema == nil {
		return errors.New("spec cannot be nil")
	}

	typ := field.Type
	switch {
	case edge.Optional && edge.Unique: // Single optional type.
		if !typ.IsGeneric() {
			return errors.Errorf("edge is optional, generic type is expected, not %q", typ.Kind)
		}

		typ = typ.GenericOf

	case edge.Optional && !edge.Unique: // Multiple optional types.
		if !typ.IsArray() {
			return errors.Errorf("edge is not unique, schema must be an array, not %q", typ.Kind)
		}

		typ = typ.Item

	case !edge.Optional && edge.Unique: // Required unique type.
		// Use field type.

	case !edge.Optional && !edge.Unique: // Required multiple types.
		if !typ.IsArray() {
			return errors.Errorf("edge is not unique, schema must be an array, not %q", typ.Kind)
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

func (m *Mapping) Comment() string {
	if len(m.EdgeMappings) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("// Following edges must be loaded:\n")
	m.comment(2, map[*gen.Type]struct{}{}, &b)
	return strings.TrimSpace(b.String())
}

func (m *Mapping) comment(indent int, walk map[*gen.Type]struct{}, b *strings.Builder) {
	wr := func(s string) {
		space := strings.Repeat(" ", indent)
		b.WriteString("// " + space + s + "\n")
	}

	for _, e := range m.EdgeMappings {
		tm, ok := m.ext.cfg.Mappings[e.From.Type]
		if !ok {
			panic("unreachable")
		}

		if len(tm.EdgeMappings) == 0 {
			wr(e.From.Name)
			continue
		}

		if _, ok := walk[e.From.Type]; ok {
			if !e.From.Optional {
				log.Fatalf("type %q edge %q infinite recursion", m.From.Name, e.From.Name)
			}

			wr(e.From.Name + "...")
			continue
		}

		func() {
			walk[e.From.Type] = struct{}{}
			defer func() { delete(walk, e.From.Type) }()

			wr(e.From.Name + ":")
			tm.comment(indent+2, walk, b)
		}()
	}
}
