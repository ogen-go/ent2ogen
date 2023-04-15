package ent2ogen

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/go-faster/errors"
	"github.com/ogen-go/ogen/gen/ir"
	"github.com/ogen-go/ogen/jsonschema"
	"github.com/ogen-go/ogen/openapi"
)

type Extension struct {
	entc.DefaultExtension
	api   *openapi.API
	index map[*jsonschema.Schema]*ir.Type
	recur map[*gen.Type]struct{}
	cfg   *Config
}

type Config struct {
	OgenPackage string
	Mappings    map[*gen.Type]*Mapping
}

func (Config) Name() string {
	return "Ent2ogen"
}

type ExtensionConfig struct {
	API         *openapi.API
	Types       map[string]*ir.Type
	OgenPackage string
}

func NewExtension(cfg ExtensionConfig) (*Extension, error) {
	if cfg.API == nil {
		return nil, errors.New("spec cannot be nil")
	}
	if cfg.Types == nil {
		return nil, errors.New("types map cannot be nil")
	}

	index := make(map[*jsonschema.Schema]*ir.Type)
	for _, t := range cfg.Types {
		if t.Schema == nil {
			continue
		}

		if _, ok := index[t.Schema]; ok {
			return nil, errors.Errorf("type map schema collision: %+v", t)
		}

		index[t.Schema] = t
	}

	return &Extension{
		api:   cfg.API,
		index: index,
		recur: map[*gen.Type]struct{}{},
		cfg: &Config{
			OgenPackage: cfg.OgenPackage,
			Mappings:    map[*gen.Type]*Mapping{},
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
				return errors.Wrapf(err, "type %q", n.Name)
			}
		}

		return next.Generate(g)
	})
}

func (ex *Extension) generateMapping(n *gen.Type) error {
	ant, err := annotation(n.Annotations)
	if err != nil {
		return errors.Wrap(err, "read annotation")
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
		return errors.Wrapf(err, "find %q schema", schemaName)
	}

	t, ok := ex.index[s]
	if !ok {
		return errors.Errorf("schema %q: ir type not found", schemaName)
	}

	return ex.createMapping(n, t)
}

func (ex *Extension) findComponent(name string) (*jsonschema.Schema, error) {
	s, ok := ex.api.Components.Schemas[name]
	if !ok {
		return nil, errors.New("component is not present in the openapi document")
	}

	return s, nil
}
