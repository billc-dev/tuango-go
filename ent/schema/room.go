package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Room holds the schema definition for the Room entity.
type Room struct {
	ent.Schema
}

// Fields of the Room.
func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}),
		field.Enum("type").Values("private", "group"),
		field.String("name").Nillable(),
		field.String("last_message").Nillable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Nillable(),
	}
}

// Edges of the Room.
func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("room_messages", Message.Type),
		edge.To("room_users", RoomUser.Type),
	}
}
