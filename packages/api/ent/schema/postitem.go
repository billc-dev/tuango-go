package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PostItem holds the schema definition for the PostItem entity.
type PostItem struct {
	ent.Schema
}

// Fields of the PostItem.
func (PostItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}),
		field.String("post_id").Nillable(),
		field.String("identifier").Nillable(),
		field.String("name").Nillable(),
		field.Float("price").Nillable(),
		field.Float("stock").Nillable(),
	}
}

// Edges of the PostItem.
func (PostItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).Ref("post_items").Unique().Field("post_id").Required(),
		edge.To("post_item", OrderItem.Type),
	}
}

func (PostItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("post_id"),
	}
}
