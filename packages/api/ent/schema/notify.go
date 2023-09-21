package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Notify holds the schema definition for the Notify entity.
type Notify struct {
	ent.Schema
}

// Fields of the Notify.
func (Notify) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}),
		field.String("user_id").Nillable(),
		field.String("line_token").Nillable(),
		field.String("fb_token").Nillable(),
		field.Time("created_at").Default(time.Now).Nillable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Nillable(),
	}
}

// Edges of the Notify.
func (Notify) Edges() []ent.Edge {
	return nil
}

func (Notify) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}
