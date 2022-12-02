package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen"
)

type SchemaB struct {
	ent.Schema
}

func (SchemaB) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
	}
}

func (SchemaB) Annotations() []schema.Annotation {
	return []schema.Annotation{
		ent2ogen.BindTo("SchemaB"),
	}
}
