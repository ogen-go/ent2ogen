package ent2ogen

import (
	"fmt"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen"
)

type Extension struct {
	entc.DefaultExtension
	spec *ogen.Spec
	cfg  *Config
}

type Config struct {
	OgenPackage string
	Mappings    []*Mapping
}

func (Config) Name() string {
	return "Ent2ogen"
}

func NewExtension(spec *ogen.Spec) (*Extension, error) {
	if spec == nil {
		return nil, fmt.Errorf("spec cannot be nil")
	}
	return &Extension{
		spec: spec,
		cfg: &Config{
			OgenPackage: "github.com/ogen-go/ent2ogen/example/openapi",
		},
	}, nil
}

// Hooks of the extension.
func (ex *Extension) Hooks() []gen.Hook {
	return []gen.Hook{
		ex.ogen,
	}
}

// Templates of the extension.
func (ex *Extension) Templates() []*gen.Template {
	return []*gen.Template{templates}
}

// Annotations of the extension.
func (ex *Extension) Annotations() []entc.Annotation {
	return []entc.Annotation{ex.cfg}
}

func (ex *Extension) ogen(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(g *gen.Graph) error {
		for _, n := range g.Nodes {
			if err := ex.generateMapping(n); err != nil {
				return fmt.Errorf("type %q: %w", n.Name, err)
			}
		}

		return next.Generate(g)
	})
}

func (ex *Extension) generateMapping(n *gen.Type) error {
	ant, err := annotation(n.Annotations)
	if err != nil {
		return fmt.Errorf("read annotation: %w", err)
	}

	if ant == nil {
		return nil
	}

	// OpenAPI schema component.
	schemaName := n.Name
	if ant.BindTo != "" {
		schemaName = ant.BindTo
	}

	s, err := ex.findComponent(schemaName)
	if err != nil {
		return fmt.Errorf("find %q schema: %w", schemaName, err)
	}

	m := &Mapping{From: n, To: s}
	if err := m.checkCompatibility(); err != nil {
		return fmt.Errorf("type %q: %w", n.Name, err)
	}

	ex.cfg.Mappings = append(ex.cfg.Mappings, m)
	return nil
}

func (ex *Extension) findComponent(name string) (*ogen.Schema, error) {
	if ex.spec.Components == nil {
		return nil, fmt.Errorf("components cannot be nil")
	}

	if ex.spec.Components.Schemas == nil {
		return nil, fmt.Errorf("schema components cannot be nil")
	}

	s, ok := ex.spec.Components.Schemas[name]
	if !ok {
		return nil, fmt.Errorf("component is not present in the openapi document")
	}

	return s, nil
}
