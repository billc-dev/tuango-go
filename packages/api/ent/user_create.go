// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/billc-dev/tuango-go/ent/comment"
	"github.com/billc-dev/tuango-go/ent/complete"
	"github.com/billc-dev/tuango-go/ent/deliver"
	"github.com/billc-dev/tuango-go/ent/like"
	"github.com/billc-dev/tuango-go/ent/message"
	"github.com/billc-dev/tuango-go/ent/order"
	"github.com/billc-dev/tuango-go/ent/post"
	"github.com/billc-dev/tuango-go/ent/roomuser"
	"github.com/billc-dev/tuango-go/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetDisplayName sets the "display_name" field.
func (uc *UserCreate) SetDisplayName(s string) *UserCreate {
	uc.mutation.SetDisplayName(s)
	return uc
}

// SetPictureURL sets the "picture_url" field.
func (uc *UserCreate) SetPictureURL(s string) *UserCreate {
	uc.mutation.SetPictureURL(s)
	return uc
}

// SetPickupNum sets the "pickup_num" field.
func (uc *UserCreate) SetPickupNum(f float64) *UserCreate {
	uc.mutation.SetPickupNum(f)
	return uc
}

// SetRole sets the "role" field.
func (uc *UserCreate) SetRole(u user.Role) *UserCreate {
	uc.mutation.SetRole(u)
	return uc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uc *UserCreate) SetNillableRole(u *user.Role) *UserCreate {
	if u != nil {
		uc.SetRole(*u)
	}
	return uc
}

// SetStatus sets the "status" field.
func (uc *UserCreate) SetStatus(u user.Status) *UserCreate {
	uc.mutation.SetStatus(u)
	return uc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uc *UserCreate) SetNillableStatus(u *user.Status) *UserCreate {
	if u != nil {
		uc.SetStatus(*u)
	}
	return uc
}

// SetNotified sets the "notified" field.
func (uc *UserCreate) SetNotified(b bool) *UserCreate {
	uc.mutation.SetNotified(b)
	return uc
}

// SetNillableNotified sets the "notified" field if the given value is not nil.
func (uc *UserCreate) SetNillableNotified(b *bool) *UserCreate {
	if b != nil {
		uc.SetNotified(*b)
	}
	return uc
}

// SetLinePay sets the "line_pay" field.
func (uc *UserCreate) SetLinePay(b bool) *UserCreate {
	uc.mutation.SetLinePay(b)
	return uc
}

// SetNillableLinePay sets the "line_pay" field if the given value is not nil.
func (uc *UserCreate) SetNillableLinePay(b *bool) *UserCreate {
	if b != nil {
		uc.SetLinePay(*b)
	}
	return uc
}

// SetFb sets the "fb" field.
func (uc *UserCreate) SetFb(b bool) *UserCreate {
	uc.mutation.SetFb(b)
	return uc
}

// SetNillableFb sets the "fb" field if the given value is not nil.
func (uc *UserCreate) SetNillableFb(b *bool) *UserCreate {
	if b != nil {
		uc.SetFb(*b)
	}
	return uc
}

// SetComment sets the "comment" field.
func (uc *UserCreate) SetComment(s string) *UserCreate {
	uc.mutation.SetComment(s)
	return uc
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (uc *UserCreate) SetNillableComment(s *string) *UserCreate {
	if s != nil {
		uc.SetComment(*s)
	}
	return uc
}

// SetDeliveredOrderCountLimit sets the "delivered_order_count_limit" field.
func (uc *UserCreate) SetDeliveredOrderCountLimit(i int) *UserCreate {
	uc.mutation.SetDeliveredOrderCountLimit(i)
	return uc
}

// SetNillableDeliveredOrderCountLimit sets the "delivered_order_count_limit" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeliveredOrderCountLimit(i *int) *UserCreate {
	if i != nil {
		uc.SetDeliveredOrderCountLimit(*i)
	}
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(s string) *UserCreate {
	uc.mutation.SetID(s)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserCreate) SetNillableID(s *string) *UserCreate {
	if s != nil {
		uc.SetID(*s)
	}
	return uc
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (uc *UserCreate) AddPostIDs(ids ...string) *UserCreate {
	uc.mutation.AddPostIDs(ids...)
	return uc
}

// AddPosts adds the "posts" edges to the Post entity.
func (uc *UserCreate) AddPosts(p ...*Post) *UserCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddPostIDs(ids...)
}

// AddUserCommentIDs adds the "user_comments" edge to the Comment entity by IDs.
func (uc *UserCreate) AddUserCommentIDs(ids ...string) *UserCreate {
	uc.mutation.AddUserCommentIDs(ids...)
	return uc
}

// AddUserComments adds the "user_comments" edges to the Comment entity.
func (uc *UserCreate) AddUserComments(c ...*Comment) *UserCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddUserCommentIDs(ids...)
}

// AddUserCompleteIDs adds the "user_completes" edge to the Complete entity by IDs.
func (uc *UserCreate) AddUserCompleteIDs(ids ...string) *UserCreate {
	uc.mutation.AddUserCompleteIDs(ids...)
	return uc
}

// AddUserCompletes adds the "user_completes" edges to the Complete entity.
func (uc *UserCreate) AddUserCompletes(c ...*Complete) *UserCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddUserCompleteIDs(ids...)
}

// AddUserDeliverIDs adds the "user_delivers" edge to the Deliver entity by IDs.
func (uc *UserCreate) AddUserDeliverIDs(ids ...string) *UserCreate {
	uc.mutation.AddUserDeliverIDs(ids...)
	return uc
}

// AddUserDelivers adds the "user_delivers" edges to the Deliver entity.
func (uc *UserCreate) AddUserDelivers(d ...*Deliver) *UserCreate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uc.AddUserDeliverIDs(ids...)
}

// AddUserLikeIDs adds the "user_likes" edge to the Like entity by IDs.
func (uc *UserCreate) AddUserLikeIDs(ids ...string) *UserCreate {
	uc.mutation.AddUserLikeIDs(ids...)
	return uc
}

// AddUserLikes adds the "user_likes" edges to the Like entity.
func (uc *UserCreate) AddUserLikes(l ...*Like) *UserCreate {
	ids := make([]string, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return uc.AddUserLikeIDs(ids...)
}

// AddUserMessageIDs adds the "user_messages" edge to the Message entity by IDs.
func (uc *UserCreate) AddUserMessageIDs(ids ...string) *UserCreate {
	uc.mutation.AddUserMessageIDs(ids...)
	return uc
}

// AddUserMessages adds the "user_messages" edges to the Message entity.
func (uc *UserCreate) AddUserMessages(m ...*Message) *UserCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uc.AddUserMessageIDs(ids...)
}

// AddUserOrderIDs adds the "user_orders" edge to the Order entity by IDs.
func (uc *UserCreate) AddUserOrderIDs(ids ...string) *UserCreate {
	uc.mutation.AddUserOrderIDs(ids...)
	return uc
}

// AddUserOrders adds the "user_orders" edges to the Order entity.
func (uc *UserCreate) AddUserOrders(o ...*Order) *UserCreate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return uc.AddUserOrderIDs(ids...)
}

// AddUserRoomIDs adds the "user_rooms" edge to the RoomUser entity by IDs.
func (uc *UserCreate) AddUserRoomIDs(ids ...string) *UserCreate {
	uc.mutation.AddUserRoomIDs(ids...)
	return uc
}

// AddUserRooms adds the "user_rooms" edges to the RoomUser entity.
func (uc *UserCreate) AddUserRooms(r ...*RoomUser) *UserCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddUserRoomIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Role(); !ok {
		v := user.DefaultRole
		uc.mutation.SetRole(v)
	}
	if _, ok := uc.mutation.Status(); !ok {
		v := user.DefaultStatus
		uc.mutation.SetStatus(v)
	}
	if _, ok := uc.mutation.Notified(); !ok {
		v := user.DefaultNotified
		uc.mutation.SetNotified(v)
	}
	if _, ok := uc.mutation.LinePay(); !ok {
		v := user.DefaultLinePay
		uc.mutation.SetLinePay(v)
	}
	if _, ok := uc.mutation.Fb(); !ok {
		v := user.DefaultFb
		uc.mutation.SetFb(v)
	}
	if _, ok := uc.mutation.Comment(); !ok {
		v := user.DefaultComment
		uc.mutation.SetComment(v)
	}
	if _, ok := uc.mutation.DeliveredOrderCountLimit(); !ok {
		v := user.DefaultDeliveredOrderCountLimit
		uc.mutation.SetDeliveredOrderCountLimit(v)
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "User.username"`)}
	}
	if _, ok := uc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`ent: missing required field "User.display_name"`)}
	}
	if _, ok := uc.mutation.PictureURL(); !ok {
		return &ValidationError{Name: "picture_url", err: errors.New(`ent: missing required field "User.picture_url"`)}
	}
	if _, ok := uc.mutation.PickupNum(); !ok {
		return &ValidationError{Name: "pickup_num", err: errors.New(`ent: missing required field "User.pickup_num"`)}
	}
	if _, ok := uc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "User.role"`)}
	}
	if v, ok := uc.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "User.role": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "User.status"`)}
	}
	if v, ok := uc.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "User.status": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Notified(); !ok {
		return &ValidationError{Name: "notified", err: errors.New(`ent: missing required field "User.notified"`)}
	}
	if _, ok := uc.mutation.LinePay(); !ok {
		return &ValidationError{Name: "line_pay", err: errors.New(`ent: missing required field "User.line_pay"`)}
	}
	if _, ok := uc.mutation.Fb(); !ok {
		return &ValidationError{Name: "fb", err: errors.New(`ent: missing required field "User.fb"`)}
	}
	if _, ok := uc.mutation.Comment(); !ok {
		return &ValidationError{Name: "comment", err: errors.New(`ent: missing required field "User.comment"`)}
	}
	if _, ok := uc.mutation.DeliveredOrderCountLimit(); !ok {
		return &ValidationError{Name: "delivered_order_count_limit", err: errors.New(`ent: missing required field "User.delivered_order_count_limit"`)}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "User.updated_at"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected User.ID type: %T", _spec.ID.Value)
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
		_node.Username = &value
	}
	if value, ok := uc.mutation.DisplayName(); ok {
		_spec.SetField(user.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = &value
	}
	if value, ok := uc.mutation.PictureURL(); ok {
		_spec.SetField(user.FieldPictureURL, field.TypeString, value)
		_node.PictureURL = &value
	}
	if value, ok := uc.mutation.PickupNum(); ok {
		_spec.SetField(user.FieldPickupNum, field.TypeFloat64, value)
		_node.PickupNum = &value
	}
	if value, ok := uc.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeEnum, value)
		_node.Role = &value
	}
	if value, ok := uc.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeEnum, value)
		_node.Status = &value
	}
	if value, ok := uc.mutation.Notified(); ok {
		_spec.SetField(user.FieldNotified, field.TypeBool, value)
		_node.Notified = &value
	}
	if value, ok := uc.mutation.LinePay(); ok {
		_spec.SetField(user.FieldLinePay, field.TypeBool, value)
		_node.LinePay = &value
	}
	if value, ok := uc.mutation.Fb(); ok {
		_spec.SetField(user.FieldFb, field.TypeBool, value)
		_node.Fb = &value
	}
	if value, ok := uc.mutation.Comment(); ok {
		_spec.SetField(user.FieldComment, field.TypeString, value)
		_node.Comment = &value
	}
	if value, ok := uc.mutation.DeliveredOrderCountLimit(); ok {
		_spec.SetField(user.FieldDeliveredOrderCountLimit, field.TypeInt, value)
		_node.DeliveredOrderCountLimit = &value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if nodes := uc.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PostsTable,
			Columns: []string{user.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserCommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserCommentsTable,
			Columns: []string{user.UserCommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserCompletesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserCompletesTable,
			Columns: []string{user.UserCompletesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(complete.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserDeliversIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserDeliversTable,
			Columns: []string{user.UserDeliversColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(deliver.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserLikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserLikesTable,
			Columns: []string{user.UserLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserMessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserMessagesTable,
			Columns: []string{user.UserMessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(message.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserOrdersTable,
			Columns: []string{user.UserOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.UserRoomsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserRoomsTable,
			Columns: []string{user.UserRoomsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(roomuser.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
