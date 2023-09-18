package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}),
		field.String("username").Unique().Nillable(),
		field.String("display_name").Nillable(),
		field.String("picture_url").Nillable(),
		field.Float("pickup_num").Nillable(),
		field.Enum("role").Values("basic", "seller", "admin").Default("basic").
			Nillable(),
		field.Enum("status").Values("registered", "approved", "blocked").Default("registered").
			Nillable(),
		field.Bool("notified").Default(false).Nillable(),
		field.Bool("line_pay").Default(false).Nillable(),
		field.Bool("fb").Default(false).Nillable(),
		field.String("comment").Default("").Nillable(),
		field.Int("delivered_order_count_limit").Default(10).Nillable(),
		field.Time("created_at").Default(time.Now).Nillable().Nillable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("user_comments", Comment.Type),
		edge.To("user_completes", Complete.Type),
		edge.To("user_delivers", Deliver.Type),
		edge.To("user_likes", Like.Type),
		edge.To("user_messages", Message.Type),
		edge.To("user_orders", Order.Type),
		edge.To("user_rooms", RoomUser.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username"),
		index.Fields("pickup_num"),
	}
}
