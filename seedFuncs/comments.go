package seedfuncs

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/billc-dev/tuango-go/ent"
	"github.com/billc-dev/tuango-go/ent/schema"
	"github.com/gofiber/fiber/v2"
)

type JSONComment struct {
	ID          string         `json:"_id"`
	UserId      string         `json:"userId"`
	PostId      string         `json:"postId"`
	DisplayName string         `json:"displayName"`
	Comment     string         `json:"comment"`
	CreatedAt   string         `json:"createdAt"`
	Replies     []schema.Reply `json:"replies"`
}

func SeedComments(client *ent.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		jsonFile, err := os.Open("seeds/comments.json")

		if err != nil {
			return err
		}
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		var orders []JSONComment
		json.Unmarshal(byteValue, &orders)

		// delete
		client.Comment.Delete().Where().ExecX(c.Context())

		users, err := client.User.Query().All(c.Context())
		if err != nil {
			return err
		}

		var bulkComments []*ent.CommentCreate

		posts, err := client.Post.Query().All(c.Context())
		if err != nil {
			return err
		}

		for _, order := range orders {
			var found bool
			for _, p := range posts {
				if p.ID == order.PostId {
					found = true
					break
				}
			}
			if !found {
				continue
			}
			t, err := time.Parse("2006-01-02T15:04:05.000Z", order.CreatedAt)
			if err != nil {
				return err
			}
			u, err := getUserById(users, order.UserId)
			if err != nil {
				return err
			}

			bulkComments = append(bulkComments,
				client.Comment.
					Create().
					SetID(order.ID).
					SetUserID(u.ID).
					SetPostID(order.PostId).
					SetComment(order.Comment).
					SetReplies(order.Replies).
					SetCreatedAt(t),
			)

		}
		batch := 2500

		for i := 0; i < len(bulkComments); i += batch {
			j := i + batch
			if j > len(bulkComments) {
				j = len(bulkComments)
			}
			_, err := client.Comment.CreateBulk(bulkComments[i:j]...).Save(c.Context())

			if err != nil {
				return err
			}
		}

		return c.JSON(fiber.Map{})
	}
}
