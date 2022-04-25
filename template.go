package ent2ogen

import (
	"embed"
	"fmt"
	"text/template"

	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ogen/gen/ir"
)

var (
	//go:embed _templates
	templateDir embed.FS
	// funcMap contains extra template functions used by ent2ogen.
	funcMap = template.FuncMap{
		"sprintf": fmt.Sprintf,
		"errorf": func(format string, args ...interface{}) (interface{}, error) {
			return nil, fmt.Errorf(format, args...)
		},
		"assign":     assign,
		"rendertype": rendertype,
	}
	// templates holds all templates used by ent2ogen.
	templates = gen.MustParse(gen.NewTemplate("ent2ogen").Funcs(funcMap).ParseFS(templateDir, "_templates/*.tmpl"))
)

func rendertype(t *ir.Type) string {
	switch t.Kind {
	case ir.KindPrimitive:
		return t.Go()
	case ir.KindArray:
		return "[]" + rendertype(t.Item)
	default:
		return "openapi." + t.Go()
	}
}
