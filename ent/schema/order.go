package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

var OrderStatusField = field.Enum("status").
	Values("ordered", "confirmed", "delivered", "completed", "missing", "canceled").
	Nillable()

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}),
		field.String("user_id").Nillable(),
		field.String("post_id").Nillable(),
		field.Int("order_num").Nillable(),
		field.String("comment").Default("").Nillable(),
		field.String("seller_comment").Default("").Nillable(),
		field.Bool("has_name").Default(false).Nillable(),
		field.Bool("is_extra").Default(false).Nillable(),
		field.Bool("fb").Default(false).Nillable(),
		field.Bool("is_in_stock").Default(false).Nillable(),
		OrderStatusField,
		field.Time("created_at").Default(time.Now).Nillable().Nillable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Nillable(),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_orders").Unique().Field("user_id").Required(),
		edge.From("post", Post.Type).Ref("post_orders").Unique().Field("post_id").Required(),
		edge.To("order_items", OrderItem.Type),
		edge.To("order_histories", OrderHistory.Type),
	}
}

func (Order) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("post_id"),
		index.Fields("created_at"),
	}
}
