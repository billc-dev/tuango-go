package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

type Image struct {
	SM string `json:"sm"`
	MD string `json:"md"`
	LG string `json:"lg"`
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().DefaultFunc(func() string {
			return primitive.NewObjectID().Hex()
		}),
		field.String("seller_id").Nillable(),
		field.Int("post_num").Nillable(),
		field.String("title").Nillable(),
		field.Text("body").Nillable(),
		field.String("deadline").MaxLen(10).Default("").Nillable(),
		field.String("delivery_date").MaxLen(10).Default("").Nillable(),
		field.Int("like_count").Default(0).Nillable(),
		field.Int("comment_count").Default(0).Nillable(),
		field.Int("order_count").Default(0).Nillable(),
		field.JSON("images", []Image{}),
		field.Enum("storage_type").Values("roomTemp", "farmGoods", "refrigerated", "frozen").
			Nillable(),
		field.Enum("status").Values("open", "closed", "completed", "canceled").Default("open").
			Nillable(),
		field.String("comment").Default("").Nillable(),
		field.Bool("delivered").Default(false).Nillable(),
		field.Bool("is_in_stock").Default(false).Nillable(),
		field.Float("normal_total").Default(0).Nillable(),
		field.Float("normal_fee").Default(0).Nillable(),
		field.Float("extra_total").Default(0).Nillable(),
		field.Float("extra_fee").Default(0).Nillable(),
		field.Time("created_at").Default(time.Now).Nillable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Nillable(),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("seller", User.Type).Ref("posts").Unique().Field("seller_id").Required(),
		edge.To("post_comments", Comment.Type),
		edge.To("post_delivers", Deliver.Type),
		edge.To("post_items", PostItem.Type),
		edge.To("post_likes", Like.Type),
		edge.To("post_orders", Order.Type),
	}
}

func (Post) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("seller_id"),
		index.Fields("post_num"),
		index.Fields("created_at"),
	}
}
