package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

type Reply struct {
	UserId      string `json:"userId"`
	DisplayName string `json:"displayName"`
	PictureUrl  string `json:"pictureUrl"`
	Reply       string `json:"reply"`
	CreatedAt   string `json:"createdAt"`
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}),
		field.String("user_id").Nillable(),
		field.String("post_id").Nillable(),
		field.Text("comment").Nillable(),
		field.JSON("replies", []Reply{}),
		field.Time("created_at").Default(time.Now).Nillable(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_comments").Unique().Field("user_id").Required(),
		edge.From("post", Post.Type).Ref("post_comments").Unique().Field("post_id").Required(),
	}
}

func (Comment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("post_id"),
		index.Fields("created_at"),
	}
}
