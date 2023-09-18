package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Complete holds the schema definition for the Complete entity.
type Complete struct {
	ent.Schema
}

type CompleteOrder struct {
	OrderId           string `json:"orderId"`
	OrderNum          uint   `json:"orderNum"`
	PostId            string `json:"postId"`
	PostNum           uint   `json:"postNum"`
	Title             string `json:"title"`
	SellerDisplayName string `json:"sellerDisplayName"`
	Order             []struct {
		Identifier string  `json:"id"`
		Name       string  `json:"item"`
		Price      float64 `json:"price"`
		Qty        float64 `json:"qty"`
		Location   string  `json:"location"`
	} `json:"order"`
}

// Fields of the Complete.
func (Complete) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}),
		field.String("user_id").Nillable(),
		field.Float("total").Nillable(),
		field.String("admin").Nillable(),
		field.Bool("line_pay").Default(false).Nillable(),
		field.Bool("confirmed").Default(false).Nillable(),
		field.JSON("orders", []CompleteOrder{}),
		field.Time("created_at").Default(time.Now).Nillable(),
	}
}

// Edges of the Complete.
func (Complete) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_completes").Unique().Field("user_id").Required(),
	}
}

func (Complete) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("created_at"),
	}
}
