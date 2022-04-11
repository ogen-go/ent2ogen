package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Comment("Time when entity was created.").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			Comment("Time when entity was updated.").
			UpdateDefault(time.Now),
	}
}
