package ent2ogen

import "entgo.io/ent/entc/gen"

type walkpath struct {
	types []*gen.Type
}

func (w *walkpath) append(t *gen.Type) *walkpath {
	return &walkpath{
		types: append(w.types[:len(w.types):len(w.types)], t),
	}
}

func (w *walkpath) has(t *gen.Type) bool {
	for _, tt := range w.types {
		if tt == t {
			return true
		}
	}
	return false
}

func (w *walkpath) typeNames() []string {
	names := make([]string, 0, len(w.types))
	for _, t := range w.types {
		names = append(names, t.Name)
	}
	return names
}
