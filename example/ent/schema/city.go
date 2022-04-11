package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen"
)

type City struct {
	ent.Schema
}

func (City) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

func (City) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		TimeMixin{},
	}
}

func (City) Annotations() []schema.Annotation {
	return []schema.Annotation{
		ent2ogen.BindTo("City"),
	}
}
