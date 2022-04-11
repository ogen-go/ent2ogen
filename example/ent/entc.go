//go:build ignore
// +build ignore

package main

import (
	"log"

	"github.com/ogen-go/ent2ogen"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"os"
	"github.com/ogen-go/ogen"
)

func main() {
	f, err := os.ReadFile("../_openapi/schema.json")
	if err != nil {
		log.Fatalf("reading openapi schema: %v", err)
	}

	spec, err := ogen.Parse(f)
	if err != nil {
		log.Fatalf("parsing openapi schema: %v", err)
	}

	ex, err := ent2ogen.NewExtension(spec)
	if err != nil {
		log.Fatalf("creating ent2ogen extension: %v", err)
	}
	
	err = entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureUpsert,
		},
	}, entc.Extensions(ex))
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
