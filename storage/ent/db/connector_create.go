// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dexidp/dex/storage/ent/db/connector"
)

// ConnectorCreate is the builder for creating a Connector entity.
type ConnectorCreate struct {
	config
	mutation *ConnectorMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (cc *ConnectorCreate) SetType(s string) *ConnectorCreate {
	cc.mutation.SetType(s)
	return cc
}

// SetName sets the "name" field.
func (cc *ConnectorCreate) SetName(s string) *ConnectorCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetResourceVersion sets the "resource_version" field.
func (cc *ConnectorCreate) SetResourceVersion(s string) *ConnectorCreate {
	cc.mutation.SetResourceVersion(s)
	return cc
}

// SetConfig sets the "config" field.
func (cc *ConnectorCreate) SetConfig(b []byte) *ConnectorCreate {
	cc.mutation.SetConfig(b)
	return cc
}

// SetID sets the "id" field.
func (cc *ConnectorCreate) SetID(s string) *ConnectorCreate {
	cc.mutation.SetID(s)
	return cc
}

// Mutation returns the ConnectorMutation object of the builder.
func (cc *ConnectorCreate) Mutation() *ConnectorMutation {
	return cc.mutation
}

// Save creates the Connector in the database.
func (cc *ConnectorCreate) Save(ctx context.Context) (*Connector, error) {
	var (
		err  error
		node *Connector
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConnectorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ConnectorCreate) SaveX(ctx context.Context) *Connector {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ConnectorCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ConnectorCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ConnectorCreate) check() error {
	if _, ok := cc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`db: missing required field "type"`)}
	}
	if v, ok := cc.mutation.GetType(); ok {
		if err := connector.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "type": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`db: missing required field "name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := connector.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`db: validator failed for field "name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.ResourceVersion(); !ok {
		return &ValidationError{Name: "resource_version", err: errors.New(`db: missing required field "resource_version"`)}
	}
	if _, ok := cc.mutation.Config(); !ok {
		return &ValidationError{Name: "config", err: errors.New(`db: missing required field "config"`)}
	}
	if v, ok := cc.mutation.ID(); ok {
		if err := connector.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`db: validator failed for field "id": %w`, err)}
		}
	}
	return nil
}

func (cc *ConnectorCreate) sqlSave(ctx context.Context) (*Connector, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(string)
	}
	return _node, nil
}

func (cc *ConnectorCreate) createSpec() (*Connector, *sqlgraph.CreateSpec) {
	var (
		_node = &Connector{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: connector.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: connector.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: connector.FieldType,
		})
		_node.Type = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: connector.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.ResourceVersion(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: connector.FieldResourceVersion,
		})
		_node.ResourceVersion = value
	}
	if value, ok := cc.mutation.Config(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: connector.FieldConfig,
		})
		_node.Config = value
	}
	return _node, _spec
}

// ConnectorCreateBulk is the builder for creating many Connector entities in bulk.
type ConnectorCreateBulk struct {
	config
	builders []*ConnectorCreate
}

// Save creates the Connector entities in the database.
func (ccb *ConnectorCreateBulk) Save(ctx context.Context) ([]*Connector, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Connector, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ConnectorMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ConnectorCreateBulk) SaveX(ctx context.Context) []*Connector {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ConnectorCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ConnectorCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
