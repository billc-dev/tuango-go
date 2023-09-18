package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Like holds the schema definition for the Like entity.
type Like struct {
	ent.Schema
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}),
		field.String("user_id").Nillable(),
		field.String("post_id").Nillable(),
		field.Time("created_at").Default(time.Now).Nillable(),
	}
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_likes").Unique().Field("user_id").Required(),
		edge.From("post", Post.Type).Ref("post_likes").Unique().Field("post_id").Required(),
	}
}

func (Like) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("post_id"),
	}
}
