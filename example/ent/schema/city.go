package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
