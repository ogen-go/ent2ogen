package ent2ogen

import (
	"encoding/json"

	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema"
)

var _ schema.Annotation = (*Annotation)(nil)

type Annotation struct {
	BindTo string
}

func Bind() Annotation {
	return Annotation{}
}

func BindTo(schema string) Annotation {
	return Annotation{
		BindTo: schema,
	}
}

// Name implements schema.Annotation interface.
func (Annotation) Name() string { return "Ent2ogen" }

// Decode unmarshalls the annotation.
func (a *Annotation) Decode(annotation interface{}) error {
	buf, err := json.Marshal(annotation)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, a)
}

func annotation(ants gen.Annotations) (*Annotation, error) {
	ant := &Annotation{}
	if ants != nil && ants[ant.Name()] != nil {
		if err := ant.Decode(ants[ant.Name()]); err != nil {
			return nil, err
		}
		return ant, nil
	}

	return nil, nil
}
