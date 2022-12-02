//go:build ignore
// +build ignore

package main

import (
	"log"
	"os"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ogen-go/ent2ogen"
	"github.com/ogen-go/ogen"
	ogengen "github.com/ogen-go/ogen/gen"
	ogengenfs "github.com/ogen-go/ogen/gen/genfs"
)

func main() {
	f, err := os.ReadFile("../_openapi/openapi.yml")
	if err != nil {
		log.Fatalf("reading openapi schema: %v", err)
	}

	spec, err := ogen.Parse(f)
	if err != nil {
		log.Fatalf("parsing openapi schema: %v", err)
	}

	g, err := ogengen.NewGenerator(spec, ogengen.Options{})
	if err != nil {
		log.Fatalf("creating ogen generator: %v", err)
	}

	if err := g.WriteSource(ogengenfs.FormattedSource{
		Format: true,
		Root:   "../api",
	}, "openapi"); err != nil {
		log.Fatalf("generating ogen sources: %v", err)
	}

	ex, err := ent2ogen.NewExtension(ent2ogen.ExtensionConfig{
		API:         g.API(),
		Types:       g.Types(),
		OgenPackage: "github.com/ogen-go/ent2ogen/internal/test/api",
	})
	if err != nil {
		log.Fatalf("creating ent2ogen extension: %v", err)
	}

	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ex))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
