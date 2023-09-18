// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/billc-dev/tuango-go/ent/post"
	"github.com/billc-dev/tuango-go/ent/schema"
	"github.com/billc-dev/tuango-go/ent/user"
)

// Post is the model entity for the Post schema.
type Post struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// SellerID holds the value of the "seller_id" field.
	SellerID *string `json:"seller_id,omitempty"`
	// PostNum holds the value of the "post_num" field.
	PostNum *int `json:"post_num,omitempty"`
	// Title holds the value of the "title" field.
	Title *string `json:"title,omitempty"`
	// Body holds the value of the "body" field.
	Body *string `json:"body,omitempty"`
	// Deadline holds the value of the "deadline" field.
	Deadline *string `json:"deadline,omitempty"`
	// DeliveryDate holds the value of the "delivery_date" field.
	DeliveryDate *string `json:"delivery_date,omitempty"`
	// LikeCount holds the value of the "like_count" field.
	LikeCount *int `json:"like_count,omitempty"`
	// CommentCount holds the value of the "comment_count" field.
	CommentCount *int `json:"comment_count,omitempty"`
	// OrderCount holds the value of the "order_count" field.
	OrderCount *int `json:"order_count,omitempty"`
	// Images holds the value of the "images" field.
	Images []schema.Image `json:"images,omitempty"`
	// StorageType holds the value of the "storage_type" field.
	StorageType *post.StorageType `json:"storage_type,omitempty"`
	// Status holds the value of the "status" field.
	Status *post.Status `json:"status,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment *string `json:"comment,omitempty"`
	// Delivered holds the value of the "delivered" field.
	Delivered *bool `json:"delivered,omitempty"`
	// IsInStock holds the value of the "is_in_stock" field.
	IsInStock *bool `json:"is_in_stock,omitempty"`
	// NormalTotal holds the value of the "normal_total" field.
	NormalTotal *float64 `json:"normal_total,omitempty"`
	// NormalFee holds the value of the "normal_fee" field.
	NormalFee *float64 `json:"normal_fee,omitempty"`
	// ExtraTotal holds the value of the "extra_total" field.
	ExtraTotal *float64 `json:"extra_total,omitempty"`
	// ExtraFee holds the value of the "extra_fee" field.
	ExtraFee *float64 `json:"extra_fee,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PostQuery when eager-loading is set.
	Edges        PostEdges `json:"-"`
	selectValues sql.SelectValues
}

// PostEdges holds the relations/edges for other nodes in the graph.
type PostEdges struct {
	// Seller holds the value of the seller edge.
	Seller *User `json:"seller,omitempty"`
	// PostComments holds the value of the post_comments edge.
	PostComments []*Comment `json:"post_comments,omitempty"`
	// PostDelivers holds the value of the post_delivers edge.
	PostDelivers []*Deliver `json:"post_delivers,omitempty"`
	// PostItems holds the value of the post_items edge.
	PostItems []*PostItem `json:"post_items,omitempty"`
	// PostLikes holds the value of the post_likes edge.
	PostLikes []*Like `json:"post_likes,omitempty"`
	// PostOrders holds the value of the post_orders edge.
	PostOrders []*Order `json:"post_orders,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// SellerOrErr returns the Seller value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) SellerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Seller == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Seller, nil
	}
	return nil, &NotLoadedError{edge: "seller"}
}

// PostCommentsOrErr returns the PostComments value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) PostCommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[1] {
		return e.PostComments, nil
	}
	return nil, &NotLoadedError{edge: "post_comments"}
}

// PostDeliversOrErr returns the PostDelivers value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) PostDeliversOrErr() ([]*Deliver, error) {
	if e.loadedTypes[2] {
		return e.PostDelivers, nil
	}
	return nil, &NotLoadedError{edge: "post_delivers"}
}

// PostItemsOrErr returns the PostItems value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) PostItemsOrErr() ([]*PostItem, error) {
	if e.loadedTypes[3] {
		return e.PostItems, nil
	}
	return nil, &NotLoadedError{edge: "post_items"}
}

// PostLikesOrErr returns the PostLikes value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) PostLikesOrErr() ([]*Like, error) {
	if e.loadedTypes[4] {
		return e.PostLikes, nil
	}
	return nil, &NotLoadedError{edge: "post_likes"}
}

// PostOrdersOrErr returns the PostOrders value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) PostOrdersOrErr() ([]*Order, error) {
	if e.loadedTypes[5] {
		return e.PostOrders, nil
	}
	return nil, &NotLoadedError{edge: "post_orders"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Post) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case post.FieldImages:
			values[i] = new([]byte)
		case post.FieldDelivered, post.FieldIsInStock:
			values[i] = new(sql.NullBool)
		case post.FieldNormalTotal, post.FieldNormalFee, post.FieldExtraTotal, post.FieldExtraFee:
			values[i] = new(sql.NullFloat64)
		case post.FieldPostNum, post.FieldLikeCount, post.FieldCommentCount, post.FieldOrderCount:
			values[i] = new(sql.NullInt64)
		case post.FieldID, post.FieldSellerID, post.FieldTitle, post.FieldBody, post.FieldDeadline, post.FieldDeliveryDate, post.FieldStorageType, post.FieldStatus, post.FieldComment:
			values[i] = new(sql.NullString)
		case post.FieldCreatedAt, post.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Post fields.
func (po *Post) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case post.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				po.ID = value.String
			}
		case post.FieldSellerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field seller_id", values[i])
			} else if value.Valid {
				po.SellerID = new(string)
				*po.SellerID = value.String
			}
		case post.FieldPostNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field post_num", values[i])
			} else if value.Valid {
				po.PostNum = new(int)
				*po.PostNum = int(value.Int64)
			}
		case post.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				po.Title = new(string)
				*po.Title = value.String
			}
		case post.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				po.Body = new(string)
				*po.Body = value.String
			}
		case post.FieldDeadline:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deadline", values[i])
			} else if value.Valid {
				po.Deadline = new(string)
				*po.Deadline = value.String
			}
		case post.FieldDeliveryDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field delivery_date", values[i])
			} else if value.Valid {
				po.DeliveryDate = new(string)
				*po.DeliveryDate = value.String
			}
		case post.FieldLikeCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field like_count", values[i])
			} else if value.Valid {
				po.LikeCount = new(int)
				*po.LikeCount = int(value.Int64)
			}
		case post.FieldCommentCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field comment_count", values[i])
			} else if value.Valid {
				po.CommentCount = new(int)
				*po.CommentCount = int(value.Int64)
			}
		case post.FieldOrderCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_count", values[i])
			} else if value.Valid {
				po.OrderCount = new(int)
				*po.OrderCount = int(value.Int64)
			}
		case post.FieldImages:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field images", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &po.Images); err != nil {
					return fmt.Errorf("unmarshal field images: %w", err)
				}
			}
		case post.FieldStorageType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field storage_type", values[i])
			} else if value.Valid {
				po.StorageType = new(post.StorageType)
				*po.StorageType = post.StorageType(value.String)
			}
		case post.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				po.Status = new(post.Status)
				*po.Status = post.Status(value.String)
			}
		case post.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				po.Comment = new(string)
				*po.Comment = value.String
			}
		case post.FieldDelivered:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field delivered", values[i])
			} else if value.Valid {
				po.Delivered = new(bool)
				*po.Delivered = value.Bool
			}
		case post.FieldIsInStock:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_in_stock", values[i])
			} else if value.Valid {
				po.IsInStock = new(bool)
				*po.IsInStock = value.Bool
			}
		case post.FieldNormalTotal:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field normal_total", values[i])
			} else if value.Valid {
				po.NormalTotal = new(float64)
				*po.NormalTotal = value.Float64
			}
		case post.FieldNormalFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field normal_fee", values[i])
			} else if value.Valid {
				po.NormalFee = new(float64)
				*po.NormalFee = value.Float64
			}
		case post.FieldExtraTotal:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field extra_total", values[i])
			} else if value.Valid {
				po.ExtraTotal = new(float64)
				*po.ExtraTotal = value.Float64
			}
		case post.FieldExtraFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field extra_fee", values[i])
			} else if value.Valid {
				po.ExtraFee = new(float64)
				*po.ExtraFee = value.Float64
			}
		case post.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				po.CreatedAt = new(time.Time)
				*po.CreatedAt = value.Time
			}
		case post.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				po.UpdatedAt = new(time.Time)
				*po.UpdatedAt = value.Time
			}
		default:
			po.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Post.
// This includes values selected through modifiers, order, etc.
func (po *Post) Value(name string) (ent.Value, error) {
	return po.selectValues.Get(name)
}

// QuerySeller queries the "seller" edge of the Post entity.
func (po *Post) QuerySeller() *UserQuery {
	return NewPostClient(po.config).QuerySeller(po)
}

// QueryPostComments queries the "post_comments" edge of the Post entity.
func (po *Post) QueryPostComments() *CommentQuery {
	return NewPostClient(po.config).QueryPostComments(po)
}

// QueryPostDelivers queries the "post_delivers" edge of the Post entity.
func (po *Post) QueryPostDelivers() *DeliverQuery {
	return NewPostClient(po.config).QueryPostDelivers(po)
}

// QueryPostItems queries the "post_items" edge of the Post entity.
func (po *Post) QueryPostItems() *PostItemQuery {
	return NewPostClient(po.config).QueryPostItems(po)
}

// QueryPostLikes queries the "post_likes" edge of the Post entity.
func (po *Post) QueryPostLikes() *LikeQuery {
	return NewPostClient(po.config).QueryPostLikes(po)
}

// QueryPostOrders queries the "post_orders" edge of the Post entity.
func (po *Post) QueryPostOrders() *OrderQuery {
	return NewPostClient(po.config).QueryPostOrders(po)
}

// Update returns a builder for updating this Post.
// Note that you need to call Post.Unwrap() before calling this method if this Post
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Post) Update() *PostUpdateOne {
	return NewPostClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the Post entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Post) Unwrap() *Post {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Post is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Post) String() string {
	var builder strings.Builder
	builder.WriteString("Post(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	if v := po.SellerID; v != nil {
		builder.WriteString("seller_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := po.PostNum; v != nil {
		builder.WriteString("post_num=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := po.Title; v != nil {
		builder.WriteString("title=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := po.Body; v != nil {
		builder.WriteString("body=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := po.Deadline; v != nil {
		builder.WriteString("deadline=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := po.DeliveryDate; v != nil {
		builder.WriteString("delivery_date=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := po.LikeCount; v != nil {
		builder.WriteString("like_count=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := po.CommentCount; v != nil {
		builder.WriteString("comment_count=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := po.OrderCount; v != nil {
		builder.WriteString("order_count=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("images=")
	builder.WriteString(fmt.Sprintf("%v", po.Images))
	builder.WriteString(", ")
	if v := po.StorageType; v != nil {
		builder.WriteString("storage_type=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := po.Status; v != nil {
		builder.WriteString("status=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := po.Comment; v != nil {
		builder.WriteString("comment=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := po.Delivered; v != nil {
		builder.WriteString("delivered=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := po.IsInStock; v != nil {
		builder.WriteString("is_in_stock=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := po.NormalTotal; v != nil {
		builder.WriteString("normal_total=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := po.NormalFee; v != nil {
		builder.WriteString("normal_fee=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := po.ExtraTotal; v != nil {
		builder.WriteString("extra_total=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := po.ExtraFee; v != nil {
		builder.WriteString("extra_fee=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := po.CreatedAt; v != nil {
		builder.WriteString("created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := po.UpdatedAt; v != nil {
		builder.WriteString("updated_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (po *Post) MarshalJSON() ([]byte, error) {
	type Alias Post
	return json.Marshal(&struct {
		*Alias
		PostEdges
	}{
		Alias:     (*Alias)(po),
		PostEdges: po.Edges,
	})
}

// Posts is a parsable slice of Post.
type Posts []*Post