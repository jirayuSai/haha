// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/jirayuSai/app/ent/customer"
	"github.com/jirayuSai/app/ent/title"
)

// TitleCreate is the builder for creating a Title entity.
type TitleCreate struct {
	config
	mutation *TitleMutation
	hooks    []Hook
}

// SetTitleName sets the TitleName field.
func (tc *TitleCreate) SetTitleName(s string) *TitleCreate {
	tc.mutation.SetTitleName(s)
	return tc
}

// AddCustomerIDs adds the customers edge to Customer by ids.
func (tc *TitleCreate) AddCustomerIDs(ids ...int) *TitleCreate {
	tc.mutation.AddCustomerIDs(ids...)
	return tc
}

// AddCustomers adds the customers edges to Customer.
func (tc *TitleCreate) AddCustomers(c ...*Customer) *TitleCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tc.AddCustomerIDs(ids...)
}

// Mutation returns the TitleMutation object of the builder.
func (tc *TitleCreate) Mutation() *TitleMutation {
	return tc.mutation
}

// Save creates the Title in the database.
func (tc *TitleCreate) Save(ctx context.Context) (*Title, error) {
	if _, ok := tc.mutation.TitleName(); !ok {
		return nil, &ValidationError{Name: "TitleName", err: errors.New("ent: missing required field \"TitleName\"")}
	}
	var (
		err  error
		node *Title
	)
	if len(tc.hooks) == 0 {
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TitleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TitleCreate) SaveX(ctx context.Context) *Title {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tc *TitleCreate) sqlSave(ctx context.Context) (*Title, error) {
	t, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	t.ID = int(id)
	return t, nil
}

func (tc *TitleCreate) createSpec() (*Title, *sqlgraph.CreateSpec) {
	var (
		t     = &Title{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: title.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: title.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.TitleName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: title.FieldTitleName,
		})
		t.TitleName = value
	}
	if nodes := tc.mutation.CustomersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   title.CustomersTable,
			Columns: []string{title.CustomersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return t, _spec
}
