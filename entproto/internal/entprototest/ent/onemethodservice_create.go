// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/onemethodservice"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OneMethodServiceCreate is the builder for creating a OneMethodService entity.
type OneMethodServiceCreate struct {
	config
	mutation *OneMethodServiceMutation
	hooks    []Hook
}

// Mutation returns the OneMethodServiceMutation object of the builder.
func (omsc *OneMethodServiceCreate) Mutation() *OneMethodServiceMutation {
	return omsc.mutation
}

// Save creates the OneMethodService in the database.
func (omsc *OneMethodServiceCreate) Save(ctx context.Context) (*OneMethodService, error) {
	return withHooks(ctx, omsc.sqlSave, omsc.mutation, omsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (omsc *OneMethodServiceCreate) SaveX(ctx context.Context) *OneMethodService {
	v, err := omsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (omsc *OneMethodServiceCreate) Exec(ctx context.Context) error {
	_, err := omsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (omsc *OneMethodServiceCreate) ExecX(ctx context.Context) {
	if err := omsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (omsc *OneMethodServiceCreate) check() error {
	return nil
}

func (omsc *OneMethodServiceCreate) sqlSave(ctx context.Context) (*OneMethodService, error) {
	if err := omsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := omsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, omsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	omsc.mutation.id = &_node.ID
	omsc.mutation.done = true
	return _node, nil
}

func (omsc *OneMethodServiceCreate) createSpec() (*OneMethodService, *sqlgraph.CreateSpec) {
	var (
		_node = &OneMethodService{config: omsc.config}
		_spec = sqlgraph.NewCreateSpec(onemethodservice.Table, sqlgraph.NewFieldSpec(onemethodservice.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// OneMethodServiceCreateBulk is the builder for creating many OneMethodService entities in bulk.
type OneMethodServiceCreateBulk struct {
	config
	err      error
	builders []*OneMethodServiceCreate
}

// Save creates the OneMethodService entities in the database.
func (omscb *OneMethodServiceCreateBulk) Save(ctx context.Context) ([]*OneMethodService, error) {
	if omscb.err != nil {
		return nil, omscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(omscb.builders))
	nodes := make([]*OneMethodService, len(omscb.builders))
	mutators := make([]Mutator, len(omscb.builders))
	for i := range omscb.builders {
		func(i int, root context.Context) {
			builder := omscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OneMethodServiceMutation)
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
					_, err = mutators[i+1].Mutate(root, omscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, omscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, omscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (omscb *OneMethodServiceCreateBulk) SaveX(ctx context.Context) []*OneMethodService {
	v, err := omscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (omscb *OneMethodServiceCreateBulk) Exec(ctx context.Context) error {
	_, err := omscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (omscb *OneMethodServiceCreateBulk) ExecX(ctx context.Context) {
	if err := omscb.Exec(ctx); err != nil {
		panic(err)
	}
}
