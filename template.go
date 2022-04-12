package ent2ogen

import (
	"embed"
	"fmt"
	"text/template"

	"entgo.io/ent/entc/gen"
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
	}
	// templates holds all templates used by ent2ogen.
	templates = gen.MustParse(gen.NewTemplate("ent2ogen").Funcs(funcMap).ParseFS(templateDir, "_templates/*.tmpl"))
)
