// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/enumwithconflictingvalue"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EnumWithConflictingValueCreate is the builder for creating a EnumWithConflictingValue entity.
type EnumWithConflictingValueCreate struct {
	config
	mutation *EnumWithConflictingValueMutation
	hooks    []Hook
}

// SetEnum sets the "enum" field.
func (ewcvc *EnumWithConflictingValueCreate) SetEnum(e enumwithconflictingvalue.Enum) *EnumWithConflictingValueCreate {
	ewcvc.mutation.SetEnum(e)
	return ewcvc
}

// Mutation returns the EnumWithConflictingValueMutation object of the builder.
func (ewcvc *EnumWithConflictingValueCreate) Mutation() *EnumWithConflictingValueMutation {
	return ewcvc.mutation
}

// Save creates the EnumWithConflictingValue in the database.
func (ewcvc *EnumWithConflictingValueCreate) Save(ctx context.Context) (*EnumWithConflictingValue, error) {
	return withHooks(ctx, ewcvc.sqlSave, ewcvc.mutation, ewcvc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ewcvc *EnumWithConflictingValueCreate) SaveX(ctx context.Context) *EnumWithConflictingValue {
	v, err := ewcvc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ewcvc *EnumWithConflictingValueCreate) Exec(ctx context.Context) error {
	_, err := ewcvc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ewcvc *EnumWithConflictingValueCreate) ExecX(ctx context.Context) {
	if err := ewcvc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ewcvc *EnumWithConflictingValueCreate) check() error {
	if _, ok := ewcvc.mutation.Enum(); !ok {
		return &ValidationError{Name: "enum", err: errors.New(`ent: missing required field "EnumWithConflictingValue.enum"`)}
	}
	if v, ok := ewcvc.mutation.Enum(); ok {
		if err := enumwithconflictingvalue.EnumValidator(v); err != nil {
			return &ValidationError{Name: "enum", err: fmt.Errorf(`ent: validator failed for field "EnumWithConflictingValue.enum": %w`, err)}
		}
	}
	return nil
}

func (ewcvc *EnumWithConflictingValueCreate) sqlSave(ctx context.Context) (*EnumWithConflictingValue, error) {
	if err := ewcvc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ewcvc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ewcvc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ewcvc.mutation.id = &_node.ID
	ewcvc.mutation.done = true
	return _node, nil
}

func (ewcvc *EnumWithConflictingValueCreate) createSpec() (*EnumWithConflictingValue, *sqlgraph.CreateSpec) {
	var (
		_node = &EnumWithConflictingValue{config: ewcvc.config}
		_spec = sqlgraph.NewCreateSpec(enumwithconflictingvalue.Table, sqlgraph.NewFieldSpec(enumwithconflictingvalue.FieldID, field.TypeInt))
	)
	if value, ok := ewcvc.mutation.Enum(); ok {
		_spec.SetField(enumwithconflictingvalue.FieldEnum, field.TypeEnum, value)
		_node.Enum = value
	}
	return _node, _spec
}

// EnumWithConflictingValueCreateBulk is the builder for creating many EnumWithConflictingValue entities in bulk.
type EnumWithConflictingValueCreateBulk struct {
	config
	err      error
	builders []*EnumWithConflictingValueCreate
}

// Save creates the EnumWithConflictingValue entities in the database.
func (ewcvcb *EnumWithConflictingValueCreateBulk) Save(ctx context.Context) ([]*EnumWithConflictingValue, error) {
	if ewcvcb.err != nil {
		return nil, ewcvcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ewcvcb.builders))
	nodes := make([]*EnumWithConflictingValue, len(ewcvcb.builders))
	mutators := make([]Mutator, len(ewcvcb.builders))
	for i := range ewcvcb.builders {
		func(i int, root context.Context) {
			builder := ewcvcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EnumWithConflictingValueMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ewcvcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ewcvcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ewcvcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ewcvcb *EnumWithConflictingValueCreateBulk) SaveX(ctx context.Context) []*EnumWithConflictingValue {
	v, err := ewcvcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ewcvcb *EnumWithConflictingValueCreateBulk) Exec(ctx context.Context) error {
	_, err := ewcvcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ewcvcb *EnumWithConflictingValueCreateBulk) ExecX(ctx context.Context) {
	if err := ewcvcb.Exec(ctx); err != nil {
		panic(err)
	}
}