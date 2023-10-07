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
	"github.com/billc-dev/tuango-go/ent/deliver"
	"github.com/billc-dev/tuango-go/ent/like"
	"github.com/billc-dev/tuango-go/ent/order"
	"github.com/billc-dev/tuango-go/ent/post"
	"github.com/billc-dev/tuango-go/ent/postitem"
	"github.com/billc-dev/tuango-go/ent/schema"
	"github.com/billc-dev/tuango-go/ent/user"
)

// PostCreate is the builder for creating a Post entity.
type PostCreate struct {
	config
	mutation *PostMutation
	hooks    []Hook
}

// SetSellerID sets the "seller_id" field.
func (pc *PostCreate) SetSellerID(s string) *PostCreate {
	pc.mutation.SetSellerID(s)
	return pc
}

// SetPostNum sets the "post_num" field.
func (pc *PostCreate) SetPostNum(i int) *PostCreate {
	pc.mutation.SetPostNum(i)
	return pc
}

// SetTitle sets the "title" field.
func (pc *PostCreate) SetTitle(s string) *PostCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetBody sets the "body" field.
func (pc *PostCreate) SetBody(s string) *PostCreate {
	pc.mutation.SetBody(s)
	return pc
}

// SetDeadline sets the "deadline" field.
func (pc *PostCreate) SetDeadline(s string) *PostCreate {
	pc.mutation.SetDeadline(s)
	return pc
}

// SetNillableDeadline sets the "deadline" field if the given value is not nil.
func (pc *PostCreate) SetNillableDeadline(s *string) *PostCreate {
	if s != nil {
		pc.SetDeadline(*s)
	}
	return pc
}

// SetDeliveryDate sets the "delivery_date" field.
func (pc *PostCreate) SetDeliveryDate(s string) *PostCreate {
	pc.mutation.SetDeliveryDate(s)
	return pc
}

// SetNillableDeliveryDate sets the "delivery_date" field if the given value is not nil.
func (pc *PostCreate) SetNillableDeliveryDate(s *string) *PostCreate {
	if s != nil {
		pc.SetDeliveryDate(*s)
	}
	return pc
}

// SetLikeCount sets the "like_count" field.
func (pc *PostCreate) SetLikeCount(i int) *PostCreate {
	pc.mutation.SetLikeCount(i)
	return pc
}

// SetNillableLikeCount sets the "like_count" field if the given value is not nil.
func (pc *PostCreate) SetNillableLikeCount(i *int) *PostCreate {
	if i != nil {
		pc.SetLikeCount(*i)
	}
	return pc
}

// SetCommentCount sets the "comment_count" field.
func (pc *PostCreate) SetCommentCount(i int) *PostCreate {
	pc.mutation.SetCommentCount(i)
	return pc
}

// SetNillableCommentCount sets the "comment_count" field if the given value is not nil.
func (pc *PostCreate) SetNillableCommentCount(i *int) *PostCreate {
	if i != nil {
		pc.SetCommentCount(*i)
	}
	return pc
}

// SetOrderCount sets the "order_count" field.
func (pc *PostCreate) SetOrderCount(i int) *PostCreate {
	pc.mutation.SetOrderCount(i)
	return pc
}

// SetNillableOrderCount sets the "order_count" field if the given value is not nil.
func (pc *PostCreate) SetNillableOrderCount(i *int) *PostCreate {
	if i != nil {
		pc.SetOrderCount(*i)
	}
	return pc
}

// SetImages sets the "images" field.
func (pc *PostCreate) SetImages(s []schema.Image) *PostCreate {
	pc.mutation.SetImages(s)
	return pc
}

// SetStorageType sets the "storage_type" field.
func (pc *PostCreate) SetStorageType(pt post.StorageType) *PostCreate {
	pc.mutation.SetStorageType(pt)
	return pc
}

// SetStatus sets the "status" field.
func (pc *PostCreate) SetStatus(po post.Status) *PostCreate {
	pc.mutation.SetStatus(po)
	return pc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pc *PostCreate) SetNillableStatus(po *post.Status) *PostCreate {
	if po != nil {
		pc.SetStatus(*po)
	}
	return pc
}

// SetComment sets the "comment" field.
func (pc *PostCreate) SetComment(s string) *PostCreate {
	pc.mutation.SetComment(s)
	return pc
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (pc *PostCreate) SetNillableComment(s *string) *PostCreate {
	if s != nil {
		pc.SetComment(*s)
	}
	return pc
}

// SetDelivered sets the "delivered" field.
func (pc *PostCreate) SetDelivered(b bool) *PostCreate {
	pc.mutation.SetDelivered(b)
	return pc
}

// SetNillableDelivered sets the "delivered" field if the given value is not nil.
func (pc *PostCreate) SetNillableDelivered(b *bool) *PostCreate {
	if b != nil {
		pc.SetDelivered(*b)
	}
	return pc
}

// SetIsInStock sets the "is_in_stock" field.
func (pc *PostCreate) SetIsInStock(b bool) *PostCreate {
	pc.mutation.SetIsInStock(b)
	return pc
}

// SetNillableIsInStock sets the "is_in_stock" field if the given value is not nil.
func (pc *PostCreate) SetNillableIsInStock(b *bool) *PostCreate {
	if b != nil {
		pc.SetIsInStock(*b)
	}
	return pc
}

// SetNormalTotal sets the "normal_total" field.
func (pc *PostCreate) SetNormalTotal(f float64) *PostCreate {
	pc.mutation.SetNormalTotal(f)
	return pc
}

// SetNillableNormalTotal sets the "normal_total" field if the given value is not nil.
func (pc *PostCreate) SetNillableNormalTotal(f *float64) *PostCreate {
	if f != nil {
		pc.SetNormalTotal(*f)
	}
	return pc
}

// SetNormalFee sets the "normal_fee" field.
func (pc *PostCreate) SetNormalFee(f float64) *PostCreate {
	pc.mutation.SetNormalFee(f)
	return pc
}

// SetNillableNormalFee sets the "normal_fee" field if the given value is not nil.
func (pc *PostCreate) SetNillableNormalFee(f *float64) *PostCreate {
	if f != nil {
		pc.SetNormalFee(*f)
	}
	return pc
}

// SetExtraTotal sets the "extra_total" field.
func (pc *PostCreate) SetExtraTotal(f float64) *PostCreate {
	pc.mutation.SetExtraTotal(f)
	return pc
}

// SetNillableExtraTotal sets the "extra_total" field if the given value is not nil.
func (pc *PostCreate) SetNillableExtraTotal(f *float64) *PostCreate {
	if f != nil {
		pc.SetExtraTotal(*f)
	}
	return pc
}

// SetExtraFee sets the "extra_fee" field.
func (pc *PostCreate) SetExtraFee(f float64) *PostCreate {
	pc.mutation.SetExtraFee(f)
	return pc
}

// SetNillableExtraFee sets the "extra_fee" field if the given value is not nil.
func (pc *PostCreate) SetNillableExtraFee(f *float64) *PostCreate {
	if f != nil {
		pc.SetExtraFee(*f)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PostCreate) SetCreatedAt(t time.Time) *PostCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableCreatedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PostCreate) SetUpdatedAt(t time.Time) *PostCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableUpdatedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PostCreate) SetID(s string) *PostCreate {
	pc.mutation.SetID(s)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PostCreate) SetNillableID(s *string) *PostCreate {
	if s != nil {
		pc.SetID(*s)
	}
	return pc
}

// SetSeller sets the "seller" edge to the User entity.
func (pc *PostCreate) SetSeller(u *User) *PostCreate {
	return pc.SetSellerID(u.ID)
}

// AddPostCommentIDs adds the "post_comments" edge to the Comment entity by IDs.
func (pc *PostCreate) AddPostCommentIDs(ids ...string) *PostCreate {
	pc.mutation.AddPostCommentIDs(ids...)
	return pc
}

// AddPostComments adds the "post_comments" edges to the Comment entity.
func (pc *PostCreate) AddPostComments(c ...*Comment) *PostCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddPostCommentIDs(ids...)
}

// AddPostDeliverIDs adds the "post_delivers" edge to the Deliver entity by IDs.
func (pc *PostCreate) AddPostDeliverIDs(ids ...string) *PostCreate {
	pc.mutation.AddPostDeliverIDs(ids...)
	return pc
}

// AddPostDelivers adds the "post_delivers" edges to the Deliver entity.
func (pc *PostCreate) AddPostDelivers(d ...*Deliver) *PostCreate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pc.AddPostDeliverIDs(ids...)
}

// AddPostItemIDs adds the "post_items" edge to the PostItem entity by IDs.
func (pc *PostCreate) AddPostItemIDs(ids ...string) *PostCreate {
	pc.mutation.AddPostItemIDs(ids...)
	return pc
}

// AddPostItems adds the "post_items" edges to the PostItem entity.
func (pc *PostCreate) AddPostItems(p ...*PostItem) *PostCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPostItemIDs(ids...)
}

// AddPostLikeIDs adds the "post_likes" edge to the Like entity by IDs.
func (pc *PostCreate) AddPostLikeIDs(ids ...string) *PostCreate {
	pc.mutation.AddPostLikeIDs(ids...)
	return pc
}

// AddPostLikes adds the "post_likes" edges to the Like entity.
func (pc *PostCreate) AddPostLikes(l ...*Like) *PostCreate {
	ids := make([]string, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return pc.AddPostLikeIDs(ids...)
}

// AddPostOrderIDs adds the "post_orders" edge to the Order entity by IDs.
func (pc *PostCreate) AddPostOrderIDs(ids ...string) *PostCreate {
	pc.mutation.AddPostOrderIDs(ids...)
	return pc
}

// AddPostOrders adds the "post_orders" edges to the Order entity.
func (pc *PostCreate) AddPostOrders(o ...*Order) *PostCreate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pc.AddPostOrderIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pc *PostCreate) Mutation() *PostMutation {
	return pc.mutation
}

// Save creates the Post in the database.
func (pc *PostCreate) Save(ctx context.Context) (*Post, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PostCreate) SaveX(ctx context.Context) *Post {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PostCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PostCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PostCreate) defaults() {
	if _, ok := pc.mutation.Deadline(); !ok {
		v := post.DefaultDeadline
		pc.mutation.SetDeadline(v)
	}
	if _, ok := pc.mutation.DeliveryDate(); !ok {
		v := post.DefaultDeliveryDate
		pc.mutation.SetDeliveryDate(v)
	}
	if _, ok := pc.mutation.LikeCount(); !ok {
		v := post.DefaultLikeCount
		pc.mutation.SetLikeCount(v)
	}
	if _, ok := pc.mutation.CommentCount(); !ok {
		v := post.DefaultCommentCount
		pc.mutation.SetCommentCount(v)
	}
	if _, ok := pc.mutation.OrderCount(); !ok {
		v := post.DefaultOrderCount
		pc.mutation.SetOrderCount(v)
	}
	if _, ok := pc.mutation.Status(); !ok {
		v := post.DefaultStatus
		pc.mutation.SetStatus(v)
	}
	if _, ok := pc.mutation.Comment(); !ok {
		v := post.DefaultComment
		pc.mutation.SetComment(v)
	}
	if _, ok := pc.mutation.Delivered(); !ok {
		v := post.DefaultDelivered
		pc.mutation.SetDelivered(v)
	}
	if _, ok := pc.mutation.IsInStock(); !ok {
		v := post.DefaultIsInStock
		pc.mutation.SetIsInStock(v)
	}
	if _, ok := pc.mutation.NormalTotal(); !ok {
		v := post.DefaultNormalTotal
		pc.mutation.SetNormalTotal(v)
	}
	if _, ok := pc.mutation.NormalFee(); !ok {
		v := post.DefaultNormalFee
		pc.mutation.SetNormalFee(v)
	}
	if _, ok := pc.mutation.ExtraTotal(); !ok {
		v := post.DefaultExtraTotal
		pc.mutation.SetExtraTotal(v)
	}
	if _, ok := pc.mutation.ExtraFee(); !ok {
		v := post.DefaultExtraFee
		pc.mutation.SetExtraFee(v)
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := post.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := post.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := post.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PostCreate) check() error {
	if _, ok := pc.mutation.SellerID(); !ok {
		return &ValidationError{Name: "seller_id", err: errors.New(`ent: missing required field "Post.seller_id"`)}
	}
	if _, ok := pc.mutation.PostNum(); !ok {
		return &ValidationError{Name: "post_num", err: errors.New(`ent: missing required field "Post.post_num"`)}
	}
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Post.title"`)}
	}
	if _, ok := pc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "Post.body"`)}
	}
	if _, ok := pc.mutation.Deadline(); !ok {
		return &ValidationError{Name: "deadline", err: errors.New(`ent: missing required field "Post.deadline"`)}
	}
	if v, ok := pc.mutation.Deadline(); ok {
		if err := post.DeadlineValidator(v); err != nil {
			return &ValidationError{Name: "deadline", err: fmt.Errorf(`ent: validator failed for field "Post.deadline": %w`, err)}
		}
	}
	if _, ok := pc.mutation.DeliveryDate(); !ok {
		return &ValidationError{Name: "delivery_date", err: errors.New(`ent: missing required field "Post.delivery_date"`)}
	}
	if v, ok := pc.mutation.DeliveryDate(); ok {
		if err := post.DeliveryDateValidator(v); err != nil {
			return &ValidationError{Name: "delivery_date", err: fmt.Errorf(`ent: validator failed for field "Post.delivery_date": %w`, err)}
		}
	}
	if _, ok := pc.mutation.LikeCount(); !ok {
		return &ValidationError{Name: "like_count", err: errors.New(`ent: missing required field "Post.like_count"`)}
	}
	if _, ok := pc.mutation.CommentCount(); !ok {
		return &ValidationError{Name: "comment_count", err: errors.New(`ent: missing required field "Post.comment_count"`)}
	}
	if _, ok := pc.mutation.OrderCount(); !ok {
		return &ValidationError{Name: "order_count", err: errors.New(`ent: missing required field "Post.order_count"`)}
	}
	if _, ok := pc.mutation.Images(); !ok {
		return &ValidationError{Name: "images", err: errors.New(`ent: missing required field "Post.images"`)}
	}
	if _, ok := pc.mutation.StorageType(); !ok {
		return &ValidationError{Name: "storage_type", err: errors.New(`ent: missing required field "Post.storage_type"`)}
	}
	if v, ok := pc.mutation.StorageType(); ok {
		if err := post.StorageTypeValidator(v); err != nil {
			return &ValidationError{Name: "storage_type", err: fmt.Errorf(`ent: validator failed for field "Post.storage_type": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Post.status"`)}
	}
	if v, ok := pc.mutation.Status(); ok {
		if err := post.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Post.status": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Comment(); !ok {
		return &ValidationError{Name: "comment", err: errors.New(`ent: missing required field "Post.comment"`)}
	}
	if _, ok := pc.mutation.Delivered(); !ok {
		return &ValidationError{Name: "delivered", err: errors.New(`ent: missing required field "Post.delivered"`)}
	}
	if _, ok := pc.mutation.IsInStock(); !ok {
		return &ValidationError{Name: "is_in_stock", err: errors.New(`ent: missing required field "Post.is_in_stock"`)}
	}
	if _, ok := pc.mutation.NormalTotal(); !ok {
		return &ValidationError{Name: "normal_total", err: errors.New(`ent: missing required field "Post.normal_total"`)}
	}
	if _, ok := pc.mutation.NormalFee(); !ok {
		return &ValidationError{Name: "normal_fee", err: errors.New(`ent: missing required field "Post.normal_fee"`)}
	}
	if _, ok := pc.mutation.ExtraTotal(); !ok {
		return &ValidationError{Name: "extra_total", err: errors.New(`ent: missing required field "Post.extra_total"`)}
	}
	if _, ok := pc.mutation.ExtraFee(); !ok {
		return &ValidationError{Name: "extra_fee", err: errors.New(`ent: missing required field "Post.extra_fee"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Post.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Post.updated_at"`)}
	}
	if _, ok := pc.mutation.SellerID(); !ok {
		return &ValidationError{Name: "seller", err: errors.New(`ent: missing required edge "Post.seller"`)}
	}
	return nil
}

func (pc *PostCreate) sqlSave(ctx context.Context) (*Post, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Post.ID type: %T", _spec.ID.Value)
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PostCreate) createSpec() (*Post, *sqlgraph.CreateSpec) {
	var (
		_node = &Post{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(post.Table, sqlgraph.NewFieldSpec(post.FieldID, field.TypeString))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.PostNum(); ok {
		_spec.SetField(post.FieldPostNum, field.TypeInt, value)
		_node.PostNum = &value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
		_node.Title = &value
	}
	if value, ok := pc.mutation.Body(); ok {
		_spec.SetField(post.FieldBody, field.TypeString, value)
		_node.Body = &value
	}
	if value, ok := pc.mutation.Deadline(); ok {
		_spec.SetField(post.FieldDeadline, field.TypeString, value)
		_node.Deadline = &value
	}
	if value, ok := pc.mutation.DeliveryDate(); ok {
		_spec.SetField(post.FieldDeliveryDate, field.TypeString, value)
		_node.DeliveryDate = &value
	}
	if value, ok := pc.mutation.LikeCount(); ok {
		_spec.SetField(post.FieldLikeCount, field.TypeInt, value)
		_node.LikeCount = &value
	}
	if value, ok := pc.mutation.CommentCount(); ok {
		_spec.SetField(post.FieldCommentCount, field.TypeInt, value)
		_node.CommentCount = &value
	}
	if value, ok := pc.mutation.OrderCount(); ok {
		_spec.SetField(post.FieldOrderCount, field.TypeInt, value)
		_node.OrderCount = &value
	}
	if value, ok := pc.mutation.Images(); ok {
		_spec.SetField(post.FieldImages, field.TypeJSON, value)
		_node.Images = value
	}
	if value, ok := pc.mutation.StorageType(); ok {
		_spec.SetField(post.FieldStorageType, field.TypeEnum, value)
		_node.StorageType = &value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(post.FieldStatus, field.TypeEnum, value)
		_node.Status = &value
	}
	if value, ok := pc.mutation.Comment(); ok {
		_spec.SetField(post.FieldComment, field.TypeString, value)
		_node.Comment = &value
	}
	if value, ok := pc.mutation.Delivered(); ok {
		_spec.SetField(post.FieldDelivered, field.TypeBool, value)
		_node.Delivered = &value
	}
	if value, ok := pc.mutation.IsInStock(); ok {
		_spec.SetField(post.FieldIsInStock, field.TypeBool, value)
		_node.IsInStock = &value
	}
	if value, ok := pc.mutation.NormalTotal(); ok {
		_spec.SetField(post.FieldNormalTotal, field.TypeFloat64, value)
		_node.NormalTotal = &value
	}
	if value, ok := pc.mutation.NormalFee(); ok {
		_spec.SetField(post.FieldNormalFee, field.TypeFloat64, value)
		_node.NormalFee = &value
	}
	if value, ok := pc.mutation.ExtraTotal(); ok {
		_spec.SetField(post.FieldExtraTotal, field.TypeFloat64, value)
		_node.ExtraTotal = &value
	}
	if value, ok := pc.mutation.ExtraFee(); ok {
		_spec.SetField(post.FieldExtraFee, field.TypeFloat64, value)
		_node.ExtraFee = &value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = &value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if nodes := pc.mutation.SellerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.SellerTable,
			Columns: []string{post.SellerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SellerID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PostCommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.PostCommentsTable,
			Columns: []string{post.PostCommentsColumn},
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
	if nodes := pc.mutation.PostDeliversIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.PostDeliversTable,
			Columns: []string{post.PostDeliversColumn},
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
	if nodes := pc.mutation.PostItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.PostItemsTable,
			Columns: []string{post.PostItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(postitem.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PostLikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.PostLikesTable,
			Columns: []string{post.PostLikesColumn},
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
	if nodes := pc.mutation.PostOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   post.PostOrdersTable,
			Columns: []string{post.PostOrdersColumn},
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
	return _node, _spec
}

// PostCreateBulk is the builder for creating many Post entities in bulk.
type PostCreateBulk struct {
	config
	err      error
	builders []*PostCreate
}

// Save creates the Post entities in the database.
func (pcb *PostCreateBulk) Save(ctx context.Context) ([]*Post, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Post, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PostCreateBulk) SaveX(ctx context.Context) []*Post {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PostCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PostCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
