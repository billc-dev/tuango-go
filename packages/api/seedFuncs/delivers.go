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

type JSONDeliver struct {
	ID           string                `json:"_id"`
	UserId       string                `json:"userId"`
	PostId       string                `json:"postId"`
	PostNum      int                   `json:"postNum"`
	NormalTotal  float64               `json:"normalTotal"`
	NormalFee    float64               `json:"normalFee"`
	ExtraTotal   float64               `json:"extraTotal"`
	ExtraFee     float64               `json:"extraFee"`
	NormalOrders []schema.DeliverOrder `json:"normalOrders"`
	ExtraOrders  []schema.DeliverOrder `json:"extraOrders"`
	CreatedAt    string                `json:"createdAt"`
}

func SeedDelivers(client *ent.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		jsonFile, err := os.Open("seeds/delivers.json")
		if err != nil {
			return err
		}
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		var delivers []JSONDeliver
		json.Unmarshal(byteValue, &delivers)

		// delete
		client.Deliver.Delete().Where().Exec(c.Context())

		users, err := client.User.Query().All(c.Context())
		if err != nil {
			return err
		}

		posts, err := client.Post.Query().All(c.Context())
		if err != nil {
			return err
		}

		var bulkDelivers []*ent.DeliverCreate

		for _, deliver := range delivers {
			u, err := getUserById(users, deliver.UserId)
			if err != nil {
				return err
			}

			var found bool
			for _, p := range posts {
				if p.ID == deliver.PostId {
					found = true
					break
				}
			}
			if !found {
				// log.Print(deliver.PostNum, " ",
				// 	deliver.NormalTotal+deliver.ExtraTotal, " ",
				// 	deliver.ExtraTotal+deliver.ExtraFee)
				continue
			}

			t, err := time.Parse("2006-01-02T15:04:05.000Z", deliver.CreatedAt)
			if err != nil {
				return err
			}

			bulkDelivers = append(bulkDelivers,
				client.Deliver.
					Create().
					SetID(deliver.ID).
					SetUserID(u.ID).
					SetPostID(deliver.PostId).
					SetNormalOrders(deliver.NormalOrders).
					SetExtraOrders(deliver.ExtraOrders).
					SetNormalTotal(deliver.NormalTotal).
					SetNormalFee(deliver.NormalFee).
					SetExtraTotal(deliver.ExtraTotal).
					SetExtraFee(deliver.ExtraFee).
					SetCreatedAt(t),
			)
		}

		batch := 2500

		for i := 0; i < len(bulkDelivers); i += batch {
			j := i + batch
			if j > len(bulkDelivers) {
				j = len(bulkDelivers)
			}
			_, err := client.Deliver.CreateBulk(bulkDelivers[i:j]...).Save(c.Context())

			if err != nil {
				return err
			}
		}

		return c.JSON(fiber.Map{})
	}
}
