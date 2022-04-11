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

func NewExtension(spec *ogen.Spec) (*Extension, error) {
	if spec == nil {
		return nil, fmt.Errorf("spec cannot be nil")
	}
	return &Extension{
		spec: spec,
		cfg: &Config{
			OgenPackage:     "github.com/ogen-go/ent2ogen/example/openapi",
			OgenPackageName: "openapi",
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
		if err := ex.generateMappers(g); err != nil {
			return fmt.Errorf("ent2ogen: %w", err)
		}

		// Let ent create all of its assets.
		if err := next.Generate(g); err != nil {
			return err
		}

		return nil
	})
}

func (ex *Extension) generateMappers(g *gen.Graph) error {
	for _, entSchema := range g.Schemas {
		oapiSchema, err := ex.findComponent(entSchema.Name)
		if err != nil {
			return fmt.Errorf("find %q schema: %w", entSchema.Name, err)
		}

		m := &Mapping{
			From:     entSchema,
			FromType: findNode(g, entSchema.Name),
			To:       oapiSchema,
		}

		if err := m.checkConvertability(); err != nil {
			return fmt.Errorf("type %q: %w", entSchema.Name, err)
		}

		ex.cfg.Mappings = append(ex.cfg.Mappings, m)
	}

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
