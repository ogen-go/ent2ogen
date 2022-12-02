package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen"
)

type SwitchModel struct {
	ent.Schema
}

func (SwitchModel) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").NotEmpty(),
		field.Enum("switch_type").Values(
			"mechanical",
			"optical",
			"electrocapacitive",
		),
	}
}

func (SwitchModel) Annotations() []schema.Annotation {
	return []schema.Annotation{
		ent2ogen.BindTo("Switches"),
	}
}