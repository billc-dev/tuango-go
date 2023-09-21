package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RoomUser holds the schema definition for the RoomUser entity.
type RoomUser struct {
	ent.Schema
}

// Fields of the RoomUser.
func (RoomUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}),
		field.String("room_id").Nillable(),
		field.String("user_id").Nillable(),
		field.String("last_read_message_id").Nillable(),
	}
}

// Edges of the RoomUser.
func (RoomUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).Ref("room_users").Unique().Field("room_id").Required(),
		edge.From("user", User.Type).Ref("user_rooms").Unique().Field("user_id").Required(),
	}
}

func (RoomUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("room_id", "user_id").Unique(),
	}
}
