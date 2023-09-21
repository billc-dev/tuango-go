package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderHistory holds the schema definition for the OrderHistory entity.
type OrderHistory struct {
	ent.Schema
}

// Fields of the OrderHistory.
func (OrderHistory) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}),
		field.String("order_id").Nillable(),
		OrderStatusField,
		field.Time("created_at").Default(time.Now).Nillable(),
	}
}

// Edges of the OrderHistory.
func (OrderHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).Ref("order_histories").Unique().Field("order_id").Required(),
	}
}

func (OrderHistory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("order_id"),
		index.Fields("created_at"),
	}
}
