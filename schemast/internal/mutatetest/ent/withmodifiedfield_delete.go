// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/schemast/internal/mutatetest/ent/predicate"
	"entgo.io/contrib/schemast/internal/mutatetest/ent/withmodifiedfield"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WithModifiedFieldDelete is the builder for deleting a WithModifiedField entity.
type WithModifiedFieldDelete struct {
	config
	hooks    []Hook
	mutation *WithModifiedFieldMutation
}

// Where appends a list predicates to the WithModifiedFieldDelete builder.
func (wmfd *WithModifiedFieldDelete) Where(ps ...predicate.WithModifiedField) *WithModifiedFieldDelete {
	wmfd.mutation.Where(ps...)
	return wmfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wmfd *WithModifiedFieldDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wmfd.sqlExec, wmfd.mutation, wmfd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wmfd *WithModifiedFieldDelete) ExecX(ctx context.Context) int {
	n, err := wmfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wmfd *WithModifiedFieldDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(withmodifiedfield.Table, sqlgraph.NewFieldSpec(withmodifiedfield.FieldID, field.TypeInt))
	if ps := wmfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wmfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wmfd.mutation.done = true
	return affected, err
}

// WithModifiedFieldDeleteOne is the builder for deleting a single WithModifiedField entity.
type WithModifiedFieldDeleteOne struct {
	wmfd *WithModifiedFieldDelete
}

// Where appends a list predicates to the WithModifiedFieldDelete builder.
func (wmfdo *WithModifiedFieldDeleteOne) Where(ps ...predicate.WithModifiedField) *WithModifiedFieldDeleteOne {
	wmfdo.wmfd.mutation.Where(ps...)
	return wmfdo
}

// Exec executes the deletion query.
func (wmfdo *WithModifiedFieldDeleteOne) Exec(ctx context.Context) error {
	n, err := wmfdo.wmfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{withmodifiedfield.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wmfdo *WithModifiedFieldDeleteOne) ExecX(ctx context.Context) {
	if err := wmfdo.Exec(ctx); err != nil {
		panic(err)
	}
}
