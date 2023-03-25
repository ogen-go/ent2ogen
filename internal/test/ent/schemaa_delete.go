// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ent2ogen/internal/test/ent/predicate"
	"github.com/ogen-go/ent2ogen/internal/test/ent/schemaa"
)

// SchemaADelete is the builder for deleting a SchemaA entity.
type SchemaADelete struct {
	config
	hooks    []Hook
	mutation *SchemaAMutation
}

// Where appends a list predicates to the SchemaADelete builder.
func (sa *SchemaADelete) Where(ps ...predicate.SchemaA) *SchemaADelete {
	sa.mutation.Where(ps...)
	return sa
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sa *SchemaADelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, SchemaAMutation](ctx, sa.sqlExec, sa.mutation, sa.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sa *SchemaADelete) ExecX(ctx context.Context) int {
	n, err := sa.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sa *SchemaADelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(schemaa.Table, sqlgraph.NewFieldSpec(schemaa.FieldID, field.TypeInt))
	if ps := sa.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sa.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sa.mutation.done = true
	return affected, err
}

// SchemaADeleteOne is the builder for deleting a single SchemaA entity.
type SchemaADeleteOne struct {
	sa *SchemaADelete
}

// Where appends a list predicates to the SchemaADelete builder.
func (sao *SchemaADeleteOne) Where(ps ...predicate.SchemaA) *SchemaADeleteOne {
	sao.sa.mutation.Where(ps...)
	return sao
}

// Exec executes the deletion query.
func (sao *SchemaADeleteOne) Exec(ctx context.Context) error {
	n, err := sao.sa.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{schemaa.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sao *SchemaADeleteOne) ExecX(ctx context.Context) {
	if err := sao.Exec(ctx); err != nil {
		panic(err)
	}
}
