package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type City struct {
	ent.Schema
}

func (City) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),

		field.String("name").NotEmpty(),

		field.Enum("required_enum").Values(
			"a",
			"b",
		),

		field.Enum("nullable_enum").Values(
			"c",
			"d",
		).Optional().Nillable(),
	}
}

func (City) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// ent2ogen.BindTo("City"),
	}
}
