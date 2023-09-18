package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/billc-dev/tuango-go/ent/order"
	"github.com/billc-dev/tuango-go/ent/post"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

type NotificationQuery struct {
	PostNum     int              `json:"post_num"`
	Status      order.Status     `json:"status"`
	StorageType post.StorageType `json:"storage_type"`
	Title       string           `json:"title"`
	DisplayName string           `json:"display_name"`
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}),
		field.String("text").Nillable(),
		field.JSON("query", NotificationQuery{}),
		field.Strings("users"),
		field.Time("created_at").Default(time.Now).Nillable(),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return nil
}

func (Notification) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
	}
}
