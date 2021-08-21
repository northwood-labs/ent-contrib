// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/implicitskippedmessage"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ImplicitSkippedMessageDelete is the builder for deleting a ImplicitSkippedMessage entity.
type ImplicitSkippedMessageDelete struct {
	config
	hooks    []Hook
	mutation *ImplicitSkippedMessageMutation
}

// Where appends a list predicates to the ImplicitSkippedMessageDelete builder.
func (ismd *ImplicitSkippedMessageDelete) Where(ps ...predicate.ImplicitSkippedMessage) *ImplicitSkippedMessageDelete {
	ismd.mutation.Where(ps...)
	return ismd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ismd *ImplicitSkippedMessageDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ismd.hooks) == 0 {
		affected, err = ismd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImplicitSkippedMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ismd.mutation = mutation
			affected, err = ismd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ismd.hooks) - 1; i >= 0; i-- {
			if ismd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ismd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ismd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ismd *ImplicitSkippedMessageDelete) ExecX(ctx context.Context) int {
	n, err := ismd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ismd *ImplicitSkippedMessageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: implicitskippedmessage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: implicitskippedmessage.FieldID,
			},
		},
	}
	if ps := ismd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ismd.driver, _spec)
}

// ImplicitSkippedMessageDeleteOne is the builder for deleting a single ImplicitSkippedMessage entity.
type ImplicitSkippedMessageDeleteOne struct {
	ismd *ImplicitSkippedMessageDelete
}

// Exec executes the deletion query.
func (ismdo *ImplicitSkippedMessageDeleteOne) Exec(ctx context.Context) error {
	n, err := ismdo.ismd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{implicitskippedmessage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ismdo *ImplicitSkippedMessageDeleteOne) ExecX(ctx context.Context) {
	ismdo.ismd.ExecX(ctx)
}
