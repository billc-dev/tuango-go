// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/billc-dev/tuango-go/ent/predicate"
	"github.com/billc-dev/tuango-go/ent/roomuser"
)

// RoomUserDelete is the builder for deleting a RoomUser entity.
type RoomUserDelete struct {
	config
	hooks    []Hook
	mutation *RoomUserMutation
}

// Where appends a list predicates to the RoomUserDelete builder.
func (rud *RoomUserDelete) Where(ps ...predicate.RoomUser) *RoomUserDelete {
	rud.mutation.Where(ps...)
	return rud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rud *RoomUserDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rud.sqlExec, rud.mutation, rud.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rud *RoomUserDelete) ExecX(ctx context.Context) int {
	n, err := rud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rud *RoomUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(roomuser.Table, sqlgraph.NewFieldSpec(roomuser.FieldID, field.TypeString))
	if ps := rud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rud.mutation.done = true
	return affected, err
}

// RoomUserDeleteOne is the builder for deleting a single RoomUser entity.
type RoomUserDeleteOne struct {
	rud *RoomUserDelete
}

// Where appends a list predicates to the RoomUserDelete builder.
func (rudo *RoomUserDeleteOne) Where(ps ...predicate.RoomUser) *RoomUserDeleteOne {
	rudo.rud.mutation.Where(ps...)
	return rudo
}

// Exec executes the deletion query.
func (rudo *RoomUserDeleteOne) Exec(ctx context.Context) error {
	n, err := rudo.rud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{roomuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rudo *RoomUserDeleteOne) ExecX(ctx context.Context) {
	if err := rudo.Exec(ctx); err != nil {
		panic(err)
	}
}
