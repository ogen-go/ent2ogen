package ent2ogen

import (
	"embed"
	"fmt"
	"text/template"

	"entgo.io/ent/entc/gen"
	"github.com/go-faster/errors"
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
		"unassign":   unassign,
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

func unassign(dst *ir.Field) (string, error) {
	assignT := "t." + dst.Name

	type assignStmt struct {
		variable string
		value    string
	}

	var stmts []assignStmt
	if dst.Type.IsGeneric() {
		if dst.Type.GenericVariant.Optional {
			// Reset ogen's generic 'Set' field.
			//
			// t.Foo.Set = false
			stmts = append(stmts, assignStmt{
				variable: fmt.Sprintf("%s.Set", assignT),
				value:    "false",
			})
		}
		if dst.Type.GenericVariant.Nullable {
			// Reset ogen's generic 'Null' field.
			//
			// t.Foo.Null = true
			stmts = append(stmts, assignStmt{
				variable: fmt.Sprintf("%s.Null", assignT),
				value:    "true",
			})
		}
	}

	// From github.com/ogen-go/ogen/gen/generics.go:boxType
	if dst.Type.GenericOf.Kind == ir.KindArray || dst.Type.GenericOf.Primitive == ir.ByteSlice {
		if dst.Type.GenericVariant.NullableOptional() {
			return "", errors.New("unexpected nullable & optional ogen array type")
		}

		// t.Foo = nil
		stmts = append(stmts, assignStmt{
			variable: assignT,
			value:    "nil",
		})
	}

	// Represent assignments in a shorter form:
	// foo, bar = true, false
	oneliner := ""
	for i, stmt := range stmts {
		oneliner += stmt.variable
		if i+1 != len(stmts) {
			oneliner += ", "
		}
	}

	oneliner += " = "
	for i, stmt := range stmts {
		oneliner += stmt.value
		if i+1 != len(stmts) {
			oneliner += ", "
		}
	}

	return oneliner, nil
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
