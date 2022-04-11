package ent2ogen

import (
	"entgo.io/ent/entc/gen"
)

func findNode(g *gen.Graph, name string) *gen.Type {
	for _, n := range g.Nodes {
		if n.Name == name {
			return n
		}
	}

	panic("unreachable")
}
