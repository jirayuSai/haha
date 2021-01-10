// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/jirayuSai/app/ent/customer"
	"github.com/jirayuSai/app/ent/gender"
	"github.com/jirayuSai/app/ent/predicate"
)

// GenderUpdate is the builder for updating Gender entities.
type GenderUpdate struct {
	config
	hooks      []Hook
	mutation   *GenderMutation
	predicates []predicate.Gender
}

// Where adds a new predicate for the builder.
func (gu *GenderUpdate) Where(ps ...predicate.Gender) *GenderUpdate {
	gu.predicates = append(gu.predicates, ps...)
	return gu
}

// SetGenderName sets the GenderName field.
func (gu *GenderUpdate) SetGenderName(s string) *GenderUpdate {
	gu.mutation.SetGenderName(s)
	return gu
}

// AddCustomerIDs adds the customers edge to Customer by ids.
func (gu *GenderUpdate) AddCustomerIDs(ids ...int) *GenderUpdate {
	gu.mutation.AddCustomerIDs(ids...)
	return gu
}

// AddCustomers adds the customers edges to Customer.
func (gu *GenderUpdate) AddCustomers(c ...*Customer) *GenderUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return gu.AddCustomerIDs(ids...)
}

// Mutation returns the GenderMutation object of the builder.
func (gu *GenderUpdate) Mutation() *GenderMutation {
	return gu.mutation
}

// RemoveCustomerIDs removes the customers edge to Customer by ids.
func (gu *GenderUpdate) RemoveCustomerIDs(ids ...int) *GenderUpdate {
	gu.mutation.RemoveCustomerIDs(ids...)
	return gu
}

// RemoveCustomers removes customers edges to Customer.
func (gu *GenderUpdate) RemoveCustomers(c ...*Customer) *GenderUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return gu.RemoveCustomerIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (gu *GenderUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(gu.hooks) == 0 {
		affected, err = gu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GenderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gu.mutation = mutation
			affected, err = gu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gu.hooks) - 1; i >= 0; i-- {
			mut = gu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GenderUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GenderUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GenderUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GenderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gender.Table,
			Columns: gender.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gender.FieldID,
			},
		},
	}
	if ps := gu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.GenderName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gender.FieldGenderName,
		})
	}
	if nodes := gu.mutation.RemovedCustomersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gender.CustomersTable,
			Columns: []string{gender.CustomersColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.CustomersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gender.CustomersTable,
			Columns: []string{gender.CustomersColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gender.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// GenderUpdateOne is the builder for updating a single Gender entity.
type GenderUpdateOne struct {
	config
	hooks    []Hook
	mutation *GenderMutation
}

// SetGenderName sets the GenderName field.
func (guo *GenderUpdateOne) SetGenderName(s string) *GenderUpdateOne {
	guo.mutation.SetGenderName(s)
	return guo
}

// AddCustomerIDs adds the customers edge to Customer by ids.
func (guo *GenderUpdateOne) AddCustomerIDs(ids ...int) *GenderUpdateOne {
	guo.mutation.AddCustomerIDs(ids...)
	return guo
}

// AddCustomers adds the customers edges to Customer.
func (guo *GenderUpdateOne) AddCustomers(c ...*Customer) *GenderUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return guo.AddCustomerIDs(ids...)
}

// Mutation returns the GenderMutation object of the builder.
func (guo *GenderUpdateOne) Mutation() *GenderMutation {
	return guo.mutation
}

// RemoveCustomerIDs removes the customers edge to Customer by ids.
func (guo *GenderUpdateOne) RemoveCustomerIDs(ids ...int) *GenderUpdateOne {
	guo.mutation.RemoveCustomerIDs(ids...)
	return guo
}

// RemoveCustomers removes customers edges to Customer.
func (guo *GenderUpdateOne) RemoveCustomers(c ...*Customer) *GenderUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return guo.RemoveCustomerIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (guo *GenderUpdateOne) Save(ctx context.Context) (*Gender, error) {

	var (
		err  error
		node *Gender
	)
	if len(guo.hooks) == 0 {
		node, err = guo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GenderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			guo.mutation = mutation
			node, err = guo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(guo.hooks) - 1; i >= 0; i-- {
			mut = guo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, guo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GenderUpdateOne) SaveX(ctx context.Context) *Gender {
	ge, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return ge
}

// Exec executes the query on the entity.
func (guo *GenderUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GenderUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GenderUpdateOne) sqlSave(ctx context.Context) (ge *Gender, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gender.Table,
			Columns: gender.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: gender.FieldID,
			},
		},
	}
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Gender.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := guo.mutation.GenderName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: gender.FieldGenderName,
		})
	}
	if nodes := guo.mutation.RemovedCustomersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gender.CustomersTable,
			Columns: []string{gender.CustomersColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.CustomersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gender.CustomersTable,
			Columns: []string{gender.CustomersColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	ge = &Gender{config: guo.config}
	_spec.Assign = ge.assignValues
	_spec.ScanValues = ge.scanValues()
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gender.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return ge, nil
}