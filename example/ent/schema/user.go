package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("first_name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.String("user_name").Unique().
			Annotations(ent2ogen.BindTo("username")),
		field.Bool("optional_nullable_bool").Optional().Nillable(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("required_city", City.Type).Unique().Required(),
		edge.To("optional_city", City.Type).Unique(),
		edge.To("friend_list", User.Type).
			Annotations(ent2ogen.BindTo("friends")),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		ent2ogen.BindTo("User"),
	}
}
