package seedfuncs

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/billc-dev/tuango-go/ent"
	"github.com/gofiber/fiber/v2"
)

type JSONNotify struct {
	ID        string `json:"_id"`
	UserId    string `json:"username"`
	LineToken string `json:"token"`
	FbToken   string `json:"fbToken"`
}

func SeedNotifies(client *ent.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		jsonFile, err := os.Open("seeds/notifies.json")
		if err != nil {
			return err
		}
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		var notifies []JSONNotify
		json.Unmarshal(byteValue, &notifies)

		// delete
		client.Notify.Delete().Where().Exec(c.Context())

		users, err := client.User.Query().All(c.Context())
		if err != nil {
			return err
		}

		var bulkNotifies []*ent.NotifyCreate

		for _, notify := range notifies {
			u, err := getUserById(users, notify.UserId)
			if err != nil {
				log.Print(notify.UserId)
				u.ID = notify.UserId
			}

			// t, err := time.Parse("2006-01-02T15:04:05.000Z", time.Now().Format(time.RFC3339))
			// if err != nil {
			// 	return err
			// }

			bulkNotifies = append(bulkNotifies,
				client.Notify.
					Create().
					SetID(notify.ID).
					SetUserID(u.ID).
					SetLineToken(notify.LineToken).
					SetFbToken(notify.FbToken),
			)
		}

		batch := 2500

		for i := 0; i < len(bulkNotifies); i += batch {
			j := i + batch
			if j > len(bulkNotifies) {
				j = len(bulkNotifies)
			}
			_, err := client.Notify.CreateBulk(bulkNotifies[i:j]...).Save(c.Context())

			if err != nil {
				return err
			}
		}

		return c.JSON(fiber.Map{})
	}
}
