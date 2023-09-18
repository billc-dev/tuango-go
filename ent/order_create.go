// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/billc-dev/tuango-go/ent/order"
	"github.com/billc-dev/tuango-go/ent/orderhistory"
	"github.com/billc-dev/tuango-go/ent/orderitem"
	"github.com/billc-dev/tuango-go/ent/post"
	"github.com/billc-dev/tuango-go/ent/user"
)

// OrderCreate is the builder for creating a Order entity.
type OrderCreate struct {
	config
	mutation *OrderMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (oc *OrderCreate) SetUserID(s string) *OrderCreate {
	oc.mutation.SetUserID(s)
	return oc
}

// SetPostID sets the "post_id" field.
func (oc *OrderCreate) SetPostID(s string) *OrderCreate {
	oc.mutation.SetPostID(s)
	return oc
}

// SetOrderNum sets the "order_num" field.
func (oc *OrderCreate) SetOrderNum(i int) *OrderCreate {
	oc.mutation.SetOrderNum(i)
	return oc
}

// SetComment sets the "comment" field.
func (oc *OrderCreate) SetComment(s string) *OrderCreate {
	oc.mutation.SetComment(s)
	return oc
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (oc *OrderCreate) SetNillableComment(s *string) *OrderCreate {
	if s != nil {
		oc.SetComment(*s)
	}
	return oc
}

// SetSellerComment sets the "seller_comment" field.
func (oc *OrderCreate) SetSellerComment(s string) *OrderCreate {
	oc.mutation.SetSellerComment(s)
	return oc
}

// SetNillableSellerComment sets the "seller_comment" field if the given value is not nil.
func (oc *OrderCreate) SetNillableSellerComment(s *string) *OrderCreate {
	if s != nil {
		oc.SetSellerComment(*s)
	}
	return oc
}

// SetHasName sets the "has_name" field.
func (oc *OrderCreate) SetHasName(b bool) *OrderCreate {
	oc.mutation.SetHasName(b)
	return oc
}

// SetNillableHasName sets the "has_name" field if the given value is not nil.
func (oc *OrderCreate) SetNillableHasName(b *bool) *OrderCreate {
	if b != nil {
		oc.SetHasName(*b)
	}
	return oc
}

// SetIsExtra sets the "is_extra" field.
func (oc *OrderCreate) SetIsExtra(b bool) *OrderCreate {
	oc.mutation.SetIsExtra(b)
	return oc
}

// SetNillableIsExtra sets the "is_extra" field if the given value is not nil.
func (oc *OrderCreate) SetNillableIsExtra(b *bool) *OrderCreate {
	if b != nil {
		oc.SetIsExtra(*b)
	}
	return oc
}

// SetFb sets the "fb" field.
func (oc *OrderCreate) SetFb(b bool) *OrderCreate {
	oc.mutation.SetFb(b)
	return oc
}

// SetNillableFb sets the "fb" field if the given value is not nil.
func (oc *OrderCreate) SetNillableFb(b *bool) *OrderCreate {
	if b != nil {
		oc.SetFb(*b)
	}
	return oc
}

// SetIsInStock sets the "is_in_stock" field.
func (oc *OrderCreate) SetIsInStock(b bool) *OrderCreate {
	oc.mutation.SetIsInStock(b)
	return oc
}

// SetNillableIsInStock sets the "is_in_stock" field if the given value is not nil.
func (oc *OrderCreate) SetNillableIsInStock(b *bool) *OrderCreate {
	if b != nil {
		oc.SetIsInStock(*b)
	}
	return oc
}

// SetStatus sets the "status" field.
func (oc *OrderCreate) SetStatus(o order.Status) *OrderCreate {
	oc.mutation.SetStatus(o)
	return oc
}

// SetCreatedAt sets the "created_at" field.
func (oc *OrderCreate) SetCreatedAt(t time.Time) *OrderCreate {
	oc.mutation.SetCreatedAt(t)
	return oc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableCreatedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetCreatedAt(*t)
	}
	return oc
}

// SetUpdatedAt sets the "updated_at" field.
func (oc *OrderCreate) SetUpdatedAt(t time.Time) *OrderCreate {
	oc.mutation.SetUpdatedAt(t)
	return oc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oc *OrderCreate) SetNillableUpdatedAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetUpdatedAt(*t)
	}
	return oc
}

// SetID sets the "id" field.
func (oc *OrderCreate) SetID(s string) *OrderCreate {
	oc.mutation.SetID(s)
	return oc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (oc *OrderCreate) SetNillableID(s *string) *OrderCreate {
	if s != nil {
		oc.SetID(*s)
	}
	return oc
}

// SetUser sets the "user" edge to the User entity.
func (oc *OrderCreate) SetUser(u *User) *OrderCreate {
	return oc.SetUserID(u.ID)
}

// SetPost sets the "post" edge to the Post entity.
func (oc *OrderCreate) SetPost(p *Post) *OrderCreate {
	return oc.SetPostID(p.ID)
}

// AddOrderItemIDs adds the "order_items" edge to the OrderItem entity by IDs.
func (oc *OrderCreate) AddOrderItemIDs(ids ...string) *OrderCreate {
	oc.mutation.AddOrderItemIDs(ids...)
	return oc
}

// AddOrderItems adds the "order_items" edges to the OrderItem entity.
func (oc *OrderCreate) AddOrderItems(o ...*OrderItem) *OrderCreate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddOrderItemIDs(ids...)
}

// AddOrderHistoryIDs adds the "order_histories" edge to the OrderHistory entity by IDs.
func (oc *OrderCreate) AddOrderHistoryIDs(ids ...string) *OrderCreate {
	oc.mutation.AddOrderHistoryIDs(ids...)
	return oc
}

// AddOrderHistories adds the "order_histories" edges to the OrderHistory entity.
func (oc *OrderCreate) AddOrderHistories(o ...*OrderHistory) *OrderCreate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddOrderHistoryIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}

// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
	oc.defaults()
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrderCreate) SaveX(ctx context.Context) *Order {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrderCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrderCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrderCreate) defaults() {
	if _, ok := oc.mutation.Comment(); !ok {
		v := order.DefaultComment
		oc.mutation.SetComment(v)
	}
	if _, ok := oc.mutation.SellerComment(); !ok {
		v := order.DefaultSellerComment
		oc.mutation.SetSellerComment(v)
	}
	if _, ok := oc.mutation.HasName(); !ok {
		v := order.DefaultHasName
		oc.mutation.SetHasName(v)
	}
	if _, ok := oc.mutation.IsExtra(); !ok {
		v := order.DefaultIsExtra
		oc.mutation.SetIsExtra(v)
	}
	if _, ok := oc.mutation.Fb(); !ok {
		v := order.DefaultFb
		oc.mutation.SetFb(v)
	}
	if _, ok := oc.mutation.IsInStock(); !ok {
		v := order.DefaultIsInStock
		oc.mutation.SetIsInStock(v)
	}
	if _, ok := oc.mutation.CreatedAt(); !ok {
		v := order.DefaultCreatedAt()
		oc.mutation.SetCreatedAt(v)
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		v := order.DefaultUpdatedAt()
		oc.mutation.SetUpdatedAt(v)
	}
	if _, ok := oc.mutation.ID(); !ok {
		v := order.DefaultID()
		oc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrderCreate) check() error {
	if _, ok := oc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Order.user_id"`)}
	}
	if _, ok := oc.mutation.PostID(); !ok {
		return &ValidationError{Name: "post_id", err: errors.New(`ent: missing required field "Order.post_id"`)}
	}
	if _, ok := oc.mutation.OrderNum(); !ok {
		return &ValidationError{Name: "order_num", err: errors.New(`ent: missing required field "Order.order_num"`)}
	}
	if _, ok := oc.mutation.Comment(); !ok {
		return &ValidationError{Name: "comment", err: errors.New(`ent: missing required field "Order.comment"`)}
	}
	if _, ok := oc.mutation.SellerComment(); !ok {
		return &ValidationError{Name: "seller_comment", err: errors.New(`ent: missing required field "Order.seller_comment"`)}
	}
	if _, ok := oc.mutation.HasName(); !ok {
		return &ValidationError{Name: "has_name", err: errors.New(`ent: missing required field "Order.has_name"`)}
	}
	if _, ok := oc.mutation.IsExtra(); !ok {
		return &ValidationError{Name: "is_extra", err: errors.New(`ent: missing required field "Order.is_extra"`)}
	}
	if _, ok := oc.mutation.Fb(); !ok {
		return &ValidationError{Name: "fb", err: errors.New(`ent: missing required field "Order.fb"`)}
	}
	if _, ok := oc.mutation.IsInStock(); !ok {
		return &ValidationError{Name: "is_in_stock", err: errors.New(`ent: missing required field "Order.is_in_stock"`)}
	}
	if _, ok := oc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Order.status"`)}
	}
	if v, ok := oc.mutation.Status(); ok {
		if err := order.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Order.status": %w`, err)}
		}
	}
	if _, ok := oc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Order.created_at"`)}
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Order.updated_at"`)}
	}
	if _, ok := oc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Order.user"`)}
	}
	if _, ok := oc.mutation.PostID(); !ok {
		return &ValidationError{Name: "post", err: errors.New(`ent: missing required edge "Order.post"`)}
	}
	return nil
}

func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Order.ID type: %T", _spec.ID.Value)
		}
	}
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		_node = &Order{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(order.Table, sqlgraph.NewFieldSpec(order.FieldID, field.TypeString))
	)
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oc.mutation.OrderNum(); ok {
		_spec.SetField(order.FieldOrderNum, field.TypeInt, value)
		_node.OrderNum = &value
	}
	if value, ok := oc.mutation.Comment(); ok {
		_spec.SetField(order.FieldComment, field.TypeString, value)
		_node.Comment = &value
	}
	if value, ok := oc.mutation.SellerComment(); ok {
		_spec.SetField(order.FieldSellerComment, field.TypeString, value)
		_node.SellerComment = &value
	}
	if value, ok := oc.mutation.HasName(); ok {
		_spec.SetField(order.FieldHasName, field.TypeBool, value)
		_node.HasName = &value
	}
	if value, ok := oc.mutation.IsExtra(); ok {
		_spec.SetField(order.FieldIsExtra, field.TypeBool, value)
		_node.IsExtra = &value
	}
	if value, ok := oc.mutation.Fb(); ok {
		_spec.SetField(order.FieldFb, field.TypeBool, value)
		_node.Fb = &value
	}
	if value, ok := oc.mutation.IsInStock(); ok {
		_spec.SetField(order.FieldIsInStock, field.TypeBool, value)
		_node.IsInStock = &value
	}
	if value, ok := oc.mutation.Status(); ok {
		_spec.SetField(order.FieldStatus, field.TypeEnum, value)
		_node.Status = &value
	}
	if value, ok := oc.mutation.CreatedAt(); ok {
		_spec.SetField(order.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := oc.mutation.UpdatedAt(); ok {
		_spec.SetField(order.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if nodes := oc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.UserTable,
			Columns: []string{order.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.PostTable,
			Columns: []string{order.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PostID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.OrderItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderItemsTable,
			Columns: []string{order.OrderItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.OrderHistoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.OrderHistoriesTable,
			Columns: []string{order.OrderHistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderhistory.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderCreateBulk is the builder for creating many Order entities in bulk.
type OrderCreateBulk struct {
	config
	builders []*OrderCreate
}

// Save creates the Order entities in the database.
func (ocb *OrderCreateBulk) Save(ctx context.Context) ([]*Order, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Order, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrderCreateBulk) SaveX(ctx context.Context) []*Order {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrderCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrderCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
