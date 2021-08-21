// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithfieldone"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithFieldOneUpdate is the builder for updating MessageWithFieldOne entities.
type MessageWithFieldOneUpdate struct {
	config
	hooks    []Hook
	mutation *MessageWithFieldOneMutation
}

// Where appends a list predicates to the MessageWithFieldOneUpdate builder.
func (mwfou *MessageWithFieldOneUpdate) Where(ps ...predicate.MessageWithFieldOne) *MessageWithFieldOneUpdate {
	mwfou.mutation.Where(ps...)
	return mwfou
}

// SetFieldOne sets the "field_one" field.
func (mwfou *MessageWithFieldOneUpdate) SetFieldOne(i int32) *MessageWithFieldOneUpdate {
	mwfou.mutation.ResetFieldOne()
	mwfou.mutation.SetFieldOne(i)
	return mwfou
}

// AddFieldOne adds i to the "field_one" field.
func (mwfou *MessageWithFieldOneUpdate) AddFieldOne(i int32) *MessageWithFieldOneUpdate {
	mwfou.mutation.AddFieldOne(i)
	return mwfou
}

// Mutation returns the MessageWithFieldOneMutation object of the builder.
func (mwfou *MessageWithFieldOneUpdate) Mutation() *MessageWithFieldOneMutation {
	return mwfou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mwfou *MessageWithFieldOneUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mwfou.hooks) == 0 {
		affected, err = mwfou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageWithFieldOneMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mwfou.mutation = mutation
			affected, err = mwfou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mwfou.hooks) - 1; i >= 0; i-- {
			if mwfou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mwfou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mwfou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mwfou *MessageWithFieldOneUpdate) SaveX(ctx context.Context) int {
	affected, err := mwfou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mwfou *MessageWithFieldOneUpdate) Exec(ctx context.Context) error {
	_, err := mwfou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwfou *MessageWithFieldOneUpdate) ExecX(ctx context.Context) {
	if err := mwfou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mwfou *MessageWithFieldOneUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messagewithfieldone.Table,
			Columns: messagewithfieldone.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messagewithfieldone.FieldID,
			},
		},
	}
	if ps := mwfou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mwfou.mutation.FieldOne(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: messagewithfieldone.FieldFieldOne,
		})
	}
	if value, ok := mwfou.mutation.AddedFieldOne(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: messagewithfieldone.FieldFieldOne,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mwfou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messagewithfieldone.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MessageWithFieldOneUpdateOne is the builder for updating a single MessageWithFieldOne entity.
type MessageWithFieldOneUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageWithFieldOneMutation
}

// SetFieldOne sets the "field_one" field.
func (mwfouo *MessageWithFieldOneUpdateOne) SetFieldOne(i int32) *MessageWithFieldOneUpdateOne {
	mwfouo.mutation.ResetFieldOne()
	mwfouo.mutation.SetFieldOne(i)
	return mwfouo
}

// AddFieldOne adds i to the "field_one" field.
func (mwfouo *MessageWithFieldOneUpdateOne) AddFieldOne(i int32) *MessageWithFieldOneUpdateOne {
	mwfouo.mutation.AddFieldOne(i)
	return mwfouo
}

// Mutation returns the MessageWithFieldOneMutation object of the builder.
func (mwfouo *MessageWithFieldOneUpdateOne) Mutation() *MessageWithFieldOneMutation {
	return mwfouo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mwfouo *MessageWithFieldOneUpdateOne) Select(field string, fields ...string) *MessageWithFieldOneUpdateOne {
	mwfouo.fields = append([]string{field}, fields...)
	return mwfouo
}

// Save executes the query and returns the updated MessageWithFieldOne entity.
func (mwfouo *MessageWithFieldOneUpdateOne) Save(ctx context.Context) (*MessageWithFieldOne, error) {
	var (
		err  error
		node *MessageWithFieldOne
	)
	if len(mwfouo.hooks) == 0 {
		node, err = mwfouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageWithFieldOneMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mwfouo.mutation = mutation
			node, err = mwfouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mwfouo.hooks) - 1; i >= 0; i-- {
			if mwfouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mwfouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mwfouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mwfouo *MessageWithFieldOneUpdateOne) SaveX(ctx context.Context) *MessageWithFieldOne {
	node, err := mwfouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mwfouo *MessageWithFieldOneUpdateOne) Exec(ctx context.Context) error {
	_, err := mwfouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwfouo *MessageWithFieldOneUpdateOne) ExecX(ctx context.Context) {
	if err := mwfouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mwfouo *MessageWithFieldOneUpdateOne) sqlSave(ctx context.Context) (_node *MessageWithFieldOne, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messagewithfieldone.Table,
			Columns: messagewithfieldone.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messagewithfieldone.FieldID,
			},
		},
	}
	id, ok := mwfouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing MessageWithFieldOne.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := mwfouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messagewithfieldone.FieldID)
		for _, f := range fields {
			if !messagewithfieldone.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != messagewithfieldone.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mwfouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mwfouo.mutation.FieldOne(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: messagewithfieldone.FieldFieldOne,
		})
	}
	if value, ok := mwfouo.mutation.AddedFieldOne(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: messagewithfieldone.FieldFieldOne,
		})
	}
	_node = &MessageWithFieldOne{config: mwfouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mwfouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messagewithfieldone.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
