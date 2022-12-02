package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen"
)

type Keyboard struct {
	ent.Schema
}

func (Keyboard) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").NotEmpty(),
		field.Int64("price"),
		field.Int64("discount").Optional().Nillable(),
	}
}

func (Keyboard) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("switches", SwitchModel.Type).Unique().Required(),
		edge.To("keycaps", KeycapModel.Type).Unique().Required(),
	}
}

func (Keyboard) Annotations() []schema.Annotation {
	return []schema.Annotation{
		ent2ogen.BindTo("Keyboard"),
	}
}
