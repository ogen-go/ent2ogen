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
	}
	// templates holds all templates used by ent2ogen.
	templates = gen.MustParse(gen.NewTemplate("ent2ogen").Funcs(funcMap).ParseFS(templateDir, "_templates/*.tmpl"))
)

type Config struct {
	OgenPackage     string
	OgenPackageName string

	Mappings []*Mapping
}

func (Config) Name() string {
	return "Ent2ogen"
}
