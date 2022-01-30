// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gqlapi/infrastructure/database/ent/categories"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoriesCreate is the builder for creating a Categories entity.
type CategoriesCreate struct {
	config
	mutation *CategoriesMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CategoriesCreate) SetName(s string) *CategoriesCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cc *CategoriesCreate) SetNillableName(s *string) *CategoriesCreate {
	if s != nil {
		cc.SetName(*s)
	}
	return cc
}

// SetDescription sets the "description" field.
func (cc *CategoriesCreate) SetDescription(s string) *CategoriesCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cc *CategoriesCreate) SetNillableDescription(s *string) *CategoriesCreate {
	if s != nil {
		cc.SetDescription(*s)
	}
	return cc
}

// Mutation returns the CategoriesMutation object of the builder.
func (cc *CategoriesCreate) Mutation() *CategoriesMutation {
	return cc.mutation
}

// Save creates the Categories in the database.
func (cc *CategoriesCreate) Save(ctx context.Context) (*Categories, error) {
	var (
		err  error
		node *Categories
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoriesMutation)
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
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
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
func (cc *CategoriesCreate) SaveX(ctx context.Context) *Categories {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CategoriesCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CategoriesCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CategoriesCreate) defaults() {
	if _, ok := cc.mutation.Name(); !ok {
		v := categories.DefaultName
		cc.mutation.SetName(v)
	}
	if _, ok := cc.mutation.Description(); !ok {
		v := categories.DefaultDescription
		cc.mutation.SetDescription(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CategoriesCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Categories.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := categories.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Categories.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Categories.description"`)}
	}
	if v, ok := cc.mutation.Description(); ok {
		if err := categories.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Categories.description": %w`, err)}
		}
	}
	return nil
}

func (cc *CategoriesCreate) sqlSave(ctx context.Context) (*Categories, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CategoriesCreate) createSpec() (*Categories, *sqlgraph.CreateSpec) {
	var (
		_node = &Categories{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: categories.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: categories.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: categories.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: categories.FieldDescription,
		})
		_node.Description = value
	}
	return _node, _spec
}

// CategoriesCreateBulk is the builder for creating many Categories entities in bulk.
type CategoriesCreateBulk struct {
	config
	builders []*CategoriesCreate
}

// Save creates the Categories entities in the database.
func (ccb *CategoriesCreateBulk) Save(ctx context.Context) ([]*Categories, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Categories, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CategoriesMutation)
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
func (ccb *CategoriesCreateBulk) SaveX(ctx context.Context) []*Categories {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CategoriesCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CategoriesCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
