package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen"
)

type KeycapModel struct {
	ent.Schema
}

func (KeycapModel) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").NotEmpty(),
		field.String("profile"),
		field.Enum("material").Values("ABS", "PBT"),
	}
}

func (KeycapModel) Annotations() []schema.Annotation {
	return []schema.Annotation{
		ent2ogen.BindTo("Keycaps"),
	}
}
