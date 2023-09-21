package seedfuncs

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/billc-dev/tuango-go/ent"
	"github.com/gofiber/fiber/v2"
)

type JSONLike struct {
	ID        string `json:"_id"`
	UserId    string `json:"userId"`
	PostId    string `json:"postId"`
	CreatedAt string `json:"createdAt"`
}

func SeedLikes(client *ent.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		jsonFile, err := os.Open("seeds/likes.json")

		if err != nil {
			return err
		}
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		var likes []JSONLike
		json.Unmarshal(byteValue, &likes)

		// delete
		client.Like.Delete().Where().ExecX(c.Context())

		users, err := client.User.Query().All(c.Context())
		if err != nil {
			return err
		}

		var bulkLikes []*ent.LikeCreate

		posts, err := client.Post.Query().All(c.Context())
		if err != nil {
			return err
		}

		for _, like := range likes {
			var found bool
			for _, p := range posts {
				if p.ID == like.PostId {
					found = true
					break
				}
			}
			if !found {
				log.Print(like.ID)
				continue
			}
			t, err := time.Parse("2006-01-02T15:04:05.000Z", like.CreatedAt)
			if err != nil {
				return err
			}
			u, err := getUserById(users, like.UserId)
			if err != nil {
				return err
			}

			bulkLikes = append(bulkLikes,
				client.Like.
					Create().
					SetID(like.ID).
					SetUserID(u.ID).
					SetPostID(like.PostId).
					SetCreatedAt(t),
			)
		}
		batch := 2500

		for i := 0; i < len(bulkLikes); i += batch {
			j := i + batch
			if j > len(bulkLikes) {
				j = len(bulkLikes)
			}
			_, err := client.Like.CreateBulk(bulkLikes[i:j]...).Save(c.Context())

			if err != nil {
				return err
			}
		}

		return c.JSON(fiber.Map{})
	}
}
