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

func assign(dst *ir.Field, src *gen.Field) (string, error) {
	var (
		assignT = "t." + dst.Name
		srcT    = "e." + src.StructField()
	)

	if src.Nillable {
		srcT = "*" + srcT
	}

	if dst.Type.IsGeneric() {
		if dst.Type.GenericOf.IsPrimitive() {
			return assignT + ".SetTo(" + srcT + ")", nil
		}

		gotyp := dst.Type.GenericOf.Go()
		return fmt.Sprintf("%s.SetTo(openapi.%s(%s))", assignT, gotyp, srcT), nil
	}

	if dst.Type.IsEnum() {
		gotyp := dst.Type.Go()
		return fmt.Sprintf("%s = openapi.%s(%s)", assignT, gotyp, srcT), nil
	}

	return fmt.Sprintf("%s = %s", assignT, srcT), nil
}

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
