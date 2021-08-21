// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithid"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithIDUpdate is the builder for updating MessageWithID entities.
type MessageWithIDUpdate struct {
	config
	hooks    []Hook
	mutation *MessageWithIDMutation
}

// Where appends a list predicates to the MessageWithIDUpdate builder.
func (mwiu *MessageWithIDUpdate) Where(ps ...predicate.MessageWithID) *MessageWithIDUpdate {
	mwiu.mutation.Where(ps...)
	return mwiu
}

// Mutation returns the MessageWithIDMutation object of the builder.
func (mwiu *MessageWithIDUpdate) Mutation() *MessageWithIDMutation {
	return mwiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mwiu *MessageWithIDUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mwiu.hooks) == 0 {
		affected, err = mwiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageWithIDMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mwiu.mutation = mutation
			affected, err = mwiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mwiu.hooks) - 1; i >= 0; i-- {
			if mwiu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mwiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mwiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mwiu *MessageWithIDUpdate) SaveX(ctx context.Context) int {
	affected, err := mwiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mwiu *MessageWithIDUpdate) Exec(ctx context.Context) error {
	_, err := mwiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwiu *MessageWithIDUpdate) ExecX(ctx context.Context) {
	if err := mwiu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mwiu *MessageWithIDUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messagewithid.Table,
			Columns: messagewithid.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: messagewithid.FieldID,
			},
		},
	}
	if ps := mwiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mwiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messagewithid.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MessageWithIDUpdateOne is the builder for updating a single MessageWithID entity.
type MessageWithIDUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageWithIDMutation
}

// Mutation returns the MessageWithIDMutation object of the builder.
func (mwiuo *MessageWithIDUpdateOne) Mutation() *MessageWithIDMutation {
	return mwiuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mwiuo *MessageWithIDUpdateOne) Select(field string, fields ...string) *MessageWithIDUpdateOne {
	mwiuo.fields = append([]string{field}, fields...)
	return mwiuo
}

// Save executes the query and returns the updated MessageWithID entity.
func (mwiuo *MessageWithIDUpdateOne) Save(ctx context.Context) (*MessageWithID, error) {
	var (
		err  error
		node *MessageWithID
	)
	if len(mwiuo.hooks) == 0 {
		node, err = mwiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageWithIDMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mwiuo.mutation = mutation
			node, err = mwiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mwiuo.hooks) - 1; i >= 0; i-- {
			if mwiuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mwiuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mwiuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (mwiuo *MessageWithIDUpdateOne) SaveX(ctx context.Context) *MessageWithID {
	node, err := mwiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mwiuo *MessageWithIDUpdateOne) Exec(ctx context.Context) error {
	_, err := mwiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwiuo *MessageWithIDUpdateOne) ExecX(ctx context.Context) {
	if err := mwiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mwiuo *MessageWithIDUpdateOne) sqlSave(ctx context.Context) (_node *MessageWithID, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   messagewithid.Table,
			Columns: messagewithid.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: messagewithid.FieldID,
			},
		},
	}
	id, ok := mwiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing MessageWithID.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := mwiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, messagewithid.FieldID)
		for _, f := range fields {
			if !messagewithid.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != messagewithid.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mwiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &MessageWithID{config: mwiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mwiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{messagewithid.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
