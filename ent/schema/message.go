package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}),
		field.String("room_id").Nillable(),
		field.String("user_id").Nillable(),
		field.String("post_id").Optional().Nillable(),
		field.String("order_id").Optional().Nillable(),
		field.String("text").Optional().Nillable(),
		field.JSON("image", Image{}).Optional(),
		field.Bool("unsent").Default(false).Nillable(),
		field.Enum("type").Values("text", "imageUrl", "post", "order", "deliver", "complete").Nillable(),
		field.Time("created_at").Default(time.Now).Nillable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Nillable(),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).Ref("room_messages").Unique().Field("room_id").Required(),
		edge.From("user", User.Type).Ref("user_messages").Unique().Field("user_id").Required(),
	}
}
