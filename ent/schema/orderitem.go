package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderItem holds the schema definition for the OrderItem entity.
type OrderItem struct {
	ent.Schema
}

// Fields of the OrderItem.
func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}),
		field.String("order_id").Nillable(),
		field.String("post_item_id").Nillable(),
		field.String("identifier").Nillable(),
		field.String("name").Nillable(),
		field.Float("price").Nillable(),
		field.Float("qty").Nillable(),
		field.String("location").Default("").Nillable(),
		field.Bool("has_name").Default(false).Nillable(),
		OrderStatusField,
	}
}

// Edges of the OrderItem.
func (OrderItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).Ref("order_items").Unique().Field("order_id").Required(),
		edge.From("post_item", PostItem.Type).Ref("post_item").Unique().Field("post_item_id").Required(),
	}
}

func (OrderItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
	}
}
