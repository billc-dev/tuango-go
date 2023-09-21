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

type JSONComplete struct {
	ID          string                 `json:"_id"`
	UserId      string                 `json:"userId"`
	DisplayName string                 `json:"displayName"`
	Total       float64                `json:"total"`
	Admin       string                 `json:"admin"`
	CreatedAt   string                 `json:"createdAt"`
	Orders      []schema.CompleteOrder `json:"orders"`
	Payment     struct {
		LinePay   bool `json:"linePay"`
		Confirmed bool `json:"confirmed"`
	} `json:"payment"`
}

func SeedCompletes(client *ent.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		jsonFile, err := os.Open("seeds/completes.json")
		if err != nil {
			return err
		}
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		var completes []JSONComplete
		json.Unmarshal(byteValue, &completes)

		// delete
		client.Complete.Delete().Where().Exec(c.Context())

		users, err := client.User.Query().All(c.Context())
		if err != nil {
			return err
		}

		var bulkCompletes []*ent.CompleteCreate

		for _, complete := range completes {
			u, err := getUserById(users, complete.UserId)
			if err != nil {
				return err
			}

			t, err := time.Parse("2006-01-02T15:04:05.000Z", complete.CreatedAt)
			if err != nil {
				return err
			}

			bulkCompletes = append(bulkCompletes,
				client.Complete.
					Create().
					SetID(complete.ID).
					SetUserID(u.ID).
					SetTotal(complete.Total).
					SetAdmin(complete.Admin).
					SetLinePay(complete.Payment.LinePay).
					SetConfirmed(complete.Payment.Confirmed).
					SetOrders(complete.Orders).
					SetCreatedAt(t),
			)
		}

		batch := 2500

		for i := 0; i < len(bulkCompletes); i += batch {
			j := i + batch
			if j > len(bulkCompletes) {
				j = len(bulkCompletes)
			}
			_, err := client.Complete.CreateBulk(bulkCompletes[i:j]...).Save(c.Context())

			if err != nil {
				return err
			}
		}

		return c.JSON(fiber.Map{})
	}
}
