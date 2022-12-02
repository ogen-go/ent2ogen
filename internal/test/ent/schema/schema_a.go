package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen"
)

type SchemaA struct {
	ent.Schema
}

func (SchemaA) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("int64"),
		field.String("string_bindto_foobar").Annotations(ent2ogen.BindTo("string_foobar_bind")),
		field.String("string_optional_nullable").Optional().Nillable(),
		field.Bool("optional_nullable_bool").Optional().Nillable(),
		field.Strings("jsontype_strings"),
		field.Strings("jsontype_strings_optional").Optional(),
		field.Enum("required_enum").Values("a", "b"),
		field.Enum("optional_nullable_enum").Values("c", "d").Optional().Nillable(),
	}
}

func (SchemaA) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("edge_schemab_unique_required", SchemaB.Type).Unique().Required(),
		edge.To("edge_schemab_unique_required_bindto_bs", SchemaB.Type).
			Unique().
			Required().
			Annotations(ent2ogen.BindTo("edge_schemab_unique_required_bs_bind")),

		edge.To("edge_schemab_unique_optional", SchemaB.Type).Unique(),
		edge.To("edge_schemab", SchemaB.Type),

		edge.To("edge_schemaa_recursive", SchemaA.Type),
	}
}

func (SchemaA) Annotations() []schema.Annotation {
	return []schema.Annotation{
		ent2ogen.BindTo("SchemaA"),
	}
}
