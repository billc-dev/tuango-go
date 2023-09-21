package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Deliver holds the schema definition for the Deliver entity.
type Deliver struct {
	ent.Schema
}

type DeliverOrder struct {
	ID         string `json:"id"`
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Price      string `json:"price"`
	Amount     string `json:"amount"`
	Qty        string `json:"qty"`
}

// Fields of the Deliver.
func (Deliver) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}),
		field.String("user_id").Nillable(),
		field.String("post_id").Nillable(),
		field.JSON("normal_orders", []DeliverOrder{}),
		field.JSON("extra_orders", []DeliverOrder{}),
		field.Float("normal_total").Default(0).Nillable(),
		field.Float("normal_fee").Default(0).Nillable(),
		field.Float("extra_total").Default(0).Nillable(),
		field.Float("extra_fee").Default(0).Nillable(),
		field.Time("created_at").Default(time.Now).Nillable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Nillable(),
	}
}

// Edges of the Deliver.
func (Deliver) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_delivers").Unique().Field("user_id").Required(),
		edge.From("post", Post.Type).Ref("post_delivers").Unique().Field("post_id").Required(),
	}
}

func (Deliver) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("post_id"),
		index.Fields("created_at"),
	}
}
