// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/billc-dev/tuango-go/ent/order"
	"github.com/billc-dev/tuango-go/ent/orderitem"
	"github.com/billc-dev/tuango-go/ent/postitem"
)

// OrderItemCreate is the builder for creating a OrderItem entity.
type OrderItemCreate struct {
	config
	mutation *OrderItemMutation
	hooks    []Hook
}

// SetOrderID sets the "order_id" field.
func (oic *OrderItemCreate) SetOrderID(s string) *OrderItemCreate {
	oic.mutation.SetOrderID(s)
	return oic
}

// SetPostItemID sets the "post_item_id" field.
func (oic *OrderItemCreate) SetPostItemID(s string) *OrderItemCreate {
	oic.mutation.SetPostItemID(s)
	return oic
}

// SetIdentifier sets the "identifier" field.
func (oic *OrderItemCreate) SetIdentifier(s string) *OrderItemCreate {
	oic.mutation.SetIdentifier(s)
	return oic
}

// SetName sets the "name" field.
func (oic *OrderItemCreate) SetName(s string) *OrderItemCreate {
	oic.mutation.SetName(s)
	return oic
}

// SetPrice sets the "price" field.
func (oic *OrderItemCreate) SetPrice(f float64) *OrderItemCreate {
	oic.mutation.SetPrice(f)
	return oic
}

// SetQty sets the "qty" field.
func (oic *OrderItemCreate) SetQty(f float64) *OrderItemCreate {
	oic.mutation.SetQty(f)
	return oic
}

// SetLocation sets the "location" field.
func (oic *OrderItemCreate) SetLocation(s string) *OrderItemCreate {
	oic.mutation.SetLocation(s)
	return oic
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableLocation(s *string) *OrderItemCreate {
	if s != nil {
		oic.SetLocation(*s)
	}
	return oic
}

// SetHasName sets the "has_name" field.
func (oic *OrderItemCreate) SetHasName(b bool) *OrderItemCreate {
	oic.mutation.SetHasName(b)
	return oic
}

// SetNillableHasName sets the "has_name" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableHasName(b *bool) *OrderItemCreate {
	if b != nil {
		oic.SetHasName(*b)
	}
	return oic
}

// SetStatus sets the "status" field.
func (oic *OrderItemCreate) SetStatus(o orderitem.Status) *OrderItemCreate {
	oic.mutation.SetStatus(o)
	return oic
}

// SetID sets the "id" field.
func (oic *OrderItemCreate) SetID(s string) *OrderItemCreate {
	oic.mutation.SetID(s)
	return oic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableID(s *string) *OrderItemCreate {
	if s != nil {
		oic.SetID(*s)
	}
	return oic
}

// SetOrder sets the "order" edge to the Order entity.
func (oic *OrderItemCreate) SetOrder(o *Order) *OrderItemCreate {
	return oic.SetOrderID(o.ID)
}

// SetPostItem sets the "post_item" edge to the PostItem entity.
func (oic *OrderItemCreate) SetPostItem(p *PostItem) *OrderItemCreate {
	return oic.SetPostItemID(p.ID)
}

// Mutation returns the OrderItemMutation object of the builder.
func (oic *OrderItemCreate) Mutation() *OrderItemMutation {
	return oic.mutation
}

// Save creates the OrderItem in the database.
func (oic *OrderItemCreate) Save(ctx context.Context) (*OrderItem, error) {
	oic.defaults()
	return withHooks(ctx, oic.sqlSave, oic.mutation, oic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oic *OrderItemCreate) SaveX(ctx context.Context) *OrderItem {
	v, err := oic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oic *OrderItemCreate) Exec(ctx context.Context) error {
	_, err := oic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oic *OrderItemCreate) ExecX(ctx context.Context) {
	if err := oic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oic *OrderItemCreate) defaults() {
	if _, ok := oic.mutation.Location(); !ok {
		v := orderitem.DefaultLocation
		oic.mutation.SetLocation(v)
	}
	if _, ok := oic.mutation.HasName(); !ok {
		v := orderitem.DefaultHasName
		oic.mutation.SetHasName(v)
	}
	if _, ok := oic.mutation.ID(); !ok {
		v := orderitem.DefaultID()
		oic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oic *OrderItemCreate) check() error {
	if _, ok := oic.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "OrderItem.order_id"`)}
	}
	if _, ok := oic.mutation.PostItemID(); !ok {
		return &ValidationError{Name: "post_item_id", err: errors.New(`ent: missing required field "OrderItem.post_item_id"`)}
	}
	if _, ok := oic.mutation.Identifier(); !ok {
		return &ValidationError{Name: "identifier", err: errors.New(`ent: missing required field "OrderItem.identifier"`)}
	}
	if _, ok := oic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "OrderItem.name"`)}
	}
	if _, ok := oic.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "OrderItem.price"`)}
	}
	if _, ok := oic.mutation.Qty(); !ok {
		return &ValidationError{Name: "qty", err: errors.New(`ent: missing required field "OrderItem.qty"`)}
	}
	if _, ok := oic.mutation.Location(); !ok {
		return &ValidationError{Name: "location", err: errors.New(`ent: missing required field "OrderItem.location"`)}
	}
	if _, ok := oic.mutation.HasName(); !ok {
		return &ValidationError{Name: "has_name", err: errors.New(`ent: missing required field "OrderItem.has_name"`)}
	}
	if _, ok := oic.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "OrderItem.status"`)}
	}
	if v, ok := oic.mutation.Status(); ok {
		if err := orderitem.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "OrderItem.status": %w`, err)}
		}
	}
	if _, ok := oic.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order", err: errors.New(`ent: missing required edge "OrderItem.order"`)}
	}
	if _, ok := oic.mutation.PostItemID(); !ok {
		return &ValidationError{Name: "post_item", err: errors.New(`ent: missing required edge "OrderItem.post_item"`)}
	}
	return nil
}

func (oic *OrderItemCreate) sqlSave(ctx context.Context) (*OrderItem, error) {
	if err := oic.check(); err != nil {
		return nil, err
	}
	_node, _spec := oic.createSpec()
	if err := sqlgraph.CreateNode(ctx, oic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected OrderItem.ID type: %T", _spec.ID.Value)
		}
	}
	oic.mutation.id = &_node.ID
	oic.mutation.done = true
	return _node, nil
}

func (oic *OrderItemCreate) createSpec() (*OrderItem, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderItem{config: oic.config}
		_spec = sqlgraph.NewCreateSpec(orderitem.Table, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeString))
	)
	if id, ok := oic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oic.mutation.Identifier(); ok {
		_spec.SetField(orderitem.FieldIdentifier, field.TypeString, value)
		_node.Identifier = &value
	}
	if value, ok := oic.mutation.Name(); ok {
		_spec.SetField(orderitem.FieldName, field.TypeString, value)
		_node.Name = &value
	}
	if value, ok := oic.mutation.Price(); ok {
		_spec.SetField(orderitem.FieldPrice, field.TypeFloat64, value)
		_node.Price = &value
	}
	if value, ok := oic.mutation.Qty(); ok {
		_spec.SetField(orderitem.FieldQty, field.TypeFloat64, value)
		_node.Qty = &value
	}
	if value, ok := oic.mutation.Location(); ok {
		_spec.SetField(orderitem.FieldLocation, field.TypeString, value)
		_node.Location = &value
	}
	if value, ok := oic.mutation.HasName(); ok {
		_spec.SetField(orderitem.FieldHasName, field.TypeBool, value)
		_node.HasName = &value
	}
	if value, ok := oic.mutation.Status(); ok {
		_spec.SetField(orderitem.FieldStatus, field.TypeEnum, value)
		_node.Status = &value
	}
	if nodes := oic.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: []string{orderitem.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OrderID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oic.mutation.PostItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.PostItemTable,
			Columns: []string{orderitem.PostItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(postitem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PostItemID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderItemCreateBulk is the builder for creating many OrderItem entities in bulk.
type OrderItemCreateBulk struct {
	config
	builders []*OrderItemCreate
}

// Save creates the OrderItem entities in the database.
func (oicb *OrderItemCreateBulk) Save(ctx context.Context) ([]*OrderItem, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oicb.builders))
	nodes := make([]*OrderItem, len(oicb.builders))
	mutators := make([]Mutator, len(oicb.builders))
	for i := range oicb.builders {
		func(i int, root context.Context) {
			builder := oicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderItemMutation)
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
					_, err = mutators[i+1].Mutate(root, oicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oicb *OrderItemCreateBulk) SaveX(ctx context.Context) []*OrderItem {
	v, err := oicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oicb *OrderItemCreateBulk) Exec(ctx context.Context) error {
	_, err := oicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oicb *OrderItemCreateBulk) ExecX(ctx context.Context) {
	if err := oicb.Exec(ctx); err != nil {
		panic(err)
	}
}
