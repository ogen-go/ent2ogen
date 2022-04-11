package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.String("username").Unique(),
		field.String("abc"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("city", City.Type).Unique().Required(),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		TimeMixin{},
	}
}
